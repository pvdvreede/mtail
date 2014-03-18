package main

import (
	"fmt"
	"github.com/ActiveState/tail"
	"os"
	"strconv"
)

func colourize(code, message string) string {
	return "\033[" + code + "m" + message + "\033[0m"
}

func embolden(message string) string {
	return "\033[1m" + message + "\033[0m"
}

func pad(message string, amount int) string {
	messageCount := len(message)
	if messageCount < amount {
		newMessage := message
		for i := 0; i < (amount - messageCount); i++ {
			newMessage += " "
		}
		return newMessage
	} else if messageCount == amount {
		return message
	} else {
		return "..." + message[(messageCount-amount)+3:]
	}
}

func main() {
	output := make(chan string)
	defer func() {
		close(output)
		tail.Cleanup()
	}()

	for i, arg := range os.Args {
		go func(fileNumber int, filename string, outputChan chan string) {
			fmt.Println("The number is " + strconv.Itoa(fileNumber+30))
			t, _ := tail.TailFile(filename, tail.Config{
				Follow: true, LimitRate: 15, MaxLineSize: 120,
			})
			for line := range t.Lines {
				message := line.Text
				output <- embolden(pad(filename, 15)+" => ") + colourize(strconv.Itoa(fileNumber+30), message)
			}
		}(i, arg, output)
	}

	for line := range output {
		fmt.Println(line)
	}
}
