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
  - function main()-> you need a main() function to run the program
