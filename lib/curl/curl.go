package curl

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Get(uri string) ([]byte, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout: 5 * time.Second, // 5s
	}

	urlArr := strings.Split(uri, "?")
	if len(urlArr) == 2 {
		//将GET请求的参数进行转义
		uri = urlArr[0] + "?" + url.PathEscape(urlArr[1])
	}
	var req *http.Request
	req, _ = http.NewRequest("GET", uri, nil)

	req.Header.Add("Authorization", "Basic YWRtaW46YWx3YXlzYmVjb2Rpbmc=")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}