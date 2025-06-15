// services/circuit_breaker_service.go
package services

import (
	"time"

	"github.com/sony/gobreaker"
)

type CircuitBreakerService struct {
	breakers map[string]*gobreaker.CircuitBreaker
}

func NewCircuitBreakerService() *CircuitBreakerService {
	return &CircuitBreakerService{
		breakers: make(map[string]*gobreaker.CircuitBreaker),
	}
}

func (s *CircuitBreakerService) GetBreaker(name string) *gobreaker.CircuitBreaker {
	if breaker, exists := s.breakers[name]; exists {
		return breaker
	}

	// Create new circuit breaker with default settings
	settings := gobreaker.Settings{
		Name:        name,
		MaxRequests: 3,
		Interval:    time.Second * 10,
		Timeout:     time.Second * 30,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			// Log state changes
		},
	}

	breaker := gobreaker.NewCircuitBreaker(settings)
	s.breakers[name] = breaker
	return breaker
}
