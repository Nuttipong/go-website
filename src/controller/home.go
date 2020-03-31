package controller

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	"github.com/nuttipong/go-website/src/viewmodel"
)

type home struct {
	homeTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/bluegreen-check", h.handleBlueGreen)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHomeViewModel()
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleBlueGreen(w http.ResponseWriter, r *http.Request) {
	envs := make(map[string]string)

	envs = getEnvWithDefault("CF_INSTANCE_ADDR", "localhost", envs)
	envs = getEnvWithDefault("VCAP_APPLICATION", "unknow", envs)

	jsonString, _ := json.Marshal(envs)

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(jsonString))
	w.WriteHeader(http.StatusOK)
}

func getEnvWithDefault(key string, defaultStr string, envs map[string]string) map[string]string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultStr
	}
	envs[key] = value
	return envs
}
