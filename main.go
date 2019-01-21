package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"xpool/controller"
	"xpool/database"
	"io"
	"net/http"
	"os"
	"unicode"
	"gopkg.in/bluesuncorp/validator.v5"
	"xpool/models"
)

const defaultPort = "8080"

var (
	msgInvalidJSON     = "Invalid JSON format"
	msgInvalidJSONType = func(e *json.UnmarshalTypeError) string {
		return "Expected " + e.Value + " but given type is " + e.Type.String() + " in JSON"
	}
	msgValidationFailed = func(e *validator.FieldError) string {
		switch e.Tag {
		case "required":
			return toSnakeCase(e.Field) + ": required"
		case "max":
			return toSnakeCase(e.Field) + ": too_long"
		default:
			return e.Error()
		}
	}
)

func main() {
	migrate()
	router := gin.Default()

	HomeGroup := router.Group("/user")
	{
		HomeGroup.POST("/createuser",apiHandle("CreateUser"), controller.User.CreateUser)
		HomeGroup.POST("/getverificationcode",apiHandle("GetVerificationCode"), controller.User.GetVerificationCode)
		HomeGroup.POST("/login",apiHandle("Login"), controller.User.Login)
		HomeGroup.POST("/getuserbytoken",apiHandle("GetUserByToken"), controller.User.GetUserByToken)
		HomeGroup.POST("/forgetpassword",apiHandle("ForgetPassword"), controller.User.ForgetPassword)
	}

	http.ListenAndServe(":"+port(), router)
}

type resultdata struct {
	Code int `json:"code"`
	Status int `json:"status"`
	Data interface{} `json:"data"`
}

func result(data interface{},status int,code int) resultdata {
	result := resultdata{}
	result.Code = code
	result.Status = status
	result.Data = data
	return result
}




func apiHandle(authority string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//需要忽略验证的模块
		ignoreValidation := map[string] int {
			"CreateUser":1 ,
			"GetVerificationCode":1,
			"Login":1,
			"ForgetPassword":1,
		}
		Token :=  c.Param("token")
		if Token != "" {
			//校验token


		} else if 1 != ignoreValidation[authority]{
			data :=result("",0,0)
			c.JSON(http.StatusOK, data)
			c.Abort()
			return
		}
		c.Next()
		errs := make([]string, 0, len(c.Errors))
		for _, e := range c.Errors {
			switch e.Err {
			case io.EOF:
				errs = append(errs, msgInvalidJSON)
			default:
				switch e.Err.(type) {
				case *json.SyntaxError:
					errs = append(errs, msgInvalidJSON)
				case *json.UnmarshalTypeError:
					errs = append(errs, msgInvalidJSONType(e.Err.(*json.UnmarshalTypeError)))
				case *validator.StructErrors:
					for _, fieldErr := range e.Err.(*validator.StructErrors).Flatten() {
						errs = append(errs, msgValidationFailed(fieldErr))
					}
				default:
					errs = append(errs, e.Error())
				}
			}
		}

		if len(c.Errors) > 0 {
			c.JSON(-1, gin.H{"errors": errs}) // -1 == not override the current error code
		}
	}
}

// https://gist.github.com/elwinar/14e1e897fdbe4d3432e1
func toSnakeCase(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	return port
}

func migrate() {
	db := database.GetDB()
	db.AutoMigrate(&models.User{},&models.VerificationCode{},&models.Token{})
}
