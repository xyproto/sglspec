package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	parser "github.com/xyproto/sglspec/parser"
)

// SGLCustomListener is a custom listener for the Sample Generation Language (SGL) that
type SGLCustomListener struct {
	parser.BaseSampleGenerationLanguageListener
	oscillators         []OscillatorData
	envelopes           []EnvelopeData
	effects             []EffectData
	fmSynths            []FMSynthData
	granularSynths      []GranularSynthData
	physicalModels      []PhysicalModelData
	additiveSynths      []AdditiveSynthData
	wavetableSynths     []WavetableSynthData
	subtractiveSynths   []SubtractiveSynthData
	karplusStrongSynths []KarplusStrongSynthData
	formantSynths       []FormantSynthData
}

func NewSGLCustomListener() *SGLCustomListener {
	return &SGLCustomListener{}
}

// EnterOscillator is called when an Oscillator rule is entered in the SGL grammar.
func (l *SGLCustomListener) EnterOscillator(ctx *parser.OscillatorContext) {
	fmt.Println("OscillatorContext", *ctx)
	oscillator := OscillatorData{
		Waveform:  ctx.Waveform().GetText(),
		Frequency: parseFloat(ctx.Frequency().GetText()),
		Amplitude: parseFloat(ctx.Mix().GetText()),
		//PhaseShift: parseFloat(ctx.PhaseShift().GetText()),
		//Duration:   ctx.Duration().GetDuration(),
		Mix: parseFloat(ctx.Mix().GetText()),
	}
	fmt.Println("Oscillator", oscillator)
	l.oscillators = append(l.oscillators, oscillator)
}

// EnterEnvelope is called when an Envelope rule is entered in the SGL grammar.
func (l *SGLCustomListener) EnterEnvelope(ctx *parser.EnvelopeContext) {
	attackTime := parseDuration(ctx.AttackTime().GetText())
	decayTime := parseDuration(ctx.DecayTime().GetText())
	sustainLevel := parseFloat(ctx.SustainLevel().GetText())
	releaseTime := parseDuration(ctx.ReleaseTime().GetText())

	l.envelopes = append(l.envelopes, EnvelopeData{
		AttackTime:   attackTime,
		DecayTime:    decayTime,
		SustainLevel: sustainLevel,
		ReleaseTime:  releaseTime,
	})
}

// Helper function to parse time.Duration from strings
func parseDuration(s string) time.Duration {
	// Check if "ms" is present
	if strings.HasSuffix(s, "ms") {
		s = strings.TrimSuffix(s, "ms")
		ms, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		return time.Duration(ms) * time.Millisecond
	}

	// If no unit is specified, treat it as milliseconds by default
	ms, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Duration(ms) * time.Millisecond
}

// Helper function to parse floats from strings
func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func (l *SGLCustomListener) GetOscillatorData() []OscillatorData {
	return l.oscillators
}

func (l *SGLCustomListener) GetEnvelopeData() []EnvelopeData {
	return l.envelopes
}
