package dao

import (
	"GoHomework/global"
	"GoHomework/model"
)

func FindQuestions(username string) []model.Question {
	var questions []model.Question
	//查询方法一
	global.GloDb.Model(&model.Question{}).Preload("Answers").Where("questioner=?", username).Find(&questions)
	//查询方法二
	//global.GlobalDb.Preload("Answers", func(d *gorm.DB) *gorm.DB {
	//	return d.Where("questioner=?", username)
	//}).Find(&questions)
	return questions
}

func FindComments(username string) []model.Comment {
	var answers []model.Comment
	global.GloDb.Model(&model.Comment{}).Where("Answerer = ?", username).Find(&answers)
	return answers
}

func QAuthIdentity(username string, ID uint) bool {
	var question model.Question
	global.GloDb.Model(&model.Question{}).Where("questioner=? AND id=?", username, ID).First(&question)
	if question.Questioner == username {
		return true
	}
	return false

}

func QModify(content, title string, id uint) {
	global.GloDb.Model(&model.Question{}).Where("id= ? ", id).
		Updates(model.Question{Content: content, Title: title})
}

func CAuthIdentity(username string, ID uint) bool {
	var comment model.Comment
	global.GloDb.Model(&model.Comment{}).Where("commenter = ? AND id = ? ", username, ID).First(&comment)
	if comment.Commenter == username {
		return true
	}
	return false
}

func CModify(content string, id uint) {
	global.GloDb.Model(&model.ModifyComment{}).Where("id = ?", id).Update("content", content)
}

func DQAuthIdentity(username string, ID uint) bool {
	var question model.Question
	global.GloDb.Model(&model.Question{}).Where("questioner=? AND id =?", username, ID).First(&question)
	if question.Questioner == username {
		return true
	}
	return false
}

func QDelete(id uint) {
	var question model.Question
	var comments []model.Comment
	global.GloDb.Model(&model.Comment{}).Unscoped().Where("questionID = ?", id).Delete(&comments)
	global.GloDb.Model(&model.Question{}).Unscoped().Where("id=?", id).Delete(&question)
}

func DCAuthIdentity(username string, ID uint) bool {
	var question model.Comment
	global.GloDb.Model(&model.Comment{}).Where("commenter=? AND id =?", username, ID).First(&question)
	if question.Commenter == username {
		return true
	}
	return false
}

func CDelete(id uint) {
	var comment model.Comment
	global.GloDb.Model(&model.Comment{}).Unscoped().Where("questionId=?", id).Delete(&comment)
}
