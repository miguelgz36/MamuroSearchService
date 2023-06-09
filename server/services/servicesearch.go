package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/miguelgz36/IndexerGolang/connection"
)

func Search(textToSearch string, page int) ([]byte, error) {

	urlSearch := "http://172.31.29.222:4080/es/_search"

	jsonBody := getJsonBodyQuery(textToSearch, page)

	reqBody, err := json.Marshal(jsonBody)
	if err != nil {
		return []byte("Error encoding params: " + err.Error()), err
	}

	fmt.Println("JSON:" + string(reqBody))

	req, err := http.NewRequest("POST", urlSearch, bytes.NewBuffer(reqBody))
	if err != nil {
		return []byte("Problem creating request: " + err.Error()), err
	}

	connection.SetHeaders(req)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return []byte("Problem doing post to zinc search: " + err.Error()), err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("Could not read body of zinc search service: " + err.Error()), err
	}

	fmt.Println("Result: " + string(body))
	return body, nil
}

func getJsonBodyQuery(textToSearch string, page int) map[string]interface{} {
	query_string := map[string]interface{}{
		"query": "\"" + textToSearch + "\"",
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
		"size":  15,
		"from":  page,
		"sort":  sort,
		"aggs":  aggs,
	}

	return result
}
