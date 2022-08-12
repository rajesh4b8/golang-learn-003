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

## Day 4

### Constants

- `const` is the keyword
- values that are constant and will not be changed

### Access Control

- naming starts with lower case -> package scope, only visisble inside the package
- naming starts with upper case -> exported from pacakge

### Functions

- first class citizens -> you pass the functions as an input
- follow the examples for different ways

### Use Case: playing cards

- Deck of cards functionality
  - New Deck
  - Print Deck
  - Shuffle
  - Deal
  - Save to File
  - Read from File

## Day 5

### Zero values

- Variables declared without an explicit initial value are given their zero value.
- `0` for numeric types,
- `false` for the boolean type
- `""` (the empty string) for strings

| Type          | Zero Value  |
| -----------   | ----------- |
| Integer       | 0           |
| Floating point| 0.0         |
| Boolean       | false       |
| String        | ""          |
| Pointer       | nil         |
| Interface     | nil         |
| Slice         | nil         |
| Map           | nil         |
| Channel       | nil         |
| Function      | nil         |

## Arrays / Slices

- An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
- Index starts with 0
- Slices are reference data types points to underlying array -> when the array is updated the slice also updated.
- A slice is formed by specifying two indices, a low and high bound, separated by a colon:
  `a[low : high]`
- Practically we use `make()` to create slices directly.

### Assignment 1

- Create and return the deck of cards with the defined suits and numbers

### GOPATH / GOROOT

- GOROOT is a variable that defines where your Go SDK is located. You do not need to change this variable, unless you plan to use different Go versions.

- GOPATH is a variable that defines the root of your workspace. By default, the workspace directory is a directory that is named go within your user home directory (~/go for Linux and MacOS, %USERPROFILE%/go for Windows). GOPATH stores your code base and all the files that are necessary for your development. You can use another directory as your workspace by configuring GOPATH for different scopes. GOPATH is the root of your workspace and contains the following folders:
  - src/: location of Go source code (for example, .go, .c, .g, .s).
  - pkg/: location of compiled package code (for example, .a).
  - bin/: location of compiled executable programs built by Go.

### Assignment 2

- Shuffle the cards after creating a deck

## Day 6

### Type conversions

- For converting compatible types you just use `new_type(old_value)`

### Assignment 3

- Deck of cards - Read from file
