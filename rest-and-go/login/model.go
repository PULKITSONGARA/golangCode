package login

type Person struct {
	Email_id string `json:"email_id"`
	Password string `json:"password"`
}

type Response struct {
	Status_code int    `json:"status_code"`
	Message     string `json:"message"`
	Data        string `json:"data"`
}
