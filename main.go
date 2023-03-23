package main

import (
	"fmt"
	"userauth/initializers"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	fmt.Println("kkk")
}
