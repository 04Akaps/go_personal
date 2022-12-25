package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// readAll은 그냥 r.Body로 들어오는 값을 다 읽어준다.
		if err := json.Unmarshal([]byte(body), x); err != nil {
			// 문제가 없을시에 x에 body값을 할당
			return
		}
	}
}
