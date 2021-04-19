package enroll

import (
	"testing"
)

func TestEscapeAttributeValue(t *testing.T) {
	cases := []struct {
		name, in, exp string
	}{
		{"backslash", "Foo\\Bar", "Foo\\\\Bar"},
		{"carriageReturn", "Foo\rBar", "Foo\\0DBar"},
		{"comma", "Foo,Bar", "Foo\\,Bar"},
		{"default", "MicroMDM", "MicroMDM"},
		{"doubleQuote", "\"Foo\"Bar", "\\\"Foo\\\"Bar"},
		{"emptyString", "", ""},
		{"equalsSign", "Foo=Bar", "Foo\\=Bar"},
		{"forwardsSlash", "Foo/Bar", "Foo\\/Bar"},
		{"hashAtStart", "#Foo", "\\#Foo"},
		{"leftAngleBracket", "Foo<Bar", "Foo\\<Bar"},
		{"lineFeed", "Foo\nBar", "Foo\\0ABar"},
		{"plusSign", "Foo+Bar", "Foo\\+Bar"},
		{"rightAngleBracket", "Foo>Bar", "Foo\\>Bar"},
		{"semicolon", "Foo;Bar", "Foo\\;Bar"},
		{"spaceAtEnd", "Foo ", "Foo\\ "},
		{"spaceAtStart", " Foo", "\\ Foo"},
		{"spaceAtStartAndEnd", " Foo ", "\\ Foo\\ "},
		{"spaceOnly", " ", "\\ "},
		{"null", "Foo\x00Bar", "Foo\\00Bar"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := escapeAttributeValue(c.in)

			if got != c.exp {
				t.Errorf("Got %s; want %s", got, c.exp)
			}
		})
	}
}
