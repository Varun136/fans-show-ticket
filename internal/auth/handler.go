package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type authHandler struct {
	db *sql.DB
}

func NewAuthHandler(db *sql.DB) *authHandler {
	return &authHandler{db: db}
}

func (h *authHandler) LoginHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var valid_req LoginRequest
	if err := json.NewDecoder(req.Body).Decode(&valid_req); err != nil {
		http.Error(res, "Invalid request", http.StatusBadRequest)
		return
	}

	err := h.GenerateAndSendOTP(valid_req.Email)
	if err != nil {
		http.Error(res, "Unable to initiate Auth at the moment", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte("OTP send to your email"))
}

func (h *authHandler) VerifyOTPHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var valid_req VerifyOTPRequest
	if err := json.NewDecoder(req.Body).Decode(&valid_req); err != nil {
		http.Error(res, "Invalid request", http.StatusBadRequest)
		return
	}

	otp_verified, err := h.VerifyOTP(valid_req.Email, valid_req.OTP)
	if err != nil {
		http.Error(res, "Failed to verify OTP", http.StatusInternalServerError)
		return
	}
	if !otp_verified {
		http.Error(res, "Invalid or expired OTP", http.StatusUnauthorized)
		return
	}

	token, err := h.GenerateJWT(valid_req.Email)
	if err != nil {
		http.Error(res, "Unable to generate tokn at the moment", http.StatusInternalServerError)
		return
	}

	response := VerifyOTPResponse{Token: token}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)

}
