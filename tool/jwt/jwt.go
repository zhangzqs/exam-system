package main

import (
	"crypto"
	"crypto/hmac"
	_ "crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

type H map[string]interface{}

func myJwt(secretKey string, uid int, exp int64) string {
	headerJson, _ := json.Marshal(H{
		"alg": "HS256", // 一种加密算法
		"typ": "JWT",
	})

	payloadJson, _ := json.Marshal(H{
		"exp": exp, // 过期时间
		"uid": 12,
	})

	headerBase64 := base64.StdEncoding.EncodeToString(headerJson)
	payloadBase64 := base64.StdEncoding.EncodeToString(payloadJson)
	signatureRaw := fmt.Sprintf("%s.%s", headerBase64, payloadBase64)
	m := crypto.SHA256
	if !m.Available() {
		// 一但执行log.Fatal程序就直接打印完这一句后退出了
		log.Fatalln("SHA256不可用")
	}
	hasher := hmac.New(m.New, []byte(secretKey))
	hasher.Write([]byte(signatureRaw))
	signature := base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))

	token := fmt.Sprintf("%s.%s", signatureRaw, signature)
	return token
}

func libJwt(secretKey string, uid int, exp int64) string {
	t, _ := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		struct {
			jwt.StandardClaims
			Uid int `json:"uid"`
		}{
			jwt.StandardClaims{
				ExpiresAt: exp,
			},
			uid,
		},
	).SignedString([]byte(secretKey))
	return t
}

func main() {
	uid := 12                               // 用户id
	exp := time.Now().Add(time.Hour).Unix() // 过期时间
	secretKey := "abc"                      // 密钥
	log.Println(myJwt(secretKey, uid, exp))
	log.Println(libJwt(secretKey, uid, exp))
}
