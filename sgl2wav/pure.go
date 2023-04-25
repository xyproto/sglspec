package main

import (
	"math"
	"time"
)

// GenerateOscillator generates an oscillator with the specified waveform, frequency, duration, and sample rate.
func GenerateOscillator(waveform string, frequency float64, duration time.Duration, sampleRate int) []float64 {
	samples := int(float64(duration) / float64(time.Second) * float64(sampleRate))
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
func ApplyAmplitudeEnvelope(data []float64, attackTime, decayTime time.Duration, sustainLevel float64, releaseTime time.Duration, sampleRate int) []float64 {
	attackSamples := int(float64(attackTime) / float64(time.Second) * float64(sampleRate))
	decaySamples := int(float64(decayTime) / float64(time.Second) * float64(sampleRate))
	releaseSamples := int(float64(releaseTime) / float64(time.Second) * float64(sampleRate))

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
	alpha := 1 / (1 + (1 / (2 * math.Pi * cutoff / float64(sampleRate))))
	filteredData := make([]float64, len(data))
	filteredData[0] = data[0]

	for i := 1; i < len(data); i++ {
		filteredData[i] = alpha * (filteredData[i-1] + data[i] - data[i-1])
	}

	return filteredData
}

// AudioDistortion applies distortion effect to the provided data.
func AudioDistortion(data []float64, drive float64, tone float64, mix float64) []float64 {
	distortedData := make([]float64, len(data))

	for i := 0; i < len(data); i++ {
		distortedData[i] = (1 - mix) * data[i]
		distortedData[i] += mix * (2 / math.Pi * math.Atan(drive*data[i]))
	}

	// Apply tone control (low-pass filter)
	distortedData = LowPassFilter(distortedData, tone*20000, 44100)

	return distortedData
}

// AudioReverb applies reverb effect to the provided data.
func AudioReverb(data []float64, mix float64, decayTime time.Duration, sampleRate int) []float64 {
	// Comb filter constants
	combCount := 8
	combDelay := []int{1617, 1557, 1491, 1422, 1277, 1116, 897, 778}
	combGain := make([]float64, combCount)
	for i := 0; i < combCount; i++ {
		combGain[i] = math.Pow(10, -3*float64(combDelay[i])/float64(sampleRate)/(decayTime.Seconds()))
	}
	// All-pass filter constants
	allPassCount := 4
	allPassDelay := []int{225, 556, 441, 341}
	allPassGain := 0.5
	// Apply comb filters
	combData := make([][]float64, combCount)
	for i := 0; i < combCount; i++ {
		combData[i] = make([]float64, len(data)+combDelay[i])
		for j := 0; j < len(data); j++ {
			combData[i][j+combDelay[i]] = data[j] + combGain[i]*combData[i][j]
		}
	}
	// Mix comb filters
	mixedData := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		for j := 0; j < combCount; j++ {
			mixedData[i] += combData[j][i+combDelay[j]]
		}
	}
	// Apply all-pass filters
	allPassData := make([][]float64, allPassCount)
	for i := 0; i < allPassCount; i++ {
		allPassData[i] = make([]float64, len(data)+allPassDelay[i])
		for j := 0; j < len(data); j++ {
			allPassData[i][j+allPassDelay[i]] = mixedData[j] + allPassGain*allPassData[i][j]
			mixedData[j] = allPassData[i][j] - allPassGain*mixedData[j]
		}
	}
	// Apply mix
	for i := 0; i < len(data); i++ {
		data[i] = (1-mix)*data[i] + mix*mixedData[i]
	}
	return data
}

// BandPassFilter applies a band-pass filter to the provided data with the specified low and high cutoff frequencies.
func BandPassFilter(data []float64, lowCutoff float64, highCutoff float64, sampleRate int) []float64 {
	lowPassedData := LowPassFilter(data, highCutoff, sampleRate)
	bandPassedData := HighPassFilter(lowPassedData, lowCutoff, sampleRate)
	return bandPassedData
}

// AudioChorus adds a chorus effect to the provided data with the specified rate, depth, and mix.
func AudioChorus(data []float64, rate float64, depth float64, mix float64, sampleRate int) []float64 {
	// This is a simple implementation of a chorus effect. You may need to adjust the parameters for your specific use case.
	chData := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		offset := int(depth * math.Sin(2*math.Pi*rate*float64(i)/float64(sampleRate)))
		if i+offset < len(data) {
			chData[i] = data[i]*(1-mix) + data[i+offset]*mix
		} else {
			chData[i] = data[i]
		}
	}
	return chData
}

