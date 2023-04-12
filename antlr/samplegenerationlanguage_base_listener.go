// Code generated from antlr/SampleGenerationLanguage.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // SampleGenerationLanguage

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseSampleGenerationLanguageListener is a complete listener for a parse tree produced by SampleGenerationLanguageParser.
type BaseSampleGenerationLanguageListener struct{}

var _ SampleGenerationLanguageListener = &BaseSampleGenerationLanguageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSampleGenerationLanguageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSampleGenerationLanguageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSampleGenerationLanguageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSampleGenerationLanguageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseSampleGenerationLanguageListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseSampleGenerationLanguageListener) ExitProg(ctx *ProgContext) {}

// EnterHeader is called when production header is entered.
func (s *BaseSampleGenerationLanguageListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BaseSampleGenerationLanguageListener) ExitHeader(ctx *HeaderContext) {}

// EnterOscillator is called when production oscillator is entered.
func (s *BaseSampleGenerationLanguageListener) EnterOscillator(ctx *OscillatorContext) {}

// ExitOscillator is called when production oscillator is exited.
func (s *BaseSampleGenerationLanguageListener) ExitOscillator(ctx *OscillatorContext) {}

// EnterEnvelope is called when production envelope is entered.
func (s *BaseSampleGenerationLanguageListener) EnterEnvelope(ctx *EnvelopeContext) {}

// ExitEnvelope is called when production envelope is exited.
func (s *BaseSampleGenerationLanguageListener) ExitEnvelope(ctx *EnvelopeContext) {}

// EnterEffect is called when production effect is entered.
func (s *BaseSampleGenerationLanguageListener) EnterEffect(ctx *EffectContext) {}

// ExitEffect is called when production effect is exited.
func (s *BaseSampleGenerationLanguageListener) ExitEffect(ctx *EffectContext) {}

// EnterFm is called when production fm is entered.
func (s *BaseSampleGenerationLanguageListener) EnterFm(ctx *FmContext) {}

// ExitFm is called when production fm is exited.
func (s *BaseSampleGenerationLanguageListener) ExitFm(ctx *FmContext) {}

// EnterGranular is called when production granular is entered.
func (s *BaseSampleGenerationLanguageListener) EnterGranular(ctx *GranularContext) {}

// ExitGranular is called when production granular is exited.
func (s *BaseSampleGenerationLanguageListener) ExitGranular(ctx *GranularContext) {}

// EnterPhysicalModel is called when production physicalModel is entered.
func (s *BaseSampleGenerationLanguageListener) EnterPhysicalModel(ctx *PhysicalModelContext) {}

// ExitPhysicalModel is called when production physicalModel is exited.
func (s *BaseSampleGenerationLanguageListener) ExitPhysicalModel(ctx *PhysicalModelContext) {}

// EnterAdditive is called when production additive is entered.
func (s *BaseSampleGenerationLanguageListener) EnterAdditive(ctx *AdditiveContext) {}

// ExitAdditive is called when production additive is exited.
func (s *BaseSampleGenerationLanguageListener) ExitAdditive(ctx *AdditiveContext) {}

// EnterWavetable is called when production wavetable is entered.
func (s *BaseSampleGenerationLanguageListener) EnterWavetable(ctx *WavetableContext) {}

// ExitWavetable is called when production wavetable is exited.
func (s *BaseSampleGenerationLanguageListener) ExitWavetable(ctx *WavetableContext) {}

// EnterSubtractive is called when production subtractive is entered.
func (s *BaseSampleGenerationLanguageListener) EnterSubtractive(ctx *SubtractiveContext) {}

// ExitSubtractive is called when production subtractive is exited.
func (s *BaseSampleGenerationLanguageListener) ExitSubtractive(ctx *SubtractiveContext) {}

// EnterKarplusStrong is called when production karplusStrong is entered.
func (s *BaseSampleGenerationLanguageListener) EnterKarplusStrong(ctx *KarplusStrongContext) {}

// ExitKarplusStrong is called when production karplusStrong is exited.
func (s *BaseSampleGenerationLanguageListener) ExitKarplusStrong(ctx *KarplusStrongContext) {}

// EnterFormant is called when production formant is entered.
func (s *BaseSampleGenerationLanguageListener) EnterFormant(ctx *FormantContext) {}

// ExitFormant is called when production formant is exited.
func (s *BaseSampleGenerationLanguageListener) ExitFormant(ctx *FormantContext) {}

