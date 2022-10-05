package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	myargs := os.Args[1:]
	fmt.Println(myargs)
	if len(myargs) == 0 {
		fmt.Println(AnalyseCfg(os.Stdin))
	} else {
		myfile, err := os.Open(myargs[0]) // first argument is filename to open, no other options yet

		if err != nil {
			fmt.Println(err)
			os.Exit(1) // returning something not zero means program failed
		}
		fmt.Println("Config-Analysis V1 of file ", myargs[0], " = ", AnalyseCfg(myfile))

	}

	os.Exit(0)
}

func AnalyseCfg(r io.Reader) int {
	myscan := bufio.NewScanner(r)
	myscan.Split(bufio.ScanLines) // (bufio.ScanWords) ScanLines is default ... so redundant but good to learn ;-)
	var lc int                    // count lines treated
	var cfgLine string
	var cfgTokens []string
	var ontCount int     // count ONTs that are added in the config
	var ontWLIDCount int // count ONTs that have a DT-LineID as a description
	var srvpCount int    // count service-ports - second token in line must be an integer number (distinguish from "service-port desc ...")
	for myscan.Scan() {
		lc++
		cfgLine = myscan.Text() // yields the line just read as a string
		cfgTokens = strings.Split(strings.TrimSpace(cfgLine), " ")
		//	fmt.Println("cfgToken[0]=", cfgTokens[0])

		switch cfgTokens[0] {
		case "ont":
			if cfgTokens[1] == "add" {
				ontCount++
				//	fmt.Println(cfgTokens[0], cfgTokens[1])
				if strings.Contains(cfgLine, "desc \"DEU.DTAG.") {
					ontWLIDCount++
				}

			}

		case "service-port":
			fmt.Println(cfgTokens[0], cfgTokens[1])

			if _, err := strconv.Atoi(cfgTokens[1]); err == nil { // check if the second token is a number
				srvpCount++
			}

		}
	}

	fmt.Println("found ", ontCount, " ONTs in the config.")
	fmt.Println("found ", ontWLIDCount, " ONTs have a DT Line-ID in their desc.")
	fmt.Println("found ", srvpCount, " Service-ports in the config.")
	return lc
}
