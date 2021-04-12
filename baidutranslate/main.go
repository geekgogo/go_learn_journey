package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	const appid string = "20200818000545881"
	const key string = "yCCXps_bkqBynm0lyAFu"
	const salt string = "1435660288"
	q := "Data Set Designer - Authorized administrators use the Data Set Designer to create data set definitions (ADS definition) as a hierarchical structure of records and their collective properties. A data set definition, with its group of records, constitutes a data set. Both record definitions and data set definitions are metadata that define the shape of the migration data."

	sign := appid + q + salt + key
	//给sign加md5值
	w := md5.New()
	io.WriteString(w, sign)
	md5sign := fmt.Sprintf("%x", w.Sum(nil))

	addr := "http://api.fanyi.baidu.com/api/trans/vip/translate?" + "q=" + url.QueryEscape(q) + "&from=en&to=zh&appid=" + appid + "&salt=" + salt + "&sign=" + md5sign

	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	//解析 resp
	statusCode := resp.StatusCode
	if statusCode == 200 {
		r, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var data map[string]interface{}
		var result string
		if err := json.Unmarshal([]byte(string(r)), &data); err == nil {
			transResult := data["trans_result"].([]interface{})
			for i := 0; i < len(transResult); i++ {
				result += fmt.Sprintf("%v\n", transResult[i].(map[string]interface{})["dst"])
			}
		} else {
			fmt.Println(err)
		}
		fmt.Println(result)
	} else {
		fmt.Println(statusCode)
	}

}
