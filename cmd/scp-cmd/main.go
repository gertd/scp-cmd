package main

import (
	"fmt"
	"os"

	"github.com/gertd/scp-cmd/cmd/scp-cmd/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
