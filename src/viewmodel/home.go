package viewmodel

type HomeViewModel struct {
	Title           string
	ApplicationId   string
	ApplicationName string
	Services        []Service
}

type Service struct {
	Name string
}

func NewHomeViewModel() *HomeViewModel {
	return &HomeViewModel{Title: "Blue-green Deployment"}
}
