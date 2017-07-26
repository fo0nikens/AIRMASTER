package domain

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
)

func ParseFile(filePath string) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return err
	}

	client := &http.Client{
		Jar: jar,
	}

	for _, url := range strings.Split(string(file), "\n") {
		if err := CheckDomain(url, client, 0); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	return nil
}