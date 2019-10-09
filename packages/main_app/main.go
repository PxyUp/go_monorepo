package main

import (
	"go_monogrepo/packages/main_app/router"
	"fmt"
)

func main() {
	fmt.Println("Run main application router")
	_ = router.GetEngine().Run(":8080")
}

func Check() {
	fmt.Println("Check")
}