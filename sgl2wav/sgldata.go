package main

import "time"

type OscillatorData struct {
	Waveform   string
	Frequency  float64
	Amplitude  float64
	PhaseShift float64
	Duration   time.Duration
	Mix        float64
}

type EnvelopeData struct {
	AttackTime   time.Duration
	DecayTime    time.Duration
	SustainLevel float64
	ReleaseTime  time.Duration
}

type EffectData struct {
	Type   string
	Params []float64
	Mix    float64
	Param1 interface{}
	Param2 interface{}
}

type FMSynthData struct {
	Carrier   OscillatorData
	Modulator OscillatorData
	ModIndex  float64
	Amplitude float64
	Envelope  EnvelopeData
}

type GranularSynthData struct {
	SampleRate int
	GrainSize  float64
	GrainCount int
}

type PhysicalModelData struct {
	ModelType string
	Params    []float64
}

func NewOscillatorData(waveform string, frequency float64, amplitude float64, phaseShift float64) *OscillatorData {
	return &OscillatorData{
		Waveform:   waveform,
		Frequency:  frequency,
		Amplitude:  amplitude,
		PhaseShift: phaseShift,
	}
}

func NewEnvelopeData(attackTime, decayTime time.Duration, sustainLevel float64, releaseTime time.Duration) *EnvelopeData {
	return &EnvelopeData{
		AttackTime:   attackTime,
		DecayTime:    decayTime,
		SustainLevel: sustainLevel,
		ReleaseTime:  releaseTime,
	}
}

func NewEffectData(effectType string, params []float64) *EffectData {
	return &EffectData{
		Type:   effectType,
		Params: params,
	}
}

func NewFMSynthData(carrier, modulator OscillatorData, modIndex, amplitude float64, envelope EnvelopeData) *FMSynthData {
	return &FMSynthData{
		Carrier:   carrier,
		Modulator: modulator,
		ModIndex:  modIndex,
		Amplitude: amplitude,
		Envelope:  envelope,
	}
}

func NewGranularSynthData(sampleRate int, grainSize float64, grainCount int) *GranularSynthData {
	return &GranularSynthData{
		SampleRate: sampleRate,
		GrainSize:  grainSize,
		GrainCount: grainCount,
	}
}

func NewPhysicalModelData(modelType string, params []float64) *PhysicalModelData {
	return &PhysicalModelData{
		ModelType: modelType,
		Params:    params,
	}
}
