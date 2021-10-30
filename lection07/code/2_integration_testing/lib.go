package lecture07

import (
	"io/ioutil"
	"net/http"
)

func HTTPReq(addr string) (string, error) {
	resp, err := http.DefaultClient.Get(addr)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
