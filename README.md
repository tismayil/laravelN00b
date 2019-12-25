# laravelN00b
Incorrect configuration allows you to access .env files or reading env variables. LaravelN00b automated scan .env files and checking debug mode in victim host.

[![asciicast](https://asciinema.org/a/ZHvUT1CUh8qekiUgcxgP0OlVB.svg)](https://asciinema.org/a/ZHvUT1CUh8qekiUgcxgP0OlVB)


## Scan rationale
- Scan host.
- Resolve IP adress and check .env file in IP Adress
- Checking debug mode Laravel ( Read .env variables )

## Installation

1 - Install with installer.sh

`chmod +x installer.sh`

`./installer.sh`

2 - Install manual

`go get github.com/briandowns/spinner`

`github.com/christophwitzko/go-curl`

`go run main.go --hostname victim.host`

or 

`go build laravelN00b main.go`

## Run

`./laravelN00b --hostname victim.host `
