package main

import (
	"testing"
	"time"
)

func TestGenerateAudioData(t *testing.T) {
	// Create a listener with a sine wave oscillator and amplitude envelope
	listener := &SGLCustomListener{
		oscillators: []OscillatorData{
			{
				Waveform:  "sine",
				Frequency: 440,
				Mix:       1,
			},
		},
		envelopes: []EnvelopeData{
			{
				AttackTime:   time.Millisecond * 10,
				DecayTime:    time.Millisecond * 10,
				SustainLevel: 0.5,
				ReleaseTime:  time.Millisecond * 10,
			},
		},
		effects: []EffectData{},
	}
	// Generate audio data
	sampleRate := 44100
	audioData, err := GenerateAudioData(listener, sampleRate)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	// Check that audio data has been generated and is not all zeroes
	hasData := false
	for _, sample := range audioData {
		if sample != 0 {
			hasData = true
			break
		}
	}
	if !hasData {
		t.Error("Generated audio data is all zeroes")
	}
}
