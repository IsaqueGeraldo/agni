package agni

import (
	"testing"
)

func TestLogger(t *testing.T) {
	t.Run("Test default style", func(t *testing.T) {
		Logger("Test message")
	})

	t.Run("Test custom style", func(t *testing.T) {
		Logger("Test message", RedText, Bold)
	})
}

func TestPrintln(t *testing.T) {
	t.Run("Test default style", func(t *testing.T) {
		Println("Test message")
	})

	t.Run("Test custom style", func(t *testing.T) {
		Println("Test message", CyanText, YellowBackground, Bold, Underline)
	})
}
