package main

import (
	"bufio"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			os.Stdout.WriteString(scanner.Text() + "\n")
		}
		os.Exit(0)
	}
	for _, filePath := range args[1:] {
		content, err := os.ReadFile(filePath)
		if err != nil {
			os.Stderr.WriteString("cat: " + filePath + ": No such file or directory\n")
			continue
		}
		os.Stdout.WriteString(string(content))
	}
}
