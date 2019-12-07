package dashboard

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto/")

	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Fprintf(w, string(body))
		}

	}
	defer resp.Body.Close()
}
