package core

const (
	APP_NAME    = "girie"
	APP_VERSION = "v1.1.0"

	DEFAULT_ETC_PATH        = "/etc/girie"
	DEFAULT_LOG_TIME_FORMAT = "02.01.2006 15:04:05.000"
	DEFAULT_LISTEN          = ":8080"
	DEFAULT_PROXY           = ""
	DEFAULT_RETRY           = 2
	DEFAULT_SPAN_THRESHOLD  = 10
	DEFAULT_TIMEOUT         = 3
	DEFAULT_USER_AGENT      = APP_NAME + " " + APP_VERSION

	LOG_CONFIG_ERROR = "config error"

	VIPER_DEFAULT_LISTEN     = "default.listen"
	VIPER_DEFAULT_PROXY      = "default.proxy"
	VIPER_DEFAULT_RETRY      = "default.retry"
	VIPER_DEFAULT_TIMEOUT    = "default.timeout"
	VIPER_DEFAULT_USER_AGENT = "default.user_agent"

	VIPER_ENV_LISTEN     = "listen"
	VIPER_ENV_PROXY      = "proxy"
	VIPER_ENV_RETRY      = "retry"
	VIPER_ENV_TIMEOUT    = "timeout"
	VIPER_ENV_USER_AGENT = "user_agent"
)
