package service

import (
	"context"
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
