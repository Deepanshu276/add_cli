#This Makefile assumes that your CLI entry point is in a file named main.go.

#The BINARY variable defines the name of the binary file that will be generated when you run make build.
BINARY=add_cli

#The default target that builds the binary.
all: build

#Builds the binary by running the go build command.
build:
	go build -o main.go

#nstalls the binary by running the go install command.
install:
	go install

#Cleans up any generated files by running the go clean command and removing the binary file.
clean:
	go clean
	rm -f $(BINARY)
#Runs the binary.
run:
	./$(BINARY)

# .PHONY target is used to specify a list of targets that should not be considered files.
.PHONY: all build install clean run
