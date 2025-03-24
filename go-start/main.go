package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(
		ABC(1, 2),
	)

	_ = gin.Default()
}

func ABC(a, b int) int {
	return a + b
}
