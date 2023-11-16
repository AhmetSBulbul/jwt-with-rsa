package main

import (
	"fmt"
	"time"

	"github.com/kataras/jwt"
)

type FooClaims struct {
	Foo string `json:"foo"`
}

var (
	encKey = jwt.MustGenerateRandom(32)
	sigKey = jwt.MustGenerateRandom(32)
)

func main() {

	// prvKey, err := ioutil.ReadFile("cert/id_rsa")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// pubKey, err := ioutil.ReadFile("cert/id_rsa.pub")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// JWE

	// Generate a token which expires at 15 minutes from now:
	claims := FooClaims{
		Foo: "bar",
	} // can be a map too.

	encrypt, decrypt, err := jwt.GCM(encKey, nil)
	if err != nil {
		fmt.Printf("failed to create new GCM: %v", err)
	}

	// Encrypt and Sign the claims:
	tokenByte, err := jwt.SignEncrypted(jwt.HS256, sigKey, encrypt, claims, jwt.MaxAge(15*time.Minute))
	if err != nil {
		fmt.Printf("failed to sign encrypted token: %v", err)
	}

	token := string(tokenByte)

	fmt.Println("TOKEN:", token)
	fmt.Println("Secret:", string(encKey))

	verifiedToken, err := jwt.VerifyEncrypted(jwt.HS256, sigKey, decrypt, tokenByte)
	if err != nil {
		fmt.Printf("failed to verify encrypted token: %v", err)
	}

	fmt.Println("VERIFIED TOKEN:", string(verifiedToken.Payload))

	// JWS

	// jwtToken := token.NewJWT(prvKey, pubKey)
	// // fmt.Println("JWT:", jwtToken)

	// // 1. Create a new JWT token.
	// tok, err := jwtToken.Generate(time.Hour, "content")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("TOKEN:", tok)

	// // 2. Validate an existing JWT token.
	// content, err := jwtToken.Validate(tok)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("CONTENT:", content)
}
