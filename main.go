package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	mr "math/rand"
	"os"
	"strings"
	"time"
)

var (
	buildVersion string

	randType   = flag.String("type", "alphanum", "alphanum, alpha, hex, num, uuid")
	randLen    = flag.Int("length", 32, "Length of the string")
	randCase   = flag.String("case", "mixed", "upper, lower, mixed")
	iterations = flag.Int("repeat", 1, "number of times to run")
	version    = flag.Bool("version", false, "current version")
)

func init() {
	if buildVersion == "" {
		buildVersion = time.Now().Format("20060102.0304PM.MST") + "-SNAPSHOT"
	}

	mr.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()

	if *version {
		fmt.Println(printVersion(buildVersion))
		os.Exit(0)
	}

	for i := 0; i < *iterations; i++ {
		output, err := processString(*randType, *randLen, *randCase)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(output)
	}
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
	b := make([]byte, 16)
	rand.Read(b)

	b[8] = (b[8] | 0x80) & 0xBF
	b[6] = (b[6] | 0x40) & 0x4F

	str := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	if strCase == "upper" {
		str = strings.ToUpper(str)
	}

	return str
}

func randomString(seed string, l int) string {
	buf := make([]byte, l)
	for i := 0; i < l; i++ {
		buf[i] = seed[mr.Intn(len(seed))]
	}
	return string(buf)
}

func printVersion(version string) string {
	return fmt.Sprintf(`
thoom.Goron - random string generator

version: %s
author: Z.d.Peacock <zdp@thoomtech.com>
link: https://github.com/thoom/goron
`, version)
}
