# golang-learn-003

GoLang training material and examples

## Day 1

### Agenda

- Intro
- Topics to be covered
- Why Go?
- History
- Go playground (https://go.dev/play/)
- Hello world example
- Hello web example

## Day 2

### Git basics

- git commit is for your local repository
- git push is to origin (remote repository)
- git clone is to clone the repo to local -- use ssh url to clone

### To do

- Create a github account
- install git in your system
- Create SSH pub / private keys

```ssh-keygen -t ed25519 -C "your_email@example.com"```

- add the pub key to github.com at settings
- cheat sheets for reference

https://training.github.com/downloads/github-git-cheat-sheet
https://about.gitlab.com/images/press/git-cheat-sheet.pdf

### Understading the basic program

- The three main areas to fous
  - package main-> every file needs to define a package, package main is mandatory for any program to run
  - import -> to get and use other packages, only needed if you are using other packages
  - function main()-> you need a main() function to run the program -> starting point for the program

## Day 3

### Understanding the packages

- main package -> Executable package, every program needs a main package if we have to run the program.
- resusable package -> these are the packages to share / reuse the code between modules
- typically the folder name will be the package name, unless it's main

### Go CLI

- Command line interface, you get this cli when go is isstalled
- Use the CLI to work with go
  - `go build` -> compile packages and dependencies and generate a binary OS's native (.exe for windows)
  - `go run` -> compile and run Go program
  - `go fmt` -> gofmt (reformat) package sources
  - `go get` -> download and install packages and dependencies
  - `go test` -> for testing the package -> running the unit tests

### Buit in types

- bool, string
- int int8 int16 int32 int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // alias for uint8
- rune // alias for int32
- float32 float64
- complex64 complex128

### Variables

- `var` is the keyword
- Variables can only be declared once in the scope
- `:=` is used to declare and initialize
- Go is strcitly type language
