package main

import (
	"fmt"
	"math"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	sglantlr "github.com/xyproto/sglspec/antlr"
)


func floatBufferToIntBuffer(buf *audio.FloatBuffer) *audio.IntBuffer {
	intData := make([]int, len(buf.Data))
	for i, v := range buf.Data {
		intData[i] = int(v * 32767)
	}

	return &audio.IntBuffer{
		Format: buf.Format,
		Data:   intData,
	}
}

func main() {
	inputFile, err := os.Open("./samples/bass01.sgl")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer inputFile.Close()

	charStream, _ := antlr.NewFileStream(inputFile.Name())
	lexer := sglantlr.NewSampleGenerationLanguageLexer(charStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := sglantlr.NewSampleGenerationLanguageParser(stream)

	_ = parser.Prog() // We're not using the tree variable, so I've replaced it with a blank identifier

	// Generate a simple sine wave example
	sampleRate := 44100
	frequency := 55.0
	duration := 1.0
	samples := int(duration * float64(sampleRate))

	buf := &audio.FloatBuffer{
		Format: &audio.Format{
			NumChannels: 1,
			SampleRate:  sampleRate,
		},
		Data: make([]float64, samples),
	}

	for i := 0; i < samples; i++ {
		buf.Data[i] = 0.5 * (1.0 + math.Sin(2*math.Pi*frequency*float64(i)/float64(sampleRate)))
	}

	// Convert the float buffer to an integer buffer and save as a .wav file
	intBuf := floatBufferToIntBuffer(buf)
	f, err := os.Create("audio/bass01.wav")
	if err != nil {
		fmt.Printf("Error creating .wav file: %v\n", err)
		return
	}
	defer f.Close()
	wavEncoder := wav.NewEncoder(f, sampleRate, 16, 1, 1)
	err = wavEncoder.Write(intBuf)
	if err != nil {
		fmt.Printf("Error saving .wav file: %v\n", err)
		return
	}
	wavEncoder.Close()

	fmt.Println("Generated audio/bass01.wav")
}
