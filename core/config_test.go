package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	defer func() {
		if rec := recover(); rec != nil {
			t.Errorf("Loading config failed: %v\n", rec)
		}
	}()

	config := LoadConfig("../.env")

	assert.IsType(t, Config{}, config)
}

func TestGetConfig(t *testing.T) {
	config := GetConfig()
	strategyName := config.Env.StrategyName
	assert.IsType(t, Config{}, config)

	config.Env.StrategyName = "undefined"
	config = GetConfig()
	assert.Equal(t, strategyName, config.Env.StrategyName)
}
