package feature

import (
	"fiber-api/databases"
	"fiber-api/routes"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_Book(t *testing.T) {
	godotenv.Load("../../.env.testing")

	db := databases.SetupDB()
	databases.Migrate(db)
	app := routes.SetupRoutes(db)

	t.Run("[POST] Book - Successful", func(t *testing.T) {
		request := strings.NewReader(`{
			"title": "book 1",
			"author": "author 1"
		}`)

		r := httptest.NewRequest("POST", "/books", request)
		r.Header.Set("Content-Type", "application/json")
		r.Header.Set("Accept", "application/json")

		resp, _ := app.Test(r)

		defer resp.Body.Close()
		bodyR, _ := ioutil.ReadAll(resp.Body)

		expected := `{"message":"Đã tạo mới book thành công","status":200}`

		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, expected, string(bodyR))
	})

	t.Run("[Get] Book - Successful", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/books", nil)
		r.Header.Set("Content-Type", "application/json")
		r.Header.Set("Accept", "application/json")

		resp, _ := app.Test(r)

		defer resp.Body.Close()
		bodyR, _ := ioutil.ReadAll(resp.Body)

		expected := `{"data":[{"id":1,"title":"book 1","author":"author 1"}],"status":200}`

		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, expected, string(bodyR))
	})
}
