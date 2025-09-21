package main

import (
	"ecommerce/cmd"
	// "ecommerce/util"
	// "fmt"
)

func main() {
	cmd.Serve()
	// jwt, err:= util.CreateJwt("mysecret", util.Payload{
	// 	Sub: 123,
	// 	FirstName: "John",
	// 	LastName: "Doe",
	// 	Email: "john@doe.com",
	// 	IsShopOwner: true,
	// })
	// if err != nil{
	// 	fmt.Println("Error creating JWT:", err)	
	// 	return
	// }
	// fmt.Println("Generated JWT:", jwt)
}
