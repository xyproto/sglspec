package main

import (
	"fmt"
	"time"
)

type SGLData struct {
	Oscillators    []OscillatorData
	Envelopes      []EnvelopeData
	Effects        []EffectData
	FMSynths       []FMSynthData
	GranularSynths []GranularSynthData
	PhysicalModels []PhysicalModelData
}

// Oscillator types
const (
	Sine = iota
	Square
	Triangle
	Sawtooth
	Noise
)

// Envelope types
const (
	Amplitude = iota
	Filter
)

// Effect types
const (
	EffectReverb = iota
	EffectDelay
	EffectChorus
	EffectDistortion
	EffectHighPass
	EffectLowPass
	EffectNoise
)

// FM types
const (
	Carrier = iota
	Modulator
)

// Granular window functions
const (
	Hanning = iota
	Hamming
	Blackman
	Triangular
)

// Physical model types
const (
	String = iota
	Membrane
	Tube
)

type Oscillator struct {
	Waveform  int
	Frequency float64
	Mix       float64
}

type Envelope struct {
	Type         int
	AttackTime   time.Duration
	DecayTime    time.Duration
	SustainLevel float64
	ReleaseTime  time.Duration
}

type Effect struct {
	Type   int
	Params map[string]float64
}

type FM struct {
	Type      int
	Waveform  int
	Frequency float64
	Mix       float64
}

type Granular struct {
	Size       float64
	Overlap    float64
	WindowFunc int
	Pitch      float64
}

type PhysicalModel struct {
	Type       int
	Params     map[string]float64
	Excitation int
	Mix        float64
}

type Additive struct {
	Frequency float64
	Amplitude float64
	Phase     float64
}

type Wavetable struct {
	Waveform int
	Index    float64
	Mix      float64
}

type Subtractive struct {
	Waveform   int
	Frequency  float64
	FilterType int
	Cutoff     float64
	Resonance  float64
}

type KarplusStrong struct {
	Frequency  float64
	Decay      float64
	FilterType int
	Cutoff     float64
}

type Formant struct {
	FilterType int
	Frequency  float64
	Bandwidth  float64
	Gain       float64
}

func GenerateAudioData(listener *SGLCustomListener, sampleRate int) ([]float64, error) {
	// Set the desired audio length in seconds
	audioLength := 5.0
	// Calculate the length in samples
	lengthInSamples := int(audioLength * float64(sampleRate))
	// Placeholder for the generated audio data
	audioData := make([]float64, lengthInSamples)

	// Process oscillators
	for i, osc := range listener.oscillators {
		oscData := GenerateOscillatorData(&osc, sampleRate)
		fmt.Printf("Oscillator %d data: %v\n", i, oscData[:10]) // Add this line

		// Ensure that both signals have the same length
		if len(oscData) < len(audioData) {
			oscData = append(oscData, make([]float64, len(audioData)-len(oscData))...)
		} else if len(oscData) > len(audioData) {
			audioData = append(audioData, make([]float64, len(oscData)-len(audioData))...)
		}

		audioData = MixAudio(audioData, oscData, osc.Mix)
	}
	// Process envelopes
	for _, env := range listener.envelopes {
		audioData = ApplyEnvelope(audioData, env, sampleRate)
	}
	// Process effects
	var err error
	for _, eff := range listener.effects {
		audioData, err = ApplyEffect(audioData, &eff, sampleRate)
		if err != nil {
			return []float64{}, err
		}
	}
	return audioData, nil
}

func GenerateOscillatorData(osc *OscillatorData, sampleRate int) []float64 {
	oscData := make([]float64, sampleRate)
	duration := time.Duration(float64(time.Second) / osc.Frequency)

	fmt.Printf("Oscillator waveform: %s, frequency: %f, duration: %s\n", osc.Waveform, osc.Frequency, duration)

	waveformIndex := -1
	switch osc.Waveform {
	case "Sine":
		waveformIndex = Sine
	case "Sawtooth":
		waveformIndex = Sawtooth
	case "Square":
		waveformIndex = Square
	case "Triangle":
		waveformIndex = Triangle
	case "Noise":
		waveformIndex = Noise
	}

	fmt.Printf("Waveform index: %d\n", waveformIndex)

	switch waveformIndex {
	case Sine:
		oscData = GenerateSineWave(osc.Frequency, duration, 1, sampleRate)
	case Sawtooth:
		oscData = GenerateSawtoothWave(osc.Frequency, duration, 1, sampleRate)
	case Square:
		oscData = GenerateSquareWave(osc.Frequency, duration, 1, sampleRate)
	case Triangle:
		oscData = GenerateTriangleWave(osc.Frequency, duration, 1, sampleRate)
	case Noise:
		oscData = GenerateNoise(osc.Waveform, duration, int(osc.Amplitude), sampleRate)
	}

	fmt.Printf("Generated oscillator data: %v\n", oscData[:10])
	return oscData
}

func GenerateNoise(noiseType string, duration time.Duration, amplitude, sampleRate int) []float64 {
	switch noiseType {
	case "white":
		return GenerateWhiteNoise(duration, float64(amplitude), sampleRate)
	case "pink":
		return GeneratePinkNoise(duration, float64(amplitude), sampleRate)
	case "brownian":
		return GenerateBrownianNoise(duration, float64(amplitude), sampleRate)
	default:
		panic("Unknown noise type: " + noiseType)
	}
}
