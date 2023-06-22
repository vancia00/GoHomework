package JwtToken

import (
	"GoHomework/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySigningKey = []byte("vancia00.com")

func GenRegisteredClaims(username string) (string, error) {
	claims := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //过期时间
			Issuer:    "vancia00",                            //发签人
		},
	}
	//生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//jwt.SigningMethodES256是使用ECDSA算法进行签名的方法，需要传递的是ECDSA私钥，而不是普通的[]byte密钥
	//所以用[]byte密钥的时候就不用ES256 而是用HS256

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("token生成失败")
		return "", err
	}
	return tokenString, nil
}
