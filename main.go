package main

import (
	"fmt"
	"os"
)

type sifreler struct {
	Uzunluk int
	Sifre   string
	Isim    string
}

func main() {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		fmt.Println("passgen <uzunluk> <isim>")
	} else if len(os.Args) == 3 {
		boyut := os.Args[1]
		isim := os.Args[2]
	}
}
