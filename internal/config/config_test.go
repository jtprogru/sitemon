package config

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildConfigFileNotFound(t *testing.T) {
	configFile := "../../config.yaml"
	c, err := BuildConfig(configFile)
	require.Error(t, err)
	assert.Nil(t, c)
}

func TestBuildConfigEmptyFile(t *testing.T) {
	configFile := "./testdata/empty.config.yml"
	c, err := BuildConfig(configFile)
	require.Error(t, err)
	assert.Nil(t, c)
}

func TestBuildConfigSuccessDecode(t *testing.T) {
	configFile := "../../configs/config.yml"
	c, err := BuildConfig(configFile)
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "DEBUG", c.Log.Level)
	assert.Equal(t, "https://1231231231231231231238cc0375b556@o412493.ingest.sentry.io/5383803", c.Sentry.Dsn)
	assert.Equal(t, "123456789:qwertyuiopasdfghjkzxcvbnm", c.Telegram.Token)
	assert.Equal(t, "-12123123123", c.Telegram.Chat)
}

func TestBuildConfigSuccessEnvProcess(t *testing.T) {
	logLevel := "INFO"
	sentryDsn := "sentry:dsn"
	tgToken := "token"
	tgChat := "-123"
	os.Clearenv()
	os.Setenv("SITEMON_LOGLEVEL", logLevel)
	os.Setenv("SITEMON_SENTRYDSN", sentryDsn)
	os.Setenv("SITEMON_TGTOKEN", tgToken)
	os.Setenv("SITEMON_TGCHAT", tgChat)
	configFile := "../../configs/config.yml"
	c, err := BuildConfig(configFile)
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, logLevel, c.Log.Level)
	assert.Equal(t, sentryDsn, c.Sentry.Dsn)
	assert.Equal(t, tgToken, c.Telegram.Token)
	assert.Equal(t, tgChat, c.Telegram.Chat)
}
