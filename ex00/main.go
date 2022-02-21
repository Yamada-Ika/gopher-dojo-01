package main

import (
	"bufio"
	"io"
	"os"
)

func trimError(err error) string {
	s := err.Error()
	for i, c := range s {
		if c == ' ' {
			return s[i+1:]
		}
	}
	return s
}

func readWrite(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		_, err := io.WriteString(w, scanner.Text()+"\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		if err := readWrite(os.Stdin, os.Stdout); err != nil {
			os.Stderr.WriteString(err.Error())
		}
	} else {
		for _, filePath := range args[1:] {
			finfo, err := os.Stat(filePath)
			if err != nil {
				os.Stderr.WriteString(err.Error())
			}
			if finfo.IsDir() {
				os.Stderr.WriteString("ft_cat: " + filePath + ": Is a directory\n")
				continue
			}
			f, err := os.Open(filePath)
			if err != nil {
				os.Stderr.WriteString("ft_cat: " + trimError(err) + "\n")
				continue
			}
			defer f.Close()
			if err := readWrite(f, os.Stdout); err != nil {
				os.Stderr.WriteString("ft_cat: " + trimError(err) + "\n")
			}
		}
	}
}
