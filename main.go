package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//kita buat 2 variable

// target ini untuk mengambil input data argumennya
var target string

// uras ini agar link yang kita kunjungi itu tidak duplikat
var urls = make(map[string]bool)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please put urls.   Example http://target.com")
		return
	}

	// untuk menganbil daa input argumen
	target = os.Args[1]
	runCrawler(target)
}

//method

func runCrawler(uri string) {
	//ga mau kalau urls nya u kosong
	//kalau kosong langsung berhenti functionnya
	if uri == "" {
		return
	}

	//jika nilai key nya tidak di temuan atau false
	//dan jika itu sudah da mka itu dplikat
	if !urls[uri] {
		urls[uri] = true
	} else {
		return
	}
	fmt.Println(uri)
	response, err := http.Get(uri)
	//kalau ada error kita return
	if err != nil {
		//kita tampilkan error nya apa
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	//kita ekstark ddocument nya
	doc, err := goquery.NewDocumentFromReader(response.Body)
	//di cek dulu ada error nya kita return
	if err != nil {
		//tampilkan errornya apa
		fmt.Println(err)
		return
	}

	//ambil tag a yang disini nanti bakal dapetin semua tag html yang ada a nya
	doc.Find("a").Each(func(i int, q *goquery.Selection) {
		//kita ambil attributnya yaitu href
		attr, exists := q.Attr("href")
		//jika ada atribut href
		if exists {

			nextLink := TrinUrl(attr)
			runCrawler(nextLink)
		}

	})
}

// function untuk dengerin kita mau linknya itu valid
func TrinUrl(uri string) string {
	//untuk menghapus kat terakhir dari urls,,misalnya
	//https://target.com/
	//https://target.com
	uri = strings.TrimSuffix(uri, "/")
	//kita mau pastikan link nya itu valid
	validUrl, err := url.Parse(uri)
	if err != nil {
		return ""
	}

	//kita mau membatasi hanya link yang ada di urls nya saja
	targetUrl, _ := url.Parse(target)
	if strings.Contains(validUrl.String(), targetUrl.Host) {
		return uri
	}
	return ""
}
