// Code generated from antlr/SampleGenerationLanguage.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // SampleGenerationLanguage

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// SampleGenerationLanguageListener is a complete listener for a parse tree produced by SampleGenerationLanguageParser.
type SampleGenerationLanguageListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterOscillator is called when entering the oscillator production.
	EnterOscillator(c *OscillatorContext)

	// EnterEnvelope is called when entering the envelope production.
	EnterEnvelope(c *EnvelopeContext)

	// EnterEffect is called when entering the effect production.
	EnterEffect(c *EffectContext)

	// EnterFm is called when entering the fm production.
	EnterFm(c *FmContext)

	// EnterGranular is called when entering the granular production.
	EnterGranular(c *GranularContext)

	// EnterPhysicalModel is called when entering the physicalModel production.
	EnterPhysicalModel(c *PhysicalModelContext)

	// EnterAdditive is called when entering the additive production.
	EnterAdditive(c *AdditiveContext)

	// EnterWavetable is called when entering the wavetable production.
	EnterWavetable(c *WavetableContext)

	// EnterSubtractive is called when entering the subtractive production.
	EnterSubtractive(c *SubtractiveContext)

	// EnterKarplusStrong is called when entering the karplusStrong production.
	EnterKarplusStrong(c *KarplusStrongContext)

	// EnterFormant is called when entering the formant production.
	EnterFormant(c *FormantContext)

	// EnterWaveform is called when entering the waveform production.
	EnterWaveform(c *WaveformContext)

	// EnterFrequency is called when entering the frequency production.
	EnterFrequency(c *FrequencyContext)

	// EnterMix is called when entering the mix production.
	EnterMix(c *MixContext)

	// EnterEnvType is called when entering the envType production.
	EnterEnvType(c *EnvTypeContext)

	// EnterAttackTime is called when entering the attackTime production.
	EnterAttackTime(c *AttackTimeContext)

	// EnterDecayTime is called when entering the decayTime production.
	EnterDecayTime(c *DecayTimeContext)

	// EnterSustainLevel is called when entering the sustainLevel production.
	EnterSustainLevel(c *SustainLevelContext)

	// EnterReleaseTime is called when entering the releaseTime production.
	EnterReleaseTime(c *ReleaseTimeContext)

	// EnterEffectType is called when entering the effectType production.
	EnterEffectType(c *EffectTypeContext)

	// EnterEffectParam is called when entering the effectParam production.
	EnterEffectParam(c *EffectParamContext)

	// EnterEffectReverb is called when entering the effectReverb production.
	EnterEffectReverb(c *EffectReverbContext)

	// EnterEffectDelay is called when entering the effectDelay production.
	EnterEffectDelay(c *EffectDelayContext)

	// EnterEffectChorus is called when entering the effectChorus production.
	EnterEffectChorus(c *EffectChorusContext)

	// EnterEffectDistortion is called when entering the effectDistortion production.
	EnterEffectDistortion(c *EffectDistortionContext)

	// EnterEffectFilter is called when entering the effectFilter production.
	EnterEffectFilter(c *EffectFilterContext)

	// EnterFmType is called when entering the fmType production.
	EnterFmType(c *FmTypeContext)

	// EnterSize is called when entering the size production.
	EnterSize(c *SizeContext)

	// EnterOverlap is called when entering the overlap production.
	EnterOverlap(c *OverlapContext)

	// EnterWindowFunction is called when entering the windowFunction production.
	EnterWindowFunction(c *WindowFunctionContext)

	// EnterPitch is called when entering the pitch production.
	EnterPitch(c *PitchContext)

	// EnterModelType is called when entering the modelType production.
	EnterModelType(c *ModelTypeContext)

	// EnterModelParam is called when entering the modelParam production.
	EnterModelParam(c *ModelParamContext)

	// EnterExcitation is called when entering the excitation production.
	EnterExcitation(c *ExcitationContext)

	// EnterAmplitude is called when entering the amplitude production.
	EnterAmplitude(c *AmplitudeContext)

	// EnterPhase is called when entering the phase production.
	EnterPhase(c *PhaseContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterFilterType is called when entering the filterType production.
	EnterFilterType(c *FilterTypeContext)

	// EnterCutoff is called when entering the cutoff production.
	EnterCutoff(c *CutoffContext)

	// EnterResonance is called when entering the resonance production.
	EnterResonance(c *ResonanceContext)

	// EnterDecay is called when entering the decay production.
	EnterDecay(c *DecayContext)

	// EnterBandwidth is called when entering the bandwidth production.
	EnterBandwidth(c *BandwidthContext)

	// EnterGain is called when entering the gain production.
	EnterGain(c *GainContext)

	// EnterDelayTime is called when entering the delayTime production.
	EnterDelayTime(c *DelayTimeContext)

	// EnterFeedback is called when entering the feedback production.
	EnterFeedback(c *FeedbackContext)

	// EnterRate is called when entering the rate production.
	EnterRate(c *RateContext)

	// EnterDepth is called when entering the depth production.
	EnterDepth(c *DepthContext)

	// EnterDrive is called when entering the drive production.
	EnterDrive(c *DriveContext)

	// EnterTone is called when entering the tone production.
	EnterTone(c *ToneContext)

	// EnterStiffness is called when entering the stiffness production.
	EnterStiffness(c *StiffnessContext)

	// EnterDamping is called when entering the damping production.
	EnterDamping(c *DampingContext)

	// EnterTension is called when entering the tension production.
	EnterTension(c *TensionContext)

	// EnterLength is called when entering the length production.
	EnterLength(c *LengthContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitOscillator is called when exiting the oscillator production.
	ExitOscillator(c *OscillatorContext)

	// ExitEnvelope is called when exiting the envelope production.
	ExitEnvelope(c *EnvelopeContext)

	// ExitEffect is called when exiting the effect production.
	ExitEffect(c *EffectContext)

	// ExitFm is called when exiting the fm production.
	ExitFm(c *FmContext)

	// ExitGranular is called when exiting the granular production.
	ExitGranular(c *GranularContext)

	// ExitPhysicalModel is called when exiting the physicalModel production.
	ExitPhysicalModel(c *PhysicalModelContext)

	// ExitAdditive is called when exiting the additive production.
	ExitAdditive(c *AdditiveContext)

	// ExitWavetable is called when exiting the wavetable production.
	ExitWavetable(c *WavetableContext)

	// ExitSubtractive is called when exiting the subtractive production.
	ExitSubtractive(c *SubtractiveContext)

	// ExitKarplusStrong is called when exiting the karplusStrong production.
	ExitKarplusStrong(c *KarplusStrongContext)

	// ExitFormant is called when exiting the formant production.
	ExitFormant(c *FormantContext)

	// ExitWaveform is called when exiting the waveform production.
	ExitWaveform(c *WaveformContext)

	// ExitFrequency is called when exiting the frequency production.
	ExitFrequency(c *FrequencyContext)

	// ExitMix is called when exiting the mix production.
	ExitMix(c *MixContext)

	// ExitEnvType is called when exiting the envType production.
	ExitEnvType(c *EnvTypeContext)

	// ExitAttackTime is called when exiting the attackTime production.
	ExitAttackTime(c *AttackTimeContext)

	// ExitDecayTime is called when exiting the decayTime production.
	ExitDecayTime(c *DecayTimeContext)

	// ExitSustainLevel is called when exiting the sustainLevel production.
	ExitSustainLevel(c *SustainLevelContext)

	// ExitReleaseTime is called when exiting the releaseTime production.
	ExitReleaseTime(c *ReleaseTimeContext)

	// ExitEffectType is called when exiting the effectType production.
	ExitEffectType(c *EffectTypeContext)

	// ExitEffectParam is called when exiting the effectParam production.
	ExitEffectParam(c *EffectParamContext)

	// ExitEffectReverb is called when exiting the effectReverb production.
	ExitEffectReverb(c *EffectReverbContext)

	// ExitEffectDelay is called when exiting the effectDelay production.
	ExitEffectDelay(c *EffectDelayContext)

	// ExitEffectChorus is called when exiting the effectChorus production.
	ExitEffectChorus(c *EffectChorusContext)

	// ExitEffectDistortion is called when exiting the effectDistortion production.
	ExitEffectDistortion(c *EffectDistortionContext)

	// ExitEffectFilter is called when exiting the effectFilter production.
	ExitEffectFilter(c *EffectFilterContext)

	// ExitFmType is called when exiting the fmType production.
	ExitFmType(c *FmTypeContext)

	// ExitSize is called when exiting the size production.
	ExitSize(c *SizeContext)

	// ExitOverlap is called when exiting the overlap production.
	ExitOverlap(c *OverlapContext)

	// ExitWindowFunction is called when exiting the windowFunction production.
	ExitWindowFunction(c *WindowFunctionContext)

	// ExitPitch is called when exiting the pitch production.
	ExitPitch(c *PitchContext)

	// ExitModelType is called when exiting the modelType production.
	ExitModelType(c *ModelTypeContext)

	// ExitModelParam is called when exiting the modelParam production.
	ExitModelParam(c *ModelParamContext)

	// ExitExcitation is called when exiting the excitation production.
	ExitExcitation(c *ExcitationContext)

	// ExitAmplitude is called when exiting the amplitude production.
	ExitAmplitude(c *AmplitudeContext)

	// ExitPhase is called when exiting the phase production.
	ExitPhase(c *PhaseContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitFilterType is called when exiting the filterType production.
	ExitFilterType(c *FilterTypeContext)

	// ExitCutoff is called when exiting the cutoff production.
	ExitCutoff(c *CutoffContext)

	// ExitResonance is called when exiting the resonance production.
	ExitResonance(c *ResonanceContext)

	// ExitDecay is called when exiting the decay production.
	ExitDecay(c *DecayContext)

	// ExitBandwidth is called when exiting the bandwidth production.
	ExitBandwidth(c *BandwidthContext)

	// ExitGain is called when exiting the gain production.
	ExitGain(c *GainContext)

	// ExitDelayTime is called when exiting the delayTime production.
	ExitDelayTime(c *DelayTimeContext)

	// ExitFeedback is called when exiting the feedback production.
	ExitFeedback(c *FeedbackContext)

	// ExitRate is called when exiting the rate production.
	ExitRate(c *RateContext)

	// ExitDepth is called when exiting the depth production.
	ExitDepth(c *DepthContext)

	// ExitDrive is called when exiting the drive production.
	ExitDrive(c *DriveContext)

	// ExitTone is called when exiting the tone production.
	ExitTone(c *ToneContext)

	// ExitStiffness is called when exiting the stiffness production.
	ExitStiffness(c *StiffnessContext)

	// ExitDamping is called when exiting the damping production.
	ExitDamping(c *DampingContext)

	// ExitTension is called when exiting the tension production.
	ExitTension(c *TensionContext)

	// ExitLength is called when exiting the length production.
	ExitLength(c *LengthContext)
}
