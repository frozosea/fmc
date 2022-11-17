package auth

type LoginUserResponse struct {
	Token               string `json:"token,omitempty" binding:"required"`
	TokenType           string `json:"token_type,omitempty" binding:"required"`
	RefreshToken        string `json:"refreshToken,omitempty" binding:"required"`
	TokenExpires        int64  `json:"tokenExpires,omitempty" binding:"required"`
	RefreshTokenExpires int64  `json:"refreshTokenExpires,omitempty" binding:"required"`
}

type BaseResponse struct {
	Success bool   `json:"success" binding:"required"`
	Error   string `json:"error" binding:"required"`
}
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
type SendRecoveryEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

type RecoveryUserRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}
