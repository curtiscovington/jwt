# JWT - Simple JWT library for Go

This library lets you get up and running with JWTs in no time.
Quickly encode a payload into a JWT, verify a JWT signiature, and Decode a payload from a JWT.

```
package main

import "github.com/curtiscovington/jwt"

type Payload struct {
    UserId string `json:"userId"`
}

func main() {

    key := []byte("password") // Or use ioutil.ReadFile to read in a private key
    token, err := jwt.Encode(Payload{
        UserId: "1",
    }, key)

    if err != nil {
        // Handle error
    }

    // ****

    if !jwt.Verify(token, key) {
        // Handle invalid token
    }

    // ****

    var p Payload
    err = jwt.Decode(token, &p, key)
    if err != nil {
        // Handle error
    }

    // Use the decoded payload!
}
```
