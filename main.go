package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	gFlag = flag.String("g", "", "Path to the input key file to encrypt and store")
	kFlag = flag.String("k", "", "Path to the encrypted key file to generate an OTP")
)

func ValidateKey(filepath string) error {
	path := strings.TrimSpace(filepath)

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("ft_otp: %v does not exist", path)
	}
	len := len(data)

	if len != 64 {
		return fmt.Errorf("ft_otp: key must be 64 hexadecimal characters")
	}

	return nil
}

func main() {
	flag.Parse()

	if *gFlag != "" {
		// validate key
		err := ValidateKey(*gFlag)
		fmt.Println("error validating : ", err)
		//fmt.Println("Encrypt and store key from ", *gFlag)
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
