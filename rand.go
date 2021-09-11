package main

import (
    "crypto/rand"
    "fmt"
)

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func shortID(length int) string {
    ll := len(chars)
    b := make([]byte, length)
    rand.Read(b) // generates len(b) random bytes
    for i := 0; i < length; i++ {
        b[i] = chars[int(b[i])%ll]
    }
    return string(b)
}

func main() {
    fmt.Println(shortID(18))
    fmt.Println(shortID(18))
    fmt.Println(shortID(18))
}
