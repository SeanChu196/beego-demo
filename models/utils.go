package models

import "regexp"

var (
	phoneRegexp     *regexp.Regexp
	emailRegexp     *regexp.Regexp
	upperCaseRegexp *regexp.Regexp
	lowerCaseRegexp *regexp.Regexp
	numRegexp       *regexp.Regexp
	usrDict         map[string]string
)

func init() {
	emailRegexp = regexp.MustCompile(`^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`)
	phoneRegexp = regexp.MustCompile(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`)
	upperCaseRegexp = regexp.MustCompile(`^(.*)[A-Z]+(.)*$`)
	lowerCaseRegexp = regexp.MustCompile(`^(.*)[a-z]+().*$`)
	numRegexp = regexp.MustCompile(`^(.)*[0-9]+(.)*$`)
	usrDict = make(map[string]string, 20)
	usrDict["18018122236"] = "12345yutGH"
}

//SignUp 用户注册,保存用户信息
func SignUp(usr, pwd string) bool {
	if usrDict[usr] == "" {
		usrDict[usr] = pwd
		return true
	}
	return false
}

//GetPwd 获取用户密码
func GetPwd(usr string) string {
	return usrDict[usr]
}

//VerifyPwd 验证用户名和密码
func VerifyPwd(usr, pwd string) bool {
	return usrDict[usr] != "" && usrDict[usr] == pwd
}

//IsValidPhone 验证手机号格式
func IsValidPhone(usr string) bool {
	return phoneRegexp.MatchString(usr)
}

//IsValidEmail 验证邮箱格式
func IsValidEmail(email string) bool {
	return emailRegexp.MatchString(email)
}

//IsValidPwd 验证密码格式
func IsValidPwd(pwd string) bool {
	if len(pwd) >= 8 && len(pwd) <= 16 {
		if upperCaseRegexp.MatchString(pwd) && lowerCaseRegexp.MatchString(pwd) && numRegexp.MatchString(pwd) {
			return true
		}
	}
	return false
}

//IsValidUsr 验证用户名格式
func IsValidUsr(usr string) bool {
	return IsValidEmail(usr) || IsValidPhone(usr)
}
