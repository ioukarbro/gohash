package curl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get(url string) (body []byte, err error) {
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		return
	}
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return
}

func PostForm(formValues map[string]string, url string) (body map[string]interface{}, err error) {
	var data map[string][]string
	for k, v := range formValues {
		data[k] = []string{v}
	}
	var resp *http.Response
	if resp, err = http.PostForm(url, data); err != nil {
		return
	}
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return
	}
	return
}

func PostJson(url string, jsonByte []byte) (body []byte, err error) {
	return PostJsonWithHeader(url, jsonByte, map[string]string{"Content-Type": "application/json"})
}

func PostJsonWithHeader(url string, jsonBytes []byte, headers map[string]string) (body []byte, err error) {
	var req *http.Request
	if req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes)); err != nil {
		return
	}
	for k, header := range headers {
		req.Header.Set(k, header)
	}
	client := &http.Client{}
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return
}
