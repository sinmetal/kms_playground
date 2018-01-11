package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	project := flag.String("project", "", "The Google Cloud Platform project ID. Required.")
	flag.Parse()
	for _, f := range []string{"project"} {
		if flag.Lookup(f).Value.String() == "" {
			log.Fatalf("The %s flag is required.", f)
		}
	}

	ctx := context.Background()
	s, err := NewKMSService(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	k := CryptKey{
		ProjectID:  *project,
		LocationID: "global",
		KeyRingID:  "sample-key-ring",
		KeyName:    "sample-key",
	}

	rct, rk, err := s.Encrypt(k, "Hello World !")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("CipherText=%s, CryptoKey=%s\n", rct, rk)

	pt, err := s.Decrypt(k, rct)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("PlainText=%s\n", pt)
}
