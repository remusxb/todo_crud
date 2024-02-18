package featureflags

import (
	"fmt"
	"log/slog"
	"strings"
)

var (
	featureFlags = make(map[string]bool)
)

func (c Config) Initialize() {
	for feature, enabled := range c {
		featureFlags[feature] = enabled
	}
}

// IsEnabled checks if a specific feature is enabled.
func IsEnabled(feature string) bool {
	var enabled, ok bool
	f := strings.ToLower(feature)

	if enabled, ok = featureFlags[f]; !ok {
		slog.Warn(fmt.Sprintf("feature `%s` not found", feature))
		return false
	}

	return enabled
}
