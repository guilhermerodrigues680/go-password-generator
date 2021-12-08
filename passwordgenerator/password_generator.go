package passwordgenerator

import (
	"crypto/rand"
	"errors"
	"math/big"
)

type Alphabet uint8

const (
	LOWERCASE_ALPHABET Alphabet = iota
	UPPERCASE_ALPHABET
	NUMBERS_ALPHABET
	SYMBOLS_ALPHABET
)

const (
	lowercaseCharacters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbersCharacters   = "0123456789"
	symbolsCharacters   = "!@#$%^&*()+_-=}{[]|:;\"/?.><,`~"
	// similarCharacters   = "ilLI|`oO0"
)

func Generate(lenght int, enableLowecase, enableUppercase, enableNumbers, enableSymbols bool) (string, error) {
	var pwd string

	for {
		pwdChars := make([]byte, lenght)
		countAlphabetCharsMap := make(map[Alphabet]int)
		for i := 0; i < lenght; i++ {
			randAlphabet, err := randByte(4) // [0,3]
			if err != nil {
				return "", err
			}

			var alphabet string
			switch Alphabet(randAlphabet) {
			case LOWERCASE_ALPHABET:
				alphabet = lowercaseCharacters
			case UPPERCASE_ALPHABET:
				alphabet = uppercaseCharacters
			case NUMBERS_ALPHABET:
				alphabet = numbersCharacters
			case SYMBOLS_ALPHABET:
				alphabet = symbolsCharacters
			default:
				return "", errors.New("alfabeto inesperado")
			}

			if _, found := countAlphabetCharsMap[Alphabet(randAlphabet)]; !found {
				countAlphabetCharsMap[Alphabet(randAlphabet)] = 0
			}
			countAlphabetCharsMap[Alphabet(randAlphabet)] += 1

			randChar, err := randByteAlphabet(alphabet)
			if err != nil {
				return "", err
			}

			pwdChars = append(pwdChars, randChar)
		}

		pwd = string(pwdChars)
		if (lenght <= 4 && lenght != len(countAlphabetCharsMap)) || (lenght > 4 && len(countAlphabetCharsMap) != 4) {
			// log.Println("pwd invalido", pwd, countAlphabetCharsMap)
			continue
		}

		break
	}

	return pwd, nil
}

func randByte(max byte) (byte, error) {
	randBigInt, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return byte(randBigInt.Int64()), nil
}

func randByteAlphabet(alphabet string) (byte, error) {
	randPos, err := randByte(byte(len(alphabet)))
	if err != nil {
		return 0, err
	}
	return alphabet[randPos], nil
}
