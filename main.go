package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	project := flag.String("project", "", "The Google Cloud Platform project ID. required.")
	cmd := flag.String("cmd", "", "cmd is required. cmd is encrypt or decrypt")
	text := flag.String("text", "", "text is required.")
	flag.Parse()
	for _, f := range []string{"project", "cmd", "text"} {
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

	switch *cmd {
	case "encrypt":
		{
			rct, rk, err := s.Encrypt(k, *text)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("CipherText=%s, CryptoKey=%s\n", rct, rk)
		}
	case "decrypt":
		{
			pt, err := s.Decrypt(k, *text)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("PlainText=%s\n", pt)
		}
	default:
		fmt.Println("The cmd flog is encrypt or decrypt")
	}
}
