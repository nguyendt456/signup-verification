package signup

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	pb "github.com/nguyendt456/signup-with-verification/pb"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const (
	verifyExpireDuration = time.Minute * 2
)

// Signup(context.Context, *User) (*SignupReponse, error)
type SignupAPI struct {
	Email    *EmailSender
	Database *gorm.DB
	Redis    *redis.Client
	pb.SignupServiceServer
}

func (s *SignupAPI) Signup(ctx context.Context, user *pb.User) (*pb.SignupReponse, error) {
	// err := ValidateUsername(user.Email)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return &pb.SignupReponse{State: "inv", User: nil}, err
	// }

	secretCode := uuid.New().String()
	res := s.Redis.Set(ctx, secretCode, user.Email, verifyExpireDuration)
	if res.Err() != nil {
		fmt.Println(res.Err().Error())
		return &pb.SignupReponse{State: "inv", User: nil}, res.Err()
	}

	u := pb.User{}
	c := s.Database.First(&u, pb.User{Email: user.Email})
	if c.Error == nil {
		user.Password = ""
		if u.IsVerified == false {
			go s.Email.SendEmailVerify(VerifyEmail{
				Email:      user.Email,
				SecretCode: secretCode,
			})
			return &pb.SignupReponse{State: "rever", User: user}, nil
		}
		return &pb.SignupReponse{State: "valid", User: user}, nil
	}

	s.Database.Create(&pb.User{
		Email:      user.Email,
		Name:       user.Name,
		Password:   user.Password,
		IsVerified: false,
	})

	go s.Email.SendEmailVerify(VerifyEmail{
		Email:      user.Email,
		SecretCode: secretCode,
	})
	user.Password = ""
	return &pb.SignupReponse{State: "verif", User: user}, nil
}

func (s *SignupAPI) Verify(ctx context.Context, v *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	email, isVerified := ValidateVerifyCode(ctx, s.Redis, v.VerificationCode)
	if !isVerified {
		return &pb.VerifyResponse{IsVerified: false}, nil
	}
	s.Database.Model(&pb.User{}).Where("email = ?", email).Update("is_verified", true)
	return &pb.VerifyResponse{IsVerified: true}, nil
}
