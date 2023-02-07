# Artithmatic CLI

This CLI tool allows you to perform basic mathematics operation like Addition, Subtraction, Multiplication and Divison. This Repo is beneficial for those who want to look at how basic cli can be built using Golang and Cobra.

# Content

* Installation
* Usage
* Commands
  * add
  * divide
  * multiply
  * subtract
* Flags
* Test

# Prerequisites
Go installed in the local computer with version >= 1.18.

In case for upgradation refer this Link [Golang](https://www.golinuxcloud.com/upgrade-go-version/)

# Installation
First clone the repo in your $Path. A common place would be within your $GOPATH

Build and copy ```arithmatic_cli ``` to your $GOPATH/bin:

```
$ go build .
```

# Running the Cli

To list all the commands and flags

```
./arithmatic_cli help
```

To run the whole command 
```
./arithmatic_cli [command] [flag] [args]
```

# Commands

## add
add command is used to add numbers. Numbers can be float typeor integer type

To perform float type addition or to return the answer in float type use ```-f``` flag.

Addition for integer numbers

```
./arithmatic_cli [add] [args]

Example: ./arithmatic_cli add 1 2

The above example will perform addition of 1 and 2 and return the ans 3

```

Addition for Float numbers
```
./arithmatic_cli [add] [-f] 1 2
 
 Example: ./arithmatic_cli add -f 1 2.1
 
 The above example will perform addition of 1 and 2 and return the ans in float number i.e 3.1
 
```




