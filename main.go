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
	"github.com/gin-contrib/cors"
)

const defaultPort = "8080"
//const defaultPort = "8081"

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

	UserGroup := router.Group("/user")
	{
		UserGroup.POST("/createuser",apiHandle("CreateUser"), controller.User.CreateUser)
		UserGroup.POST("/getverificationcode",apiHandle("GetVerificationCode"), controller.User.GetVerificationCode)
		UserGroup.POST("/login",apiHandle("Login"), controller.User.Login)
		UserGroup.POST("/getuserbytoken",apiHandle("GetUserByToken"), controller.User.GetUserByToken)
		UserGroup.POST("/forgetpassword",apiHandle("ForgetPassword"), controller.User.ForgetPassword)
		UserGroup.POST("/resetpassword",apiHandle("ResetPassword"), controller.User.ResetPassword)
	}

	DepositGroup := router.Group("/deposit")
	{
		DepositGroup.POST("/adddeposit",apiHandle("AddDeposit"), controller.Deposit.AddDeposit)
		DepositGroup.POST("/getdepositlist",apiHandle("GetDepositList"), controller.Deposit.GetDepositList)
		DepositGroup.POST("/depositreview",apiHandle("DepositReview"), controller.Deposit.DepositReview)
		DepositGroup.POST("/extractdeposit",apiHandle("ExtractDeposit"), controller.Deposit.ExtractDeposit)
		DepositGroup.POST("/extractdepositreview",apiHandle("ExtractDepositReview"), controller.Deposit.ExtractDepositReview)
		DepositGroup.POST("/getextractdepositlist",apiHandle("GetExtractDepositList"), controller.Deposit.GetExtractDepositList)
	}


	MiningGroup := router.Group("/mining")
	{
		MiningGroup.POST("/loanmining",apiHandle("LoanMining"), controller.Mining.LoanMining)
		MiningGroup.POST("/loanminingreview",apiHandle("LoanMiningReview"), controller.Mining.LoanMiningReview)
		MiningGroup.POST("/isbindingminingaddress",apiHandle("IsBindingMiningAddress"), controller.Mining.IsBindingMiningAddress)
		MiningGroup.POST("/extractloanmining",apiHandle("ExtractLoanMining"), controller.Mining.ExtractLoanMining)
		MiningGroup.POST("/extractloanminingreview",apiHandle("ExtractLoanMiningReview"), controller.Mining.ExtractLoanMiningReview)
		MiningGroup.POST("/getloanmininglist",apiHandle("GetLoanMiningList"), controller.Mining.GetLoanMiningList)
		MiningGroup.POST("/getextractloanmininglist",apiHandle("GetExtractLoanMiningList"), controller.Mining.GetExtractLoanMiningList)

	}

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

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
		Token :=  c.PostForm("token")
		if "" != Token {
			//校验token
			userinfo := controller.GetUserInfoByToken(Token)
			if "" == userinfo.Email {
				c.JSON(http.StatusOK, controller.ResponseFun("token无效请登录",401))
				c.Abort()
				return
			}
		} else if 1 != ignoreValidation[authority]{
			c.JSON(http.StatusOK, controller.ResponseFun("请登录",401))
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
	db.AutoMigrate(&models.User{},&models.VerificationCode{},&models.Token{},&models.Deposit{},
	&models.UserDepositBalance{}, &models.DepositOperatingLog{},&models.ExtractDeposit{},
	&models.LoanMining{},&models.LoanMiningLog{},&models.UserLoanMiningBalance{},&models.ExtractLoanMiningBalance{})
}
