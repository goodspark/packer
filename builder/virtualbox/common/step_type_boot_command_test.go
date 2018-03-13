package common

import (
	"testing"
)

func TestScancodes(t *testing.T) {
	var bootcommand = []string{
		"1234567890-=<enter><wait>",
		"!@#$%^&*()_+<enter>",
		"qwertyuiop[]<enter>",
		"QWERTYUIOP{}<enter>",
		"asdfghjkl;'`<enter>",
		`ASDFGHJKL:"~<enter>`,
		"\\zxcvbnm,./<enter>",
		"|ZXCVBNM<>?<enter>",
		"<enter>",
		"<leftAltOn><esc><leftAltOff><wait5>",
		"<tab><tab><tab><tab><tab><spacebar><wait5>",
	}

	var expected = [][]string{
		{"02", "82", "03", "83", "04", "84", "05", "85", "06", "86", "07", "87", "08", "88", "09", "89", "0a", "8a", "0b", "8b", "0c", "8c", "0d", "8d", "1c", "9c", "wait"},
		{"2a", "02", "aa", "82", "2a", "03", "aa", "83", "2a", "04", "aa", "84", "2a", "05", "aa", "85", "2a", "06", "aa", "86", "2a", "07", "aa", "87", "2a", "08", "aa", "88", "2a", "09", "aa", "89", "2a", "0a", "aa", "8a", "2a", "0b", "aa", "8b", "2a", "0c", "aa", "8c", "2a", "0d", "aa", "8d", "1c", "9c"},
		{"10", "90", "11", "91", "12", "92", "13", "93", "14", "94", "15", "95", "16", "96", "17", "97", "18", "98", "19", "99", "1a", "9a", "1b", "9b", "1c", "9c"},
		{"2a", "10", "aa", "90", "2a", "11", "aa", "91", "2a", "12", "aa", "92", "2a", "13", "aa", "93", "2a", "14", "aa", "94", "2a", "15", "aa", "95", "2a", "16", "aa", "96", "2a", "17", "aa", "97", "2a", "18", "aa", "98", "2a", "19", "aa", "99", "2a", "1a", "aa", "9a", "2a", "1b", "aa", "9b", "1c", "9c"},
		{"1e", "9e", "1f", "9f", "20", "a0", "21", "a1", "22", "a2", "23", "a3", "24", "a4", "25", "a5", "26", "a6", "27", "a7", "28", "a8", "29", "a9", "1c", "9c"},
		{"2a", "1e", "aa", "9e", "2a", "1f", "aa", "9f", "2a", "20", "aa", "a0", "2a", "21", "aa", "a1", "2a", "22", "aa", "a2", "2a", "23", "aa", "a3", "2a", "24", "aa", "a4", "2a", "25", "aa", "a5", "2a", "26", "aa", "a6", "2a", "27", "aa", "a7", "2a", "28", "aa", "a8", "2a", "29", "aa", "a9", "1c", "9c"},
		{"2b", "ab", "2c", "ac", "2d", "ad", "2e", "ae", "2f", "af", "30", "b0", "31", "b1", "32", "b2", "33", "b3", "34", "b4", "35", "b5", "1c", "9c"},
		{"2a", "2b", "aa", "ab", "2a", "2c", "aa", "ac", "2a", "2d", "aa", "ad", "2a", "2e", "aa", "ae", "2a", "2f", "aa", "af", "2a", "30", "aa", "b0", "2a", "31", "aa", "b1", "2a", "32", "aa", "b2", "2a", "33", "aa", "b3", "2a", "34", "aa", "b4", "2a", "35", "aa", "b5", "1c", "9c"},
		{"1c", "9c"},
		{"38", "01", "81", "b8", "wait5"},
		{"0f", "8f", "0f", "8f", "0f", "8f", "0f", "8f", "0f", "8f", "39", "b9", "wait5"},
	}

	for i, command := range bootcommand {
		for j, code := range scancodes(command) {
			if code != expected[i][j] {
				t.Fatalf("%#v should have become %#v", scancodes(command), expected[i])
			}
		}
	}
}