package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	sglantlr "github.com/xyproto/sglspec/antlr"
)

func main() {
	// Parse command line arguments
	inputFilename := flag.String("i", "samples/bass01.sgl", "Input .sgl file")
	outputFilename := flag.String("o", "audio/bass01.wav", "Output .wav file")
	sampleRate := flag.Int("sampleRate", 44100, "Sample rate (default: 44100)")

	flag.Parse()

	if *inputFilename == "" {
		fmt.Println("Error: Input file not specified.")
		flag.Usage()
		os.Exit(1)
	}

	charStream, _ := antlr.NewFileStream(*inputFilename)
	lexer := sglantlr.NewSampleGenerationLanguageLexer(charStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := sglantlr.NewSampleGenerationLanguageParser(stream)

	// Create the custom listener
	listener := NewSGLCustomListener()

	// Walk the parse tree with the custom listener
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Prog())

	// Generate audio data
	audioData := GenerateAudioData(listener, *sampleRate)

	// Write audio data to WAV file
	err := WriteWAVFile(*outputFilename, audioData, *sampleRate)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully generated audio and saved to", *outputFilename)
}
