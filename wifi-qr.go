package main

import (
	"github.com/mdp/qrterminal"
	"os"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()
	ssid := args[0]
	pass := args[1]

	str := fmt.Sprintf("WIFI:S:%s;T:WPA;P:%s;;", ssid, pass)

	// WIFI:S:MyWi-Fi;T:WPA;P:PassPhrase;;
// Generate a 'dense' qrcode with the 'Low' level error correction and write it to Stdout
	qrterminal.Generate(str, qrterminal.L, os.Stdout)

}