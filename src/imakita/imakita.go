package imakita

import (
	"fmt"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"net/http"
)
var yahooApiKey = ""
var a3rtApiKey = ""

type A3rt struct {
	Status     int      `json:"status"`
	Message    string   `json:"message"`
	Suggestion []string `json:"suggestion"`
}

func getKeyPhrase(format string, sentence string) (map[string]uint, error) {
	//【注意】apikeyをコミットしないように！！
	url := fmt.Sprintf("https://jlp.yahooapis.jp/KeyphraseService/V1/extract?appid=%s&output=%s&sentence=%s", yahooApiKey, format, url.QueryEscape(sentence))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := new(http.Client)
	resp, _ := client.Do(req)
	body, _:= ioutil.ReadAll(resp.Body)
	var keyWords map[string]uint
	if err := json.Unmarshal(body, &keyWords); err != nil {
		return nil, err
    }
	return keyWords, nil
}
//TODO:2行以下の場合はなんか付け足す。
func Imakita(sentence string) ([]string, error) {
	var imakita3 []string
	keyWords, err := getKeyPhrase("json", sentence)
	if err != nil {
		return nil, err
	}
	client := new(http.Client)
	for word, _ := range keyWords {
		sentence, err := getSentence(word, client)
		if err != nil {
			return nil, err
		}
		sentence = word + sentence
		imakita3 = append(imakita3, sentence)
	}
	return imakita3, nil
}

func getSentence(word string, client *http.Client) (string, error) { 
	//【注意】apikeyをコミットしないように！!
	url := fmt.Sprintf("https://api.a3rt.recruit-tech.co.jp/text_suggest/v2/predict?apikey=%s&previous_description=%s", a3rtApiKey,  url.QueryEscape(word))
	var v A3rt
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	resp, _ := client.Do(req)
	body, _:= ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &v); err != nil {
		return "", err
	}
	return v.Suggestion[0], nil
}