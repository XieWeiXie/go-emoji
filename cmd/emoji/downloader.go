package emoji_fetch

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func DownLoader(url string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("you should add url")
	}
	request, _ := http.NewRequest("GET", url, nil)
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New("connect to internet fail")
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
