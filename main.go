package main

import "net/http"

func main() {
	req, err := createRequest()
	checkErr(err)

	res, err := http.DefaultClient.Do(req)
	checkErr(err)

	handleResponse(res)
}
