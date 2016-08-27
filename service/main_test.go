package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type respData struct {
	Code     int      `json:"code"`
	Error    string   `json:"error"`
	Mess     string   `json:"mess"`
	Keywords []string `json:"keywords"`
	Text     string   `json:"text"`
}

func TestFailToAddBlackWords(t *testing.T) {
	postURL := "http://127.0.0.1:8080/v1/black_words"
	resp, err := http.PostForm(postURL, url.Values{})
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected HTTP status is 200, but get %d\n", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	data := respData{}
	json.Unmarshal(body, &data)

	t.Log(string(body))
	if data.Code != 0 {
		t.Errorf("code: %d, error: %s, mess: %s", data.Code, data.Error, data.Mess)
	}
}

func TestAddBlackWordsSuccessfully(t *testing.T) {
	postURL := "http://127.0.0.1:8080/v1/black_words"
	resp, err := http.PostForm(postURL, url.Values{"q": {"测试"}})
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected HTTP status is 200, but get %d\n", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	data := respData{}
	json.Unmarshal(body, &data)

	t.Log(string(body))
	if data.Code == 0 {
		t.Errorf("code: %d, error: %s, mess: %s", data.Code, data.Error, data.Mess)
	}
}

func TestDeleteBlackWordsSuccess(t *testing.T) {
	delURL := "http://127.0.0.1:8080/v1/black_words"

	r, err := http.NewRequest("DELETE", delURL, strings.NewReader(`{"q": "测试,test02"}`))
	r.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(r)

	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected HTTP status is 200, but get %d\n", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	data := respData{}
	json.Unmarshal(body, &data)

	t.Log(string(body))
	if data.Code == 0 {
		t.Errorf("code: %d, error: %s, mess: %s", data.Code, data.Error, data.Mess)
	}
}

func TestQuerySuccessfullyWithoutBlackWords(t *testing.T) {
	queryURL := "http://127.0.0.1:8080/v1/query"
	resp, err := http.PostForm(queryURL, url.Values{"q": {"完全和谐的文本"}})
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected HTTP status is 200, but get %d\n", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	data := respData{}
	json.Unmarshal(body, &data)

	if data.Code != 1 {
		t.Errorf("code: %d, error: %s, mess: %s", data.Code, data.Error, data.Mess)
	}

	t.Log(data.Keywords)
	t.Log(data.Text)
	if len(data.Keywords) != 0 {
		t.Error("过滤失败")
	}
}

func TestQuerySuccessfullyWithBlackWords(t *testing.T) {
	queryURL := "http://127.0.0.1:8080/v1/query"
	resp, err := http.PostForm(queryURL, url.Values{"q": {"对于1989年诺贝尔和平奖有什么看法"}})
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected HTTP status is 200, but get %d\n", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	data := respData{}
	json.Unmarshal(body, &data)

	if data.Code == 0 {
		t.Errorf("code: %d, error: %s, mess: %s", data.Code, data.Error, data.Mess)
	}

	t.Log(data.Keywords)
	t.Log(data.Text)
	if data.Keywords[0] != "1989年诺贝尔和平奖" {
		t.Error("过滤失败")
	}
}
