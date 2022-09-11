package handler

import (
	"login-app/internal/core/domain"
	"login-app/internal/core/services"
	"login-app/platform/logger"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockUserRepository struct {
}

func (r MockUserRepository) Insert(user *domain.User) (string, error) {
	return "id", nil
}

func TestBadRequestWhenPostMethod(t *testing.T) {
	// expectBody := `{"code":"bad_request","message":"Bad Request","data":{}}`

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/v1/users", strings.NewReader("invalid-json"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	h := &Handler{}

	if assert.NoError(t, h.Post(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		// assert.Equal(t, expectBody, strings.Trim(rec.Body.String(), "\n"))
	}
}

func TestSuccessCreatedWhenPostMethod(t *testing.T) {
	expectBody := `{"code":"created","message":"Created","data":{"id":"id"}}`

	requestBody := `{"name":"john","email":"johndoe@email.com","password":"12345678","password_confirmation":"12345678"}`

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	h := &Handler{services.NewService(logger.LoggerMock{}, MockUserRepository{})}

	if assert.NoError(t, h.Post(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectBody, strings.Trim(rec.Body.String(), "\n"))
	}
}
