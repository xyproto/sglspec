// Code generated from SampleGenerationLanguage.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // SampleGenerationLanguage
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by SampleGenerationLanguageParser.
type SampleGenerationLanguageVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SampleGenerationLanguageParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#header.
	VisitHeader(ctx *HeaderContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#oscillator.
	VisitOscillator(ctx *OscillatorContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#envelope.
	VisitEnvelope(ctx *EnvelopeContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#effect.
	VisitEffect(ctx *EffectContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#fm.
	VisitFm(ctx *FmContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#granular.
	VisitGranular(ctx *GranularContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#physicalModel.
	VisitPhysicalModel(ctx *PhysicalModelContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#additive.
	VisitAdditive(ctx *AdditiveContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#wavetable.
	VisitWavetable(ctx *WavetableContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#subtractive.
	VisitSubtractive(ctx *SubtractiveContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#karplusStrong.
	VisitKarplusStrong(ctx *KarplusStrongContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#formant.
	VisitFormant(ctx *FormantContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#waveform.
	VisitWaveform(ctx *WaveformContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#frequency.
	VisitFrequency(ctx *FrequencyContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#mix.
	VisitMix(ctx *MixContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#envType.
	VisitEnvType(ctx *EnvTypeContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#attackTime.
	VisitAttackTime(ctx *AttackTimeContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#decayTime.
	VisitDecayTime(ctx *DecayTimeContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#sustainLevel.
	VisitSustainLevel(ctx *SustainLevelContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#releaseTime.
	VisitReleaseTime(ctx *ReleaseTimeContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#effectType.
	VisitEffectType(ctx *EffectTypeContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#effectParam.
	VisitEffectParam(ctx *EffectParamContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#effectReverb.
	VisitEffectReverb(ctx *EffectReverbContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#effectDelay.
	VisitEffectDelay(ctx *EffectDelayContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#effectChorus.
	VisitEffectChorus(ctx *EffectChorusContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#effectDistortion.
	VisitEffectDistortion(ctx *EffectDistortionContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#effectFilter.
	VisitEffectFilter(ctx *EffectFilterContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#fmType.
	VisitFmType(ctx *FmTypeContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#size.
	VisitSize(ctx *SizeContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#overlap.
	VisitOverlap(ctx *OverlapContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#windowFunction.
	VisitWindowFunction(ctx *WindowFunctionContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#pitch.
	VisitPitch(ctx *PitchContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#modelType.
	VisitModelType(ctx *ModelTypeContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#modelParam.
	VisitModelParam(ctx *ModelParamContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#excitation.
	VisitExcitation(ctx *ExcitationContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#amplitude.
	VisitAmplitude(ctx *AmplitudeContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#phase.
	VisitPhase(ctx *PhaseContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#filterType.
	VisitFilterType(ctx *FilterTypeContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#cutoff.
	VisitCutoff(ctx *CutoffContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#resonance.
	VisitResonance(ctx *ResonanceContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#decay.
	VisitDecay(ctx *DecayContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#bandwidth.
	VisitBandwidth(ctx *BandwidthContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#gain.
	VisitGain(ctx *GainContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#delayTime.
	VisitDelayTime(ctx *DelayTimeContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#feedback.
	VisitFeedback(ctx *FeedbackContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#rate.
	VisitRate(ctx *RateContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#depth.
	VisitDepth(ctx *DepthContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#drive.
	VisitDrive(ctx *DriveContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#tone.
	VisitTone(ctx *ToneContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#stiffness.
	VisitStiffness(ctx *StiffnessContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#damping.
	VisitDamping(ctx *DampingContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#tension.
	VisitTension(ctx *TensionContext) interface{}

	// Visit a parse tree produced by SampleGenerationLanguageParser#length.
	VisitLength(ctx *LengthContext) interface{}
}
