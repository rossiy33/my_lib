package api

import (
	"bytes"
	"encoding/json"
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
			//log.Fatal(err)
			return
		}
		resp, err = http.Post(url, "application/json", bytes.NewBuffer(j))
	}
	defer resp.Body.Close()
	if err != nil {
		//log.Fatal(err)
		return
	}

	// レスポンスのBodyを読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		//log.Fatal(err)
		return
	}

	// JSONをmapにデコード、できなかったらbodyに入れる
	//var result JsonMap
	//if err := json.Unmarshal(body, &result); err != nil {
	//	log.Fatal(err)
	//}
	err = json.Unmarshal(body, &result)
	if err != nil {
		result = JsonMap{}
		result["body"] = string(body)
		err = nil
	}

	//return result
	return
}
