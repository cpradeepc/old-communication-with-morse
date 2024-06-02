package main

import (
	mc "app/morsechar"
	"bufio"
	"fmt"
	"log"
	"os"
)

//var str string

func main() {
	fmt.Println("write here any string.")
	reader := bufio.NewReader(os.Stdin)
	str, errRead := reader.ReadString('\n')
	if errRead != nil {
		log.Fatalln("errRead = ", errRead)
	}

	fmt.Println("typing chars= ", str)
	myCode, err := mc.StrToMorseStr(str)
	if err != nil {
		log.Fatalln("err is : ", err.Error())
	}
	fmt.Println("mycode = ", myCode)
}
