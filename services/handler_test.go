package services

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Coderx44/s3_store/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuit struct {
	suite.Suite
	service *mocks.Service
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuit))
}

func (suite *HandlerTestSuit) SetupTest() {
	suite.service = &mocks.Service{}
}

func (suite *HandlerTestSuit) TearDownSuite() {
	suite.service.AssertExpectations(suite.T())
}

func (s *HandlerTestSuit) TestGetPreSignedPutUrl() {
	type args struct {
		method   string
		endpoint string
	}
	tests := []struct {
		name       string
		args       args
		wantResp   string
		wantStatus int
		prepare    func(args, *mocks.Service, string)
	}{
		// POSITIVE || GetPresignedPutURL returns 200 status
		{
			name: "GetPresignedPutURL returns 200 status",
			args: args{
				method:   http.MethodGet,
				endpoint: "/put-presigned-url?object_name=test",
			},
			wantResp:   "https://exampleurl/hello/signed_url",
			wantStatus: http.StatusOK,
			prepare: func(a args, s *mocks.Service, wantUrl string) {
				s.On("GetPreSignedPutUrl", context.TODO(), "test").Return(wantUrl, nil).Once()
			},
		},
		// NEGATIVE || when object name is empty
		{
			name: "when object name is empty",
			args: args{
				method:   http.MethodGet,
				endpoint: "/put-presigned-url",
			},
			wantResp:   "object_name parameter is required\n",
			wantStatus: http.StatusBadRequest,
			prepare: func(a args, s *mocks.Service, wantUrl string) {
				// s.On("GetPreSignedPutUrl", context.TODO(), "test").Return(wantUrl, nil).Once()
				return
			},
		},
		// NEGATIVE || when error occurred at service layer
		{
			name: "when error occurred at service layer",
			args: args{
				method:   http.MethodGet,
				endpoint: "/put-presigned-url?object_name=test",
			},
			wantResp:   "Error generating signed URL\n",
			wantStatus: http.StatusInternalServerError,
			prepare: func(a args, s *mocks.Service, wantUrl string) {
				s.On("GetPreSignedPutUrl", context.TODO(), "test").Return("", errors.New("custom-error")).Once()
				return
			},
		},
	}
	t := s.T()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.args.method, tt.args.endpoint, nil)
			w := httptest.NewRecorder()
			tt.prepare(tt.args, s.service, tt.wantResp)

			got := GetPreSignedPutUrl(s.service)
			got.ServeHTTP(w, r)
			assert.Equal(t, w.Result().StatusCode, tt.wantStatus)
			assert.Equal(t, tt.wantResp, w.Body.String())
		})
	}
}

func (s *HandlerTestSuit) TestGetPreSignedGetUrl() {
	type args struct {
		method   string
		endpoint string
	}
	tests := []struct {
		name       string
		args       args
		wantResp   string
		wantStatus int
		prepare    func(args, *mocks.Service, string)
	}{
		// POSITIVE || GetPresignedGetURL returns 200 status
		{
			name: "GetPresignedGetURL returns 200 status",
			args: args{
				method:   http.MethodGet,
				endpoint: "/get-presigned-url?object_name=test",
			},
			wantResp:   "https://exampleurl/hello/signed_url",
			wantStatus: http.StatusOK,
			prepare: func(a args, s *mocks.Service, wantUrl string) {
				s.On("GetPreSignedGetUrl", context.TODO(), "test").Return(wantUrl, nil).Once()
			},
		},
		// NEGATIVE || when object name is empty
		{
			name: "when object name is empty",
			args: args{
				method:   http.MethodGet,
				endpoint: "/get-presigned-url",
			},
			wantResp:   "object_name parameter is required\n",
			wantStatus: http.StatusBadRequest,
			prepare: func(a args, s *mocks.Service, wantUrl string) {
				// s.On("GetPreSignedPutUrl", context.TODO(), "test").Return(wantUrl, nil).Once()
				return
			},
		},
		// NEGATIVE || when error occurred at service layer
		{
			name: "when error occurred at service layer",
			args: args{
				method:   http.MethodGet,
				endpoint: "/get-presigned-url?object_name=test",
			},
			wantResp:   "Error generating signed URL\n",
			wantStatus: http.StatusInternalServerError,
			prepare: func(a args, s *mocks.Service, wantUrl string) {
				s.On("GetPreSignedGetUrl", context.TODO(), "test").Return("", errors.New("custom-error")).Once()
				return
			},
		},
	}
	t := s.T()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.args.method, tt.args.endpoint, nil)
			w := httptest.NewRecorder()
			tt.prepare(tt.args, s.service, tt.wantResp)

			got := GetPreSignedGetUrl(s.service)
			got.ServeHTTP(w, r)
			assert.Equal(t, w.Result().StatusCode, tt.wantStatus)
			assert.Equal(t, tt.wantResp, w.Body.String())
		})
	}
}
