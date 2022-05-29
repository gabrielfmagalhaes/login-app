package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBadRequestWhenPostMethod(t *testing.T) {
	expectBody := "{\"code\":\"bad_request\",\"message\":\"Bad Request\",\"data\":{}}\n"

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("invalid-json"))
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	c.SetPath("/v1/users")

	h := &Handler{}

	if assert.NoError(t, h.Post(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, expectBody, rec.Body.String())
	}
}

func TestSuccessCreatedWhenPostMethod(t *testing.T) {
	expectBody := "{\"code\":\"created\",\"message\":\"Created\",\"data\":{}}\n"

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	c.SetPath("/v1/users")

	h := &Handler{}

	if assert.NoError(t, h.Post(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectBody, rec.Body.String())
	}
}
