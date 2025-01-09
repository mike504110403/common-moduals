package ilog

import (
	"bytes"
	"os"
	"unicode/utf8"
)

// configDefault : 預設設定
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}
	// Override default config
	cfg := config[0]

	if len(cfg.ReportedErrorEvent) == 0 {
		cfg.ReportedErrorEvent = ConfigDefault.ReportedErrorEvent
	}

	if len(cfg.ReportLevel) == 0 {
		cfg.ReportLevel = ConfigDefault.ReportLevel
	}
	return cfg
}

func parseLevel(slice []severity) map[severity]bool {
	levelMap := make(map[severity]bool)
	for _, s := range slice {
		levelMap[s] = true
	}
	return levelMap
}

func fixUTF8(s string) string {
	if utf8.ValidString(s) {
		return s
	}

	// Otherwise time to build the sequence.
	buf := new(bytes.Buffer)
	buf.Grow(len(s))
	for _, r := range s {
		if utf8.ValidRune(r) {
			buf.WriteRune(r)
		} else {
			buf.WriteRune('\uFFFD')
		}
	}
	return buf.String()
}

func isArgGet(str string) bool {
	for _, arg := range os.Args {
		if arg == str {
			return true
		}
	}
	return false
}
