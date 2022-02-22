package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func esGet(uri string) []byte {
	esURL := fmt.Sprintf("%s://%s:%s/%s", esScheme, esHost, esPort, uri)
	resp, err := http.Get(esURL)
	errorExit(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errorExit(err)

	return body
}
