package internal_test

import (
	"fmt"
	"testing"

	"github.com/edualb/ghype/internal"
)

func TestIsValidToken(t *testing.T) {
	type input struct {
		from string
	}

	type output struct {
		IsValid bool
	}

	type test struct {
		name string
		i    input
		o    output
	}

	tests := []test{
		{
			name: "IsValidToken with an empty value should return false",
			i: input{
				from: "",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with ( value should return false",
			i: input{
				from: "(",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with ) value should return false",
			i: input{
				from: ")",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with < value should return false",
			i: input{
				from: "<",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with > value should return false",
			i: input{
				from: ">",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with @ value should return false",
			i: input{
				from: "@",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with , value should return false",
			i: input{
				from: ",",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with ; value should return false",
			i: input{
				from: ";",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with : value should return false",
			i: input{
				from: ":",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with \\ value should return false",
			i: input{
				from: "\\",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with / value should return false",
			i: input{
				from: "/",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with DoubleQuoteMark value should return false",
			i: input{
				from: string(internal.DoubleQuoteMark.ToBytes()),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with [ value should return false",
			i: input{
				from: "[",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with ] value should return false",
			i: input{
				from: "]",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with ? value should return false",
			i: input{
				from: "?",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with = value should return false",
			i: input{
				from: "=",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with } value should return false",
			i: input{
				from: "}",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with { value should return false",
			i: input{
				from: "{",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with SP value should return false",
			i: input{
				from: string(internal.SP.ToBytes()),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with HT value should return false",
			i: input{
				from: string(internal.HT.ToBytes()),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with more than one wrong value should return false",
			i: input{
				from: "POS>T",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with GET value should return true",
			i: input{
				from: "GET",
			},
			o: output{
				IsValid: true,
			},
		},
	}

	for _, tt := range tests {
		out := internal.IsValidToken([]byte(tt.i.from))

		if out != tt.o.IsValid {
			t.Error(fmt.Sprintf("[%s]: want %v, got %v", tt.name, tt.o.IsValid, out))
			return
		}

	}
}

func TestIsCTLValid(t *testing.T) {
	type input struct {
		from string
	}

	type output struct {
		IsValid bool
	}

	type test struct {
		name string
		i    input
		o    output
	}

	tests := []test{
		{
			name: "IsValidToken with an empty value should return false",
			i: input{
				from: "",
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with NUL value should return false",
			i: input{
				from: string(rune(0)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with SOH value should return false",
			i: input{
				from: string(rune(1)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with STX value should return false",
			i: input{
				from: string(rune(2)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with ETX value should return false",
			i: input{
				from: string(rune(3)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with EOT value should return false",
			i: input{
				from: string(rune(4)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with ENQ value should return false",
			i: input{
				from: string(rune(5)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with ACK value should return false",
			i: input{
				from: string(rune(6)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with BEL value should return false",
			i: input{
				from: string(rune(7)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with BS value should return false",
			i: input{
				from: string(rune(8)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with TAB value should return false",
			i: input{
				from: string(rune(9)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with LF value should return false",
			i: input{
				from: string(rune(10)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with VT value should return false",
			i: input{
				from: string(rune(11)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with FF value should return false",
			i: input{
				from: string(rune(12)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with CR value should return false",
			i: input{
				from: string(rune(13)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with SO value should return false",
			i: input{
				from: string(rune(14)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with SI value should return false",
			i: input{
				from: string(rune(15)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with DLE value should return false",
			i: input{
				from: string(rune(16)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with DC1 value should return false",
			i: input{
				from: string(rune(17)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with DC2 value should return false",
			i: input{
				from: string(rune(18)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with DC3 value should return false",
			i: input{
				from: string(rune(19)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with DC4 value should return false",
			i: input{
				from: string(rune(20)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with NAK value should return false",
			i: input{
				from: string(rune(21)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with SYN value should return false",
			i: input{
				from: string(rune(22)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with ETB value should return false",
			i: input{
				from: string(rune(23)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with CAN value should return false",
			i: input{
				from: string(rune(24)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with EM value should return false",
			i: input{
				from: string(rune(25)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with SUB value should return false",
			i: input{
				from: string(rune(26)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with ESC value should return false",
			i: input{
				from: string(rune(27)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with FS value should return false",
			i: input{
				from: string(rune(28)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with GS value should return false",
			i: input{
				from: string(rune(29)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with RS value should return false",
			i: input{
				from: string(rune(30)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with US value should return false",
			i: input{
				from: string(rune(31)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with DEL value should return false",
			i: input{
				from: string(rune(127)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with GET+NUL value should return false",
			i: input{
				from: "GET" + string(rune(0)),
			},
			o: output{
				IsValid: false,
			},
		},
		{
			name: "IsValidToken with GET value should return true",
			i: input{
				from: "GET",
			},
			o: output{
				IsValid: true,
			},
		},
	}

	for _, tt := range tests {
		out := internal.IsCTLValid([]byte(tt.i.from))

		if out != tt.o.IsValid {
			t.Error(fmt.Sprintf("[%s]: want %v, got %v", tt.name, tt.o.IsValid, out))
			return
		}

	}
}

func TestIsByteAlpha(t *testing.T) {
	type input struct {
		from byte
	}

	type output struct {
		isValid bool
	}

	type test struct {
		name string
		i    input
		o    output
	}

	tests := []test{
		{
			name: "IsByteAlpha with h value should return isValid = true",
			i: input{
				from: []byte{'h'}[0],
			},
			o: output{
				isValid: true,
			},
		},
		{
			name: "IsByteAlpha with H value should return isValid = true",
			i: input{
				from: []byte{'H'}[0],
			},
			o: output{
				isValid: true,
			},
		},
		{
			name: "IsByteAlpha with * value should return isValid = false",
			i: input{
				from: []byte{'*'}[0],
			},
			o: output{
				isValid: false,
			},
		},
		{
			name: "IsByteAlpha with / value should return isValid = false",
			i: input{
				from: []byte{'/'}[0],
			},
			o: output{
				isValid: false,
			},
		},
	}

	for _, tt := range tests {
		isValid := internal.IsByteAlpha(tt.i.from)
		if isValid != tt.o.isValid {
			t.Error(fmt.Sprintf("[%s]: want isValid %v, got %v", tt.name, tt.o.isValid, isValid))
		}
	}
}

func TestIsByteDigit(t *testing.T) {
	type input struct {
		from byte
	}

	type output struct {
		isValid bool
	}

	type test struct {
		name string
		i    input
		o    output
	}

	tests := []test{
		{
			name: "IsByteDigit with 0 value should return isValid = true",
			i: input{
				from: []byte{'0'}[0],
			},
			o: output{
				isValid: true,
			},
		},
		{
			name: "IsByteDigit with 9 value should return isValid = true",
			i: input{
				from: []byte{'9'}[0],
			},
			o: output{
				isValid: true,
			},
		},
		{
			name: "IsByteDigit with h value should return isValid = false",
			i: input{
				from: []byte{'h'}[0],
			},
			o: output{
				isValid: false,
			},
		},
		{
			name: "IsByteDigit with : value should return isValid = false",
			i: input{
				from: []byte{':'}[0],
			},
			o: output{
				isValid: false,
			},
		},
	}

	for _, tt := range tests {
		isValid := internal.IsByteDigit(tt.i.from)
		if isValid != tt.o.isValid {
			t.Error(fmt.Sprintf("[%s]: want isValid %v, got %v", tt.name, tt.o.isValid, isValid))
		}
	}
}
