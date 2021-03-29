package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	AuthController "github.com/Kourin1996/go-crud-api-sample/api/controllers/v1/auth"
	"github.com/Kourin1996/go-crud-api-sample/api/models/auth"
	"github.com/Kourin1996/go-crud-api-sample/tests/helper"
	"github.com/Kourin1996/go-crud-api-sample/tests/helper/mock/services"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	type Args struct {
		body map[string]interface{}
		res  *auth.AuthResult
		err  error
	}
	type Expected struct {
		isServiceCalled bool
		dto             *auth.SignUpDto
		status          int
		err             error
		body            map[string]interface{}
	}

	tests := []struct {
		name     string
		args     Args
		expected Expected
	}{
		{
			name: "should be success",
			args: Args{
				body: map[string]interface{}{
					"username": "Test Username",
					"password": "Test Password",
				},
				res: &auth.AuthResult{Token: "Test Token"},
				err: nil,
			},
			expected: Expected{
				isServiceCalled: true,
				dto: &auth.SignUpDto{
					Username: "Test Username",
					Password: "Test Password",
				},
				status: http.StatusOK,
				err:    nil,
				body: map[string]interface{}{
					"token": "Test Token",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := helper.MustMarshalJSON(t, tt.args.body)
			resBody := helper.MustMarshalJSON(t, tt.expected.body)

			e := helper.NewTestEcho()
			g := e.Group("/v1")
			req := httptest.NewRequest(http.MethodPost, "/auth/signup", strings.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			mockAuthService := &services.MockAuthService{}
			mockAuthService.On("SignUp", tt.expected.dto).Return(tt.args.res, tt.args.err)
			c := AuthController.NewAuthController(g, mockAuthService)

			err := c.SignUp(ctx)
			if err != nil {
				e.HTTPErrorHandler(err, ctx)
			}

			if tt.expected.isServiceCalled {
				mockAuthService.AssertExpectations(t)
			} else {
				mockAuthService.AssertNotCalled(t, "SignUp", tt.expected.dto)
			}
			if tt.expected.err == nil {
				assert.NoError(t, err)
				mockAuthService.AssertExpectations(t)
			} else {
				assert.Error(t, err)
				assert.IsType(t, tt.expected.err, err)
				assert.Equal(t, tt.expected.err.Error(), err.Error())
			}
			assert.Equal(t, tt.expected.status, rec.Code)
			assert.JSONEq(t, resBody, rec.Body.String())
		})
	}
}

func TestSignIn(t *testing.T) {
	type Args struct {
		body map[string]interface{}
		res  *auth.AuthResult
		err  error
	}
	type Expected struct {
		isServiceCalled bool
		dto             *auth.SignInDto
		status          int
		err             error
		body            map[string]interface{}
	}

	tests := []struct {
		name     string
		args     Args
		expected Expected
	}{
		{
			name: "should be success",
			args: Args{
				body: map[string]interface{}{
					"username": "Test Username",
					"password": "Test Password",
				},
				res: &auth.AuthResult{Token: "Test Token"},
				err: nil,
			},
			expected: Expected{
				isServiceCalled: true,
				dto: &auth.SignInDto{
					Username: "Test Username",
					Password: "Test Password",
				},
				status: http.StatusOK,
				err:    nil,
				body: map[string]interface{}{
					"token": "Test Token",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := helper.MustMarshalJSON(t, tt.args.body)
			resBody := helper.MustMarshalJSON(t, tt.expected.body)

			e := helper.NewTestEcho()
			g := e.Group("/v1")
			req := httptest.NewRequest(http.MethodPost, "/auth/signin", strings.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			mockAuthService := &services.MockAuthService{}
			mockAuthService.On("SignIn", tt.expected.dto).Return(tt.args.res, tt.args.err)
			c := AuthController.NewAuthController(g, mockAuthService)

			err := c.SignIn(ctx)
			if err != nil {
				e.HTTPErrorHandler(err, ctx)
			}

			if tt.expected.isServiceCalled {
				mockAuthService.AssertExpectations(t)
			} else {
				mockAuthService.AssertNotCalled(t, "SignIn", tt.expected.dto)
			}
			if tt.expected.err == nil {
				assert.NoError(t, err)
				mockAuthService.AssertExpectations(t)
			} else {
				assert.Error(t, err)
				assert.IsType(t, tt.expected.err, err)
				assert.Equal(t, tt.expected.err.Error(), err.Error())
			}
			assert.Equal(t, tt.expected.status, rec.Code)
			assert.JSONEq(t, resBody, rec.Body.String())
		})
	}
}
