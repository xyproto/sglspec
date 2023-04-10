package main

import (
	"math"
)

// GenerateOscillator generates an oscillator with the specified waveform, frequency, duration, and sample rate.
func GenerateOscillator(waveform string, frequency float64, duration float64, sampleRate int) []float64 {
	samples := int(duration * float64(sampleRate))
	data := make([]float64, samples)

	switch waveform {
	case "Sine":
		for i := 0; i < samples; i++ {
			data[i] = math.Sin(2 * math.Pi * frequency * float64(i) / float64(sampleRate))
		}
	case "Square":
		for i := 0; i < samples; i++ {
			data[i] = math.Copysign(1, math.Sin(2*math.Pi*frequency*float64(i)/float64(sampleRate)))
		}
	case "Triangle":
		period := float64(sampleRate) / frequency
		for i := 0; i < samples; i++ {
			data[i] = 4*math.Abs(math.Mod(float64(i)/period, 1)-0.5) - 1
		}
	case "Sawtooth":
		period := float64(sampleRate) / frequency
		for i := 0; i < samples; i++ {
			data[i] = 2*(float64(i)/period-math.Floor(0.5+float64(i)/period)) - 1
		}
	}

	return data
}

// ApplyAmplitudeEnvelope applies an amplitude envelope to the provided data.
func ApplyAmplitudeEnvelope(data []float64, attackTime float64, decayTime float64, sustainLevel float64, releaseTime float64, sampleRate int) []float64 {
	attackSamples := int(attackTime * float64(sampleRate))
	decaySamples := int(decayTime * float64(sampleRate))
	releaseSamples := int(releaseTime * float64(sampleRate))

	// Apply attack
	for i := 0; i < attackSamples; i++ {
		data[i] *= float64(i) / float64(attackSamples)
	}

	// Apply decay
	for i := attackSamples; i < attackSamples+decaySamples; i++ {
		data[i] *= sustainLevel + (1-sustainLevel)*(1-float64(i-attackSamples)/float64(decaySamples))
	}

	// Apply release
	for i := len(data) - releaseSamples; i < len(data); i++ {
		data[i] *= sustainLevel * (1 - float64(i-len(data)+releaseSamples)/float64(releaseSamples))
	}

	return data
}

// LowPassFilter applies a low-pass filter to the provided data.
func LowPassFilter(data []float64, cutoff float64, sampleRate int) []float64 {
	alpha := math.Exp(-2 * math.Pi * cutoff / float64(sampleRate))
	filteredData := make([]float64, len(data))
	filteredData[0] = data[0]

	for i := 1; i < len(data); i++ {
		filteredData[i] = alpha*filteredData[i-1] + (1-alpha)*data[i]
	}

	return filteredData
}

// HighPassFilter applies a high-pass filter to the provided data.
func HighPassFilter(data []float64, cutoff float64, sampleRate int) []float64 {
	// TO BE IMPLEMENTED
	return nil
}
