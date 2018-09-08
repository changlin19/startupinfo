package startupinfo

import (
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	configInfo := map[string]interface{}{"a": 123}
	Print("test.yaml", configInfo, time.Now().String(), "test")
}
