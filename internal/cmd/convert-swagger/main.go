package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/go-openapi/spec"
)

func fixMethod(path string, op *spec.Operation) {
	if op != nil {
		r := regexp.MustCompile(`(delete|get|head|options|patch|post|put)\b`)
		if r.MatchString(strings.ToLower(op.ID)) {
			op.ID = fmt.Sprintf("%s%s", op.ID, op.Tags[0])
		}
	}
}

func main() {
	resp, err := http.Get("https://api.crayon.com/swagger/v1/swagger.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	api := spec.Swagger{}
	err = json.Unmarshal(body, &api)
	if err != nil {
		panic(err)
	}

	for i, h := range api.Paths.Paths {
		fixMethod(i, h.Delete)
		fixMethod(i, h.Get)
		fixMethod(i, h.Head)
		fixMethod(i, h.Options)
		fixMethod(i, h.Patch)
		fixMethod(i, h.Post)
		fixMethod(i, h.Put)
	}

	raw, err := json.MarshalIndent(api, "", "   ")
	if err != nil {
		panic(err)
	}
	sw, err := os.Create("./swagger.json")
	if err != nil {
		panic(err)
	}
	_, err = sw.Write(raw)
	if err != nil {
		panic(err)
	}
	defer sw.Close()
}
