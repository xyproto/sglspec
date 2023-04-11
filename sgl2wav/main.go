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

	fmt.Printf("Parsing input file: %s\n", *inputFilename)

	// Create the ANTLR input stream
	charStream, err := antlr.NewFileStream(*inputFilename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Create the lexer
	lexer := sglantlr.NewSampleGenerationLanguageLexer(charStream)

	// Create the token stream
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the parser
	parser := sglantlr.NewSampleGenerationLanguageParser(tokenStream)

	// Parse the input file
	parseTree := parser.Prog()

	// Create the custom listener
	listener := NewSGLCustomListener()

	// Walk the parse tree with the custom listener
	antlr.ParseTreeWalkerDefault.Walk(listener, parseTree)

	// Generate audio data
	audioData, err := GenerateAudioData(listener, *sampleRate)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generating audio with sample rate: %d\n", *sampleRate)

	// Write audio data to WAV file
	if err := WriteWAVFile(*outputFilename, audioData, *sampleRate); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully generated audio and saved to", *outputFilename)
}
