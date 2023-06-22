package api

import (
	"GoHomework/dao"
	"GoHomework/global"
	"GoHomework/model"
	"GoHomework/tools/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteComment(c *gin.Context) {
	var deletecomment model.Delete
	err := c.ShouldBindJSON(&deletecomment)
	if err != nil {
		Response.ResFail(c, "参数绑定失败")
		return
	}
	//进行身份验证
	Rusername, flag := c.Get("username")
	if !flag {
		Response.ResFail(c, "无法获取用户名!")
		return
	}
	username, flag := Rusername.(string)
	if !flag {
		Response.ResFail(c, "用户名类型断言失败!")
		return
	}
	flag = dao.DCAuthIdentity(username, deletecomment.ID)
	if !flag {
		Response.ResFail(c, "无权限进行此操作!")
		return
	}
	dao.CDelete(deletecomment.ID)
	Response.ResSuccess(c, "删除回答成功!")

}

func DeleteQuestion(c *gin.Context) {
	var deletequestion model.Delete
	err := c.ShouldBindJSON(&deletequestion)
	if err != nil {
		Response.ResFail(c, "参数绑定失败")
		return
	}
	//进行身份验证
	Rusername, flag := c.Get("username")
	if !flag {
		Response.ResFail(c, "无法获取用户名")
		return
	}
	username, flag := Rusername.(string)
	if !flag {
		Response.ResFail(c, "用户名类型断言失败")
		return
	}
	flag = dao.DQAuthIdentity(username, deletequestion.ID)
	if !flag {
		Response.ResFail(c, "无权限进行该操作！")
		return
	}
	dao.QDelete(deletequestion.ID)
	Response.ResSuccess(c, "删除成功！")
}

func ModifyComment(c *gin.Context) {
	var modifycomment model.ModifyComment
	err := c.ShouldBind(&modifycomment)
	if err != nil {
		Response.ResFail(c, "参数绑定失败！")
		return
	}
	Rusername, flag := c.Get("username")
	if !flag {
		Response.ResFail(c, "未能获取用户名！")
		return
	}
	username, flag := Rusername.(string)
	if !flag {
		Response.ResFail(c, "断言失败！")
		return
	}
	flag = dao.CAuthIdentity(username, modifycomment.ID)
	if !flag {
		Response.ResFail(c, "无权限修改此回答！")
		return
	}
	dao.CModify(modifycomment.Content, modifycomment.ID)
	Response.ResSuccess(c, "修改成功！")

}

func ModifyQuestion(c *gin.Context) {
	var modifyquestion model.ModifyQuestion
	err := c.ShouldBind(&modifyquestion)
	if err != nil {
		Response.ResFail(c, "参数绑定失败！")
		return
	}
	Rusername, flag := c.Get("username")
	if !flag {
		Response.ResFail(c, "无法获取用户名！")
		return
	}
	username, flag := Rusername.(string)
	if !flag {
		Response.ResFail(c, "断言失败！")
		return
	}
	flag = dao.QAuthIdentity(username, modifyquestion.ID)
	if !flag {
		Response.ResFail(c, "无权限进行此操作！")
		return
	}
	dao.QModify(modifyquestion.Content, modifyquestion.Title, modifyquestion.ID)
	Response.ResSuccess(c, "问题修改成功！")
}

func ViewQuestions(c *gin.Context) {
	Rusername, flag := c.Get("username")
	if !flag {
		Response.ResFail(c, "无法获取用户名！")
		return
	}
	username, flag := Rusername.(string)
	if !flag {
		Response.ResFail(c, "断言失败！")
		return
	}
	questions := dao.FindQuestions(username)
	comments := dao.FindComments(username)
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"questions": questions,
		"comments":  comments,
	})
}

func CreateComment(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBind(&comment); err != nil {
		Response.ResFail(c, "参数绑定失败！")
		return
	}
	Rusername, flag := c.Get("username")
	if !flag {
		Response.ResFail(c, "未能获取用户名！")
		return
	}
	username, flag := Rusername.(string)
	if !flag {
		Response.ResFail(c, "断言失败！")
		return
	}
	depositcomment := model.Comment{
		Commenter:  username,
		Content:    comment.Content,
		QuestionID: comment.QuestionID,
	}
	global.GloDb.Model(&model.Question{}).Create(&depositcomment)
	Response.ResSuccess(c, "回答成功！")

}

func CreatQuestion(c *gin.Context) {
	var Cquestion model.Question
	err := c.ShouldBind(&Cquestion)
	if err != nil {
		Response.ResFail(c, "参数绑定失败！")
		return
	}
	//questioner, err := dao.ExtractUsernameFromJWT() //传递令牌
	Rusername, flag := c.Get("username")
	if !flag {
		Response.ResFail(c, "未能获取用户名！")
		return
	}
	username, flag := Rusername.(string)
	if !flag {
		Response.ResFail(c, "断言失败！")
		return
	}
	depositCquestion := model.Question{
		Questioner: username,
		Title:      Cquestion.Title,
		Content:    Cquestion.Content,
	}
	global.GloDb.Model(&model.Question{}).Create(&depositCquestion)
	Response.ResSuccess(c, "提问成功！")
}
