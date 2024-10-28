package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/akptools/go-services-template/internal/util"
)

func Healthcheck(w http.ResponseWriter, _ *http.Request) error {
	healthCheckResponse := map[string]string{
		"pid":    strconv.Itoa(os.Getpid()),
		"status": "success",
		"env":    os.Getenv("APP_ENV"),
	}
	return util.WriteJSON(w, http.StatusOK, healthCheckResponse)
}
