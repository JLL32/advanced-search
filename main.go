package main

import (
	"github.com/saferwall/advanced-search/repl"
	// "fmt"
	"os"
	// "os/user"
)

func main() {
	// user, err := user.Current()
	// if err != nil {
	// 	panic(err)
	// }
	repl.Start(os.Stdin, os.Stdout)
}
