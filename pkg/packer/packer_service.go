package packer

import (
	"fmt"
	"pack-svc/pkg/repository"
	"sort"
)

type Service interface {
	AddPackSize(size int)
	RemovePackSize(size int)
	GetPackSizes() []int
	CalculatePacks(size int, packsizes []int) map[int]int
}

type PackerService struct {
	repository repository.PackSizesRepository
}

func (ps *PackerService) AddPackSize(size int) {
	ps.repository.AddPackSize(size)
}
func (ps *PackerService) RemovePackSize(size int) {
	ps.repository.RemovePackSize(size)
}
func (ps *PackerService) GetPackSizes() []int {
	return ps.repository.GetPackSizes()
}
func (ps *PackerService) CalculatePacks(orderSize int, packSizes []int) map[int]int {
	return ps.calculatePacks(orderSize, packSizes)
}

func (ph *PackerService) calculatePacks(orderSize int, availablePackSizes []int) map[int]int {
	packSizes := make([]int, len(availablePackSizes))
	copy(packSizes, availablePackSizes)

	fmt.Println(fmt.Sprintf("calculating for order size : %d and pack sizes : %v", orderSize, packSizes))
	sort.Sort(sort.IntSlice(packSizes))
	result := make(map[int]int)

	for _, size := range packSizes {
		count := 0
		for orderSize >= size {
			orderSize -= size
			count++
		}
		if count > 0 {
			result[size] = count
		}
	}

	if orderSize > 0 {
		for _, size := range packSizes {
			if size >= orderSize {
				result[size]++
				break
			}
		}
	}

	// Compress the pack count to higher variations without increasing the total items sent
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	for i := 0; i < len(packSizes)-1; i++ {
		currentSize := packSizes[i]
		for j := i + 1; j < len(packSizes); j++ {
			smallerSize := packSizes[j]
			compressibleCount := currentSize / smallerSize

			if result[smallerSize] >= compressibleCount {
				times := result[smallerSize] / compressibleCount

				result[smallerSize] -= times * compressibleCount
				result[currentSize] += times

				if result[smallerSize] == 0 {
					break
				}
			}
		}
	}

	for size, count := range result {
		if count == 0 {
			delete(result, size)
		}
	}

	return result
}

func NewPackerService(repository repository.PackSizesRepository) Service {
	return &PackerService{
		repository: repository,
	}
}
