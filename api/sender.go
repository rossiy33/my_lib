package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JsonMap map[string]interface{}

func JsonToJson(url string, jm JsonMap) (result JsonMap, err error) {
	var resp *http.Response
	if jm == nil {
		resp, err = http.Get(url)
	} else {
		var j []byte
		j, err = json.Marshal(jm)
		if err != nil {
			return
		}
		resp, err = http.Post(url, "application/json", bytes.NewBuffer(j))
	}
	if resp == nil {
		err = errors.New("null resp")
		return
	}
	if resp.Body == nil {
		err = errors.New("null resp.Body")
		return
	}
	defer resp.Body.Close()
	if err != nil {
		return
	}

	// レスポンスのBodyを読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if body == nil {
		err = errors.New("null body")
		return
	}

	// JSONをmapにデコード、できなかったらbodyに入れる
	err = json.Unmarshal(body, &result)
	if err != nil {
		if body == nil {
			return
		}
		result = JsonMap{}
		result["body"] = string(body)
		err = nil
	}

	return
}
