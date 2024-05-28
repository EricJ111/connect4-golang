package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Read() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func ReadInt() int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	parsedInt, _ := strconv.ParseInt(text, 10, 64)
	return int(parsedInt)
}
