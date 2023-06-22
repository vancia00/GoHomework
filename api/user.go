package api

import (
	"GoHomework/boot"
	"GoHomework/dao"
	"GoHomework/global"
	"GoHomework/model"
	"GoHomework/tools/JwtToken"
	"GoHomework/tools/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PhoneLogin(c *gin.Context) {
	var user model.PhoneLogin
	err := c.ShouldBind(&user)
	if err != nil {
		Response.ResFail(c, "参数绑定失败！")
		return
	}
	flag := dao.CheckPhoneNum(user.PhoneNumber)
	if !flag {
		Response.ResFail(c, "手机号格式错误！")
		return
	}
	flag = dao.CheckPhoneNumExist(user.PhoneNumber)
	if !flag {
		Response.ResFail(c, "手机号未注册！")
		return
	}
	flag = dao.IsPhoneNumAPassMatched(user.PhoneNumber, user.Password)
	if !flag {
		Response.ResFail(c, "密码错误！")
		return
	}
	//JWT认证
	userName := dao.GetUsernameFromEmail(user.PhoneNumber)
	tokenString, err := JwtToken.GenRegisteredClaims(userName)
	if err != nil {
		Response.ResFail(c, "生成JWT失败!")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "success",
		"username": userName,
		"data":     gin.H{"token": tokenString},
	})
	Response.ResSuccess(c, "登陆成功！")

}

func EmailLogin(c *gin.Context) {
	var user model.EmailLogin
	err := c.ShouldBind(&user)
	if err != nil {
		Response.ResFail(c, "参数绑定失败！")
		return
	}
	flag := dao.CheckEmailExist(user.Email)
	if !flag {
		Response.ResFail(c, "邮箱不存在,请重新输入！")
		return
	}

	flag = dao.IsEmailAPassMatched(user.Email, user.Password)
	if !flag {
		Response.ResFail(c, "密码错误，请重新输入！")
		return
	}

	//JWT认证
	userName := dao.GetUsernameFromEmail(user.Email)
	tokenString, err := JwtToken.GenRegisteredClaims(userName)
	if err != nil {
		Response.ResFail(c, "生成JWT失败!")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "success",
		"username": userName,
		"data":     gin.H{"token": tokenString},
	})

	Response.ResSuccess(c, "登陆成功！")
}

func AccountLogin(c *gin.Context) {
	var userform model.UserNameLogin
	err := c.ShouldBind(&userform)
	if err != nil {
		Response.ResFail(c, "参数绑定失败！")
		return
	}

	flag := dao.CheckUsername(userform.UserName)
	if !flag {
		Response.ResFail(c, "该用户不存在！")
		return
	}
	//用加盐后的密码来进行比对
	flag = dao.IsUnameAPassMatched(userform.UserName, userform.Password)
	if !flag {
		Response.ResFail(c, "密码不正确,请重新输入！")
		return
	}

	tokenString, err := JwtToken.GenRegisteredClaims(userform.UserName)
	if err != nil {
		Response.ResFail(c, "生成JWT失败!")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "success",
		"username": userform.UserName,
		"data":     gin.H{"token": tokenString},
	})
	Response.ResSuccess(c, "登陆成功！")
}

// 没有对密码进行要求，只是再model层进行了长度限制
func UserRegister(c *gin.Context) {
	//获取form表单数据
	var registerForm model.UserRegister
	err := c.ShouldBind(&registerForm)
	if err != nil {
		Response.ResFail(c, "参数绑定失败！")
		return
	}

	//返回false就会执行
	flag := dao.CheckPhoneNum(registerForm.PhoneNumber)
	if !flag {
		Response.ResFail(c, "手机号必须全为数字！")
		return
	}

	flag = dao.IsPhoneNumRepeat(registerForm.PhoneNumber)
	if !flag {
		Response.ResFail(c, "手机号已被注册！")
		return
	}

	flag = dao.CheckEmail(registerForm.Email)
	if !flag {
		Response.ResFail(c, "邮箱格式不正确！")
		return
	}

	flag = dao.IsEmailRepeat(registerForm.Email)
	if !flag {
		Response.ResFail(c, "邮箱已注册！")
		return
	}
	randomNumber := boot.GenerateRandomNumber() //生成6-10的随机数
	salt := boot.GenValidateCode(randomNumber)
	registerForm.Password = boot.MakePassword(registerForm.Password, salt)
	//存入数据库
	user := model.User{
		UserName:    registerForm.UserName,
		Password:    registerForm.Password,
		Email:       registerForm.Email,
		PhoneNumber: registerForm.PhoneNumber,
		Salt:        salt,
	}
	global.GloDb.Model(&model.User{}).Create(&user)
	Response.ResSuccess(c, "注册成功！")

}
