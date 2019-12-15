package main

import (
	"fmt"

	"github.com/featTheB/anifan-card/pkg/clientGraphQl"
)

func main() {
	query := `query ($id: Int) { # Define which variables will be used in the query (id)
		Media (id: $id, type: ANIME) { # Insert our variables into the query arguments (id) (type: ANIME is hard-coded in the query)
		  id
		  title {
			romaji
			english
			native
		  }
		}
	  }`
	variables := make(map[string]string)
	variables["id"] = "15125"
	url := "https://graphql.anilist.co"
	res := make(map[string]interface{})
	m, err := clientGraphQl.Fetch(url, query, variables, &res)
	if err != nil {
		fmt.Println(err)
		fmt.Println(m)

	}
	fmt.Print(res)

}
