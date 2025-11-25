package services

import (
	"crypto/rand"
	"math/big"
	"strings"
)

type PasswordConfig struct {
	Length     int
	UseLetters bool
	UseDigits  bool
	UseSymbols bool
	UseUpper   bool
}

var CharsetMap = map[string]string{
	"lower":   "abcdefghijklmnopqrstuvwxyz",
	"upper":   "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"digits":  "0123456789",
	"symbols": "!@#$%^&*()_+-=[]{}|;:,.<>?",
}

func GeneratePassword(config PasswordConfig) (string, error) {
	var charsetBuilder strings.Builder
	charsetBuilder.Grow(100)

	rules := []struct {
		Enabled bool
		Key     string
	}{
		{config.UseLetters, "lower"},
		{config.UseUpper, "upper"},
		{config.UseDigits, "digits"},
		{config.UseSymbols, "symbols"},
	}

	for _, rule := range rules {
		if rule.Enabled {
			charsetBuilder.WriteString(CharsetMap[rule.Key])
		}
	}

	if charsetBuilder.Len() == 0 {
		charsetBuilder.WriteString(CharsetMap["lower"])
	}

	fullCharset := charsetBuilder.String()
	fullCharsetLen := big.NewInt(int64(len(fullCharset)))
	password := make([]byte, config.Length)

	for i := 0; i < config.Length; i++ {
		num, err := rand.Int(rand.Reader, fullCharsetLen)
		if err != nil {
			return "", err
		}
		password[i] = fullCharset[num.Int64()]
	}

	return string(password), nil
}
