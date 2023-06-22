package api

import (
	"GoHomework/tools/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/api/register", UserRegister)       //注册
	r.POST("/api/accountLogin", AccountLogin)   //账号密码登录
	r.POST("/api/emailLogin", EmailLogin)       //手机号登录
	r.POST("/api/phoneNumberLogin", PhoneLogin) //手机号码密码登录

	//问答社区
	QA := r.Group("/api/qa")
	{
		QA.Use(middleware.JWTAuthMiddleware())
		QA.POST("/createQst", CreatQuestion)    //提出问题
		QA.POST("/createAns", CreateComment)    //发起回答
		QA.GET("/ViewOwnQst", ViewQuestions)    //查看发布的问题和回答
		QA.PUT("/modifyQst", ModifyQuestion)    //修改问题
		QA.PUT("/modifyAns", ModifyComment)     //修改回答
		QA.DELETE("/deleteQst", DeleteQuestion) //删除问题
		QA.DELETE("/deleteAns", DeleteComment)  //删除回答
	}
	r.Run(":8080")

}
