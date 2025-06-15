// services/data_service.go
package services

import (
	"math/rand"
	"time"
	"trading-platform-backend/models"
)

type DataService struct{}

func NewDataService() *DataService {
	return &DataService{}
}

// GetHoldings returns mock holdings data
func (s *DataService) GetHoldings(userID uint) *models.HoldingsResponse {
	holdings := []models.Holdings{
		{
			Symbol:       "RELIANCE",
			Quantity:     10,
			AveragePrice: 2450.50,
			CurrentPrice: 2485.20,
			PNL:          347.00,
			PNLPercent:   1.42,
		},
		{
			Symbol:       "TCS",
			Quantity:     5,
			AveragePrice: 3820.75,
			CurrentPrice: 3795.30,
			PNL:          -127.25,
			PNLPercent:   -0.67,
		},
		{
			Symbol:       "HDFCBANK",
			Quantity:     8,
			AveragePrice: 1675.40,
			CurrentPrice: 1702.80,
			PNL:          219.20,
			PNLPercent:   1.64,
		},
		{
			Symbol:       "INFY",
			Quantity:     15,
			AveragePrice: 1834.60,
			CurrentPrice: 1856.90,
			PNL:          334.50,
			PNLPercent:   1.22,
		},
		{
			Symbol:       "HINDUNILVR",
			Quantity:     6,
			AveragePrice: 2720.30,
			CurrentPrice: 2698.45,
			PNL:          -131.10,
			PNLPercent:   -0.80,
		},
	}

	pnlCard := models.PNLCard{
		TotalPNL:        642.35,
		TotalPNLPercent: 0.96,
		DayPNL:          234.50,
		DayPNLPercent:   0.35,
		RealizedPNL:     1250.75,
		UnrealizedPNL:   642.35,
	}

	return &models.HoldingsResponse{
		Holdings: holdings,
		PNLCard:  pnlCard,
	}
}

// GetOrderbook returns mock orderbook data
func (s *DataService) GetOrderbook(userID uint) *models.OrderbookResponse {
	orders := []models.Order{
		{
			ID:        "ORD001",
			Symbol:    "RELIANCE",
			OrderType: "BUY",
			Quantity:  10,
			Price:     2450.50,
			Status:    "COMPLETED",
			OrderTime: time.Now().Add(-2 * time.Hour),
			ExecutedTime: func() *time.Time {
				t := time.Now().Add(-2 * time.Hour).Add(5 * time.Minute)
				return &t
			}(),
		},
		{
			ID:        "ORD002",
			Symbol:    "TCS",
			OrderType: "SELL",
			Quantity:  3,
			Price:     3825.00,
			Status:    "COMPLETED",
			OrderTime: time.Now().Add(-1 * time.Hour),
			ExecutedTime: func() *time.Time {
				t := time.Now().Add(-1 * time.Hour).Add(2 * time.Minute)
				return &t
			}(),
		},
		{
			ID:        "ORD003",
			Symbol:    "HDFCBANK",
			OrderType: "BUY",
			Quantity:  5,
			Price:     1680.25,
			Status:    "PENDING",
			OrderTime: time.Now().Add(-30 * time.Minute),
		},
		{
			ID:        "ORD004",
			Symbol:    "INFY",
			OrderType: "BUY",
			Quantity:  8,
			Price:     1840.00,
			Status:    "CANCELLED",
			OrderTime: time.Now().Add(-45 * time.Minute),
		},
		{
			ID:        "ORD005",
			Symbol:    "ITC",
			OrderType: "SELL",
			Quantity:  12,
			Price:     415.75,
			Status:    "COMPLETED",
			OrderTime: time.Now().Add(-3 * time.Hour),
			ExecutedTime: func() *time.Time {
				t := time.Now().Add(-3 * time.Hour).Add(1 * time.Minute)
				return &t
			}(),
		},
	}

	pnlCard := models.PNLCard{
		TotalPNL:        1893.10,
		TotalPNLPercent: 2.84,
		DayPNL:          567.25,
		DayPNLPercent:   0.85,
		RealizedPNL:     1250.75,
		UnrealizedPNL:   642.35,
	}

	return &models.OrderbookResponse{
		Orders:  orders,
		PNLCard: pnlCard,
	}
}

// GetPositions returns mock positions data
func (s *DataService) GetPositions(userID uint) *models.PositionsResponse {
	positions := []models.Position{
		{
			Symbol:               "BHARTIARTL",
			Quantity:             20,
			AveragePrice:         965.30,
			CurrentPrice:         972.85,
			UnrealizedPNL:        151.00,
			UnrealizedPNLPercent: 0.78,
			PositionType:         "LONG",
		},
		{
			Symbol:               "SBIN",
			Quantity:             25,
			AveragePrice:         587.40,
			CurrentPrice:         582.15,
			UnrealizedPNL:        -131.25,
			UnrealizedPNLPercent: -0.89,
			PositionType:         "LONG",
		},
		{
			Symbol:               "KOTAKBANK",
			Quantity:             12,
			AveragePrice:         1789.60,
			CurrentPrice:         1795.20,
			UnrealizedPNL:        67.20,
			UnrealizedPNLPercent: 0.31,
			PositionType:         "LONG",
		},
		{
			Symbol:               "ITC",
			Quantity:             30,
			AveragePrice:         418.75,
			CurrentPrice:         415.30,
			UnrealizedPNL:        -103.50,
			UnrealizedPNLPercent: -0.82,
			PositionType:         "SHORT",
		},
	}

	pnlCard := models.PNLCard{
		TotalPNL:        -16.55,
		TotalPNLPercent: -0.02,
		DayPNL:          -16.55,
		DayPNLPercent:   -0.02,
		RealizedPNL:     0.00,
		UnrealizedPNL:   -16.55,
	}

	return &models.PositionsResponse{
		Positions: positions,
		PNLCard:   pnlCard,
	}
}

// GenerateRandomPrices generates random price fluctuations for demonstration
func (s *DataService) GenerateRandomPrices() {
	// This could be used to simulate real-time price updates
	// For now, it's a placeholder for future implementation
	rand.Seed(time.Now().UnixNano())
}
