package simpletoken

import "testing"

const (
	tokenString = "81a8757365726e616d65a661686d64727ae5da89071bea20ecf97e961fa665aae395761451"
)

func TestGenerate(t *testing.T) {
	simpleToken, _ := New("sha1", []byte("hello-world"))
	token, err := simpleToken.Generate(map[string]interface{}{
		"username": "ahmdrz",
	})
	if err != nil {
		t.Fatal(err)
		return
	}
	if tokenString != token.String() {
		t.Fatalf("invalid token generated. got %s", token.String())
	}
	t.Log(token)
}

func TestParseString(t *testing.T) {
	simpleToken, _ := New("sha1", []byte("hello-world"))

	output := map[string]interface{}{}
	err := simpleToken.ParseString(tokenString, &output)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(output)
}

func TestInvalidParseString(t *testing.T) {
	simpleToken, _ := New("sha1", []byte("hello-world"))
	invalidTokenString := "81a8757365726e616d65a661686d64727a25f7c1173c6c5dabbcca6b094775eac2c681c70a"

	output := map[string]interface{}{}
	err := simpleToken.ParseString(invalidTokenString, &output)
	if err == nil {
		t.Fatal("invalid token is validated by hmac.")
		return
	}
}
