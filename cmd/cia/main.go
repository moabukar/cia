package main

import (
	"os"

	"github.com/moabukar/cia/internal/cli"
)

var version = ""

func main() {
	// if err := cli.Execute(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	cli.Execute(version, os.Args)
}
