package notfound

// Config defines the config for middleware.
type Config struct {
	// ResponseTraceHeader : 回應 cloud trace header
	LogNotFound bool
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	LogNotFound: false,
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
