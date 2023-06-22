package dao

import (
	"GoHomework/boot"
	"GoHomework/global"
	"GoHomework/model"
	"regexp"
	"unicode"
)

func judgePassword() {

}

//以下是注册时所需函数

// 检查电话号码是否全为数字(长度在定义时已经做了限定     后续还需添加长度不正确的时候的提示代码)
func CheckPhoneNum(phoneNum string) bool {
	for _, char := range phoneNum {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

// 检验手机号是否已经注册过了(数据库里面查询)
func IsPhoneNumRepeat(phoneNum string) bool {
	var user model.User
	global.GloDb.Model(&model.User{}).Where("phoneNumber=?", phoneNum).Find(&user)
	if user.PhoneNumber == "" {
		return true
	}
	return false
}

// 检察邮箱格式是否正确  (还需要修改)
func CheckEmail(email string) bool {
	// 定义邮箱的正则表达式模式
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// 使用正则表达式匹配邮箱格式
	matched, _ := regexp.MatchString(pattern, email) //忽略了错误处理

	return matched
}

// 检查邮箱是否重复
func IsEmailRepeat(email string) bool {
	var user model.User
	global.GloDb.Model(&model.User{}).Where("email=?", email).Find(&user)
	if user.Email == "" {
		return true
	}
	return false
}

//以下是登陆时所需函数

func CheckUsername(username string) bool {
	var user model.User
	global.GloDb.Model(&model.User{}).Where("username=?", username).Find(&user)
	if user.UserName == "" {
		return false
	}
	return true
}

// 检查用户名和密码是否匹配
func IsUnameAPassMatched(username, password string) bool {
	var user model.User
	global.GloDb.Model(&model.User{}).Where("username = ?", username).Find(&user)
	flag := boot.ValidPassword(password, user.Salt, user.Password)
	return flag
}

func CheckEmailExist(email string) bool {
	var user model.User
	global.GloDb.Model(&model.User{}).Where("email=?", email).Find(&user)
	if user.Email == "" {
		return false
	}
	return true

}

func IsEmailAPassMatched(email, password string) bool {
	var user model.User
	global.GloDb.Model(&model.User{}).Where("email=?", email).Find(&user)
	if user.Password == password {
		return true
	}
	return false

}

func CheckPhoneNumExist(phoneNum string) bool {
	var user model.User
	global.GloDb.Model(&model.User{}).Where("phoneNumber=?", phoneNum).Find(&user)
	if user.PhoneNumber == "" {
		return false
	}
	return true
}

func IsPhoneNumAPassMatched(phoneNum, password string) bool {
	var user model.User
	global.GloDb.Model(&model.User{}).Where("phoneNumber=?", phoneNum).Find(&user)
	if user.Password == password {
		return true
	}
	return false

}

// 从邮箱中获取姓名
func GetUsernameFromEmail(email string) string {
	var user model.User
	global.GloDb.Model(&model.User{}).Where("email=?", email).Find(&user)
	return user.UserName
}

func GetUsernameFromPhoneNumber(phoneNumber string) string {
	var user model.User
	global.GloDb.Model(&model.User{}).Where("phoneNumber=?", phoneNumber).Find(&user)
	return user.UserName
}
