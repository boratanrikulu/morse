package morse

import (
	"fmt"
	"strings"
	"testing"
)

var (
	testValues = map[string]string{
		"F":                  "..-.",
		"Bora Tanrıkulu":     "-... --- .-. .- / - .- -. .-. .. -.- ..- .-.. ..-",
		"Amatör Telsizcilik": ".- -- .- - ---. .-. / - . .-.. ... .. --.. -.-. .. .-.. .. -.-",
		"SOS":                "... --- ...",
	}
)

func TestMorseEncoding(t *testing.T) {
	for input, want := range testValues {
		m := NewMorse()

		result, err := m.Encode(strings.NewReader(input))
		if err != nil {
			t.Error(err)
		}

		if result != want {
			t.Fatalf("Encoding failed!\nResult: %s\nWanted: %s\n", result, want)
		}

		fmt.Printf("-\tResult: %s\n\tWanted: %s\n", result, want)
	}
}

func TestMorseDecoding(t *testing.T) {
	for want, input := range testValues {
		m := NewMorse()

		result, err := m.Decode(strings.NewReader(input))
		if err != nil {
			t.Error(err)
		}

		want = strings.ToUpper(want)
		if result != want {
			t.Fatalf("Decoding failed!\nResult: %s\nWanted: %s\n", result, want)
		}

		fmt.Printf("-\tResult: %s\n\tWanted: %s\n", result, want)
	}
}
