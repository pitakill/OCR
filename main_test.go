package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_format(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			` _  _  _  _  _  _  _  _  _ 
| || || || || || || || || |
|_||_||_||_||_||_||_||_||_|

`,
			[]string{
				zero, zero, zero, zero, zero, zero, zero, zero, zero,
			},
		},
		{
			`                           
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |

`,
			[]string{
				one, one, one, one, one, one, one, one, one,
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
|_ |_ |_ |_ |_ |_ |_ |_ |_ 

`,
			[]string{
				two, two, two, two, two, two, two, two, two,
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
 _| _| _| _| _| _| _| _| _|

`,
			[]string{
				three, three, three, three, three, three, three, three, three,
			},
		},
		{
			`                           
|_||_||_||_||_||_||_||_||_|
  |  |  |  |  |  |  |  |  |

`,
			[]string{
				four, four, four, four, four, four, four, four, four,
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
|_ |_ |_ |_ |_ |_ |_ |_ |_ 
 _| _| _| _| _| _| _| _| _|

`,
			[]string{
				five, five, five, five, five, five, five, five, five,
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
|_ |_ |_ |_ |_ |_ |_ |_ |_ 
|_||_||_||_||_||_||_||_||_|

`,
			[]string{
				six, six, six, six, six, six, six, six, six,
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |

`,
			[]string{
				seven, seven, seven, seven, seven, seven, seven, seven, seven,
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
|_||_||_||_||_||_||_||_||_|
|_||_||_||_||_||_||_||_||_|

`,
			[]string{
				eight, eight, eight, eight, eight, eight, eight, eight, eight,
			},
		},
		{
			` _  _  _  _  _  _  _  _  _ 
|_||_||_||_||_||_||_||_||_|
 _| _| _| _| _| _| _| _| _|

`,
			[]string{
				nine, nine, nine, nine, nine, nine, nine, nine, nine,
			},
		},
	}

	for _, test := range tests {
		output := format(strings.Split(test.input, "\n"))

		if !reflect.DeepEqual(output, test.want) {
			t.Errorf("output: %v, but want: %v", output, test.want)
		}
	}
}
