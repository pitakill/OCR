package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

const (
	zero  = " _ | ||_|"
	one   = "     |  |"
	two   = " _  _||_ "
	three = " _  _| _|"
	four  = "   |_|  |"
	five  = " _ |_  _|"
	six   = " _ |_ |_|"
	seven = " _   |  |"
	eight = " _ |_||_|"
	nine  = " _ |_| _|"

	// width of the digit
	chunk = 3
	// width of the account
	size = 9
	// height of the account number
	upToLines = 4
	// unknown number
	unknown = '?'
	// padding to correct ascii characters for numbers
	padding = 48
	// lowest valid rune
	floor = 0
	// heighest valid rune
	ceil = 9
)

var (
	str2digits = map[string]int32{
		zero:  0,
		one:   1,
		two:   2,
		three: 3,
		four:  4,
		five:  5,
		six:   6,
		seven: 7,
		eight: 8,
		nine:  9,
	}
	// illegible data
	ErrIllData = errors.New("Illegible data")
)

func format(input []string) []rune {
	output := make([]rune, size)
	str := strings.Join(input, "")
	str = strings.TrimSuffix(str, "\n")

	// output[0] = str[0:3] + str[27:30] + str[54:57]
	// output[1] = str[3:6] + str[30:33] + str[57:60]
	// ...
	// output[8] = str[24:27] + str[51:54] + str[78:81]
	// The abstraction of the code above is:
	// output[i] = str[i*chunk:i*chunk+chunk] + str[i*chunk+(size*chunk):i*chunk+(size*chunk)+chunk] + str[i*chunk+(size*chunk*2):i*chunk+(size*chunk*2)+chunk]
	// and finally
	for i := range output {
		step := i * chunk
		delta := size * chunk

		candidate := str[step:step+chunk] + str[step+delta:step+delta+chunk] + str[step+(delta*2):step+(delta*2)+chunk]

		if digit, ok := str2digits[candidate]; ok {
			output[i] = digit
		} else {
			output[i] = unknown
		}
	}

	return output
}

func checksum(input []rune) (bool, error) {
	if len(input) != size {
		return false, ErrIllData
	}

	var total int
	for i := len(input) - 1; i >= 0; i-- {
		if runeInRange(input[i]) {
			total += int(input[i]) * (len(input) - i)
		} else {
			return false, ErrIllData
		}
	}

	return total%11 == 0, nil
}

func runeInRange(input rune) bool {
	if floor <= input && input <= ceil {
		return true
	}

	return false
}

func fixPadding(input rune) rune {
	if runeInRange(input) {
		return input + padding
	}

	return input
}

func main() {
	input := flag.String("input", "", "File to get the info")
	output := flag.String("output", "output.txt", "File to store the processed info")
	flag.Parse()

	if *input == "" {
		log.Fatal("Please provide a valid filename to extract the information with the flag filename")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	inputFile, err := os.Open(path.Join(cwd, *input))
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	bucket := make([]string, upToLines)

	// for tracking the lines read
	lines := 1

	outputFile, err := os.Create(path.Join(cwd, *output))
	if err != nil {
		log.Fatalln(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		if lines == upToLines {
			lines = 1
			data := format(bucket)
			for i, rune := range data {
				outputFile.WriteString(string(fixPadding(rune))) // nolint
				if i == len(data)-1 {
					str := "OK"
					ok, err := checksum(data)
					if err != nil {
						str = "ILL"
					} else if !ok {
						str = "ERR"
					}
					outputFile.WriteString(fmt.Sprintf(" %s\n", str)) // nolint
				}
			}
			bucket = make([]string, upToLines)
			continue
		}

		bucket[lines] = scanner.Text()
		lines++
	}

	outputFile.Sync() // nolint
}
