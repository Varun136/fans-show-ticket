package auth

import (
	"fmt"
	"math/rand"
	"time"
)

func (h *authHandler) GenerateAndSendOTP(email string) error {
	otp, err := h.generateOTP(email)
	if err != nil {
		return err
	}
	err2 := h.sendOTP(email, otp)
	if err2 != nil {
		return err
	}

	return nil
}

func (h *authHandler) generateOTP(email string) (string, error) {

	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	expiresAt := time.Now().Add(5 * time.Minute)

	query := "INSERT INTO otps (email, otp_code, expires_at) VALUES ($1, $2, $3)"
	_, err := h.db.Exec(query, email, otp, expiresAt)
	if err != nil {
		return "", err
	}

	return otp, nil
}

func (h *authHandler) sendOTP(email string, otp string) error {
	fmt.Printf("Send email to %v: otp %v", email, otp)
	return nil
}

func (h *authHandler) VerifyOTP(email string, otp string) (bool, error) {
	query := "SELECT COUNT(*) FROM otps WHERE email=$1 AND otp_code=$2 AND expires_at > NOW()"
	var count int
	err := h.db.QueryRow(query, email, otp).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (h *authHandler) CreateUser(email string) error {
	query := "INSERT INTO users(email) values ($1)"
	_, err := h.db.Exec(query, email)
	if err != nil {
		return err
	}
	return nil
}
