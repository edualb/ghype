package ghype

import (
	"fmt"
	"testing"
)

func TestBuildFromRequestURI(t *testing.T) {
	type input struct {
		from string
	}

	type output struct {
		rURI     string
		gHypeErr error
	}

	type test struct {
		name string
		i    input
		o    output
	}

	tests := []test{
		{
			name: "buildFromRequestURI with * value should return * string and err = nil",
			i: input{
				from: "*",
			},
			o: output{
				rURI:     "*",
				gHypeErr: nil,
			},
		},
		{
			name: "buildFromRequestURI with something:something value should return something:blublu string and err = nil",
			i: input{
				from: "something:something",
			},
			o: output{
				rURI:     "something:something",
				gHypeErr: nil,
			},
		},
		{
			name: "buildFromRequestURI with something: value should return empty string and err = ErrRequestURIParse",
			i: input{
				from: "something:",
			},
			o: output{
				rURI:     "something:",
				gHypeErr: nil,
			},
		},
		{
			name: "buildFromRequestURI with : value should return empty string and err = ErrRequestURIParse",
			i: input{
				from: ":",
			},
			o: output{
				rURI:     "",
				gHypeErr: ErrRequestURIParse,
			},
		},
		{
			name: "buildFromRequestURI with :something value should return empty string and err = ErrRequestURIParse",
			i: input{
				from: ":something",
			},
			o: output{
				rURI:     "",
				gHypeErr: ErrRequestURIParse,
			},
		},
		{
			name: "buildFromRequestURI with something/:something value should return empty string and err = ErrRequestURIParse",
			i: input{
				from: "something/:something",
			},
			o: output{
				rURI:     "",
				gHypeErr: ErrRequestURIParse,
			},
		},
	}

	for _, tt := range tests {
		rURI, err := buildFromRequestURI([]byte(tt.i.from))
		if err != tt.o.gHypeErr {
			t.Error(fmt.Sprintf("[%s]: want err %v, got %v", tt.name, tt.o.gHypeErr, err))
			return
		}
		if tt.o.rURI != rURI {
			t.Error(fmt.Sprintf("[%s]: want rURI %s, got %v", tt.name, tt.o.rURI, rURI))
		}
	}
}

func TestIsSchemeValid(t *testing.T) {
	type input struct {
		scheme string
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
			name: "isSchemeValid with http value should return isValid = true",
			i: input{
				scheme: "http",
			},
			o: output{
				isValid: true,
			},
		},
		{
			name: "isSchemeValid with h2tp value should return isValid = true",
			i: input{
				scheme: "h2tp",
			},
			o: output{
				isValid: true,
			},
		},
		{
			name: "isSchemeValid with http+ value should return isValid = true",
			i: input{
				scheme: "http+",
			},
			o: output{
				isValid: true,
			},
		},
		{
			name: "isSchemeValid with http- value should return isValid = true",
			i: input{
				scheme: "http-",
			},
			o: output{
				isValid: true,
			},
		},
		{
			name: "isSchemeValid with http. value should return isValid = true",
			i: input{
				scheme: "http.",
			},
			o: output{
				isValid: true,
			},
		},
		{
			name: "isSchemeValid with h22p+.-9 value should return isValid = true",
			i: input{
				scheme: "h22p+.-9",
			},
			o: output{
				isValid: true,
			},
		},
		{
			name: "isSchemeValid with h22p+./9 value should return isValid = false",
			i: input{
				scheme: "h22p+./9",
			},
			o: output{
				isValid: false,
			},
		},
		{
			name: "isSchemeValid with empty value should return isValid = false",
			i: input{
				scheme: "",
			},
			o: output{
				isValid: false,
			},
		},
		{
			name: "isSchemeValid with 8ttp value should return isValid = false",
			i: input{
				scheme: "8ttp",
			},
			o: output{
				isValid: false,
			},
		},
	}

	for _, tt := range tests {
		isValid := isSchemeValid([]byte(tt.i.scheme))
		if isValid != tt.o.isValid {
			t.Error(fmt.Sprintf("[%s]: want isValid %v, got %v", tt.name, tt.o.isValid, isValid))
		}
	}
}
