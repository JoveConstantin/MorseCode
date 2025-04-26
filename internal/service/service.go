package service

import (
	"errors"
	"strings"

	morse "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseOrTextConverter(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("Empty file transferred")
	}
	if isMorse(s) {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}

func isMorse(s string) bool {

	for _, r := range s {
		if !strings.ContainsRune(".- /_", r) {
			return false
		}
	}
	return true

}
