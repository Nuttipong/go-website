package viewmodel

type HomeViewModel struct {
	Title string
}

func NewHomeViewModel() *HomeViewModel {
	return &HomeViewModel{Title: "Blue-green Deployment"}
}
