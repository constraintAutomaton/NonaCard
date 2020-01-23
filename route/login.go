package route

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type oauthInfo struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	Code         string `json:"code"`
}

const oauthLink string = "https://anilist.co/api/v2/oauth/token"

// Login log or register the user
func Login(w http.ResponseWriter, r *http.Request) {
	defer http.Redirect(w, r, "", http.StatusSeeOther)
	token := r.URL.Query()["code"][0]
	jwt, err := getJwt(token)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(jwt)
		getUserInfo(jwt)
	}
}

func getJwt(pToken string) (string, error) {
	info := oauthInfo{GrantType: "authorization_code",
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURI:  os.Getenv("REDIRECT_URL"),
		Code:         pToken}
	infoJSON, err := json.Marshal(info)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(oauthLink, "application/json", bytes.NewBuffer(infoJSON))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}
func getUserInfo(jwt string) (string, error) {
	return "", nil
}
