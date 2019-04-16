## Simple Token !
> Generating HMAC based token is so easy with `simpletoken`..

#### Installation

```bash
$ go get -u github.com/ahmdrz/simpletoken
```

#### Examples

```go
package main

import (
	"fmt"
	"time"

	"github.com/ahmdrz/simpletoken"
)

type Payload struct {
	Username  string
	CreatedAt time.Time
}

func main() {
	simpleToken, err := simpletoken.New("md5", []byte("long-long-secret-key"))
	if err != nil {
		panic(err)
	}

	myPayload := Payload{
		Username:  "ahmdrz",
		CreatedAt: time.Now(),
	}
	token, err := simpleToken.Generate(myPayload)
	if err != nil {
		panic(err)
	}

	fmt.Printf("token is %s\n", token)

	output := Payload{}
	err = simpleToken.ParseString(token.String(), &output)
	if err != nil {
		panic(err)
	}

	fmt.Printf("payload is %v\n", output)
}
```

Find `expirableToken` example in `example` directory.