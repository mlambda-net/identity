package conf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_Load(t *testing.T) {
	config := LoadConfig()
	assert.Equal(t, config.App.Name, "")
}
