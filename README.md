[![GoDoc](https://godoc.org/github.com/lucasmenendez/gopaillier?status.svg)](https://godoc.org/github.com/lucasmenendez/gopaillier)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucasmenendez/gopaillier)](https://goreportcard.com/report/github.com/lucasmenendez/gopaillier)
[![test](https://github.com/lucasmenendez/gopaillier/workflows/test/badge.svg)](https://github.com/lucasmenendez/gopaillier/actions?query=workflow%3Atest)
[![license](https://img.shields.io/github/license/lucasmenendez/gopaillier)](LICENSE)

# GoPaillier
Extended version of a Paillier cryptosystem implementation in Go. 

## Features
- Extended Paillier cryptosystem implementation with negative number support (read more [here](./pkg/paillier/)).
- Uses Standard Form notation to encode numbers allowing to use Paillier encryption scheme over integer and floating points numbers (read more about [number package here](./pkg/number/number.go)).
- Allows four different operations:
  - Addition between encrypted and plain numbers: `A' + B`.
  - subtraction between encrypted and plain numbers: `A' + (-B)`.
  - Multiplication between encrypted and plain numbers: `A' * B`.
  - Division between encrypted and plain numbers: `A' * 1/B`.

### Installation
```sh
go get github.com/lucasmenendez/gopaillier@latest
```

### Examples
There are three basic examples ready to help starting with the library:
- Basic Paillier example: [Source code](./examples/basic/main.go).
- Median example: [Source code](./examples/median/main.go).
- SDK example: [Source code](./examples/sdk/main.go).