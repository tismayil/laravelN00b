package main

import (
	"flag"
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/christophwitzko/go-curl"
)

func sendReq(hostname string) string {
	err, str, _ := curl.String(hostname)
	if err != nil {
		return ""
	}
	//fmt.Println(str)
	//fmt.Println(resp.Header["Server"]) // access response header
	err, _, _ = curl.Bytes(hostname)
	if err != nil {
		return ""
	}
	return str
}

func envOrNot(response string) bool {

	envFound := regexp.MustCompile("APP([a-zA-Z_]+)=")

	if len(envFound.FindStringSubmatch(response)) > 0 {

		return true

	}

	return false
}

func main() {

	hostname := flag.String("hostname", "", "Please input hostname")
	flag.Parse()

	fmt.Println(`
	_                              _ _   _  ___   ___  _     
	| |    __ _ _ __ __ ___   _____| | \ | |/ _ \ / _ \| |__  
	| |   / _  | '__/ _  \ \ / / _ \ |  \| | | | | | | | '_ \ 
	| |__| (_| | | | (_| |\ V /  __/ | |\  | |_| | |_| | |_) |
	|_____\__,_|_|  \__,_| \_/ \___|_|_| \_|\___/ \___/|_.__/ 
    	Laravel .env file checker. ( Scan domain and IP )                                     
	Host : ` + *hostname + `
	Sponsored : www.bekchy.com
	`)

	addresses, _ := net.LookupIP("www.bekchy.com")

	ip := ""
	for i := 0; i < len(addresses); i++ {
		segments := strings.SplitAfter(addresses[i].String(), " ")
		ip = segments[0]
	}

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	s.Restart()
	fmt.Printf("\033[1;34m%s\033[0m", "Checking .env in : "+"http://"+*hostname+"/ \n")

	firstReq := envOrNot(sendReq("http://" + *hostname + "/.env"))

	if firstReq == false {
		s.Restart()
		fmt.Printf("\033[1;34m%s\033[0m", "Checking .env in : "+"https://"+*hostname+"/ \n")

		firstReq := envOrNot(sendReq("https://" + *hostname + "/.env"))
		if firstReq == false {
			s.Restart()
			fmt.Printf("\033[1;34m%s\033[0m", "Checking .env in : "+"http://"+ip+"/ \n")

			firstReq := envOrNot(sendReq("http://" + ip + "/.env"))
			if firstReq == false {

				s.Restart()
				fmt.Printf("\033[1;34m%s\033[0m", "Checking .env in : "+"http://www."+*hostname+"/ \n")
				firstReq := envOrNot(sendReq("http://www." + ip + "/.env"))
				if firstReq == false {

					s.Restart()
					fmt.Printf("\033[1;34m%s\033[0m", "Checking .env in : "+"https://www."+*hostname+"/ \n")
					firstReq := envOrNot(sendReq("https://www." + ip + "/.env"))
					if firstReq == false {
						fmt.Printf("\033[1;31m%s\033[0m", ".env Not Found \n")
					} else {
						s.Restart()
						fmt.Printf("\033[1;36m%s\033[0m", ".env Found : "+"https://www."+ip+"/.env \n")
					}

				} else {
					s.Restart()
					fmt.Printf("\033[1;36m%s\033[0m", ".env Found : "+"http://www."+ip+"/.env \n")
				}

			} else {
				s.Restart()
				fmt.Printf("\033[1;36m%s\033[0m", ".env Found : "+"http://"+ip+"/.env \n")
			}
		} else {
			s.Restart()
			fmt.Printf("\033[1;36m%s\033[0m", ".env Found : "+"https://"+*hostname+"/.env \n")
		}
	} else {
		s.Restart()
		fmt.Printf("\033[1;36m%s\033[0m", ".env Found : "+"http://"+*hostname+"/.env \n")
	}

	s.Restart()
	s.Stop()

}
