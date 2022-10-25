package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_format(t *testing.T) {
	tests := []struct {
		input string
		want  []rune
	}{
		{
			` _  _  _  _  _  _  _  _  _ 
| || || || || || || || || |
|_||_||_||_||_||_||_||_||_|

`,
			[]rune{
				'0', '0', '0', '0', '0', '0', '0', '0', '0',
			},
		},
		{
			`                           
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |

`,
			[]rune{
				'1', '1', '1', '1', '1', '1', '1', '1', '1',
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
|_ |_ |_ |_ |_ |_ |_ |_ |_ 

`,
			[]rune{
				'2', '2', '2', '2', '2', '2', '2', '2', '2',
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
 _| _| _| _| _| _| _| _| _|

`,
			[]rune{
				'3', '3', '3', '3', '3', '3', '3', '3', '3',
			},
		},
		{
			`                           
|_||_||_||_||_||_||_||_||_|
  |  |  |  |  |  |  |  |  |

`,
			[]rune{
				'4', '4', '4', '4', '4', '4', '4', '4', '4',
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
|_ |_ |_ |_ |_ |_ |_ |_ |_ 
 _| _| _| _| _| _| _| _| _|

`,
			[]rune{
				'5', '5', '5', '5', '5', '5', '5', '5', '5',
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
|_ |_ |_ |_ |_ |_ |_ |_ |_ 
|_||_||_||_||_||_||_||_||_|

`,
			[]rune{
				'6', '6', '6', '6', '6', '6', '6', '6', '6',
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |

`,
			[]rune{
				'7', '7', '7', '7', '7', '7', '7', '7', '7',
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
|_||_||_||_||_||_||_||_||_|
|_||_||_||_||_||_||_||_||_|

`,
			[]rune{
				'8', '8', '8', '8', '8', '8', '8', '8', '8',
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
|_||_||_||_||_||_||_||_||_|
 _| _| _| _| _| _| _| _| _|

`,
			[]rune{
				'9', '9', '9', '9', '9', '9', '9', '9', '9',
			},
		},
		{
			` _  _  _  _  _  _  _  _    
|_||_||_||_||_||_||_||_|   
 _| _| _| _| _| _| _| _|   

`,
			[]rune{
				'9', '9', '9', '9', '9', '9', '9', '9', '?' + padding,
			},
		},
	}

	for _, test := range tests {
		output := format(strings.Split(test.input, "\n"))

		// Correct padding
		for i := range test.want {
			test.want[i] -= padding
		}

		if !reflect.DeepEqual(output, test.want) {
			t.Errorf("output: %v, but want: %v", output, test.want)
		}
	}
}

func Test_checksum(t *testing.T) {
	tests := []struct {
		input []rune
		want  bool
	}{
		{
			[]rune{'0'},
			false,
		},
		{
			[]rune{'0', '1'},
			false,
		},
		{
			[]rune{'0', '0', '0', '0', '0', '0', '0', '0', '0'},
			true,
		},
		{
			[]rune{'3', '4', '5', '8', '8', '2', '8', '6', '5'},
			true,
		},
		{
			[]rune{'4', '5', '7', '5', '0', '8', '0', '0', '0'},
			true,
		},
		{
			[]rune{'6', '6', '4', '3', '7', '1', '4', '9', '5'},
			false,
		},
		{
			[]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'},
			false,
		},
	}

	for _, test := range tests {
		// Correct padding
		for i := range test.input {
			test.input[i] -= padding
		}

		output, _ := checksum(test.input)

		if output != test.want {
			t.Errorf("input: %v, output: %v, but want: %v", test.input, output, test.want)
		}
	}
}

func Test_runeInRange(t *testing.T) {
	tests := []struct {
		input rune
		want  bool
	}{
		{'0', true},
		{'1', true},
		{'2', true},
		{'3', true},
		{'4', true},
		{'5', true},
		{'6', true},
		{'7', true},
		{'8', true},
		{'9', true},
		{'?', false},
		{'a', false},
		{'>', false},
	}

	for _, test := range tests {
		// Correct padding
		if test.want {
			test.input -= padding
		}

		output := runeInRange(test.input)

		if output != test.want {
			t.Errorf("output: %v, but wants: %v", output, test.want)
		}
	}
}

func Test_fixPadding(t *testing.T) {
	tests := []struct {
		input rune
		want  rune
	}{
		{'0', '0'},
		{'1', '1'},
		{'2', '2'},
		{'3', '3'},
		{'4', '4'},
		{'5', '5'},
		{'6', '6'},
		{'7', '7'},
		{'8', '8'},
		{'9', '9'},
		{'\b', '8'},
		{'\t', '9'},
		{'?', '?'},
		{'a', 'a'},
		{'>', '>'},
	}

	for _, test := range tests {
		// Correct padding
		if runeInRange(test.want) {
			test.input -= padding
		}

		output := fixPadding(test.input)

		if output != test.want {
			t.Errorf("output: %v, but wants: %v", output, test.want)
		}
	}
}
