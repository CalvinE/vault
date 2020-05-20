package env

import "os"

// Get attempts to read a variable from environment variables
// If not present the fallback value is returned.
//
// env is the environment variable to look for.
//
// fallback is the string to use in the event that the environment variable is not set.
func Get(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
