package ios

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Simple simple usage Write and Read
func Simple() {
	var b bytes.Buffer
	b.Write([]byte("Eko Kurniawan "))
	fmt.Fprintf(&b, "%s", "Khannedy")

	fmt.Println(string(b.Bytes()))

	b.WriteTo(os.Stdout)
}

// Curl simple curl
func Curl() {
	response, err := http.Get("https://www.blibli.com/")
	if err != nil {
		fmt.Println(err)
		return
	}

	dest := io.MultiWriter(os.Stdout)

	io.Copy(dest, response.Body)
	if err := response.Body.Close(); err != nil {
		log.Println(err)
	}
}
