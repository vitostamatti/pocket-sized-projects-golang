package main

import "testing"

func TestGreet_English(t *testing.T) {
	lang := language("en")
	want := "Hello world"

	got := greet(lang)
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
func TestGreet_French(t *testing.T) {
	lang := language("fr")
	want := "Bonjour le monde"

	got := greet(lang)
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
func TestGreet_Default(t *testing.T) {
	lang := language("de")
	want := "unsupported language: \"de\""

	got := greet(lang)
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Greek": {
			lang: "el",
			want: "Γεια σου κόσμε",
		},
		"Default": {
			lang: "de",
			want: "unsupported language: \"de\"",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)
			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
