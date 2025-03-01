package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCustomerRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockCustomerRepository(ctrl)
	ctx := context.Background()

	tests := []struct {
		name      string
		mock      func()
		method    func() (interface{}, error)
		expect    interface{}
		expectErr bool
	}{
		{
			name: "Save success",
			mock: func() {
				category := domain.Customer{
					Id:      1,
					Name:    "Prabowo Subianto",
					Email:   "psubianto@mamsiang.gratis",
					Phone:   "+628123456789",
					Address: "Istana Negara",
				}
				repo.EXPECT().Save(ctx, domain.Customer{
					Id:      1,
					Name:    "Prabowo Subianto",
					Email:   "psubianto@mamsiang.gratis",
					Phone:   "+628123456789",
					Address: "Istana Negara",
				}).Return(category, nil)
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Customer{
					Id:      1,
					Name:    "Prabowo Subianto",
					Email:   "psubianto@mamsiang.gratis",
					Phone:   "+628123456789",
					Address: "Istana Negara",
				})
			},
			expect: domain.Customer{
				Id:      1,
				Name:    "Prabowo Subianto",
				Email:   "psubianto@mamsiang.gratis",
				Phone:   "+628123456789",
				Address: "Istana Negara",
			},
			expectErr: false,
		},
		{
			name: "Save failure",
			mock: func() {
				repo.EXPECT().Save(ctx, gomock.Any()).Return(domain.Customer{}, errors.New("error saving"))
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Customer{Name: "Invalid"})
			},
			expect:    domain.Customer{},
			expectErr: true,
		},
		{
			name: "Update success",
			mock: func() {
				customer := domain.Customer{
					Id:      1,
					Name:    "Prabowo Subianto",
					Email:   "psubianto@mamsiang.gratis",
					Phone:   "+628123456789",
					Address: "Istana Nusantara",
				}
				repo.EXPECT().Update(ctx, customer).Return(customer, nil)
			},
			method: func() (interface{}, error) {
				return repo.Update(ctx, domain.Customer{
					Id:      1,
					Name:    "Prabowo Subianto",
					Email:   "psubianto@mamsiang.gratis",
					Phone:   "+628123456789",
					Address: "Istana Nusantara",
				})
			},
			expect: domain.Customer{
				Id:      1,
				Name:    "Prabowo Subianto",
				Email:   "psubianto@mamsiang.gratis",
				Phone:   "+628123456789",
				Address: "Istana Nusantara",
			},
			expectErr: false,
		}, {
			name: "Delete success",
			mock: func() {
				repo.EXPECT().Delete(ctx, domain.Customer{Id: 1}).Return(nil)
			},
			method: func() (interface{}, error) {
				return nil, repo.Delete(ctx, domain.Customer{Id: 1})
			},
			expect:    nil,
			expectErr: false,
		}, {
			name: "Find by id success",
			mock: func() {
				repo.EXPECT().FindById(ctx, uint64(1)).Return(domain.Customer{
					Id:      1,
					Name:    "Prabowo Subianto",
					Email:   "psubianto@mamsiang.gratis",
					Phone:   "+628123456789",
					Address: "Istana Negara",
				}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 1)
			},
			expect: domain.Customer{
				Id:      1,
				Name:    "Prabowo Subianto",
				Email:   "psubianto@mamsiang.gratis",
				Phone:   "+628123456789",
				Address: "Istana Negara",
			},
			expectErr: false,
		},
		{
			name: "Find by if not found",
			mock: func() {
				repo.EXPECT().FindById(ctx, uint64(999)).Return(domain.Customer{}, errors.New("not found"))
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 999)
			},
			expect:    domain.Customer{},
			expectErr: true,
		},
		{
			name: "Find all success",
			mock: func() {
				repo.EXPECT().FindAll(ctx).Return([]domain.Customer{{
					Id:      1,
					Name:    "Prabowo Subianto",
					Email:   "psubianto@mamsiang.gratis",
					Phone:   "+628123456789",
					Address: "Istana Negara",
				}}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindAll(ctx)
			},
			expect: []domain.Customer{{
				Id:      1,
				Name:    "Prabowo Subianto",
				Email:   "psubianto@mamsiang.gratis",
				Phone:   "+628123456789",
				Address: "Istana Negara",
			}},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			result, err := tt.method()

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, result)
			}
		})
	}
}
