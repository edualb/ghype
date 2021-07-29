package internal

import (
	"bytes"
	"strings"
)

type Rule rune

const (
	CR              = Rule(13)
	LF              = Rule(10)
	HT              = Rule(9)
	SP              = Rule(32)
	DoubleQuoteMark = Rule(34)
)

func (r Rule) ToBytes() []byte {
	return []byte(string(r))
}

// IsValidToken is responsable to check if the token is valid, following this rule:
// token = 1*<any CHAR except CTLs or separators>
// separators = "(" | ")" | "<" | ">" | "@" | "," | ";" | ":" | "\" | <"> | "/" | "[" | "]" | "?" | "=" | "{" | "}" | SP | HT
func IsValidToken(from []byte) bool {
	if len(from) <= 0 {
		return false
	}
	if !IsCTLValid(from) {
		return false
	}
	return !bytes.Contains(from, []byte("(")) &&
		!bytes.Contains(from, []byte(")")) &&
		!bytes.Contains(from, []byte("<")) &&
		!bytes.Contains(from, []byte(">")) &&
		!bytes.Contains(from, []byte("@")) &&
		!bytes.Contains(from, []byte(",")) &&
		!bytes.Contains(from, []byte(";")) &&
		!bytes.Contains(from, []byte(":")) &&
		!bytes.Contains(from, []byte("\\")) &&
		!bytes.Contains(from, DoubleQuoteMark.ToBytes()) &&
		!bytes.Contains(from, []byte("/")) &&
		!bytes.Contains(from, []byte("[")) &&
		!bytes.Contains(from, []byte("]")) &&
		!bytes.Contains(from, []byte("?")) &&
		!bytes.Contains(from, []byte("=")) &&
		!bytes.Contains(from, []byte("{")) &&
		!bytes.Contains(from, []byte("}")) &&
		!bytes.Contains(from, SP.ToBytes()) &&
		!bytes.Contains(from, HT.ToBytes())
}

// IsValidToken is responsable to check if the CTL is valid, following this rule:
// CTL = <any US-ASCII control character(octets 0 - 31) and DEL (127)>
func IsCTLValid(from []byte) bool {
	if len(from) <= 0 {
		return false
	}
	return !bytes.ContainsRune(from, rune(0)) &&
		!bytes.ContainsRune(from, rune(1)) &&
		!bytes.ContainsRune(from, rune(2)) &&
		!bytes.ContainsRune(from, rune(3)) &&
		!bytes.ContainsRune(from, rune(4)) &&
		!bytes.ContainsRune(from, rune(5)) &&
		!bytes.ContainsRune(from, rune(6)) &&
		!bytes.ContainsRune(from, rune(7)) &&
		!bytes.ContainsRune(from, rune(8)) &&
		!bytes.ContainsRune(from, rune(9)) &&
		!bytes.ContainsRune(from, rune(10)) &&
		!bytes.ContainsRune(from, rune(11)) &&
		!bytes.ContainsRune(from, rune(12)) &&
		!bytes.ContainsRune(from, rune(13)) &&
		!bytes.ContainsRune(from, rune(14)) &&
		!bytes.ContainsRune(from, rune(15)) &&
		!bytes.ContainsRune(from, rune(16)) &&
		!bytes.ContainsRune(from, rune(17)) &&
		!bytes.ContainsRune(from, rune(18)) &&
		!bytes.ContainsRune(from, rune(19)) &&
		!bytes.ContainsRune(from, rune(20)) &&
		!bytes.ContainsRune(from, rune(21)) &&
		!bytes.ContainsRune(from, rune(22)) &&
		!bytes.ContainsRune(from, rune(23)) &&
		!bytes.ContainsRune(from, rune(24)) &&
		!bytes.ContainsRune(from, rune(25)) &&
		!bytes.ContainsRune(from, rune(26)) &&
		!bytes.ContainsRune(from, rune(27)) &&
		!bytes.ContainsRune(from, rune(28)) &&
		!bytes.ContainsRune(from, rune(29)) &&
		!bytes.ContainsRune(from, rune(30)) &&
		!bytes.ContainsRune(from, rune(31)) &&
		!bytes.ContainsRune(from, rune(127))
}

func IsByteAlpha(from byte) bool {
	b := []byte{from}
	return IsAlpha(b)
}

func IsAlpha(from []byte) bool {
	return isAlpha(from, 65)
}

func isAlpha(from []byte, runeValue int8) bool {
	if runeValue == 91 {
		return false
	}

	r := rune(runeValue)
	rString := string(r)

	if bytes.ContainsAny(from, rString) || bytes.ContainsAny(from, strings.ToLower(rString)) {
		return true
	}

	return false || isAlpha(from, runeValue+1)
}

func IsByteDigit(from byte) bool {
	b := []byte{from}
	return IsDigit(b)
}

func IsDigit(from []byte) bool {
	return isDigit(from, 48)
}

func isDigit(from []byte, runeValue int8) bool {
	if runeValue == 58 {
		return false
	}

	if bytes.ContainsRune(from, rune(runeValue)) {
		return true
	}

	return false || isDigit(from, runeValue+1)
}
