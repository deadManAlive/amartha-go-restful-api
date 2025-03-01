package service

import (
	"context"
	"errors"
	"testing"

	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCustomerRepository(ctrl)
	mockValidator := validator.New()
	customerService := NewCustomerService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.CustomerCreateRequest
		mock      func()
		expect    web.CustomerResponse
		expectErr bool
	}{
		{
			name: "success",
			input: web.CustomerCreateRequest{
				Name:    "Prabowo Subianto",
				Email:   "psubianto@mamsiang.gratis",
				Phone:   "+628123456789",
				Address: "Istana Negara",
			},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Customer{
					Id:         1,
					Name:       "Prabowo Subianto",
					Email:      "psubianto@mamsiang.gratis",
					Phone:      "+628123456789",
					Address:    "Istana Negara",
					LoyaltyPts: 0,
				}, nil)
			},
			expect: web.CustomerResponse{
				Id:         1,
				Name:       "Prabowo Subianto",
				Email:      "psubianto@mamsiang.gratis",
				Phone:      "+628123456789",
				Address:    "Istana Negara",
				LoyaltyPts: 0,
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := customerService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestDeleteCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCustomerRepository(ctrl)
	customerService := NewCustomerService(mockRepo, validator.New())

	tests := []struct {
		name       string
		customerId uint64
		mock       func()
		expectErr  bool
	}{
		{
			name:       "Success",
			customerId: 1,
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), uint64(1)).Return(domain.Customer{
					Id:         1,
					Name:       "Prabowo Subianto",
					Email:      "psubianto@mamsiang.gratis",
					Phone:      "+628123456789",
					Address:    "Istana Negara",
					LoyaltyPts: 0,
				}, nil)
				mockRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectErr: false,
		},
		{
			name:       "Not found",
			customerId: 999,
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), uint64(999)).Return(domain.Customer{}, errors.New("not found"))
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := customerService.Delete(context.Background(), tt.customerId)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateCustomer(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(*mocks.MockCustomerRepository)
		input   web.CustomerUpdateRequest
		expects error
	}{
		{
			name: "Success",
			mock: func(mcr *mocks.MockCustomerRepository) {
				mcr.EXPECT().FindById(gomock.Any(), uint64(1)).Return(domain.Customer{
					Id:         1,
					Name:       "Prabowo Subianto",
					Email:      "psubianto@mamsiang.gratis",
					Phone:      "+628123456789",
					Address:    "Istana Negara",
					LoyaltyPts: 0,
				}, nil)
				mcr.EXPECT().Update(gomock.Any(), gomock.Any()).Return(domain.Customer{
					Id:         1,
					Name:       "Prabowo Subianto",
					Email:      "psubianto@mamsiang.gratis",
					Phone:      "+628123456789",
					Address:    "Istana Nusantara",
					LoyaltyPts: 0,
				}, nil)
			},
			input: web.CustomerUpdateRequest{
				Id:      1,
				Name:    "Prabowo Subianto",
				Email:   "psubianto@mamsiang.gratis",
				Phone:   "+628123456789",
				Address: "Istana Nusantara",
			},
			expects: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockCustomerRepo := mocks.NewMockCustomerRepository(ctrl)
			tt.mock(mockCustomerRepo)

			service := NewCustomerService(mockCustomerRepo, validator.New())
			_, err := service.Update(context.Background(), tt.input)

			if tt.expects != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expects.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFindByIdCustomer(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(*mocks.MockCustomerRepository)
		input   uint64
		expects web.CustomerResponse
		err     error
	}{
		{
			name: "Success",
			mock: func(mcr *mocks.MockCustomerRepository) {
				mcr.EXPECT().FindById(gomock.Any(), uint64(1)).Return(domain.Customer{
					Id:         1,
					Name:       "Prabowo Subianto",
					Email:      "psubianto@mamsiang.gratis",
					Phone:      "+628123456789",
					Address:    "Istana Negara",
					LoyaltyPts: 0,
				}, nil)
			},
			input: 1,
			expects: web.CustomerResponse{
				Id:         1,
				Name:       "Prabowo Subianto",
				Email:      "psubianto@mamsiang.gratis",
				Phone:      "+628123456789",
				Address:    "Istana Negara",
				LoyaltyPts: 0,
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCustomerRepo := mocks.NewMockCustomerRepository(ctrl)
			tt.mock(mockCustomerRepo)

			service := NewCustomerService(mockCustomerRepo, validator.New())
			result, err := service.FindById(context.Background(), tt.input)
			assert.Equal(t, tt.expects, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestFindAllCustomers(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(*mocks.MockCustomerRepository)
		expects []web.CustomerResponse
		err     error
	}{
		{
			name: "Success",
			mock: func(mcr *mocks.MockCustomerRepository) {
				mcr.EXPECT().FindAll(gomock.Any()).Return([]domain.Customer{{
					Id:         1,
					Name:       "Prabowo Subianto",
					Email:      "psubianto@mamsiang.gratis",
					Phone:      "+628123456789",
					Address:    "Istana Negara",
					LoyaltyPts: 0,
				}}, nil)
			},
			expects: []web.CustomerResponse{{
				Id:         1,
				Name:       "Prabowo Subianto",
				Email:      "psubianto@mamsiang.gratis",
				Phone:      "+628123456789",
				Address:    "Istana Negara",
				LoyaltyPts: 0,
			}},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCustomerRepo := mocks.NewMockCustomerRepository(ctrl)
			tt.mock(mockCustomerRepo)

			service := NewCustomerService(mockCustomerRepo, validator.New())
			result, err := service.FindAll(context.Background())
			assert.Equal(t, tt.expects, result)
			assert.Equal(t, tt.err, err)
		})
	}
}
