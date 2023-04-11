package main

import (
	"os"
	"testing"
	"time"
)

func TestWriteWAVFile(t *testing.T) {
	filename := "test.wav"
	duration := 1 * time.Second
	amplitude := 0.8
	sampleRate := 44100

	// Generate white noise
	audioData := GenerateWhiteNoise(duration, amplitude, sampleRate)

	// Write WAV file
	err := WriteWAVFile(filename, audioData, sampleRate)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filename)

	// Read WAV file
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	// Check if file contains any audio data at all
	if len(data) < 44 { // WAV header is 44 bytes
		t.Errorf("WAV file %s contains no audio data", filename)
	}
}
