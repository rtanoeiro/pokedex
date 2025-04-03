package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRequest(url string) ([]byte, error) {

	response, httpErr := http.Get(url)
		
	if response.StatusCode != 200 {
		return []byte{}, fmt.Errorf("Unable to get data, try again.")
	}

	if httpErr != nil {
		fmt.Println()
		return []byte{}, fmt.Errorf("Error getting data, try again.")
	}

	resData, ioError := io.ReadAll(response.Body)

	if ioError != nil {
		return []byte{}, fmt.Errorf("Found error when reading Body from HTTP Get Response")
	}

	return resData, nil

}