package main

import (
	"fmt"
	"time"
)

func Unit() {
	fmt.Println("Unit OK")
}
func main() {
	fmt.Println(time.Now().Format(time.RFC3339))
}
