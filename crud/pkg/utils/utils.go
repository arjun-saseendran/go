package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, user interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	} else {
		json.Unmarshal([]byte(body), user)
	}

}
