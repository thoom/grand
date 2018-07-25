package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/pborman/uuid"
)

var (
	randType = flag.String("type", "alphanum", "alphanum, alpha, hex, num, uuid")
	randLen  = flag.Int("length", 32, "Length of the string")
	randCase = flag.String("case", "mixed", "upper, lower, mixed")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()

	output, err := processString(*randType, *randLen, *randCase)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(output)
}

func processString(randType string, randLen int, randCase string) (string, error) {
	switch randType {
	case "alphanum":
		return alphanum(randLen, randCase), nil
	case "alpha":
		return alpha(randLen, randCase), nil
	case "hex":
		return hexEncode(randLen, randCase), nil
	case "num":
		return num(randLen), nil
	case "uuid":
		return uuidEncode(randCase), nil
	}

	return "", fmt.Errorf("Invalid -type")
}

func alphanum(length int, strCase string) string {
	str := randomString("abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ0123456789", length)
	if strCase == "upper" {
		str = strings.ToUpper(str)
	} else if strCase == "lower" {
		str = strings.ToLower(str)
	}

	return str
}

func alpha(length int, strCase string) string {
	str := randomString("abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ", length)
	if strCase == "upper" {
		str = strings.ToUpper(str)
	} else if strCase == "lower" {
		str = strings.ToLower(str)
	}

	return str
}

func num(length int) string {
	return randomString("0123456789", length)
}

func hexEncode(length int, strCase string) string {
	str := randomString("0123456789abcdefABCDEF", length)
	if strCase == "upper" {
		str = strings.ToUpper(str)
	} else if strCase == "lower" {
		str = strings.ToLower(str)
	}

	return str
}

func uuidEncode(strCase string) string {
	str := uuid.New()
	if strCase == "upper" {
		str = strings.ToUpper(str)
	}

	return str
}

func randomString(seed string, l int) string {
	buf := make([]byte, l)
	for i := 0; i < l; i++ {
		buf[i] = seed[rand.Intn(len(seed))]
	}
	return string(buf)
}