// EnterWaveform is called when production waveform is entered.
func (s *BaseSampleGenerationLanguageListener) EnterWaveform(ctx *WaveformContext) {}

// ExitWaveform is called when production waveform is exited.
func (s *BaseSampleGenerationLanguageListener) ExitWaveform(ctx *WaveformContext) {}

// EnterFrequency is called when production frequency is entered.
func (s *BaseSampleGenerationLanguageListener) EnterFrequency(ctx *FrequencyContext) {}

// ExitFrequency is called when production frequency is exited.
func (s *BaseSampleGenerationLanguageListener) ExitFrequency(ctx *FrequencyContext) {}

// EnterMix is called when production mix is entered.
func (s *BaseSampleGenerationLanguageListener) EnterMix(ctx *MixContext) {}

// ExitMix is called when production mix is exited.
func (s *BaseSampleGenerationLanguageListener) ExitMix(ctx *MixContext) {}

// EnterEnvType is called when production envType is entered.
func (s *BaseSampleGenerationLanguageListener) EnterEnvType(ctx *EnvTypeContext) {}

// ExitEnvType is called when production envType is exited.
func (s *BaseSampleGenerationLanguageListener) ExitEnvType(ctx *EnvTypeContext) {}

// EnterAttackTime is called when production attackTime is entered.
func (s *BaseSampleGenerationLanguageListener) EnterAttackTime(ctx *AttackTimeContext) {}

// ExitAttackTime is called when production attackTime is exited.
func (s *BaseSampleGenerationLanguageListener) ExitAttackTime(ctx *AttackTimeContext) {}

// EnterDecayTime is called when production decayTime is entered.
func (s *BaseSampleGenerationLanguageListener) EnterDecayTime(ctx *DecayTimeContext) {}

// ExitDecayTime is called when production decayTime is exited.
func (s *BaseSampleGenerationLanguageListener) ExitDecayTime(ctx *DecayTimeContext) {}

// EnterSustainLevel is called when production sustainLevel is entered.
func (s *BaseSampleGenerationLanguageListener) EnterSustainLevel(ctx *SustainLevelContext) {}

// ExitSustainLevel is called when production sustainLevel is exited.
func (s *BaseSampleGenerationLanguageListener) ExitSustainLevel(ctx *SustainLevelContext) {}

// EnterReleaseTime is called when production releaseTime is entered.
func (s *BaseSampleGenerationLanguageListener) EnterReleaseTime(ctx *ReleaseTimeContext) {}

// ExitReleaseTime is called when production releaseTime is exited.
func (s *BaseSampleGenerationLanguageListener) ExitReleaseTime(ctx *ReleaseTimeContext) {}

// EnterEffectType is called when production effectType is entered.
func (s *BaseSampleGenerationLanguageListener) EnterEffectType(ctx *EffectTypeContext) {}

// ExitEffectType is called when production effectType is exited.
func (s *BaseSampleGenerationLanguageListener) ExitEffectType(ctx *EffectTypeContext) {}

// EnterEffectParam is called when production effectParam is entered.
func (s *BaseSampleGenerationLanguageListener) EnterEffectParam(ctx *EffectParamContext) {}

// ExitEffectParam is called when production effectParam is exited.
func (s *BaseSampleGenerationLanguageListener) ExitEffectParam(ctx *EffectParamContext) {}

// EnterEffectReverb is called when production effectReverb is entered.
func (s *BaseSampleGenerationLanguageListener) EnterEffectReverb(ctx *EffectReverbContext) {}

// ExitEffectReverb is called when production effectReverb is exited.
func (s *BaseSampleGenerationLanguageListener) ExitEffectReverb(ctx *EffectReverbContext) {}

// EnterEffectDelay is called when production effectDelay is entered.
func (s *BaseSampleGenerationLanguageListener) EnterEffectDelay(ctx *EffectDelayContext) {}

// ExitEffectDelay is called when production effectDelay is exited.
func (s *BaseSampleGenerationLanguageListener) ExitEffectDelay(ctx *EffectDelayContext) {}

// EnterEffectChorus is called when production effectChorus is entered.
func (s *BaseSampleGenerationLanguageListener) EnterEffectChorus(ctx *EffectChorusContext) {}

// ExitEffectChorus is called when production effectChorus is exited.
func (s *BaseSampleGenerationLanguageListener) ExitEffectChorus(ctx *EffectChorusContext) {}

