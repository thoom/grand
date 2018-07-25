package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessStringAlphaNumMixed(t *testing.T) {
	assert := assert.New(t)

	str, err := processString("alphanum", 10, "mixed")
	assert.Empty(err)
	assert.Len(str, 10)
	assert.Regexp(regexp.MustCompile("^[a-zA-Z0-9]+$"), str)
}

func TestProcessStringAlphaMixed(t *testing.T) {
	assert := assert.New(t)

	str, err := processString("alpha", 10, "mixed")
	assert.Empty(err)
	assert.Len(str, 10)
	assert.Regexp(regexp.MustCompile("^[a-zA-Z]+$"), str)
}

func TestProcessStringHexMixed(t *testing.T) {
	assert := assert.New(t)

	str, err := processString("hex", 10, "mixed")
	assert.Empty(err)
	assert.Len(str, 10)
	assert.Regexp(regexp.MustCompile("^[a-fA-F0-9]+$"), str)
}

func TestProcessStringNumMixed(t *testing.T) {
	assert := assert.New(t)

	str, err := processString("num", 10, "mixed")
	assert.Empty(err)
	assert.Len(str, 10)
	assert.Regexp(regexp.MustCompile("^[0-9]+$"), str)
}

func TestProcessStringUUID(t *testing.T) {
	assert := assert.New(t)

	str, err := processString("uuid", 10, "mixed")
	assert.Empty(err)
	assert.Len(str, 36)
	assert.Regexp(regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[8|9|a|b][a-f0-9]{3}-[a-f0-9]{12}$"), str)
}

func TestProcessStringErr(t *testing.T) {
	assert := assert.New(t)
	str, err := processString("invalid", 10, "upper")
	assert.Empty(str)
	assert.Error(err)
}

func TestAlphaNumMixed(t *testing.T) {
	assert := assert.New(t)

	resp := alphanum(100, "mixed")
	assert.Len(resp, 100)
	assert.Regexp(regexp.MustCompile("^[a-zA-Z0-9]+$"), resp)
}

func TestAlphaNumUpper(t *testing.T) {
	assert := assert.New(t)

	resp := alphanum(100, "upper")
	assert.Len(resp, 100)
	assert.Regexp(regexp.MustCompile("^[A-Z0-9]+$"), resp)
}

func TestAlphaNumLower(t *testing.T) {
	assert := assert.New(t)

	resp := alphanum(100, "lower")
	assert.Len(resp, 100)
	assert.Regexp(regexp.MustCompile("^[a-z0-9]+$"), resp)
}

func TestAlphaMixed(t *testing.T) {
	assert := assert.New(t)

	resp := alpha(100, "mixed")
	assert.Len(resp, 100)
	assert.Regexp(regexp.MustCompile("^[a-zA-Z]+$"), resp)
}

func TestAlphaUpper(t *testing.T) {
	assert := assert.New(t)

	resp := alpha(100, "upper")
	assert.Len(resp, 100)
	assert.Regexp(regexp.MustCompile("^[A-Z]+$"), resp)
}

func TestAlphaLower(t *testing.T) {
	assert := assert.New(t)

	resp := alpha(100, "lower")
	assert.Len(resp, 100)
	assert.Regexp(regexp.MustCompile("^[a-z]+$"), resp)
}

func TestHexMixed(t *testing.T) {
	assert := assert.New(t)

	resp := hexEncode(100, "mixed")
	assert.Len(resp, 100)
	assert.Regexp(regexp.MustCompile("^[a-fA-F0-9]+$"), resp)
}

func TestHexUpper(t *testing.T) {
	assert := assert.New(t)

	resp := hexEncode(100, "upper")
	assert.Len(resp, 100)
	assert.Regexp(regexp.MustCompile("^[A-F0-9]+$"), resp)
}

func TestHexLower(t *testing.T) {
	assert := assert.New(t)

	resp := hexEncode(100, "lower")
	assert.Len(resp, 100)
	assert.Regexp(regexp.MustCompile("^[a-f0-9]+$"), resp)
}

func TestNum(t *testing.T) {
	assert := assert.New(t)

	resp := num(100)
	assert.Len(resp, 100)
	assert.Regexp(regexp.MustCompile("^[0-9]+$"), resp)
}

func TestUuid(t *testing.T) {
	assert := assert.New(t)
	assert.Regexp(regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[8|9|a|b][a-f0-9]{3}-[a-f0-9]{12}$"), uuidEncode("lower"))
}

func TestUuidUpper(t *testing.T) {
	assert := assert.New(t)
	assert.Regexp(regexp.MustCompile("^[A-F0-9]{8}-[A-F0-9]{4}-4[A-F0-9]{3}-[8|9|A|B][A-F0-9]{3}-[A-F0-9]{12}$"), uuidEncode("upper"))
}
