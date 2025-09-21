package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)


func (m *Middlewares) AutenticateJWT(next http.Handler) http.Handler{
 return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	header := r.Header.Get("Authorization")
	if header == ""{
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	headerArr := strings.Split(header, " ")
	if len(headerArr) != 2 || headerArr[0] != "Bearer"{
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	accessToken := headerArr[1]
	//fmt.Println("Access Token: ", accessToken)
	tokenPrats := strings.Split(accessToken, ".")
	//fmt.Println("Token Parts: ", tokenPrats)
	if len(tokenPrats) != 3{
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	jwtHeader := tokenPrats[0]
	jwtPayload := tokenPrats[1]
	jwtSignature := tokenPrats[2]
	message := jwtHeader + "." + jwtPayload

	// cng := config.GetConfig()
	
	byteArrScret := []byte(m.cnf.JwtSecretKey)

	byteArrMessage := []byte(message)

	//fmt.Println("Byte array message: ", byteArrMessage)
	h := hmac.New(sha256.New,byteArrScret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)
	newSignature := base64UrlEncode(signature)
	//fmt.Println("Signature B64: ", signatureB64)
	if newSignature != jwtSignature{
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	next.ServeHTTP(w, r)
 })
}
func base64UrlEncode(data []byte) string {
		return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
	}