// EnterEffectDistortion is called when production effectDistortion is entered.
func (s *BaseSampleGenerationLanguageListener) EnterEffectDistortion(ctx *EffectDistortionContext) {}

// ExitEffectDistortion is called when production effectDistortion is exited.
func (s *BaseSampleGenerationLanguageListener) ExitEffectDistortion(ctx *EffectDistortionContext) {}

// EnterEffectFilter is called when production effectFilter is entered.
func (s *BaseSampleGenerationLanguageListener) EnterEffectFilter(ctx *EffectFilterContext) {}

// ExitEffectFilter is called when production effectFilter is exited.
func (s *BaseSampleGenerationLanguageListener) ExitEffectFilter(ctx *EffectFilterContext) {}

// EnterFmType is called when production fmType is entered.
func (s *BaseSampleGenerationLanguageListener) EnterFmType(ctx *FmTypeContext) {}

// ExitFmType is called when production fmType is exited.
func (s *BaseSampleGenerationLanguageListener) ExitFmType(ctx *FmTypeContext) {}

// EnterSize is called when production size is entered.
func (s *BaseSampleGenerationLanguageListener) EnterSize(ctx *SizeContext) {}

// ExitSize is called when production size is exited.
func (s *BaseSampleGenerationLanguageListener) ExitSize(ctx *SizeContext) {}

// EnterOverlap is called when production overlap is entered.
func (s *BaseSampleGenerationLanguageListener) EnterOverlap(ctx *OverlapContext) {}

// ExitOverlap is called when production overlap is exited.
func (s *BaseSampleGenerationLanguageListener) ExitOverlap(ctx *OverlapContext) {}

// EnterWindowFunction is called when production windowFunction is entered.
func (s *BaseSampleGenerationLanguageListener) EnterWindowFunction(ctx *WindowFunctionContext) {}

// ExitWindowFunction is called when production windowFunction is exited.
func (s *BaseSampleGenerationLanguageListener) ExitWindowFunction(ctx *WindowFunctionContext) {}

// EnterPitch is called when production pitch is entered.
func (s *BaseSampleGenerationLanguageListener) EnterPitch(ctx *PitchContext) {}

// ExitPitch is called when production pitch is exited.
func (s *BaseSampleGenerationLanguageListener) ExitPitch(ctx *PitchContext) {}

// EnterModelType is called when production modelType is entered.
func (s *BaseSampleGenerationLanguageListener) EnterModelType(ctx *ModelTypeContext) {}

// ExitModelType is called when production modelType is exited.
func (s *BaseSampleGenerationLanguageListener) ExitModelType(ctx *ModelTypeContext) {}

// EnterModelParam is called when production modelParam is entered.
func (s *BaseSampleGenerationLanguageListener) EnterModelParam(ctx *ModelParamContext) {}

// ExitModelParam is called when production modelParam is exited.
func (s *BaseSampleGenerationLanguageListener) ExitModelParam(ctx *ModelParamContext) {}

// EnterExcitation is called when production excitation is entered.
func (s *BaseSampleGenerationLanguageListener) EnterExcitation(ctx *ExcitationContext) {}

// ExitExcitation is called when production excitation is exited.
func (s *BaseSampleGenerationLanguageListener) ExitExcitation(ctx *ExcitationContext) {}

// EnterAmplitude is called when production amplitude is entered.
func (s *BaseSampleGenerationLanguageListener) EnterAmplitude(ctx *AmplitudeContext) {}

// ExitAmplitude is called when production amplitude is exited.
func (s *BaseSampleGenerationLanguageListener) ExitAmplitude(ctx *AmplitudeContext) {}

// EnterPhase is called when production phase is entered.
func (s *BaseSampleGenerationLanguageListener) EnterPhase(ctx *PhaseContext) {}

// ExitPhase is called when production phase is exited.
func (s *BaseSampleGenerationLanguageListener) ExitPhase(ctx *PhaseContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseSampleGenerationLanguageListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseSampleGenerationLanguageListener) ExitIndex(ctx *IndexContext) {}

// EnterFilterType is called when production filterType is entered.
func (s *BaseSampleGenerationLanguageListener) EnterFilterType(ctx *FilterTypeContext) {}

// ExitFilterType is called when production filterType is exited.
func (s *BaseSampleGenerationLanguageListener) ExitFilterType(ctx *FilterTypeContext) {}

// EnterCutoff is called when production cutoff is entered.
func (s *BaseSampleGenerationLanguageListener) EnterCutoff(ctx *CutoffContext) {}

