package main

import (
	"fmt"

	"github.com/user_authentication_golang/initializers"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	fmt.Println("yaao")
}
