package main

import (
	"strconv"
	"fmt"
)

func main(){
	s := strconv.FormatFloat(3.1215, 'f', 1, 64)

	fmt.Println(s)
}


