package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user model
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	Password  string         `json:"-" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// RefreshToken represents refresh token for JWT
type RefreshToken struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Token     string    `json:"token" gorm:"not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"-" gorm:"foreignKey:UserID"`
}

// Holdings represents user's stock holdings
type Holdings struct {
	Symbol       string  `json:"symbol"`
	Quantity     int     `json:"quantity"`
	AveragePrice float64 `json:"average_price"`
	CurrentPrice float64 `json:"current_price"`
	PNL          float64 `json:"pnl"`
	PNLPercent   float64 `json:"pnl_percent"`
}

// Order represents order data
type Order struct {
	ID           string     `json:"id"`
	Symbol       string     `json:"symbol"`
	OrderType    string     `json:"order_type"` // BUY or SELL
	Quantity     int        `json:"quantity"`
	Price        float64    `json:"price"`
	Status       string     `json:"status"`
	OrderTime    time.Time  `json:"order_time"`
	ExecutedTime *time.Time `json:"executed_time,omitempty"`
}

// Position represents user's positions
type Position struct {
	Symbol               string  `json:"symbol"`
	Quantity             int     `json:"quantity"`
	AveragePrice         float64 `json:"average_price"`
	CurrentPrice         float64 `json:"current_price"`
	UnrealizedPNL        float64 `json:"unrealized_pnl"`
	UnrealizedPNLPercent float64 `json:"unrealized_pnl_percent"`
	PositionType         string  `json:"position_type"` // LONG or SHORT
}

// PNLCard represents PNL summary
type PNLCard struct {
	TotalPNL        float64 `json:"total_pnl"`
	TotalPNLPercent float64 `json:"total_pnl_percent"`
	DayPNL          float64 `json:"day_pnl"`
	DayPNLPercent   float64 `json:"day_pnl_percent"`
	RealizedPNL     float64 `json:"realized_pnl"`
	UnrealizedPNL   float64 `json:"unrealized_pnl"`
}

// API Request/Response structures
type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type HoldingsResponse struct {
	Holdings []Holdings `json:"holdings"`
	PNLCard  PNLCard    `json:"pnl_card"`
}

type OrderbookResponse struct {
	Orders  []Order `json:"orders"`
	PNLCard PNLCard `json:"pnl_card"`
}

type PositionsResponse struct {
	Positions []Position `json:"positions"`
	PNLCard   PNLCard    `json:"pnl_card"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
