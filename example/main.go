package main

import (
	"log"
	"time"

	"github.com/ahmdrz/simpletoken"
)

var tokenGenerator *simpletoken.SimpleToken

func init() {
	simpleToken, err := simpletoken.New("sha1", []byte("long-long-secret-key"))
	if err != nil {
		panic(err)
	}
	tokenGenerator = simpleToken
}

func generateToken(input ExpirableToken) *simpletoken.Token {
	token, err := tokenGenerator.Generate(input)
	if err != nil {
		panic(err)
	}
	return token
}

func parseToken(input string, output *ExpirableToken) {
	err := tokenGenerator.ParseString(input, output)
	if err != nil {
		panic(err)
	}
}

type ExpirableToken struct {
	ExpiresAt time.Time
	Payload   map[string]interface{}
}

func (t ExpirableToken) IsExpired() bool {
	return t.ExpiresAt.After(time.Now())
}

func main() {
	token := ExpirableToken{
		ExpiresAt: time.Now().Add(1 * time.Second),
		Payload: map[string]interface{}{
			"username": "ahmdrz",
		},
	}

	tokenString := generateToken(token).String()
	log.Printf("token is %s", tokenString)

	parseToken(tokenString, &token)
	log.Printf("token state: %t", token.IsExpired())

	time.Sleep(2 * time.Second)

	parseToken(tokenString, &token)
	log.Printf("token state: %t", token.IsExpired())
}
