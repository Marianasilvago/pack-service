package repository

import "sync"

//go:generate mockery --name=PackSizesRepository --filename=pack_repository.go
type PackSizesRepository interface {
	AddPackSize(size int)
	RemovePackSize(size int)
	GetPackSizes() []int
}

type packSizesRepository struct {
	sync.RWMutex
	Sizes []int
}

func NewPackRepository() PackSizesRepository {
	return &packSizesRepository{
		Sizes: []int{},
	}
}

// addPackSize adds a new pack size, ensuring no duplicates
func (p *packSizesRepository) AddPackSize(size int) {
	p.Lock()
	defer p.Unlock()
	for _, s := range p.Sizes {
		if s == size {
			return // Avoid adding duplicates
		}
	}
	p.Sizes = append(p.Sizes, size)
}

// removePackSize removes a pack size if it exists
func (p *packSizesRepository) RemovePackSize(size int) {
	p.Lock()
	defer p.Unlock()
	for i, s := range p.Sizes {
		if s == size {
			p.Sizes = append(p.Sizes[:i], p.Sizes[i+1:]...)
			break
		}
	}
}

// getPackSizes returns the current list of pack sizes
func (p *packSizesRepository) GetPackSizes() []int {
	p.RLock()
	defer p.RUnlock()
	return p.Sizes
}
