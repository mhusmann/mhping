// mhping my last attempt to check which hosts are alive in my network
// Solution now:
// create an array of channels
// fill the channels by a go routine
// then i check whether a channel is != ""
// mhusmann 2016-12-05
package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const lim = 100 // number of hosts to check
const ipStart = "192.168.178."
const extern = "www.google.com"

// test whether the internet works
func checkExtern(adr string) {
	fmt.Printf("First pinging %s\n", adr)
	out0, err := exec.Command("ping", "-c1", adr).Output()
	if err != nil {
		log.Fatal("extern is not accessible")
	}
	s := strings.SplitAfter(string(out0), "--- ")
	fmt.Printf("%s\n", s[0])
}

// fillChannel put fpings result in each channel
// if the host is not alive, fill that channel with an empty string
func fillChannel(c chan string, ip string, retries *string) {
	out, _ := exec.Command("fping", "-r"+*retries,
		"-de", ip).CombinedOutput()
	res := string(out)
	if strings.Contains(res, "alive") {
		split := strings.Split(res, " ")
		c <- fmt.Sprintf("%17s  %30s % 5s ms)\n", ip, split[0],
			split[3])
	}
	c <- ""
}

func main() {
	retries := flag.String("retry", "3", "Number of retries")
	flag.Parse()
	var channels [lim]chan string
	var slice = &channels
	var ip string
	count := 0

	checkExtern(extern)
	fmt.Println("Now running fping as GO-Routines")
	for i := range slice {
		slice[i] = make(chan string)
	}
	for i := 0; i < lim; i++ {
		ip = ipStart + strconv.Itoa(i+1)
		go fillChannel(slice[i], ip, retries)
	}
	for _, c := range slice {
		if res := <-c; len(res) > 0 {
			count++
			fmt.Printf("%3d  %s", count, res)
		}
	}
	fmt.Printf("Found %d hosts\a\n", count)
}
