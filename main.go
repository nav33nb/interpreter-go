package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome %v,\nThis is monke v1.0 REPL, write 'bye' to exit\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
