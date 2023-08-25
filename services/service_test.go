package services

import (
	"context"
	"errors"
	"testing"

	"github.com/Coderx44/s3_store/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuit struct {
	suite.Suite
	Service
	storage *mocks.Storage
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuit))
}

func (suite *ServiceTestSuit) SetupTest() {
	suite.storage = &mocks.Storage{}
	suite.Service = NewStorageService(suite.storage)
}

func (suite *ServiceTestSuit) TearDownSuite() {
	suite.storage.AssertExpectations(suite.T())
}

func (s *ServiceTestSuit) Test_service_GetPreSignedPutUrl() {

	type args struct {
		ctx        context.Context
		objectName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		prepare func(args, *mocks.Storage, string)
	}{
		// POSITIVE | when put url gets generated
		{
			name: "when put url gets generated",
			args: args{
				ctx:        context.TODO(),
				objectName: "hello.txt",
			},
			want:    "https://exampleurl/hello/signed_url",
			wantErr: false,
			prepare: func(a args, s *mocks.Storage, want string) {
				s.On("GetPreSignedPutUrl", a.ctx, a.objectName).Return(want, nil).Once()
			},
		},
		// NEGATIVE | when storage returns an error
		{
			name: "when storage returns an error",
			args: args{
				ctx:        context.TODO(),
				objectName: "hello.txt",
			},
			want:    "",
			wantErr: true,
			prepare: func(a args, s *mocks.Storage, want string) {
				s.On("GetPreSignedPutUrl", a.ctx, a.objectName).Return(want, errors.New("custom error")).Once()
			},
		},
	}
	for _, tt := range tests {
		t := s.T()
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, s.storage, tt.want)
			got, err := s.Service.GetPreSignedPutUrl(tt.args.ctx, tt.args.objectName)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetPreSignedPutUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.GetPreSignedPutUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *ServiceTestSuit) Test_service_GetPreSignedGetUrl() {
	type args struct {
		ctx        context.Context
		objectName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		prepare func(args, *mocks.Storage, string)
	}{
		// POSITIVE | when get url gets generated
		{
			name: "when get url gets generated",
			args: args{
				ctx:        context.TODO(),
				objectName: "hello.txt",
			},
			want:    "https://exampleurl/hello/signed_url",
			wantErr: false,
			prepare: func(a args, s *mocks.Storage, want string) {
				s.On("GetPreSignedGetUrl", a.ctx, a.objectName).Return(want, nil).Once()
			},
		},
		// NEGATIVE | when storage returns an error
		{
			name: "when storage returns an error",
			args: args{
				ctx:        context.TODO(),
				objectName: "hello.txt",
			},
			want:    "",
			wantErr: true,
			prepare: func(a args, s *mocks.Storage, want string) {
				s.On("GetPreSignedGetUrl", a.ctx, a.objectName).Return(want, errors.New("custom error")).Once()
			},
		},
	}
	for _, tt := range tests {
		t := s.T()
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, s.storage, tt.want)
			got, err := s.Service.GetPreSignedGetUrl(tt.args.ctx, tt.args.objectName)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetPreSignedPutUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.GetPreSignedPutUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
