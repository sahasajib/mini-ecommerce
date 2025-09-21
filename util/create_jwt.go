package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct{
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct{
	Sub int `json:"sub"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	IsShopOwner bool `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error){
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	jsonHeader, err := json.Marshal(header)
	if err != nil{
		return  "", err
	}
	headerB64 := base64UrlEncode(jsonHeader)
	jsonPayload, err := json.Marshal(data)
	if err != nil{
		return  "", err
	}
	payloadB64 := base64UrlEncode(jsonPayload)

	message := headerB64 + "." + payloadB64
	byteArrScret := []byte(secret)

	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrScret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)
	signatureB64 := base64UrlEncode(signature)

	jwt := message + "." + signatureB64
	return jwt, nil

}
func base64UrlEncode(data []byte) string {
		return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
	}