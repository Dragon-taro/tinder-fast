package types

type User struct {
	User struct {
		ID       string `json:"_id"`
		Name     string `json:"name"`
		APIToken string `json:"api_token"`
	} `json:"data"`
}

type ResultUser struct {
	ID string `json:"_id"`
}

type Users struct {
	Users []ResultUser `json:"results"`
}
