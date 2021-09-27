package utils

import "os"

func GetNormalizedValues(headers map[string][]string) map[string]string {
	headersNorm := make(map[string]string)
	for key, val := range headers {
		if len(val) > 0 {
			headersNorm[key] = val[0]
		}
	}
	return headersNorm
}

func GetEnvOrDefault(envVar string, defaultValue string) string {
	if v := os.Getenv(envVar); v != "" {
		return v
	}
	return defaultValue
}