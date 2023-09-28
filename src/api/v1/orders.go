package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ReturnError(w http.ResponseWriter, errString string) {
	w.Header().Set("Content-Type", "application/json")
	str := fmt.Sprintf("{\"error\" : %v, \"message\" : \"%v\"}", true, errString)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(str))
}

func GetParamId(URL string) string {
	id := strings.Split(URL, "/")
	return id[len(id)-1]
}

func ReturnResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	js, _ := json.Marshal(data)
	w.Write(js)
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	id := GetParamId(r.URL.Path)
	ReturnResponse(w, "{\"status\":\"ok\", \"id\":\""+id+"\"}")
}
