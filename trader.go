package main

import (
	"fmt"
	"hexnet/trader/core"
)

func main() {
	config := core.LoadConfig("")
	fmt.Printf("trader : %v\n", config.Env.StrategyName)
}
