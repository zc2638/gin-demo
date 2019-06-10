package curl

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	METHOD_POST = "POST"
	METHOD_GET  = "GET"
)

type M map[string]interface{}

type HttpReq struct {
	Url      string
	Method   string
	Header   map[string]string
	Params   M
	Body     []byte
	CertFile string
	KeyFile  string
	Timeout  time.Duration
}

func (h *HttpReq) buildBody() {
	if h.Body != nil {
		return
	}

	var data string
	for k, v := range h.Params {
		if data != "" {
			data += "&"
		}
		data += k + "=" + v.(string)
	}

	switch h.Method {
	case METHOD_POST:
		h.Body = []byte(data)
		break
	case METHOD_GET:
		urlArr := strings.Split(h.Url, "?")
		if len(urlArr) == 2 {
			if data != "" {
				urlArr[1] = urlArr[1] + "&" + data
			}
			//将GET请求的参数进行转义
			h.Url = urlArr[0] + "?" + url.PathEscape(urlArr[1])
		}
		break
	}
}

func (h *HttpReq) Do() ([]byte, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if h.CertFile != "" {
		cert, err := tls.LoadX509KeyPair(h.CertFile, h.KeyFile)
		if err != nil {
			return nil, err
		}
		tr.DisableCompression = true
		tr.TLSClientConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	}
	var client = &http.Client{
		Transport: tr,
		Timeout:   h.Timeout,
	}
	h.buildBody()

	req, err := http.NewRequest(h.Method, h.Url, bytes.NewReader(h.Body))
	if err != nil {
		return nil, err
	}

	if h.Header != nil {
		for k, v := range h.Header {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (h *HttpReq) Get() ([]byte, error) {

	h.Method = METHOD_GET
	return h.Do()
}