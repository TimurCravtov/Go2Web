package printer

import (
	"encoding/json"
	"go2web/internal/request"

	"github.com/TylerBrock/colorjson"
)

func JsonPrinter(urlPath string, response *request.HttpResponse) (string, error) {

	var obj map[string]interface{}
	err := json.Unmarshal(response.Body, &obj)
	if err != nil {
		return "", err 
	}

	f := colorjson.NewFormatter()
	f.Indent = 4
	
	coloredBytes, err := f.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(coloredBytes), nil
}