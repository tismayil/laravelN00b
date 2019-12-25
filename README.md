# laravelN00b
Incorrect configuration allows you to access .env files. LaravelN00b automated scan .env files in victim host.

## Scan rationale
- Scan host.
- Resolve IP adress and check .env file IP Adress

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
