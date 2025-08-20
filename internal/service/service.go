package service

import (
	"errors"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	if input == "" {
		return "", errors.New("пустая строка")
	}

	if isMorse(input) {
		return morse.ToText(input), nil
	}

	return morse.ToMorse(input), nil
}

func isMorse(s string) bool {

	for _, c := range s {
		if c != '.' && c != '-' && c != ' ' {
			return false
		}
	}

	return true
}
