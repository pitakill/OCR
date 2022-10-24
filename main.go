package main

import (
	"bufio"
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
)

var str2digits = map[string]int32{
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

func format(input []string) []string {
	output := make([]string, size)
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

		output[i] = str[step:step+chunk] + str[step+delta:step+delta+chunk] + str[step+(delta*2):step+(delta*2)+chunk]
	}

	return output
}

func main() {
	filename := flag.String("filename", "", "File to get the info")
	flag.Parse()

	if *filename == "" {
		log.Fatal("Please provide a valid filename to extract the information with the flag filename")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path.Join(cwd, *filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bucket := make([]string, upToLines)

	// for tracking the lines read
	lines := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if lines == upToLines {
			lines = 1
			data := format(bucket)
			for i, digit := range data {
				fmt.Print(str2digits[digit])
				if i == len(data)-1 {
					fmt.Printf("\n")
				}
			}
			bucket = make([]string, upToLines)
			continue
		}

		bucket[lines] = scanner.Text()
		lines++
	}
}
