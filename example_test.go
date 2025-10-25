package sentencepiece_test

import (
	"fmt"
	"log"
	"os"

	"github.com/SpringMT/go-sentencepiece"
)

func ExampleProcessor_Encode() {
	protoFile := os.Getenv("MODELPATH")
	if protoFile == "" {
		log.Println("Need MODELPATH env var to run example")
		return
	}

	proc, err := sentencepiece.NewProcessorFromPath(protoFile)
	if err != nil {
		log.Fatal(err)
	}

	text := "Encoding produces tokens that LLMs can learn and understand"
	tokens := proc.Encode(text)

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func ExampleProcessor_Decode() {
	protoFile := os.Getenv("MODELPATH")
	if protoFile == "" {
		log.Println("Need MODELPATH env var to run example")
		return
	}

	proc, err := sentencepiece.NewProcessorFromPath(protoFile)
	if err != nil {
		log.Fatal(err)
	}

	ids := []int{17534, 2134}
	text := proc.Decode(ids)

	fmt.Println(text)
}
