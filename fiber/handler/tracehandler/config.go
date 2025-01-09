package tracehandler

// Config : 暫時沒有可自訂項目
type Config struct {
	TraceHaderKey    string
	AllResAddTraceID bool
}

var ConfigDefault = Config{
	TraceHaderKey:    XCloudTraceContext,
	AllResAddTraceID: false,
}

var traceKey = XCloudTraceContext

func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]
	if len(cfg.TraceHaderKey) == 0 {
		cfg.TraceHaderKey = ConfigDefault.TraceHaderKey
	}

	return cfg
}
