package service

import (
	"log"
	"time"

	"github.com/thinhcompany/ecommerce-ver-2/internal/repo"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/response"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/utils/crypto"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/utils/random"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/utils/redisutil"
	sendto "github.com/thinhcompany/ecommerce-ver-2/pkg/utils/send_to"
)

const otpExpiration = 5 * time.Minute // you can change this duration

type IUserService interface {
	Register(email string, purpose string) response.ResponseData
	CheckUserExists(email string) response.ResponseData
}

type userService struct {
	userRepo     repo.IUserRepo
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(
	userRepo repo.IUserRepo,
	userAuthRepo repo.IUserAuthRepository) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

func (s *userService) CheckUserExists(email string) response.ResponseData {
	exists := s.userRepo.GetUserByEmail(email)
	if exists {
		log.Printf("User already exists: %s", email)
		return response.SuccessResponse("User exists")
	}
	return response.ErrorResponse(response.ErrorCodeNotFound, nil)
}

func (s *userService) Register(email string, purpose string) response.ResponseData {
	// 0. Hash email
	hashedEmail := crypto.HashSHA256(email)
	log.Println("Email has been hash:", hashedEmail)

	// 1. Check if OTP is still valid in Redis
	//c623eeaccf18df2ba50d855a138ede0a19c6844d48cdd263152cff6f78c2c36e
	valid, errOTP := redisutil.IsOtpStillValid(hashedEmail)
	if errOTP != nil {
		log.Printf("Redis EXISTS error: %v", errOTP)
		return response.ErrorResponse(response.ErrorCodeRedisError, nil)
	}
	if valid {
		log.Println("OTP still valid. Skipping re-send.")
		return response.ErrorResponse(response.ErrorCodeOtpStillValid, nil)
	}

	// 2. Optional: add rate limiting logic if needed

	// 3. Check if user already exists
	if s.userRepo.GetUserByEmail(email) {
		log.Printf("User already exists: %s", email)
		return response.ErrorResponse(response.ErrorCodeUserHasExists, nil)
	}

	// 4. Generate OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = "999999"
	}
	log.Println("Generated OTP:", otp)
	log.Printf("Generated OTP: ****%s", otp[len(otp)-2:])

	// 5. Save OTP to Redis
	saveErr := s.userAuthRepo.AddOTP(hashedEmail, otp, otpExpiration)
	if saveErr != nil {
		log.Printf("Error saving OTP to Redis: %v", saveErr)
		return response.ErrorResponse(response.ErrorCodeRedisError, nil)
	}

	// 6. TODO: Send OTP to email or SMS (not implemented yet)
	err := sendto.SendTemplateOtp(
		[]string{email},
		"thinhproee@gmail.com",
		"otp.html",
		map[string]any{
			"OTP":  otp,
			"Year": time.Now().Year(),
		},
	)
	if err != nil {
		log.Println("Failed to send email:", err)
		return response.ErrorResponse(response.ErrorCodeEmailSend, nil)
	}
	log.Println("OTP sent successfully to:", email)
	return response.SuccessResponse(map[string]string{
		"email": email,
		"otp":   otp, // ⚠️ only return for testing; remove in production
	})
}
