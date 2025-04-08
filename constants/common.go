package constants

type Item struct {
	Title    string `json:"title"`
	Price    string `json:"price"`
	Rating   string `json:"rating"`
	Reviews  string `json:"reviews"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
}

type Resp struct {
	Items []Item `json:"items"`
}

type User struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
