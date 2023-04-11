package main

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	sglantlr "github.com/xyproto/sglspec/antlr"
)

// SGLCustomListener is a custom listener for the Sample Generation Language (SGL) that
type SGLCustomListener struct {
	sglantlr.BaseSampleGenerationLanguageListener
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
func (l *SGLCustomListener) EnterOscillator(ctx *sglantlr.OscillatorContext) {
	oscillator := OscillatorData{
		Waveform:  ctx.Waveform().GetText(),
		Frequency: parseFloat(ctx.Frequency().GetText()),
		Amplitude: parseFloat(ctx.Mix().GetText()),
	}
	l.oscillators = append(l.oscillators, oscillator)
}

// EnterEnvelope is called when an Envelope rule is entered in the SGL grammar.
func (l *SGLCustomListener) EnterEnvelope(ctx *sglantlr.EnvelopeContext) {
	attackTime := parseFloat(ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText())
	decayTime := parseFloat(ctx.GetChild(3).(*antlr.TerminalNodeImpl).GetText())
	sustainLevel := parseFloat(ctx.GetChild(5).(*antlr.TerminalNodeImpl).GetText())
	releaseTime := parseFloat(ctx.GetChild(7).(*antlr.TerminalNodeImpl).GetText())

	l.envelopes = append(l.envelopes, EnvelopeData{
		AttackTime:   attackTime,
		DecayTime:    decayTime,
		SustainLevel: sustainLevel,
		ReleaseTime:  releaseTime,
	})
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
