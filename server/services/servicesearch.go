package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/miguelgz36/IndexerGolang/connection"
)

func Search(textToSearch string) []byte {

	urlSearch := "http://localhost:4080/es/_search"

	jsonBody := getJsonBodyQuery(textToSearch)

	reqBody, err := json.Marshal(jsonBody)
	if err != nil {
		fmt.Println("Error encoding params:", err)
	}

	fmt.Println("JSON:" + string(reqBody))

	req, err := http.NewRequest("POST", urlSearch, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err.Error())
	}

	connection.SetHeaders(req)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("ZINC: " + string(body))
	return body
}

func getJsonBodyQuery(textToSearch string) map[string]interface{} {
	query_string := map[string]interface{}{
		"query": "message:equals('" + textToSearch + "')",
	}

	must := []map[string]interface{}{
		{"query_string": query_string},
	}

	bool0 := map[string]interface{}{
		"must": must,
	}

	query := map[string]interface{}{
		"bool": bool0,
	}

	auto_date_histogram := map[string]interface{}{
		"bucket": 100,
		"field":  "@timestamp",
	}

	histogram := map[string]interface{}{
		"auto_date_histogram": auto_date_histogram,
	}

	aggs := map[string]interface{}{
		"histogram": histogram,
	}

	sort := []string{
		"-@timestamp",
	}

	result := map[string]interface{}{
		"query": query,
		"size":  10,
		"from":  0,
		"sort":  sort,
		"aggs":  aggs,
	}

	return result
}
