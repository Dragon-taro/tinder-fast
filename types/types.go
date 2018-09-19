package types

type User struct {
	User struct {
		ID       string `json:"_id"`
		Name     string `json:"name"`
		APIToken string `json:"api_token"`
	} `json:"user"`
}

type ResultUser struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type Users struct {
	Users []ResultUser `json:"results"`
}
