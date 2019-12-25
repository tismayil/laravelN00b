#!/bin/bash
echo "Installing laravelNoob"
go get github.com/briandowns/spinner
echo "GoLang spinner Pack Installed"
go get github.com/christophwitzko/go-curl
echo "GoLang go-curl Pack Installed"
go build -o laravelNoob main.go
echo "laravelN00b Builded."
echo -e "Usage: ./laravelN00b --hostname victim.host"