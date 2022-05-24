package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"

	lofo "github.com/zhandosmd/go-final-project"
	"github.com/zhandosmd/go-final-project/pkg/service"
	mocks "github.com/zhandosmd/go-final-project/pkg/service/mocks"
)

func TestHandler_signUp(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mocks.MockAuthorization, user lofo.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            lofo.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username": "person", "name": "person name", "password": "qwerty123"}`,
			inputUser: lofo.User{
				Username: "username",
				Name:     "Test Name",
				Password: "qwerty",
			},
			mockBehavior: func(r *mocks.MockAuthorization, user lofo.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"username": "person"}`,
			inputUser:            lofo.User{},
			mockBehavior:         func(r *mocks.MockAuthorization, user lofo.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"username": "person", "name": "person name", "password": "qwerty123"}`,
			inputUser: lofo.User{
				Username: "username",
				Name:     "Test Name",
				Password: "qwerty123",
			},
			mockBehavior: func(r *mocks.MockAuthorization, user lofo.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			repo := mocks.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
