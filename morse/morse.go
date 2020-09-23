package morse

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// Morse is a interface that you encode and decode
// with your input.
type Morse interface {
	Encode(r io.Reader) (string, error)
	Decode(r io.Reader) (string, error)
}

// NewMorse returns a new morse struct.
func NewMorse() Morse {
	return &morse{}
}

type morse struct{}

// Encode returns encoded morse code for the given input.
func (m *morse) Encode(r io.Reader) (string, error) {
	rr, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("There is an issue with input. Err: %s", err)
	}

	data := strings.ToUpper(strings.TrimSpace(string(rr)))
	words := strings.Split(data, " ")

	var resultCodes []string
	var foundLetterCount int
	for _, word := range words {
		var wordCodes []string

		for _, letter := range word {
			if code, ok := letterToCode[letter]; ok {
				wordCodes = append(wordCodes, code)
				foundLetterCount++
			}
		}

		wResult := strings.Join(wordCodes, " ")
		if wResult == "" {
			continue
		}

		resultCodes = append(resultCodes, wResult)
	}

	result := strings.Join(resultCodes, " / ")
	return result, nil
}

// Decode returns decoded morse code result for the given input.
func (m *morse) Decode(r io.Reader) (string, error) {
	rr, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("There is an issue with input. Err: %s", err)
	}

	data := string(rr)
	resultCodes := strings.Split(data, "/")

	var result []string
	for _, resultCode := range resultCodes {
		var word string

		resultCode = strings.TrimSpace(resultCode)
		for _, letterCode := range strings.Split(resultCode, " ") {
			if letter, ok := codeToLetter[letterCode]; ok {
				word += string(letter)
			}
		}

		result = append(result, word)
	}

	v := strings.Join(result, " ")
	return v, nil
}

// Letters for long and short morse signals.
const (
	S = "." // Short signal.
	L = "-" // Long signal.
)

// Morse Codes
//
// Sources:
// - https://morsecode.world/international/morse.html
// - https://morsedecoder.com/tr/
var letterToCode = map[rune]string{
	'A':  S + L,
	'B':  L + S + S + S,
	'C':  L + S + L + S,
	'D':  L + S + S,
	'E':  S,
	'F':  S + S + L + S,
	'G':  L + L + S,
	'H':  S + S + S + S,
	'I':  S + S,
	'J':  S + L + L + L,
	'K':  L + S + L,
	'L':  S + L + S + S,
	'M':  L + L,
	'N':  L + S,
	'O':  L + L + L,
	'P':  S + L + L + S,
	'Q':  L + L + S + L,
	'R':  S + L + S,
	'S':  S + S + S,
	'T':  L,
	'U':  S + S + L,
	'V':  S + S + S + L,
	'W':  S + L + L,
	'X':  L + S + S + L,
	'Y':  L + S + L + L,
	'Z':  L + L + S + S,
	'Ç':  L + S + L + S + S, // Turkish Special Letters
	'Ğ':  L + L + S + L + S, //
	'İ':  S + L + S + S + L, //
	'Ö':  L + L + L + S,     //
	'Ş':  S + L + L + S + S, //
	'Ü':  S + S + L + L,     //
	'Á':  S + L + L + S + L, // Continental (Gerke)
	'Ä':  S + L + S + L,     //
	'É':  S + S + L + S + S, //
	'Ñ':  L + L + S + L + L, //
	'1':  S + L + L + L + L,
	'2':  S + S + L + L + L,
	'3':  S + S + S + L + L,
	'4':  S + S + S + S + L,
	'5':  S + S + S + S + S,
	'6':  L + S + S + S + S,
	'7':  L + L + S + S + S,
	'8':  L + L + L + S + S,
	'9':  L + L + L + L + S,
	'0':  L + L + L + L + L,
	'.':  S + L + S + L + S + L,
	':':  L + L + L + S + S + S,
	',':  L + L + S + S + L + L,
	';':  L + S + L + S + L,
	'?':  S + S + L + L + S + S,
	'=':  L + S + S + S + L,
	'\'': S + L + L + L + L + S,
	'/':  L + S + S + L + S,
	'!':  L + S + L + S + L + L,
	'-':  L + S + S + S + S + L,
	'_':  S + S + L + L + S + L,
	'"':  S + L + S + S + L + S,
	'(':  L + S + L + L + S,
	')':  L + S + L + L + S + L,
	'$':  S + S + S + L + S + S + L,
	'&':  S + L + S + S + S,
	'@':  S + L + L + S + L + S,
	'+':  S + L + S + L + S,
}

// codeToLetter is auto-created while initializing the package.
var codeToLetter map[string]rune

func init() {
	codeToLetter = make(map[string]rune)
	for k, v := range letterToCode {
		codeToLetter[v] = k
	}
}
