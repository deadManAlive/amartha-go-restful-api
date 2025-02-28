package repository

import (
	"context"
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
					Name:    "Budi",
					Email:   "budi@pekerti.com",
					Phone:   "62812345678910",
					Address: "Sumatra",
				}
				repo.EXPECT().Save(ctx, domain.Customer{
					Id:      1,
					Name:    "Budi",
					Email:   "budi@pekerti.com",
					Phone:   "62812345678910",
					Address: "Sumatra",
				}).Return(category, nil)
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Customer{
					Id:      1,
					Name:    "Budi",
					Email:   "budi@pekerti.com",
					Phone:   "62812345678910",
					Address: "Sumatra",
				})
			},
			expect: domain.Customer{
				Id:      1,
				Name:    "Budi",
				Email:   "budi@pekerti.com",
				Phone:   "62812345678910",
				Address: "Sumatra",
			},
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
