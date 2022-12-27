package helpers

import (
	"testing"
)

func HandleTestError(errorString string, err error, t *testing.T) {
	if err != nil {
		t.Log(errorString+":", err)
	}
}
