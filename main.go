package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	req, err := createRequest()
	checkErr(err)

	res, err := http.DefaultClient.Do(req)
	checkErr(err)

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	checkErr(err)

	if res.StatusCode != http.StatusOK {
		checkErr(fmt.Errorf("%s %s", res.Status, b))
	}

	log.Printf("%s %s", res.Status, b)}
