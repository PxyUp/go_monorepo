package main

import (
	"go_monogrepo/packages/second/router"
	"fmt"
)

func main() {
	fmt.Println("Run main second router")
	_ = router.GetEngine().Run(":8081")
}