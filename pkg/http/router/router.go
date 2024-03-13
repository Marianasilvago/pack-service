package router

import (
	"net/http"
	"pack-svc/pkg/http/internal/handler"
	"pack-svc/pkg/http/internal/middleware"
	"pack-svc/pkg/packer"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewRouter(lgr *zap.Logger, packService packer.Service) http.Handler {
	router := mux.NewRouter()
	router.Use(handlers.RecoveryHandler())

	packHandler := handler.NewPackHandler(lgr, packService)

	router.HandleFunc("/", serveIndexHTML).Methods(http.MethodGet)
	router.HandleFunc("/pack-sizes", withMiddlewares(lgr, middleware.WithErrorHandler(lgr, packHandler.HandleAddPackSize))).Methods("POST")
	router.HandleFunc("/pack-sizes", withMiddlewares(lgr, middleware.WithErrorHandler(lgr, packHandler.HandleRemovePackSize))).Methods("DELETE")
	router.HandleFunc("/pack-sizes", withMiddlewares(lgr, middleware.WithErrorHandler(lgr, packHandler.HandleGetPackSizes))).Methods("GET")

	router.HandleFunc("/calculate-packs", withMiddlewares(lgr, middleware.WithErrorHandler(lgr, packHandler.HandleCalculatePacks))).Methods("POST")

	return router
}

func withMiddlewares(lgr *zap.Logger, hnd http.HandlerFunc) http.HandlerFunc {
	return middleware.WithSecurityHeaders(middleware.WithReqResLog(lgr, hnd))
}

func serveIndexHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./frontend/index.html")
}
