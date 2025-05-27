package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchAndUnmarshal(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in GET: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in read body: %v\n", err)
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in unmarshal: %v\n", err)
		return err
	}
	return nil
}
