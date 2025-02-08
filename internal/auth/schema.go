package auth

type LoginRequest struct {
	Email string `json:"email"`
}

type VerifyOTPRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type VerifyOTPResponse struct {
	Token string `json:"token"`
}
