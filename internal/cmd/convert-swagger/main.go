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

func replaceMethodName(from string) string {
	methods := map[string]string{
		"post":   "Create",
		"put":    "Update",
		"patch":  "Update",
		"get":    "Get",
		"delete": "Delete",
	}

	if methods[from] != "" {
		return methods[from]
	}

	return from
}

func fixMethod(path string, op *spec.Operation) {
	if op != nil {
		r := regexp.MustCompile(`(delete|get|head|options|patch|post|put)\b`)
		id := strings.ToLower(op.ID)
		if r.MatchString(id) {
			op.ID = fmt.Sprintf("%s%s", replaceMethodName(id), op.Tags[0])
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

	// Add missing scheme: https
	api.Schemes = []string{"https"}

	// Add missing parameters for token
	parameters := []spec.Parameter{
		{
			ParamProps: spec.ParamProps{
				Name:     "username",
				In:       "body",
				Required: true,
			},
		},
		{
			ParamProps: spec.ParamProps{
				Name:     "password",
				In:       "body",
				Required: true,
			},
		},
		{
			ParamProps: spec.ParamProps{
				Name:     "grant_type",
				In:       "body",
				Required: true,
			},
		},
		{
			ParamProps: spec.ParamProps{
				Name:     "scope",
				In:       "body",
				Required: true,
			},
		},
	}

	api.Paths.Paths["/api/v1/connect/token"].Post.Parameters = parameters
	api.Paths.Paths["/api/v1/connect/token"].Post.Consumes = []string{"application/x-www-form-urlencoded"}

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
