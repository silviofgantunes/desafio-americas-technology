// ./order-service/models/model.go

package models

import "time"

// Direction representa a direção da ordem
type Direction string

const (
	DirectionBuy  Direction = "buy"
	DirectionSell Direction = "sell"
)

// OrderType representa o tipo de ordem
type OrderType string

const (
	OrderTypeMarket OrderType = "market"
	OrderTypeLimit  OrderType = "limit"
)

// Pair representa o par de moedas
type Pair string

const (
	PairUSDBTC Pair = "USD/BTC"
	PairUSDADA Pair = "USD/ADA"
	PairUSDETH Pair = "USD/ETH"
	PairBTCUSD Pair = "BTC/USD"
	PairETHUSD Pair = "ETH/USD"
	PairADAUSD Pair = "ADA/USD"

	// Adicione outros pares conforme necessário
)

type Order struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	UserID    string    `gorm:"type:char(36);not null" json:"user_id,omitempty"`
	UserName  string    `gorm:"type:varchar(255);not null" json:"user_name,omitempty"`
	Pair      Pair      `gorm:"type:varchar(255);not null" json:"pair,omitempty"`
	Amount    float64   `gorm:"not null" json:"amount,omitempty"`
	Direction Direction `gorm:"type:varchar(255);not null" json:"direction,omitempty"`
	Type      OrderType `gorm:"type:varchar(255);not null" json:"type,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

// CreateOrderRequest representa a estrutura da solicitação para criar uma nova ordem
type CreateOrderRequest struct {
	UserID    string    `json:"user_id" binding:"required"`
	UserName  string    `json:"user_name"`
	Pair      Pair      `json:"pair" binding:"required"`
	Amount    float64   `json:"amount" binding:"required"`
	Direction Direction `json:"direction" binding:"required"`
	Type      OrderType `json:"type" binding:"required"`
}
