package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/cocoasterr/net_http/services"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}
func (c *AuthController) createToken(username string, expiration time.Time) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}
	secretKey := os.Getenv("SECRET_KEY")
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expiration.Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func (c *AuthController) createRefreshToken(username string) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}
	secretKey := os.Getenv("SECRET_KEY")
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed!", 400)
		return
	}
	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid Request Body!", 500)
		return
	}
	if err := c.AuthService.PayloadRegisterCheck(payload); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	payload["password"] = c.AuthService.HashPassword(payload["password"].(string))
	ctx := r.Context()
	if err := c.AuthService.CreateService(ctx, payload); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	resp := map[string]interface{}{"Message": "Success!"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed!", 400)
		return
	}
	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid Request Body!", 500)
		return
	}
	ctx := r.Context()
	user, err := c.AuthService.FindByService(ctx, "email", payload["email"])
	if err != nil {
		http.Error(w, "Email not Found!", 404)
		return
	}
	userResp := user[0]
	checkPass := c.AuthService.CheckPasswordHash(payload["password"].(string), userResp["password"].(string))
	if !checkPass {
		if err != nil {
			http.Error(w, "Wrong Password!", 400)
			return
		}
		return
	}
	username := userResp["username"].(string)

	accessToken, err := c.createToken(username, time.Now().Add(time.Minute*15))
	if err != nil {
		http.Error(w, "Failed to create access token", http.StatusInternalServerError)
		return
	}
	refreshToken, err := c.createRefreshToken(username)
	if err != nil {
		http.Error(w, "Failed to create refresh token", http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{"Access Token": accessToken, "Refresh Token": refreshToken}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
