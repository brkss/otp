package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	gFlag = flag.String("g", "", "Path to the input key file to encrypt and store")
	kFlag = flag.String("k", "", "Path to the encrypted key file to generate an OTP")
)

func main() {
	flag.Parse()

	if *gFlag != "" {
		fmt.Println("Encrypt and store key from ", *gFlag)
	}

	if *kFlag != "" {
		fmt.Println("Generate OTP from ", *kFlag)
	}

	if *kFlag == "" && *gFlag == "" {
		fmt.Println("No flags provided.\nUse -g to specify key file for encryption or -k to specify file to generate OTP.")
	}

	t := time.Now().Unix()
	fmt.Print("unix time : ", t)
}
