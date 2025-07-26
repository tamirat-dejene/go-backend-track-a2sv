package bootstrap

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadEnv_Success(t *testing.T) {
	// Write a temporary .env file
	content := `
APP_ENV=development
SERVER_ADDRESS=localhost:8080
MONGO_URI=mongodb://localhost:27017
DB_NAME=testdb
REFRESH_TOKEN_SECRET=refreshsecret
ACCESS_TOKEN_SECRET=accesssecret
REFRESH_TOKEN_EXPIRY_HOUR=24
ACCESS_TOKEN_EXPIRY_HOUR=1
CTX_TIMEOUT=30
`
	tmpFile := ".test.env"
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	require.NoError(t, err)
	defer os.Remove(tmpFile)

	env, err := NewEnv(tmpFile)
	require.NoError(t, err)
	require.Equal(t, "development", env.AppEnv)
	require.Equal(t, "localhost:8080", env.ServerAddress)
}
