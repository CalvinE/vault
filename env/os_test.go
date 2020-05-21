package env

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	testEnvVar := "TEST_ENV_VAR_NAME"
	fallbackTestValue, testValue := "fallback", "set_value"
	notSetValue := Get(testEnvVar, fallbackTestValue)
	if fallbackTestValue != notSetValue {
		t.Errorf("fallback value not returned properly: expected %v got %v\n", fallbackTestValue, notSetValue)
	}
	os.Setenv(testEnvVar, testValue)
	setTestValue := Get(testEnvVar, testValue)
	if testValue != setTestValue {
		t.Errorf("test value not returned properly: expected %v got %v\n", testValue, setTestValue)
	}
}
