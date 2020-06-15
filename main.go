package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ImbaCow/lexer/app"
)

func main() {
	appArgs := os.Args[1:]
	if len(appArgs) != 1 {
		fmt.Printf("Invalid program arguments count: %v", len(appArgs))
		return
	}
	filePath := appArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Cannot open file: %s", filePath)
		return
	}

	writer := bufio.NewWriter(os.Stdout)

	lexer := app.NewLexer()
	lexer.TokenizeFile(file, writer)
	writer.Flush()
}