// ExitCutoff is called when production cutoff is exited.
func (s *BaseSampleGenerationLanguageListener) ExitCutoff(ctx *CutoffContext) {}

// EnterResonance is called when production resonance is entered.
func (s *BaseSampleGenerationLanguageListener) EnterResonance(ctx *ResonanceContext) {}

// ExitResonance is called when production resonance is exited.
func (s *BaseSampleGenerationLanguageListener) ExitResonance(ctx *ResonanceContext) {}

// EnterDecay is called when production decay is entered.
func (s *BaseSampleGenerationLanguageListener) EnterDecay(ctx *DecayContext) {}

// ExitDecay is called when production decay is exited.
func (s *BaseSampleGenerationLanguageListener) ExitDecay(ctx *DecayContext) {}

// EnterBandwidth is called when production bandwidth is entered.
func (s *BaseSampleGenerationLanguageListener) EnterBandwidth(ctx *BandwidthContext) {}

// ExitBandwidth is called when production bandwidth is exited.
func (s *BaseSampleGenerationLanguageListener) ExitBandwidth(ctx *BandwidthContext) {}

// EnterGain is called when production gain is entered.
func (s *BaseSampleGenerationLanguageListener) EnterGain(ctx *GainContext) {}

// ExitGain is called when production gain is exited.
func (s *BaseSampleGenerationLanguageListener) ExitGain(ctx *GainContext) {}

// EnterDelayTime is called when production delayTime is entered.
func (s *BaseSampleGenerationLanguageListener) EnterDelayTime(ctx *DelayTimeContext) {}

// ExitDelayTime is called when production delayTime is exited.
func (s *BaseSampleGenerationLanguageListener) ExitDelayTime(ctx *DelayTimeContext) {}

// EnterFeedback is called when production feedback is entered.
func (s *BaseSampleGenerationLanguageListener) EnterFeedback(ctx *FeedbackContext) {}

// ExitFeedback is called when production feedback is exited.
func (s *BaseSampleGenerationLanguageListener) ExitFeedback(ctx *FeedbackContext) {}

// EnterRate is called when production rate is entered.
func (s *BaseSampleGenerationLanguageListener) EnterRate(ctx *RateContext) {}

// ExitRate is called when production rate is exited.
func (s *BaseSampleGenerationLanguageListener) ExitRate(ctx *RateContext) {}

// EnterDepth is called when production depth is entered.
func (s *BaseSampleGenerationLanguageListener) EnterDepth(ctx *DepthContext) {}

// ExitDepth is called when production depth is exited.
func (s *BaseSampleGenerationLanguageListener) ExitDepth(ctx *DepthContext) {}

// EnterDrive is called when production drive is entered.
func (s *BaseSampleGenerationLanguageListener) EnterDrive(ctx *DriveContext) {}

// ExitDrive is called when production drive is exited.
func (s *BaseSampleGenerationLanguageListener) ExitDrive(ctx *DriveContext) {}

// EnterTone is called when production tone is entered.
func (s *BaseSampleGenerationLanguageListener) EnterTone(ctx *ToneContext) {}

// ExitTone is called when production tone is exited.
func (s *BaseSampleGenerationLanguageListener) ExitTone(ctx *ToneContext) {}

// EnterStiffness is called when production stiffness is entered.
func (s *BaseSampleGenerationLanguageListener) EnterStiffness(ctx *StiffnessContext) {}

// ExitStiffness is called when production stiffness is exited.
func (s *BaseSampleGenerationLanguageListener) ExitStiffness(ctx *StiffnessContext) {}

// EnterDamping is called when production damping is entered.
func (s *BaseSampleGenerationLanguageListener) EnterDamping(ctx *DampingContext) {}

// ExitDamping is called when production damping is exited.
func (s *BaseSampleGenerationLanguageListener) ExitDamping(ctx *DampingContext) {}

// EnterTension is called when production tension is entered.
func (s *BaseSampleGenerationLanguageListener) EnterTension(ctx *TensionContext) {}

// ExitTension is called when production tension is exited.
func (s *BaseSampleGenerationLanguageListener) ExitTension(ctx *TensionContext) {}

// EnterLength is called when production length is entered.
func (s *BaseSampleGenerationLanguageListener) EnterLength(ctx *LengthContext) {}

// ExitLength is called when production length is exited.
func (s *BaseSampleGenerationLanguageListener) ExitLength(ctx *LengthContext) {}
