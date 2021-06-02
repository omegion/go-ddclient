package provider

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProvider(t *testing.T) {
	_ = os.Setenv("CF_API_KEY", "X")

	expectedProviders := []string{
		"cloudflare",
	}
	for _, name := range expectedProviders {
		_, err := GetProvider(name)
		assert.NoError(t, err)
	}
}
