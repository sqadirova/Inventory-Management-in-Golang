package errors

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Response struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

type HttpError struct {
	Code     int `json:"code"`
	Response *Response
}

func (he HttpError) Error() string {
	return fmt.Sprintf("%s - %s", he.Response.Key, he.Response.Message)
}

func NewHttpError(status int, response *Response) HttpError {
	return HttpError{
		Code:     status,
		Response: response,
	}
}

var repo map[string]map[string]string

func init() {
	repo = readErrorFile()
}

func getErrorResponse(name, language string) string {
	if _, ok := repo[name][language]; ok {
		return repo[name][language]
	}
	return repo[name]["en"]
}

func NewResponseByKey(key, language string) *Response {
	return &Response{
		Key:     key,
		Message: getErrorResponse(key, language),
	}
}

func readErrorFile() map[string]map[string]string {
	byteValue, err := os.ReadFile("errors/errorResponse.json")

	if err != nil {
		log.Fatal(err)
	}

	var result map[string]map[string]string
	err = json.Unmarshal(byteValue, &result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}
