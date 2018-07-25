# goron [![Build Status](https://travis-ci.org/thoom/goron.svg?branch=master)](https://travis-ci.org/thoom/goron) [![Go Report Card](https://goreportcard.com/badge/github.com/thoom/goron)](https://goreportcard.com/report/github.com/thoom/goron) [![codecov](https://codecov.io/gh/thoom/goron/branch/master/graph/badge.svg)](https://codecov.io/gh/thoom/goron) [![GoDoc](https://godoc.org/github.com/thoom/goron?status.svg)](https://godoc.org/github.com/thoom/goron)

goron is a random string generator.

It supports:

* alphanumeric (uppercase, lowercase, mixed)
* alphabetic (uppercase, lowercase, mixed)
* hex (uppercase, lowercase, mixed)
* numeric
* uuid (uppercase, lowercase)

Usage:

`goron -type [alphanum, alpha, hex, num, uuid] -length [int] -case [upper, lower, mixed] -times [int]`