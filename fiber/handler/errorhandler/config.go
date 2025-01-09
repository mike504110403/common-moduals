package errorhandler

// Config defines the config for middleware.
type Config struct {
	// ResponseTraceHeader : 回應 cloud trace header
	ResCloudTraceHeader bool
	Log                 bool
	LogRetStatus        bool
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	ResCloudTraceHeader: true,
	Log:                 true,
	LogRetStatus:        false,
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}
	// Override default config
	cfg := config[0]
	return cfg
}
