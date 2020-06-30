package main

import (
	"fmt"

	"example.org/luksam/kiwi-server/cmd"
)

func main() {
	fmt.Println("--- application started ---")

	cmd.Execute()

}
