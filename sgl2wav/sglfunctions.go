package main

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
	AttackTime   float64
	DecayTime    float64
	SustainLevel float64
	ReleaseTime  float64
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

func GenerateAudioData(listener *SGLCustomListener, sampleRate int) []float64 {
	// Placeholder for the generated audio data
	audioData := make([]float64, sampleRate)

	// Process oscillators
	for _, osc := range listener.oscillators {
		oscData := GenerateOscillatorData(&osc, sampleRate)
		audioData = MixAudio(audioData, oscData, osc.Mix)
	}

	// Process envelopes
	for _, env := range listener.envelopes {
		audioData = ApplyEnvelope(audioData, env, sampleRate)
	}

	// Process effects
	for _, eff := range listener.effects {
		audioData = ApplyEffect(audioData, &eff, sampleRate)
	}

	return audioData
}

func GenerateOscillatorData(osc *OscillatorData, sampleRate int) []float64 {
	oscData := make([]float64, sampleRate)
	duration := float64(sampleRate) / osc.Frequency

	switch osc.Waveform {
	case "sine":
		oscData = GenerateSineWave(osc.Frequency, duration, 1, sampleRate)
	case "sawtooth":
		oscData = GenerateSawtoothWave(osc.Frequency, duration, 1, sampleRate)
	case "square":
		oscData = GenerateSquareWave(osc.Frequency, duration, 1, sampleRate)
	case "riangle":
		oscData = GenerateTriangleWave(osc.Frequency, duration, 1, sampleRate)
	case "noise":
		oscData = GenerateNoise(osc.Waveform, duration, osc.Amplitude, sampleRate)
	}
	return oscData
}

func GenerateNoise(noiseType string, duration float64, amplitude float64, sampleRate int) []float64 {
	switch noiseType {
	case "white":
		return GenerateWhiteNoise(duration, amplitude, sampleRate)
	case "pink":
		return GeneratePinkNoise(duration, amplitude, sampleRate)
	case "brownian":
		return GenerateBrownianNoise(duration, amplitude, sampleRate)
	default:
		panic("Unknown noise type: " + noiseType)
	}
}