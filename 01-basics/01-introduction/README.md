# Chapter 1 – Introduction to Go

## Why Go?

- Designed at Google for large scale systems
- Fast compilation
- Built-in concurrency
- Simple syntax
- Produces single static binary

## Go Philosophy

- Simplicity > cleverness
- Composition over inheritance
- Explicit error handling
- Built for concurrency
- Developer productivity first

## Go Toolchain

- go run : Compile + run
- go build : compile
- go install : Compile + Place binary
- go test : Run tests
- go mod : Dependency management
- go fmt : Formatting
- go vet : Static Analysis

## Hello World Anatomy

package main
import "fmt"
func main()

## Commands

### go run

Compiles and runs the program temporarily.

### go build

Builds an executable in the current directory.

### go install

Builds and installs binary to GOPATH/bin.
