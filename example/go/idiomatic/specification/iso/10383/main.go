package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	url := `https://www.iso20022.org/sites/default/files/ISO10383_MIC/ISO10383_MIC.xml`
	path := `iso-10383.xml`
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	file, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		panic(err)
	}
}
