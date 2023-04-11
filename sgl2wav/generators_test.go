package main

import (
	"math"
	"testing"
	"time"
)

func TestGenerateSawtoothWave(t *testing.T) {
	frequency := 440.0
	duration := time.Millisecond * 100
	amplitude := 1.0
	sampleRate := 44100

	data := GenerateSawtoothWave(frequency, duration, amplitude, sampleRate)

	if len(data) != sampleRate/10 {
		t.Errorf("Incorrect length of sawtooth wave data: got %v, expected %v", len(data), sampleRate/10)
	}

	// Find two consecutive peaks in the waveform
	peak1 := -1
	peak2 := -1
	for i := 1; i < len(data)-1; i++ {
		if data[i] > data[i-1] && data[i] > data[i+1] {
			if peak1 == -1 {
				peak1 = i
			} else {
				peak2 = i
				break
			}
		}
	}

	if peak1 == -1 || peak2 == -1 {
		t.Error("Could not find two consecutive peaks in the sawtooth wave")
	} else {
		expectedDistance := int(float64(sampleRate) / frequency)
		distance := peak2 - peak1
		if distance != expectedDistance {
			t.Errorf("Unexpected distance between sawtooth wave peaks: got %v, expected %v", distance, expectedDistance)
		}
	}
}

// Test the white noise generation to make sure that the output values are within the expected amplitude range
func TestGenerateWhiteNoise(t *testing.T) {
	duration := time.Millisecond * 100
	amplitude := 1.0
	sampleRate := 44100

	data := GenerateWhiteNoise(duration, amplitude, sampleRate)

	if len(data) != sampleRate/10 {
		t.Errorf("Incorrect length of white noise data: got %v, expected %v", len(data), sampleRate/10)
	}

	for i, sample := range data {
		if math.Abs(sample) > amplitude {
			t.Errorf("White noise sample %v has amplitude outside of the expected range: %v", i, sample)
		}
	}
}

// Test the pink noise generation to make sure that the output values are within the expected amplitude range
func TestGeneratePinkNoise(t *testing.T) {
	duration := time.Millisecond * 100
	amplitude := 1.0
	sampleRate := 44100

	data := GeneratePinkNoise(duration, amplitude, sampleRate)

	if len(data) != sampleRate/10 {
		t.Errorf("Incorrect length of pink noise data: got %v, expected %v", len(data), sampleRate/10)
	}

	for i, sample := range data {
		if math.Abs(sample) > amplitude {
			t.Errorf("Pink noise sample %v has amplitude outside of the expected range: %v", i, sample)
		}
	}
}

// Test the brownian noise generation to make sure that the output values are within the expected amplitude range
func TestGenerateBrownianNoise(t *testing.T) {
	duration := time.Millisecond * 100
	amplitude := 1.0
	sampleRate := 44100

	data := GenerateBrownianNoise(duration, amplitude, sampleRate)

	if len(data) != sampleRate/10 {
		t.Errorf("Incorrect length of brownian noise data: got %v, expected %v", len(data), sampleRate/10)
	}

	for i, sample := range data {
		if math.Abs(sample) > amplitude {
			t.Errorf("Brownian noise sample %v has amplitude outside of the expected range: %v", i, sample)
		}
	}
}
