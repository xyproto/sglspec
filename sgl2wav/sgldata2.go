package main

type AdditiveSynthData struct {
	Oscillators []OscillatorData
	Envelope    EnvelopeData
}

type WavetableSynthData struct {
	Wavetable []float64
	Frequency float64
	Amplitude float64
	Envelope  EnvelopeData
}

type SubtractiveSynthData struct {
	Oscillator OscillatorData
	FilterType string
	Cutoff     float64
	Resonance  float64
	Envelope   EnvelopeData
}

type KarplusStrongSynthData struct {
	Frequency    float64
	DampingRatio float64
}

type FormantSynthData struct {
	FundamentalFrequency float64
	Formants             []float64
	Amplitude            float64
}

func NewAdditiveSynthData(oscillators []OscillatorData, envelope EnvelopeData) *AdditiveSynthData {
	return &AdditiveSynthData{
		Oscillators: oscillators,
		Envelope:    envelope,
	}
}

func NewWavetableSynthData(wavetable []float64, frequency float64, amplitude float64, envelope EnvelopeData) *WavetableSynthData {
	return &WavetableSynthData{
		Wavetable: wavetable,
		Frequency: frequency,
		Amplitude: amplitude,
		Envelope:  envelope,
	}
}

func NewSubtractiveSynthData(oscillator OscillatorData, filterType string, cutoff float64, resonance float64, envelope EnvelopeData) *SubtractiveSynthData {
	return &SubtractiveSynthData{
		Oscillator: oscillator,
		FilterType: filterType,
		Cutoff:     cutoff,
		Resonance:  resonance,
		Envelope:   envelope,
	}
}

func NewKarplusStrongSynthData(frequency float64, dampingRatio float64) *KarplusStrongSynthData {
	return &KarplusStrongSynthData{
		Frequency:    frequency,
		DampingRatio: dampingRatio,
	}
}

func NewFormantSynthData(fundamentalFrequency float64, formants []float64, amplitude float64) *FormantSynthData {
	return &FormantSynthData{
		FundamentalFrequency: fundamentalFrequency,
		Formants:             formants,
		Amplitude:            amplitude,
	}
}
