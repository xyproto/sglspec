package main

import (
	"testing"
	"time"
)

func TestGeneratePulseWave(t *testing.T) {
	frequency := 440.0
	duration := time.Millisecond * 100
	amplitude := 1.0
	pulseWidth := 25.0
	sampleRate := 44100
	data := GeneratePulseWave(frequency, duration, amplitude, pulseWidth, sampleRate)

	if len(data) == 0 {
		t.Errorf("Pulse wave has no data")
	}

	// Check if waveform is flatlining
	sum := 0.0
	for _, sample := range data {
		sum += sample
	}
	if sum/float64(len(data)) == 0 {
		t.Errorf("Pulse wave is flatlining")
	}
}

func TestGenerateModulatedSineWave(t *testing.T) {
	frequency := 440.0
	duration := time.Millisecond * 100
	amplitude := 0.8
	modulationFrequency := 5.0
	modulationAmplitude := 0.5
	sampleRate := 44100
	data := GenerateModulatedSineWave(frequency, duration, amplitude, modulationFrequency, modulationAmplitude, sampleRate)

	if len(data) == 0 {
		t.Errorf("Modulated sine wave has no data")
	}

	// Check if waveform is flatlining
	sum := 0.0
	for _, sample := range data {
		sum += sample
	}
	if sum/float64(len(data)) == 0 {
		t.Errorf("Modulated sine wave is flatlining")
	}
}

func TestGenerateTriangleWave(t *testing.T) {
	frequency := 440.0
	duration := time.Second
	amplitude := 0.8
	sampleRate := 44100
	samples := GenerateTriangleWave(frequency, duration, amplitude, sampleRate)

	if len(samples) == 0 {
		t.Errorf("Triangle wave has no data")
	}

	// Check if waveform is flatlining
	sum := 0.0
	for _, sample := range samples {
		sum += sample
	}
	if sum/float64(len(samples)) == 0 {
		t.Errorf("Triangle wave is flatlining")
	}
}
