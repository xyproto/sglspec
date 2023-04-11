package main

import (
	"math"
	"time"
)

func GeneratePulseWave(frequency float64, duration time.Duration, amplitude, pulseWidth float64, sampleRate int) []float64 {
	numSamples := int(float64(duration) * float64(sampleRate) / float64(time.Second))
	samples := make([]float64, numSamples)
	pulseWidthRatio := pulseWidth / 100.0

	for i := 0; i < numSamples; i++ {
		time := float64(i) / float64(sampleRate)
		positionInPeriod := math.Mod(time*frequency, 1)
		if positionInPeriod < pulseWidthRatio {
			samples[i] = amplitude
		} else {
			samples[i] = -amplitude
		}
	}

	return samples
}

func GenerateModulatedSineWave(frequency float64, duration time.Duration, amplitude, modulationFrequency, modulationAmplitude float64, sampleRate int) []float64 {
	numSamples := int(float64(duration) * float64(sampleRate) / float64(time.Second))
	samples := make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		time := float64(i) / float64(sampleRate)
		modulation := math.Sin(2*math.Pi*modulationFrequency*time) * modulationAmplitude
		samples[i] = amplitude * math.Sin(2*math.Pi*(frequency+modulation)*time)
	}

	return samples
}

func GenerateTriangleWave(frequency float64, duration time.Duration, amplitude float64, sampleRate int) []float64 {
	numSamples := int(float64(duration) * float64(sampleRate) / float64(time.Second))
	samples := make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		time := float64(i) / float64(sampleRate)
		positionInPeriod := math.Mod(time*frequency, 1)
		if positionInPeriod < 0.25 {
			samples[i] = 4 * positionInPeriod * amplitude
		} else if positionInPeriod < 0.75 {
			samples[i] = 2*amplitude*(0.5-positionInPeriod) + amplitude
		} else {
			samples[i] = 4 * (positionInPeriod - 1) * amplitude
		}
	}

	return samples
}
