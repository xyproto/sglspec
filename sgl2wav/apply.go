package main

import (
	"fmt"
	"time"
)

// ApplyEnvelope applies the amplitude envelope to the audio data.
func ApplyEnvelope(data []float64, env EnvelopeData, sampleRate int) []float64 {
	// Call the ApplyAmplitudeEnvelope function with the envelope data
	return ApplyAmplitudeEnvelope(data, env.AttackTime, env.DecayTime, env.SustainLevel, env.ReleaseTime, sampleRate)
}

func ApplyEffect(data []float64, effect *EffectData, sampleRate int) ([]float64, error) {
	var duration time.Duration = time.Duration(float64(len(data)) / float64(sampleRate) * float64(time.Second))
	switch effect.Type {
	case "EffectReverb":
		param, ok := effect.Param1.(time.Duration)
		if !ok {
			return nil, fmt.Errorf("effect param1 is not a duration")
		}
		return AudioReverb(data, effect.Mix, param, sampleRate), nil
	case "EffectDelay":
		param1, ok1 := effect.Param1.(time.Duration)
		param2, ok2 := effect.Param2.(float64)
		if !ok1 {
			return nil, fmt.Errorf("effect param1 is not a duration")
		}
		if !ok2 {
			return nil, fmt.Errorf("effect param2 is not a float64")
		}
		return AudioDelay(data, param1, param2, effect.Mix, sampleRate), nil
	case "EffectChorus":
		param1, ok1 := effect.Param1.(float64)
		param2, ok2 := effect.Param2.(float64)
		if !ok1 {
			return nil, fmt.Errorf("effect param1 is not a float64")
		}
		if !ok2 {
			return nil, fmt.Errorf("effect param2 is not a float64")
		}
		return AudioChorus(data, param1, param2, effect.Mix, sampleRate), nil
	case "EffectDistortion":
		param1, ok1 := effect.Param1.(float64)
		param2, ok2 := effect.Param2.(float64)
		if !ok1 {
			return nil, fmt.Errorf("effect param1 is not a float64")
		}
		if !ok2 {
			return nil, fmt.Errorf("effect param2 is not a float64")
		}
		return AudioDistortion(data, param1, param2, effect.Mix), nil
	case "EffectHighPass":
		param1, ok1 := effect.Param1.(float64)
		if !ok1 {
			return nil, fmt.Errorf("effect param1 is not a float64")
		}
		return HighPassFilter(data, param1, sampleRate), nil
	case "EffectLowPass":
		param1, ok1 := effect.Param1.(float64)
		if !ok1 {
			return nil, fmt.Errorf("effect param1 is not a float64")
		}
		return LowPassFilter(data, param1, sampleRate), nil
	case "EffectNoise":
		param1, ok1 := effect.Param1.(float64)
		if !ok1 {
			return nil, fmt.Errorf("effect param1 is not a float64")
		}
		noise := GenerateWhiteNoise(duration, param1, sampleRate)
		return MixAudio(data, noise, effect.Mix), nil
	default:
		return data, nil
	}
}

func GenerateSquareWave(frequency float64, duration time.Duration, amplitude float64, sampleRate int) []float64 {
	return GenerateOscillator("square", frequency, duration, sampleRate)
}

func ApplyOscillator(osc *OscillatorData, sampleRate int) []float64 {
	return GenerateOscillator(osc.Waveform, osc.Frequency, osc.Duration, sampleRate)
}
