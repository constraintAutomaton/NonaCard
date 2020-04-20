package remotecommuncation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GraphQlClient struct{}

func (g GraphQlClient) Fetch(query ParameterQuery) error {

	b, err := formatQuery(query.Query, query.Variables)
	if err != nil {
		return nil
	}
	resp, err := http.Post(query.Url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = getResponse(resp, query.Out)
	if err != nil {
		return err
	}
	return nil
}

func (g GraphQlClient) formatQuery(pQuery string, pVariables *map[string]string) ([]byte, error) {
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
func (g GraphQlClient) getResponse(resp *http.Response, pRepInterface interface{}) error {
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
