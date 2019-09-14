package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
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

const defaultPort = "8081"
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
	router.Use(Cors())

	UserGroup := router.Group("/user")
	{
		UserGroup.POST("/createuser",apiHandle("CreateUser"), controller.User.CreateUser)
		UserGroup.POST("/getverificationcode",apiHandle("GetVerificationCode"), controller.User.GetVerificationCode)
		UserGroup.POST("/login",apiHandle("Login"), controller.User.Login)
		UserGroup.POST("/getuserbytoken",apiHandle("GetUserByToken"), controller.User.GetUserByToken)
		UserGroup.POST("/forgetpassword",apiHandle("ForgetPassword"), controller.User.ForgetPassword)
		UserGroup.POST("/resetpassword",apiHandle("ResetPassword"), controller.User.ResetPassword)
		UserGroup.POST("/logout",apiHandle("Logout"), controller.User.Logout)
	}

	DepositGroup := router.Group("/deposit")
	{
		DepositGroup.POST("/adddeposit",apiHandle("AddDeposit"), controller.Deposit.AddDeposit)
		DepositGroup.POST("/getdepositlist",apiHandle("GetDepositList"), controller.Deposit.GetDepositList)
		DepositGroup.POST("/admingetdepositlist",apiHandle("AdminGetDepositList"), controller.Deposit.AdminGetDepositList)
		DepositGroup.POST("/depositreview",apiHandle("DepositReview"), controller.Deposit.DepositReview)
		DepositGroup.POST("/extractdeposit",apiHandle("ExtractDeposit"), controller.Deposit.ExtractDeposit)
		DepositGroup.POST("/extractdepositreview",apiHandle("ExtractDepositReview"), controller.Deposit.ExtractDepositReview)
		DepositGroup.POST("/getextractdepositlist",apiHandle("GetExtractDepositList"), controller.Deposit.GetExtractDepositList)
		DepositGroup.POST("/admingetextractdepositlist",apiHandle("AdminGetExtractDepositList"), controller.Deposit.AdminGetExtractDepositList)
		DepositGroup.POST("/deposit_balance",apiHandle("DepositBalance"), controller.Deposit.DepositBalance)
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
		MiningGroup.POST("/admingetloanmininglist",apiHandle("AdminGetLoanMiningList"), controller.Mining.AdminGetLoanMiningList)
		MiningGroup.POST("/admingetextractloanmininglist",apiHandle("AdminGetExtractLoanMiningList"), controller.Mining.AdminGetExtractLoanMiningList)
		MiningGroup.POST("/user_loan_mining_balance",apiHandle("UserLoanMiningBalance"), controller.Mining.UserLoanMiningBalance)
	}

	BalanceGroup := router.Group("/balance")
	{
		BalanceGroup.POST("/extractbalance",apiHandle("ExtractBalance"), controller.Balance.ExtractBalance)
		BalanceGroup.POST("/extractbalancereview",apiHandle("ExtractBalanceReview"), controller.Balance.ExtractBalanceReview)
		BalanceGroup.POST("/extractbalancelist",apiHandle("ExtractBalanceList"), controller.Balance.ExtractBalanceList)
		BalanceGroup.POST("/adminextractbalancelist",apiHandle("AdminExtractBalanceList"), controller.Balance.AdminExtractBalanceList)
	}

	IncomeGroup := router.Group("/income")
	{
		IncomeGroup.POST("/income_total",apiHandle("IncomeTotal"), controller.Income.IncomeTotal)
		IncomeGroup.POST("/income_balance",apiHandle("IncomeBalance"), controller.Income.IncomeBalance)
		IncomeGroup.POST("/extract_income_balance",apiHandle("ExtractIncomeBalance"), controller.Income.ExtractIncomeBalance)
		IncomeGroup.POST("/extract_income_list",apiHandle("ExtractIncomeList"), controller.Income.ExtractIncomeList)
		IncomeGroup.POST("/admin_extract_income_list",apiHandle("ExtractIncomeList"), controller.Income.AdminExtractIncomeList)
		IncomeGroup.POST("/extract_income_review",apiHandle("ExtractIncomeReview"), controller.Income.ExtractIncomeReview)
	}

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	go controller.BenefitCalculation()
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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method      //请求方法
		origin := c.Request.Header.Get("Origin")        //请求头部
		var headerKeys []string                             // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")       // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()        //  处理请求
	}
}

func migrate() {
	db := database.GetDB()
	db.AutoMigrate(&models.User{},&models.VerificationCode{},&models.Token{},&models.Deposit{},
	&models.UserDepositBalance{}, &models.DepositOperatingLog{},&models.ExtractDeposit{},
	&models.LoanMining{},&models.LoanMiningLog{},&models.UserLoanMiningBalance{},&models.ExtractLoanMiningBalance{},
	&models.UserBalance{},&models.ExtractBalance{},&models.ExtractBalanceLog{},
)
}
