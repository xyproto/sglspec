package main

import (
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func floatBufferToIntBuffer(floatBuffer []float64) *audio.IntBuffer {
	intBuffer := &audio.IntBuffer{
		Format: &audio.Format{
			NumChannels: 1,
			SampleRate:  44100,
		},
		Data: make([]int, len(floatBuffer)),
	}

	for i, sample := range floatBuffer {
		intBuffer.Data[i] = int(sample * 32767)
	}

	return intBuffer
}

func WriteWAVFile(filename string, audioData []float64, sampleRate int) error {
	// Create output file
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Initialize WAV encoder
	encoder := wav.NewEncoder(outFile, sampleRate, 16, 1, 1)

	// Convert float64 audio data to int16 in an IntBuffer
	intBuffer := floatBufferToIntBuffer(audioData)

	// Write audio data to the WAV file
	err = encoder.Write(intBuffer)
	if err != nil {
		return err
	}

	// Close the WAV encoder
	err = encoder.Close()
	if err != nil {
		return err
	}

	return nil
}
