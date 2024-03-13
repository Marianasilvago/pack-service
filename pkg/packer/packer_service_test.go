package packer_test

import (
	"pack-svc/pkg/packer"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"pack-svc/pkg/repository/mocks"
)

func TestAddPackSize(t *testing.T) {
	mockRepo := new(mocks.PackSizesRepository)
	mockRepo.On("AddPackSize", mock.AnythingOfType("int")).Once()

	service := packer.NewPackerService(mockRepo)
	service.AddPackSize(100)

	mockRepo.AssertExpectations(t)
}

func TestRemovePackSize(t *testing.T) {
	mockRepo := new(mocks.PackSizesRepository)
	mockRepo.On("RemovePackSize", mock.AnythingOfType("int")).Once()

	service := packer.NewPackerService(mockRepo)
	service.RemovePackSize(100)

	mockRepo.AssertExpectations(t)
}

func TestGetPackSizes(t *testing.T) {
	expected := []int{100, 200}
	mockRepo := new(mocks.PackSizesRepository)
	mockRepo.On("GetPackSizes").Return(expected)

	service := packer.NewPackerService(mockRepo)
	result := service.GetPackSizes()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	mockRepo.AssertExpectations(t)
}

func TestCalculatePacks(t *testing.T) {
	tests := []struct {
		name          string
		orderSize     int
		packSizes     []int
		expectedPacks map[int]int
	}{
		{
			name:          "Order 1 item",
			orderSize:     1,
			packSizes:     []int{250, 1000, 2000, 500, 5000},
			expectedPacks: map[int]int{250: 1},
		},
		{
			name:          "Order 250 items",
			orderSize:     250,
			packSizes:     []int{250, 2000, 500, 1000, 5000},
			expectedPacks: map[int]int{250: 1},
		},
		{
			name:          "Order 251 items",
			orderSize:     251,
			packSizes:     []int{250, 500, 1000, 2000, 5000},
			expectedPacks: map[int]int{500: 1},
		},
		{
			name:          "Order 501 items",
			orderSize:     501,
			packSizes:     []int{250, 500, 5000, 1000, 2000},
			expectedPacks: map[int]int{500: 1, 250: 1},
		},
		{
			name:          "Order 12001 items",
			orderSize:     12001,
			packSizes:     []int{250, 500, 1000, 2000, 5000},
			expectedPacks: map[int]int{5000: 2, 2000: 1, 250: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := packer.NewPackerService(nil)

			result := svc.CalculatePacks(tt.orderSize, tt.packSizes)

			if !reflect.DeepEqual(result, tt.expectedPacks) {
				t.Errorf("CalculatePacks(%d, %v) got %v, want %v", tt.orderSize, tt.packSizes, result, tt.expectedPacks)
			}
		})
	}
}
