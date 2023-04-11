package main

// ApplyEnvelope applies the amplitude envelope to the audio data.
func ApplyEnvelope(data []float64, env EnvelopeData, sampleRate int) []float64 {
	// Call the ApplyAmplitudeEnvelope function with the envelope data
	return ApplyAmplitudeEnvelope(data, env.AttackTime, env.DecayTime, env.SustainLevel, env.ReleaseTime, sampleRate)
}

func ApplyEffect(data []float64, effect *EffectData, sampleRate int) []float64 {
	duration := float64(len(data)) / float64(sampleRate)

	switch effect.Type {
	case "EffectReverb":
		return AudioReverb(data, effect.Mix, effect.Param1, sampleRate)
	case "EffectDelay":
		return AudioDelay(data, effect.Param1, effect.Param2, effect.Mix, sampleRate)
	case "EffectChorus":
		return AudioChorus(data, effect.Param1, effect.Param2, effect.Mix, sampleRate)
	case "EffectDistortion":
		return AudioDistortion(data, effect.Param1, effect.Param2, effect.Mix)
	case "EffectHighPass":
		return HighPassFilter(data, effect.Param1, sampleRate)
	case "EffectLowPass":
		return LowPassFilter(data, effect.Param1, sampleRate)
	case "EffectNoise":
		noise := GenerateWhiteNoise(duration, effect.Param1, sampleRate)
		return MixAudio(data, noise, effect.Mix)
	default:
		return data
	}
}

func GenerateSquareWave(frequency float64, duration float64, amplitude float64, sampleRate int) []float64 {
	return GenerateOscillator("square", frequency, duration, sampleRate)
}

func ApplyOscillator(osc *OscillatorData, sampleRate int) []float64 {
	return GenerateOscillator(osc.Waveform, osc.Frequency, osc.Duration, sampleRate)
}
