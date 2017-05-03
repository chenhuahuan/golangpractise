package main

import (
    	"fmt"
	"golangpractise/myemail/libofm"
)

func main() {
    mycontent := " my dear令"

    email := libofm.NewEmail("hunterchen@.com",
        "test golang email", mycontent)

    err := libofm.SendEmail(email)

    fmt.Println(err)

}