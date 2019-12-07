package clientGraphQl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Body struct {
	query     string
	variables map[string]string
}

func Fetch(pUrl string, pQuery string, pVariables map[string]string, pRepInterface interface{}) (interface{}, error) {
	if nbVariable := NbOccurence(pQuery, "$"); nbVariable != len(pVariables) {
		return -1, errors.New("the number of variable in query don't match the variables pass")
	}
	for k := range pVariables {
		if index := strings.Index(pQuery, k); index == -1 {
			return -1, errors.New("variable don't match query")
		}
	}
	body := Body{pQuery, pVariables}
	b, err := json.Marshal(body)
	if err != nil {
		return -1, errors.New("Unable to create the JSON from the variables and the query")
	}
	resp, err := http.Post(pUrl, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return -1, errors.New("Unable to post to the api")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		m, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return -1, errors.New("The Api send an error")
		} else {
			return m, errors.New("The Api send an error")
		}
	}
	m, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, errors.New("Unable to read the response")
	}
	err = json.Unmarshal(m, &pRepInterface)
	if err != nil {
		return -1, errors.New("Unable to convert response to JSON")
	}
	return 1, nil

}

/**
NbOccurence(s string, pSubstring string) int
find the number of occurence of the substring in the string
*/
func NbOccurence(s string, pSubstring string) int {
	index := strings.Index(s, pSubstring)
	if index == -1 || len(pSubstring) == 0 || len(s) == 0 {
		return 0
	} else {
		if index+len(pSubstring) <= len(s) {
			return 1 + NbOccurence(s[index+len(pSubstring):], pSubstring)
		} else {
			return 0
		}
	}
}
