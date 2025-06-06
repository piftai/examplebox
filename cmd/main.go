package main

import (
	"fmt"
	"github.com/piftai/examplebox/internal/auth"
	"time"
)

func main() {
	manager := auth.NewJWTManager("test", time.Duration(1*time.Minute))
	fmt.Println(manager.GenerateToken("test"))
}
