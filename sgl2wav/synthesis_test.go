package main

import (
	"math"
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

	if len(data) != sampleRate/10 {
		t.Errorf("Incorrect length of pulse wave data: got %v, expected %v", len(data), sampleRate/10)
	}

	positiveCount := 0
	for _, sample := range data {
		if sample == amplitude {
			positiveCount++
		} else if sample != -amplitude {
			t.Errorf("Unexpected pulse wave value: got %v, expected %v or %v", sample, amplitude, -amplitude)
		}
	}

	expectedPositiveCount := int(float64(sampleRate) * pulseWidth / 100 / 10)
	if positiveCount != expectedPositiveCount {
		t.Errorf("Unexpected number of positive samples in pulse wave: got %v, expected %v", positiveCount, expectedPositiveCount)
	}
}

func TestGenerateModulatedSineWave(t *testing.T) {
	frequency := 440.0
	duration := time.Millisecond * 100
	amplitude := 1.0
	modulationFrequency := 5.0
	modulationAmplitude := 10.0
	sampleRate := 44100

	data := GenerateModulatedSineWave(frequency, duration, amplitude, modulationFrequency, modulationAmplitude, sampleRate)

	if len(data) != sampleRate/10 {
		t.Errorf("Incorrect length of modulated sine wave data: got %v, expected %v", len(data), sampleRate/10)
	}

	for _, sample := range data {
		if math.Abs(sample) > amplitude {
			t.Errorf("Modulated sine wave value exceeds expected amplitude: got %v, max expected %v", sample, amplitude)
		}
	}
}

func TestGenerateTriangleWave(t *testing.T) {
	frequency := 440.0
	duration := time.Millisecond * 100
	amplitude := 1.0
	sampleRate := 44100

	data := GenerateTriangleWave(frequency, duration, amplitude, sampleRate)

	if len(data) != sampleRate/10 {
		t.Errorf("Incorrect length of triangle wave data: got %v, expected %v", len(data), sampleRate/10)
	}

	for i := 1; i < len(data)-1; i++ {
		if data[i] > amplitude || data[i] < -amplitude {
			t.Errorf("Triangle wave value exceeds expected amplitude: got %v, expected between %v and %v", data[i], -amplitude, amplitude)
		}
		if (data[i]-data[i-1])*(data[i]-data[i+1]) <= 0 {
			t.Errorf("Triangle wave not strictly increasing or decreasing at sample %v", i)
		}
	}
}
