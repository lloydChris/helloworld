package stuff

import (
	"fmt"
	"testing"
)

func TestGreeting(t *testing.T) {
	names := []string{
		"Charlotte",
		"Lorelei",
	}

	for _, v := range names {
		t.Run(v, func(t *testing.T) {
			result := SayHello(v)
			if result != fmt.Sprintf("Hello, %v", v) {
				t.Errorf("expected: something, actual: %v", result)
			}
		})
	}
}
