package packer

import (
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
func (ps *PackerService) calculatePacks(orderSize int, availablePackSizes []int) map[int]int {
	if len(availablePackSizes) == 0 || orderSize <= 0 {
		return nil
	}

	packSizes := make([]int, len(availablePackSizes))
	copy(packSizes, availablePackSizes)

	// Phase 1: Attempt to find an exact fit
	exactFit, exactFitFound := findExactFit(orderSize, packSizes)
	if exactFitFound {
		return exactFit
	}

	// Phase 2: Minimize the total number of items if exact fit is not found
	minimumItemsCombination := minimizeItems(orderSize, packSizes)

	return reducePacksCount(packSizes, minimumItemsCombination)
}

func findExactFit(orderSize int, availablePackSizes []int) (map[int]int, bool) {
	packSizes := make([]int, len(availablePackSizes))
	copy(packSizes, availablePackSizes)

	// Sort pack sizes in ascending order to start with the smallest packs
	sort.Ints(packSizes)

	// Initialize the result map to store the final pack selection
	var bestCombination map[int]int
	var bestTotalItems int = int(^uint(0) >> 1) // Set to maximum int initially

	// Recursive function to find the best combination of packs
	var findCombination func(int, map[int]int, int)
	findCombination = func(remainingOrderSize int, currentCombination map[int]int, startIndex int) {
		if remainingOrderSize == 0 {
			totalItems := 0
			for size, count := range currentCombination {
				totalItems += size * count
			}
			if totalItems < bestTotalItems {
				bestTotalItems = totalItems
				bestCombination = copyMap(currentCombination)
			}
			return
		}

		for i := startIndex; i < len(packSizes); i++ {
			size := packSizes[i]
			if size <= remainingOrderSize {
				newCombination := copyMap(currentCombination)
				newCombination[size]++
				findCombination(remainingOrderSize-size, newCombination, i)
			}
		}
	}

	findCombination(orderSize, make(map[int]int), 0)

	return bestCombination, getTotalItemCount(bestCombination) == orderSize
}

func getTotalItemCount(selectedCombination map[int]int) int {
	sum := 0
	for k, v := range selectedCombination {
		sum += k * v
	}
	return sum
}

func minimizeItems(orderSize int, availablePackSizes []int) map[int]int {
	packSizes := make([]int, len(availablePackSizes))
	copy(packSizes, availablePackSizes)

	// Sort pack sizes in ascending order to facilitate the selection process
	sort.Ints(packSizes)

	// Result map to hold the number of packs of each size
	result := make(map[int]int)

	// Iterate over the pack sizes to find the best fit for the order size
	for i := len(packSizes) - 1; i >= 0; i-- {
		packSize := packSizes[i]

		// Check if the current pack size is smaller than the remaining order size
		if orderSize >= packSize {
			// Calculate how many of this pack size are needed
			count := orderSize / packSize
			result[packSize] = count

			// Decrease the order size by the total number of items in the selected packs
			orderSize -= count * packSize
		}

		// If the order size is exactly met, break the loop
		if orderSize == 0 {
			break
		}
	}

	// If there is a remaining order size, use smaller packs to fulfill it
	if orderSize > 0 {
		for _, packSize := range packSizes {
			if orderSize <= packSize {
				result[packSize]++
				break
			}
		}
	}

	return result
}

func copyMap(original map[int]int) map[int]int {
	newMap := make(map[int]int)
	for k, v := range original {
		newMap[k] = v
	}
	return newMap
}

func reducePacksCount(availablePackSizes []int, combination map[int]int) map[int]int {
	packSizes := make([]int, len(availablePackSizes))
	copy(packSizes, availablePackSizes)

	sort.Ints(packSizes)

	for i := len(packSizes) - 1; i >= 0; i-- {
		largerPackSize := packSizes[i]

		for j := 0; j < i; j++ {
			smallerPackSize := packSizes[j]
			combineCount := largerPackSize / smallerPackSize

			// Check if we can replace smaller packs with a larger one
			if combination[smallerPackSize] >= combineCount {
				// Calculate how many larger packs we can use instead of smaller ones
				numReplacements := combination[smallerPackSize] / combineCount

				// Update the pack count in the combination
				combination[smallerPackSize] -= numReplacements * combineCount
				if combination[largerPackSize] == 0 {
					combination[largerPackSize] = numReplacements
				} else {
					combination[largerPackSize] += numReplacements
				}

				// If after replacement the count of smaller packs becomes zero, remove it from the map
				if combination[smallerPackSize] == 0 {
					delete(combination, smallerPackSize)
				}
			}
		}
	}

	return combination
}
func NewPackerService(repository repository.PackSizesRepository) Service {
	return &PackerService{
		repository: repository,
	}
}