// AudioPhaser adds a phaser effect to the provided data with the specified rate, depth, feedback, and mix.
func AudioPhaser(data []float64, rate float64, depth float64, feedback float64, mix float64, sampleRate int) []float64 {
	// This is a simple implementation of a phaser effect. You may need to adjust the parameters for your specific use case.
	phData := make([]float64, len(data))
	feedbackData := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		offset := int(depth * math.Sin(2*math.Pi*rate*float64(i)/float64(sampleRate)))
		if i+offset < len(data) {
			phData[i] = data[i]*(1-mix) + (data[i+offset]+feedbackData[i])*mix
			feedbackData[i] = phData[i] * feedback
		} else {
			phData[i] = data[i]
		}
	}
	return phData
}

// AudioFlanger adds a flanger effect to the provided data with the specified rate, depth, feedback, and mix.
func AudioFlanger(data []float64, rate float64, depth float64, feedback float64, mix float64, sampleRate int) []float64 {
	// This is a simple implementation of a flanger effect. You may need to adjust the parameters for your specific use case.
	flData := make([]float64, len(data))
	feedbackData := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		offset := int(depth * math.Sin(2*math.Pi*rate*float64(i)/float64(sampleRate)))
		if i+offset < len(data) {
			flData[i] = data[i]*(1-mix) + (data[i+offset]+feedbackData[i])*mix
			feedbackData[i] = flData[i] * feedback
		} else {
			flData[i] = data[i]
		}
	}
	return flData
}

// AudioDelay adds a delay effect to the provided data with the specified delay time, feedback, and mix.
func AudioDelay(data []float64, delayTime time.Duration, feedback, mix float64, sampleRate int) []float64 {
	delayData := make([]float64, len(data))
	delaySamples := int(delayTime.Seconds() * float64(sampleRate))
	for i := 0; i < len(data); i++ {
		if i-delaySamples >= 0 {
			delayData[i] = data[i]*(1-mix) + (data[i-delaySamples]+delayData[i-delaySamples]*feedback)*mix
		} else {
			delayData[i] = data[i]
		}
	}
	return delayData
}

// Normalize normalizes the provided data to the specified amplitude.
func Normalize(data []float64, amplitude float64) []float64 {
	normalizedData := make([]float64, len(data))
	max := 0.0

	for _, sample := range data {
		if math.Abs(sample) > max {
			max = math.Abs(sample)
		}
	}

	if max > 0 {
		scale := amplitude / max
		for i := range data {
			normalizedData[i] = data[i] * scale
		}
	}

	return normalizedData
}

// GenerateSineWave generates a sine wave of the specified frequency, duration, and amplitude.
func GenerateSineWave(frequency float64, duration time.Duration, amplitude float64, sampleRate int) []float64 {
	samples := int(float64(duration) / float64(time.Second) * float64(sampleRate))
	data := make([]float64, samples)
	angularFrequency := 2 * math.Pi * frequency

	for i := 0; i < samples; i++ {
		data[i] = amplitude * math.Sin(angularFrequency*float64(i)/float64(sampleRate))
	}

	return data
}

// GenerateFilteredSineWave generates a sine wave with a low-pass or high-pass filter applied.
func GenerateFilteredSineWave(frequency float64, duration time.Duration, amplitude float64, filterType string, cutoff float64, sampleRate int) []float64 {
	numSamples := int(duration.Seconds() * float64(sampleRate))
	sineWave := make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate)
		sineWave[i] = amplitude * math.Sin(2*math.Pi*frequency*t)
	}

	switch filterType {
	case "Low-pass":
		return LowPassFilter(sineWave, cutoff, sampleRate)
	case "High-pass":
		return HighPassFilter(sineWave, cutoff, sampleRate)
	default:
		return sineWave
	}
}

// GenerateDistortedSineWave generates a sine wave with a distortion effect applied.
func GenerateDistortedSineWave(frequency float64, duration time.Duration, amplitude float64, drive float64, tone float64, mix float64, sampleRate int) []float64 {
	sineWave := GenerateSineWave(frequency, duration, amplitude, sampleRate)
	distortedSineWave := make([]float64, len(sineWave))

	for i, sample := range sineWave {
		distortedSineWave[i] = math.Tanh(sample*drive) * amplitude
	}

	toneFilteredWave := LowPassFilter(distortedSineWave, tone, sampleRate)

	return MixAudio(sineWave, toneFilteredWave, mix)
}

// MixAudio mixes two signals with the specified mix ratio.
func MixAudio(signal1, signal2 []float64, mix float64) []float64 {
	if len(signal1) != len(signal2) {
		panic("Input signals have different lengths")
	}

	mixedSignal := make([]float64, len(signal1))
	for i := 0; i < len(signal1); i++ {
		mixedSignal[i] = signal1[i]*(1-mix) + signal2[i]*mix
	}
	return mixedSignal
}
