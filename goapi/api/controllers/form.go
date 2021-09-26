package controllers

// Receiver ; get key from POST method
type PostReceiver struct {
	// Check api key
	Username string `json:"uname"`
	APIkey   string `json:"api_key"`

	// Control robot
	Movement string `json:"movement"`
}
