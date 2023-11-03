# gRand ![Builds](https://github.com/thoom/grand/actions/workflows/main.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/thoom/grand)](https://goreportcard.com/report/github.com/thoom/grand) [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=grand&metric=coverage)](https://sonarcloud.io/summary/overall?id=grand) [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=grand&metric=security_rating)](https://sonarcloud.io/summary/overall?id=grand) [![GoDoc](https://godoc.org/github.com/thoom/grand?status.svg)](https://godoc.org/github.com/thoom/grand)
 

gRand is a random string generator.

It supports:

* alphanumeric (uppercase, lowercase, mixed)
* alphabetic (uppercase, lowercase, mixed)
* hex (uppercase, lowercase, mixed)
* numeric
* uuid (uppercase, lowercase)

## Installation

There are several ways to download and install the `grand` client.

### Binary Releases

The preferred method is to download the latest binary release for your platform from the [Github Releases](https://github.com/thoom/grand/releases) section.

### Using Docker

If you already use Docker, gRand is packaged into a very small, OS-less image (~4 Mb compressed, 7.5 Mb uncompressed). To learn more about the Docker image, see the [Github Packages](https://github.com/users/thoom/packages/container/package/grand) section.

**Basic usage**

```
docker run --rm -it ghcr.io/thoom/grand
```

### Using Go

Finally, you can also just build it directly on your machine if you already have Go installed:

```
go get github.com/thoom/grand
```

## Usage

```
-case string
    upper, lower, mixed (default "mixed")
-length int
    Length of the string (default 32)
-repeat int
    number of times to run (default 1)
-type string
    alphanum, alpha, hex, num, uuid (default "alphanum")
-version
    current version
```