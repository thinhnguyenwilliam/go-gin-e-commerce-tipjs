package service

import (
	"log"
	"time"

	"github.com/thinhcompany/ecommerce-ver-2/internal/repo"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/response"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/utils/crypto"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/utils/random"
	sendto "github.com/thinhcompany/ecommerce-ver-2/pkg/utils/send_to"
)

const otpExpiration = 5 * time.Minute // you can change this duration

type IUserService interface {
	Register(email string, purpose string) int
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

func (s *userService) Register(email string, purpose string) int {
	// 0. Hash email
	hashedEmail := crypto.HashSHA256(email)
	log.Println("Email has been hash:", hashedEmail)

	// 1. Check if OTP is still valid in Redis
	// otpKey := fmt.Sprintf("otp:%s:%s", purpose, hashedEmail)
	// exists, err := global.Rdb.Exists(context.Background(), otpKey).Result()
	// if err != nil {
	// 	log.Printf("Redis EXISTS error: %v", err)
	// 	return response.ErrorCodeRedisError // define this constant
	// }
	// if exists > 0 {
	// 	return response.ErrorCodeOtpStillValid // define this constant
	// }

	// 2. Optional: add rate limiting logic if needed

	// 3. Check if user already exists
	if s.userRepo.GetUserByEmail(email) {
		return response.ErrorCodeUserHasExists
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
		return response.ErrorCodeRedisError
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
		return response.ErrorCodeEmailSend
	}

	return response.ErrorCodeSuccess
}
