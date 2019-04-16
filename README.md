## Simple Token !
> Generate HMAC based token so easy with simpletoken.

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

func main() {
	simpleToken, err := simpletoken.New("md5", []byte("long-long-secret-key"))
	if err != nil {
		panic(err)
	}

	token, err := simpleToken.Generate(map[string]interface{}{
		"username":   "ahmdrz",
		"created_at": time.Now(),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("token is %s\n", token)

	output := map[string]interface{}{}
	err = simpleToken.ParseString(token.String(), &output)
	if err != nil {
		panic(err)
	}

	fmt.Printf("payload is %v\n", output)
}
```

Find `expirableToken` example in `example` directory.