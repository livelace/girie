package core

import (
	log "github.com/livelace/logrus"
	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
	// Read generated/existed configuration.
	v := viper.New()
	v.SetConfigName("config.toml")
	v.SetConfigType("toml")
	v.AddConfigPath(DEFAULT_ETC_PATH)

	if err := v.ReadInConfig(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn(LOG_CONFIG_ERROR)
	}

	// Set defaults.
	v.SetDefault(VIPER_DEFAULT_LISTEN, DEFAULT_LISTEN)
	v.SetDefault(VIPER_DEFAULT_PROXY, DEFAULT_PROXY)
	v.SetDefault(VIPER_DEFAULT_RETRY, DEFAULT_RETRY)
	v.SetDefault(VIPER_DEFAULT_TIMEOUT, DEFAULT_TIMEOUT)
	v.SetDefault(VIPER_DEFAULT_USER_AGENT, DEFAULT_USER_AGENT)

	// Environment variables have higher priority over config parameters.
	v.SetEnvPrefix(APP_NAME)
	_ = v.BindEnv(VIPER_ENV_LISTEN)
	_ = v.BindEnv(VIPER_ENV_PROXY)
	_ = v.BindEnv(VIPER_ENV_RETRY)
	_ = v.BindEnv(VIPER_ENV_TIMEOUT)
	_ = v.BindEnv(VIPER_ENV_USER_AGENT)

	if v.IsSet(VIPER_ENV_LISTEN) {
		v.Set(VIPER_DEFAULT_LISTEN, v.GetString(VIPER_ENV_LISTEN))
	}

	if v.IsSet(VIPER_ENV_PROXY) {
		v.Set(VIPER_DEFAULT_PROXY, v.GetString(VIPER_ENV_PROXY))
	}

	if v.IsSet(VIPER_ENV_RETRY) {
		v.Set(VIPER_DEFAULT_RETRY, v.GetInt(VIPER_ENV_RETRY))
	}

	if v.IsSet(VIPER_ENV_TIMEOUT) {
		v.Set(VIPER_DEFAULT_TIMEOUT, v.GetInt(VIPER_ENV_TIMEOUT))
	}

	if v.IsSet(VIPER_ENV_USER_AGENT) {
		v.Set(VIPER_DEFAULT_USER_AGENT, v.GetString(VIPER_ENV_USER_AGENT))
	}

	return v
}
