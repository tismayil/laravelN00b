package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
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
	err, _, _ = curl.Bytes(hostname)
	if err != nil {
		return ""
	}
	return str
}

func otherMethods(hostname string) string {

	cb := func(st curl.IoCopyStat) error {
		if st.Response != nil {
		}
		return nil
	}
	_, str, _ := curl.String(
		hostname, cb, "method=", "POST",
		"data=", strings.NewReader("{\"asd\": \"test\"}"),
		"disablecompression=", true,
		"header=", http.Header{"X-My-Header": {"laravelN00b"}},
	)

	return str
}

func envOrNot(response string) bool {
	envFound := regexp.MustCompile("APP([a-zA-Z_]+)=")
	if len(envFound.FindStringSubmatch(response)) > 0 {
		return true
	}
	return false
}

func debugOrNot(response string) bool {
	debugFound := regexp.MustCompile("APP([a-zA-Z_]+)</td>")
	if len(debugFound.FindStringSubmatch(response)) > 0 {
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
	Sponsored by : bekchy.com
	`)

	addresses, _ := net.LookupIP(*hostname)

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
	fmt.Printf("\033[1;33m%s\033[0m", "--------------------------------------------------------- \n")
	fmt.Printf("\033[1;33m%s\033[0m", "Checking .env Debug Mode \n")

	denugCheck := debugOrNot(otherMethods("http://" + *hostname))

	if denugCheck == false {

		denugCheck := debugOrNot(otherMethods("https://" + *hostname))
		if denugCheck == false {
			s.Restart()
			fmt.Printf("\033[1;31m%s\033[0m", "Debug mode not active \n")
		} else {
			s.Restart()
			fmt.Printf("\033[1;36m%s\033[0m", "* Laravel Debug Mode Active send POST or PUT and read .env Variables. \n\n")
			fmt.Printf("\033[1;36m%s\033[0m", "Example curl Command:\ncurl -X POST https://"+*hostname+" | grep APP_\n\n\n")
		}

	} else {
		s.Restart()
		fmt.Printf("\033[1;36m%s\033[0m", "* Laravel Debug Mode Active send POST or PUT and read .env Variables. \n\n")
		fmt.Printf("\033[1;36m%s\033[0m", "Example curl Command:\ncurl -X POST http://"+*hostname+" | grep APP_\n\n\n")
	}

	s.Restart()
	s.Stop()

}
