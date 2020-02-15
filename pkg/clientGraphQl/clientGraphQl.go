package clientGraphQl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
func Fetch(pURL string, pQuery string, pVariables map[string]string, pRepInterface *map[string]interface{}) (int, error)
Fetch a graphQl Api
*/
func Fetch(pURL string, pQuery string, pVariables *map[string]string, pRepInterface interface{}, pAuthorization ...string) error {

	b, err := formatQuery(pQuery, pVariables)
	if err != nil {
		return nil
	}
	resp, err := http.Post(pURL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = getResponse(resp, pRepInterface)
	if err != nil {
		return err
	}
	return nil
}

func formatQuery(pQuery string, pVariables *map[string]string) ([]byte, error) {
	body := map[string]interface{}{
		"query":     pQuery,
		"variables": *pVariables,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return make([]byte, 0), err
	}
	return b, nil
}
func getResponse(resp *http.Response, pRepInterface interface{}) error {
	if resp.StatusCode != 200 {
		return fmt.Errorf("API send http error %d", resp.StatusCode)
	}
	m, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(m, pRepInterface)
	if err != nil {
		return err
	}
	return nil
}
