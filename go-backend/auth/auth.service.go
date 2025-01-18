package auth

import (
	"fmt"
	"go-todo-app/user"
	"go-todo-app/utils/envutil"
	"go-todo-app/utils/tokenutil"
	"log"
	"strconv"
)

type (
	config struct {
		secret   string
		expiraty int
	}
	service struct {
		config      *config
		userService *user.Service
	}
)

func genClaim(payload *TokenPayload) *map[string]any {
	return &map[string]interface{}{
		"payload": payload,
		"https://hasura.io/jwt/claims": map[string]any{
			"x-hasura-default-role": "user",
			"x-hasura-user-id":      fmt.Sprintf("%d", payload.Id),
		},
	}
}

func (s *service) login(request LoginRequest) (*LoginResponse, error) {
	fmt.Println("login called")

	user, err := s.userService.FindByUsername(request.UserName)
	if err != nil {
		return nil, err
	}

	err = user.CheckPassword(request.PassWord)
	if err != nil {
		return nil, &Errors.IncorrectPassword
	}

	claims := *genClaim(&TokenPayload{user.ID})
	accessToken, err := tokenutil.CreateToken(
		claims,
		s.config.secret,
		s.config.expiraty,
	)
	return &LoginResponse{AccessToken: accessToken}, err
}

func (s *service) verifyToken(token string) (*TokenPayload, error) {
	rawPayload, err := tokenutil.VerifyToken(token, s.config.secret)
	if err != nil {
		return nil, &Errors.InvalidToken
	}

	payload, err := toPayload(rawPayload)
	if err != nil {
		return nil, &Errors.InvalidToken
	}

	return payload, nil
}

func loadConfig() *config {
	secret := envutil.GetEvnVariable("SECRET")
	expiratyString := envutil.GetEvnVariable("TOKEN_EXPIRATION")

	expiraty, err := strconv.Atoi(expiratyString)
	if err != nil {
		log.Fatalf("Invalid TOKEN_EXPIRATION value: %v", err)
	}

	return &config{
		secret, expiraty,
	}
}

func NewService() *service {
	userService := user.NewService()
	config := loadConfig()
	service := service{config, userService}

	return &service
}
