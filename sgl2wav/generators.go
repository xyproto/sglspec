package main

import (
	"math"
	"math/rand"
	"time"
)

type PinkNoise struct {
	amplitude     float64
	values        []float64
	probabilities []float64
}

func NewPinkNoise(amplitude float64) *PinkNoise {
	numRows := 12
	values := make([]float64, numRows)
	probabilities := make([]float64, numRows)

	for i := 0; i < numRows; i++ {
		values[i] = rand.Float64()*2 - 1
		probabilities[i] = math.Pow(2, -float64(i))
	}

	return &PinkNoise{
		amplitude:     amplitude,
		values:        values,
		probabilities: probabilities,
	}
}

func (pn *PinkNoise) Next() float64 {
	var sum float64

	for i := 0; i < len(pn.values); i++ {
		if rand.Float64() < pn.probabilities[i] {
			pn.values[i] = rand.Float64()*2 - 1
		}
		sum += pn.values[i]
	}

	return pn.amplitude * (sum / float64(len(pn.values)))
}

type BrownianNoise struct {
	amplitude float64
	value     float64
	stepSize  float64
}

func NewBrownianNoise(amplitude float64) *BrownianNoise {
	return &BrownianNoise{
		amplitude: amplitude,
		value:     0.0,
		stepSize:  0.01, // You can adjust the step size to control the smoothness of the brownian noise
	}
}

func (bn *BrownianNoise) Next() float64 {
	step := (rand.Float64()*2 - 1) * bn.stepSize
	bn.value += step

	if bn.value > 1 {
		bn.value = 1
	} else if bn.value < -1 {
		bn.value = -1
	}

	return bn.amplitude * bn.value
}

func GenerateSawtoothWave(frequency float64, duration time.Duration, amplitude float64, sampleRate int) []float64 {
	numSamples := int(float64(duration) / float64(time.Second) * float64(sampleRate))
	data := make([]float64, numSamples)
	period := float64(sampleRate) / frequency
	for i := range data {
		data[i] = amplitude * (2 * (float64(i)/period - math.Floor(0.5+float64(i)/period)))
	}
	return data
}

// GenerateWhiteNoise generates white noise with the given length, amplitude, and sample rate.
func GenerateWhiteNoise(duration time.Duration, amplitude float64, sampleRate int) []float64 {
	numSamples := int(float64(duration) / float64(time.Second) * float64(sampleRate))
	data := make([]float64, numSamples)
	for i := range data {
		data[i] = amplitude * (rand.Float64()*2 - 1)
	}
	return data
}

// GeneratePinkNoise generates pink noise with the given duration, amplitude, and sample rate.
func GeneratePinkNoise(duration time.Duration, amplitude float64, sampleRate int) []float64 {
	numSamples := int(float64(duration) / float64(time.Second) * float64(sampleRate))
	data := make([]float64, numSamples)
	pink := NewPinkNoise(amplitude)
	for i := range data {
		data[i] = pink.Next()
	}
	return data
}

// GenerateBrownianNoise generates brownian noise with the given duration, amplitude, and sample rate.
func GenerateBrownianNoise(duration time.Duration, amplitude float64, sampleRate int) []float64 {
	numSamples := int(float64(duration) / float64(time.Second) * float64(sampleRate))
	data := make([]float64, numSamples)
	brown := NewBrownianNoise(amplitude)
	for i := range data {
		data[i] = brown.Next()
	}
	return data
}
