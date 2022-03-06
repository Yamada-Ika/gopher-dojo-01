package imgconv

import (
	"errors"
	"flag"
	"os"
)

func validateFlag() (from, to string, err error) {
	switch inputExtFlag {
	case "jpg", "jpeg", "png", "gif":
		break
	default:
		return "", "", errors.New("error: invalid extension")
	}
	switch outputExtFlag {
	case "jpg", "jpeg", "png", "gif":
		break
	default:
		return "", "", errors.New("error: invalid extension")
	}
	return inputExtFlag, outputExtFlag, nil
}

var inputExtFlag string
var outputExtFlag string

func init() {
	flag.StringVar(&inputExtFlag, "i", "jpg", "input file extension")
	flag.StringVar(&outputExtFlag, "o", "png", "input file extension")
}

func ValidateArgs() (dirs []string, from, to string, err error) {
	if len(os.Args) == 1 {
		return nil, "", "", errors.New("error: invalid argument\n")
	}
	flag.Parse()
	from, to, err = validateFlag()
	if err != nil {
		return nil, "", "", err
	}
	return flag.Args(), from, to, nil
}
