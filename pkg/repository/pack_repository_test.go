package repository_test

import (
	"pack-svc/pkg/repository"
	"reflect"
	"testing"
)

func TestPackSizesRepository(t *testing.T) {
	tests := []struct {
		name            string
		addPackSizes    []int
		removePackSizes []int
		expectedSizes   []int
	}{
		{
			name:            "Adding unique sizes",
			addPackSizes:    []int{250, 500},
			removePackSizes: []int{},
			expectedSizes:   []int{250, 500},
		},
		{
			name:            "Adding with duplicates",
			addPackSizes:    []int{1000, 1000},
			removePackSizes: []int{},
			expectedSizes:   []int{1000},
		},
		{
			name:            "Removing existing size",
			addPackSizes:    []int{250, 500, 1000},
			removePackSizes: []int{500},
			expectedSizes:   []int{250, 1000},
		},
		{
			name:            "Removing non-existent size",
			addPackSizes:    []int{250, 1000},
			removePackSizes: []int{500},
			expectedSizes:   []int{250, 1000},
		},
		{
			name:            "Multiple operations",
			addPackSizes:    []int{250, 500, 1000, 2000, 500},
			removePackSizes: []int{1000, 250},
			expectedSizes:   []int{500, 2000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewPackRepository()

			for _, size := range tt.addPackSizes {
				repo.AddPackSize(size)
			}

			for _, size := range tt.removePackSizes {
				repo.RemovePackSize(size)
			}

			actualSizes := repo.GetPackSizes()
			if !reflect.DeepEqual(actualSizes, tt.expectedSizes) {
				t.Errorf("After %s, expected pack sizes %v, got %v", tt.name, tt.expectedSizes, actualSizes)
			}
		})
	}
}
