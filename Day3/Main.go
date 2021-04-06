package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	weird    = "Weird"
	notweird = "Not Weird"
)

func calcNumber(input int32) string {

	if input%2 == 0 {
		if input <= 5 && input >= 2 {
			return notweird
		} else if input >= 6 && input <= 20 {
			return weird
		} else if input > 20 {
			return notweird
		} else {
			return weird
		}

	} else {
		return weird
	}

}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)
	fmt.Println(calcNumber(N))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
