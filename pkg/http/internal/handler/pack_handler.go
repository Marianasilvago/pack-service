package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"pack-svc/pkg/http/internal/utils"
	"pack-svc/pkg/packer"
)

type PackHandler struct {
	lgr *zap.Logger
	svc packer.Service
}

func NewPackHandler(lgr *zap.Logger, svc packer.Service) *PackHandler {
	return &PackHandler{
		lgr: lgr,
		svc: svc,
	}
}
func (ph *PackHandler) HandleAddPackSize(w http.ResponseWriter, r *http.Request) error {
	var request struct {
		Size int `json:"size"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return err
	}

	ph.svc.AddPackSize(request.Size)
	utils.WriteSuccessResponse(w, http.StatusNoContent, nil)
	return nil
}

func (ph *PackHandler) HandleRemovePackSize(w http.ResponseWriter, r *http.Request) error {
	var request struct {
		Size int `json:"size"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return err
	}

	ph.svc.RemovePackSize(request.Size)
	utils.WriteSuccessResponse(w, http.StatusNoContent, nil)
	return nil
}

func (ph *PackHandler) HandleGetPackSizes(w http.ResponseWriter, r *http.Request) error {
	sizes := ph.svc.GetPackSizes()
	utils.WriteSuccessResponse(w, http.StatusOK, sizes)
	return nil
}

func (ph *PackHandler) HandleCalculatePacks(w http.ResponseWriter, r *http.Request) error {
	var request struct {
		OrderSize int `json:"orderSize"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return err
	}
	packSizes := ph.svc.GetPackSizes()
	packs := ph.svc.CalculatePacks(request.OrderSize, packSizes)

	utils.WriteSuccessResponse(w, http.StatusOK, packs)
	return nil
}
