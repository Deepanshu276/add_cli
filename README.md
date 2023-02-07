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
## subtract
subtract command is used to subtract numbers. Numbers can be float typeor integer type

To perform float type subtraction or to return the answer in float type use ```-s``` flag.

Subtraction for integer numbers

```
./arithmatic_cli [subtract] [args]

Example: ./arithmatic_cli subtract 2 1

The above example will perform subtraction of 2 and 1 and return the ans 1

```

Subraction for Float numbers
```
./arithmatic_cli [subtract] [-s] 2.1 1
 
 Example: ./arithmatic_cli subtract -s 2.1 1
 
 The above example will perform subtraction of 2.1 and 1 and return the ans in float number i.e 1.2
 
```

## multiply
multiply command is used to multiply numbers. Numbers can be float typeor integer type

To perform float type multiplication or to return the answer in float type use ```-m``` flag.

Multiplication for integer numbers

```
./arithmatic_cli [multiply] [args]

Example: ./arithmatic_cli multiply 1 2

The above example will perform multiplication of 1 and 2 and return the ans 2

```

Multiplication for Float numbers
```
./arithmatic_cli [multiply] [-m] 1 2.1
 
 Example: ./arithmatic_cli multiply -m 1 2.1
 
 The above example will perform multiplication of 1 and 2.1 and return the ans in float number i.e 2.1
 
```

## divide
divide command is used to divide numbers. Numbers can be float typeor integer type

To perform float type division or to return the answer in float type use ```-d``` flag.

Division for integer numbers

```
./arithmatic_cli [divide] [args]

Example: ./arithmatic_cli divide 2 1

The above example will perform division of 2 and 1 and return the ans 2

```

Divison for Float numbers
```
./arithmatic_cli [divide] [-d] 2 1
 
 Example: ./arithmatic_cli divide -d 2 1
 
 The above example will perform division of 2 and 1 and return the ans in float number i.e 2.000
 
```

# Test
To test the cli using different test cases 
1. cd into test folder using ```cd test```.
2. To run all test file all at once use 
```
go test
```
3. To run a particular test file use :
```
go test -run=TestName path/to/testfile.go

Example : To run test for addition

go test -run= TestAdder path/to/testfile.go
```






