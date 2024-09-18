package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func HttpGet2Obj(url string, header map[string]string, v any) error {
	response, err := HttpGet(url, header)
	if err != nil {
		return err
	}
	respBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}
	err = json.Unmarshal(respBytes, v)
	return err
}
func HttpGet(url string, header map[string]string) (map[string]interface{}, error) {
	var ret map[string]interface{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// LogOut("error", fmt.Sprintf("http get new failed, url: %s, err: %s", url, err.Error()))
		return ret, err
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Type", "en_US")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{
		Timeout: 600 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		// LogOut("error", fmt.Sprintf("http get send failed, url: %s, err: %s", url, err.Error()))
		return ret, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// LogOut("error", "http get resp failed, url:"+url+", Body :"+lib1.Json_Package(resp)+"result:"+lib1.Json_Package(ret))
		return ret, errors.New("wrong http code" + strconv.Itoa(resp.StatusCode))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ret)
	if err != nil {
		// LogOut("error", fmt.Sprintf("http get unmarshal failed, url: %s, err: %s", url, err.Error()))
		return ret, err
	}

	//LogOut("info", fmt.Sprintf("http get success, url: %s, resp: %v", url, header, ret))
	return ret, nil
}

func HttpPost(url string, header map[string]string) (map[string]interface{}, error) {
	var ret map[string]interface{}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		// LogOut("error", fmt.Sprintf("http get new failed, url: %s, err: %s", url, err.Error()))
		return ret, err
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Type", "en_US")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{
		Timeout: 600 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		// LogOut("error", fmt.Sprintf("http get send failed, url: %s, err: %s", url, err.Error()))
		return ret, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// LogOut("error", "http get resp failed, url:"+url+", Body :"+lib1.Json_Package(resp)+"result:"+lib1.Json_Package(ret))
		return ret, errors.New("wrong http code" + strconv.Itoa(resp.StatusCode))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ret)
	if err != nil {
		// LogOut("error", fmt.Sprintf("http get unmarshal failed, url: %s, err: %s", url, err.Error()))
		return ret, err
	}

	//LogOut("info", fmt.Sprintf("http get success, url: %s, resp: %v", url, header, ret))
	return ret, nil
}
