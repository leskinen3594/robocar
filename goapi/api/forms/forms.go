package forms

// Receiver ; get key from POST method
type APIkeyReceiver struct {
	// Check api key
	Username string `json:"uname"`
	APIkey   string `json:"api_key"`
}

type ControlReceiver struct {
	// Control robot
	Username string `json:"uname"`
	Movement string `json:"movement"`
}
