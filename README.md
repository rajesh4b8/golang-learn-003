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
- Slices are reference data types points to underlying array -> when the slice is updated the underlying array also updated.
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

## Day 8

### Structs

- What is Struct?
- Definition
- Declaration without initialization -> init with zero values
- Declaration with initialization
- Embedded structs -> one struct inside another struct

### Pass by Value

- When we pass any arg to a function or method the values will be copied in memory

### Pointers

- Use `&` to access the pointer (address)
- Use `*` to access the value from pointer
- `*` is used infront of a type to indicate it's a pointer type
- pointer indirection

### Assignment 4

- Update the deck of cards from `[]string` to `struct`

## Day 9

### Value Types / Reference Types

- When you pass the value type as param -> data is duplicated and any modifications will not update the original
- When you pass the reference type as param -> ref is duplicated but the underlying data is same, any modifications will result in updating the original

| Value Type       | Referece Type |
| ---------------- | ------------- |
| int              | slices        |
| float            | maps          |
| string           | channels      |
| bool             | pointers      |
| structs          | functions     |
| arrays           |      |

### Maps

- Map data structure -> a list of key / value pairs
- A key / value pair is called Entity
- All the keys are of same type
- All the valyes are of same type

| Map                 | Struct               |
|-------------------- |----------------------|
| All keys - same type| different types      |
| related properties  | represent something    |
| All values - same type|can be different |
| keys - can be added | Define them at compile time |
| keys are indexed    | keys not indexed |
| Reference Type      | Value Type |

### make vs new

- The new built-in function allocates memory, the value returned is a pointer to a newly allocated zero value of that type.
- The make built-in function allocates and initializes an object of type slice, map, or chan (only).

### Assignment 5

- Print numbers from 1 to 10 and print whether it's even or odd
- Sample output

``` text
1 - odd
2 - even
3 - odd
4 - even
...
```

## Day 10

### Interfaces

- An interface type in Go is kind of like a definition. It defines and describes the exact methods that some other type must have.
- To implement an interface
  - No need to explicitly mention it implements
  - No `implmenents` or `entends` keywords to say the struct implement the interface
  - Interfaces are implemented implicitly, A type implements an interface by implementing its methods.
- What is the interface in OOP?
An interface defines the behavior of an object. It only specifies what the object is supposed to do. It is actually a concept of abstraction and encapsulation

### Assignment 6

- New Income stream rental income inside the `./examples/interfaces`

## Day 11

- More on interfaces

### Unit Testing

- To test a peiece of code in a func or method
- using `testing` package

### Assignment 7

- Write more unit tests for deck of cards
