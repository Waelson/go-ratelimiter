package config_test

import (
	"errors"
	"github.com/Waelson/go-ratelimit/internal/config"
	"testing"
)

// Aqui vai a definição de MockEnvLoader
type MockEnvLoader struct {
}

func (m MockEnvLoader) Load() error {
	return errors.New("failed to load .env file")
}

func TestLoadConfigFailure(t *testing.T) {
	mockLoaderFail := &MockEnvLoader{}

	// Chama LoadConfig com o loader que falha
	_, err := config.LoadConfig(mockLoaderFail)
	if err == nil {
		t.Errorf("Exptected failure")
	}
}
