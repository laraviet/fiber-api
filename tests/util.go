package tests

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
)

func ApiRoute(method string, uri string, body io.Reader) *http.Request {
	request := httptest.NewRequest(method, uri, body)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	return request
}

func GetBody(app *fiber.App, r *http.Request) (*http.Response, string, error) {
	resp, _ := app.Test(r)
	defer resp.Body.Close()
	bodyR, err := ioutil.ReadAll(resp.Body)
	return resp, string(bodyR), err
}
