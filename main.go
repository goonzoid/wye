package main

import (
	"fmt"
	"os"

	"github.com/go-audio/wav"
)

func main() {
	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	decoder := wav.NewDecoder(inputFile)
	if !decoder.IsValidFile() {
		fmt.Println("file invalid - only wav files are supported")
		os.Exit(1)
	}

	format := decoder.Format()

	if format.NumChannels == 1 {
		fmt.Println("this file is mono")
		os.Exit(0)
	} else if format.NumChannels > 2 {
		fmt.Printf("this file has %d channels - i don't know how to handle that\n", format.NumChannels)
	} else {
		buf, err := decoder.FullPCMBuffer()
		if err != nil {
			panic(err)
		}

		frameCount := buf.NumFrames()
		for frame := 0; frame < frameCount; frame++ {
			index := frame * 2
			if buf.Data[index] != buf.Data[index+1] {
				fmt.Println("this file is stereo, and contains stereo information")
				os.Exit(0)
			}
		}
		fmt.Println("this file is stereo, but contains mono information")
	}
}
