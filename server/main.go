package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ahsayde/yapl/yapl"
	"gopkg.in/yaml.v3"
)

type Request struct {
	Input  string `json:"input"`
	Policy string `json:"policy"`
	Params string `json:"params"`
}

type Error struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Errors      []string `json:"errors"`
}

func EvalHandler(rw http.ResponseWriter, rq *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Content-Type", "application/json")

	var request Request
	err := json.NewDecoder(rq.Body).Decode(&request)
	if err != nil {
		http.Error(rw, fmt.Sprintf("invalid request, error: %v", err), http.StatusBadRequest)
		return
	}

	var input map[string]interface{}
	if err = json.Unmarshal([]byte(request.Input), &input); err != nil {
		if err = yaml.Unmarshal([]byte(request.Input), &input); err != nil {
			e := Error{
				Type:        "invalid_input",
				Description: "input must be a valid json or yaml",
				Errors:      []string{err.Error()},
			}
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(e)
			return
		}
	}

	var params map[string]interface{}
	if err = json.Unmarshal([]byte(request.Params), &params); err != nil {
		if err = yaml.Unmarshal([]byte(request.Params), &params); err != nil {
			e := Error{
				Type:        "invalid_params",
				Description: "parameters must be a valid json or yaml",
				Errors:      []string{err.Error()},
			}
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(e)
			return
		}
	}

	policy, err := yapl.Parse([]byte(request.Policy))
	if err != nil {
		var errObject Error

		switch e := err.(type) {
		case yapl.ParserError:
			errObject = Error{
				Type:        "invalid_policy",
				Description: "invalid policy syntax",
			}
			for i := range e {
				errObject.Errors = append(errObject.Errors, e[i].Error())
			}
		default:
			errObject = Error{
				Type:        "invalid_policy",
				Description: "policy must be a valid yaml",
				Errors:      []string{err.Error()},
			}
		}

		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(errObject)
		return
	}

	result, err := policy.Eval(input, params)
	if err != nil {
		http.Error(rw, fmt.Sprintf("failed to validate input, error: %v", err), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(result)
}

func main() {
	http.HandleFunc("/eval", EvalHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic(err)
	}
}
