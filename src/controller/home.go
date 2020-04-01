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
	// http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello Blue-green"))
	})
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/bluegreen-check", h.handleBlueGreen)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHomeViewModel()

	vcap := map[string]string{}
	vservice := map[string]string{}

	j1 := []byte(os.Getenv("VCAP_APPLICATION"))
	j2 := []byte(os.Getenv("VCAP_SERVICES"))

	_ = json.Unmarshal(j1, &vcap)
	_ = json.Unmarshal(j2, &vservice)

	vm.ApplicationId = vcap["application_id"]
	vm.ApplicationName = vcap["application_name"]

	for k, _ := range vservice {
		svr := viewmodel.Service{Name: k}
		vm.Services = append(vm.Services, svr)
	}

	h.homeTemplate.Execute(w, vm)
}

func (h home) handleBlueGreen(w http.ResponseWriter, r *http.Request) {
	envs := make(map[string]string)
	envs = getEnvWithDefault("CF_INSTANCE_ADDR", "localhost", envs)
	envs = getEnvWithDefault("VCAP_APPLICATION", "{\"application_id\":\"-\",\"application_name\":\"localhost\"}", envs)

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
