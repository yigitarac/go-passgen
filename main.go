package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
)

type sifreler struct {
	Uzunluk int    `json:"uzunluk"`
	Sifre   string `json:"sifre"`
	Isim    string `json:"isim"`
}

func main() {
	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Home dizini bulunamadı")
		return
	}
	path := filepath.Join(homePath, ".passDB.json")
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("passgen new <uzunluk> <isim>")
		fmt.Println("passgen list")
	} else if len(os.Args) == 4 {
		if os.Args[1] == "new" {
			boyut := os.Args[2]
			isim := os.Args[3]
			boyutTamSayi, err := strconv.Atoi(boyut)
			if err != nil {
				fmt.Println("Şifre tam sayıya dönüştürülemedi")
				return
			}
			sifreOlustur(boyutTamSayi, isim, path)
		}
	} else if len(os.Args) == 2 {
		if os.Args[1] == "list" {
			sifreleriListele(path)
		}
	}
}

func sifreOlustur(boyut int, isim string, yol string) {
	karakterHavuzu := "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuopasdfghjklizxcvbnm123456789!+-?&/"
	var sifre []byte
	var slice []sifreler
	for i := 0; i < boyut; i++ {
		index := rand.Intn(len(karakterHavuzu))
		sifre = append(sifre, karakterHavuzu[index])
	}
	yeniSifre := sifreler{
		Uzunluk: boyut,
		Isim:    isim,
		Sifre:   string(sifre),
	}
	okunanListe, err := os.ReadFile(yol)
	if err != nil {
		fmt.Println("Liste okunamadı.")
	}
	if err == nil {
		err = json.Unmarshal(okunanListe, &slice)
	}
	slice = append(slice, yeniSifre)
	listelenmisSifreOzellikleri, err := json.Marshal(slice)
	os.WriteFile(yol, listelenmisSifreOzellikleri, 0644)
}
func sifreleriListele(yol string) {
	var kayitliSifreler sifreler
	err := json.Unmarshal([]byte(yol), &kayitliSifreler)
	if err != nil {
		fmt.Println("Şifre dosyası okunamadı")
		return
	}
}
