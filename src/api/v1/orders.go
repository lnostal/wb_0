package v1

import (
	"L0/src/cache"
	"L0/src/db/model"
	"L0/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func ReturnError(w http.ResponseWriter, errString string) {
	w.Header().Set("Content-Type", "application/json")
	str := fmt.Sprintf("{\"error\" : %v, \"message\" : \"%v\"}", true, errString)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(str))
}

func ReturnResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	js, _ := json.Marshal(data)
	w.Write(js)
}

// -------

func getParamId(URL string) string {
	id := strings.Split(URL, "/")
	return id[len(id)-1]
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	idStr := getParamId(r.URL.Path)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ReturnError(w, err.Error())
		return
	}

	order, err := cache.AtIndex(id - 1)
	if err != nil {
		ReturnError(w, err.Error())
		return
	}

	var o model.Order

	ReturnResponse(w, utils.ConvertToModel(order, o))
}
