package feature

import (
	"fiber-api/databases"
	"fiber-api/routes"
	"fiber-api/tests"
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
		payload := strings.NewReader(`{
			"title": "book 1",
			"author": "author 1"
		}`)

		r := tests.ApiRoute("POST", "/books", payload)

		resp, bodyR, _ := tests.GetBody(app, r)

		expected := `{"message":"Đã tạo mới book thành công","status":200}`

		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, expected, string(bodyR))
	})

	t.Run("[Get] Book - Successful", func(t *testing.T) {
		r := tests.ApiRoute("GET", "/books", nil)

		resp, bodyR, _ := tests.GetBody(app, r)

		expected := `{"data":[{"id":1,"title":"book 1","author":"author 1"}],"status":200}`

		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, expected, string(bodyR))
	})

	t.Run("[PUT] Book - Successful", func(t *testing.T) {
		payload := strings.NewReader(`{
			"title": "book 1 edit",
			"author": "author 1 edit"
		}`)

		r := tests.ApiRoute("PUT", "/books/1", payload)

		resp, bodyR, _ := tests.GetBody(app, r)

		expected := `{"message":"Đã cập nhật book thành công","status":200}`

		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, expected, string(bodyR))
	})

	t.Run("[DELETE] Book - Successful", func(t *testing.T) {
		r := tests.ApiRoute("DELETE", "/books/1", nil)

		resp, bodyR, _ := tests.GetBody(app, r)

		expected := `{"message":"Đã xóa book thành công","status":200}`

		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, expected, string(bodyR))
	})
}
