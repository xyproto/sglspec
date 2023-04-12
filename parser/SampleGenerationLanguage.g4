// SampleGenerationLanguage.g4
grammar SampleGenerationLanguage;

prog: header (oscillator | envelope | effect | fm | granular | physicalModel | additive | wavetable | subtractive | karplusStrong | formant)+;

header: '#' 'SampleGenerationLanguage' 'v' NUMBER;

oscillator: 'Oscillator:' waveform COMMA frequency COMMA 'Mix:' mix;

envelope: 'Envelope:' envType COMMA 'Attack:' attackTime COMMA 'Decay:' decayTime COMMA 'Sustain:' sustainLevel COMMA 'Release:' releaseTime;

effect: 'Effect:' effectType (COMMA effectParam)+;

fm: 'FM:' fmType COMMA waveform COMMA frequency (COMMA 'Mix:' mix)?;

granular: 'Granular:' 'Size:' size COMMA 'Overlap:' overlap COMMA 'Window:' windowFunction COMMA 'Pitch:' pitch;

physicalModel: 'PhysicalModel:' modelType COMMA modelParam* COMMA 'Excitation:' excitation COMMA 'Mix:' mix;

additive: 'Additive:' 'Frequency:' frequency COMMA 'Amplitude:' amplitude COMMA 'Phase:' phase;

wavetable: 'Wavetable:' 'Waveform:' waveform COMMA 'Index:' index (COMMA 'Mix:' mix)?;

subtractive: 'Subtractive:' 'Waveform:' waveform COMMA 'Frequency:' frequency COMMA 'Filter:' filterType COMMA 'Cutoff:' cutoff COMMA 'Resonance:' resonance;

karplusStrong: 'KarplusStrong:' 'Frequency:' frequency COMMA 'Decay:' decay COMMA 'Filter:' filterType COMMA 'Cutoff:' cutoff;

formant: 'Formant:' 'Filter:' filterType COMMA 'Frequency:' frequency COMMA 'Bandwidth:' bandwidth COMMA 'Gain:' gain;

waveform: ('Sine' | 'Square' | 'Triangle' | 'Sawtooth' | 'Noise');
frequency: NUMBER 'Hz'? ;
mix: NUMBER;
envType: ('Amplitude' | 'Filter');
attackTime: NUMBER ('s' | 'ms')?;
decayTime: NUMBER ('s' | 'ms')?;
sustainLevel: NUMBER;
releaseTime: NUMBER ('s' | 'ms')?;
effectType: ('Reverb' | 'Delay' | 'Chorus' | 'Distortion' | 'High-pass' | 'Low-pass' | 'Noise');
effectParam: effectReverb | effectDelay | effectChorus | effectDistortion | effectFilter;
effectReverb: 'Decay:' decayTime | 'Mix:' mix;
effectDelay: 'Time:' delayTime COMMA 'Feedback:' feedback COMMA 'Mix:' mix;
effectChorus: 'Rate:' rate COMMA 'Depth:' depth COMMA 'Mix:' mix;
effectDistortion: 'Drive:' drive COMMA 'Tone:' tone COMMA 'Mix:' mix;
effectFilter: 'Cutoff:' cutoff COMMA 'Mix:' mix;
fmType: ('Carrier' | 'Modulator');
size: NUMBER ('s' | 'ms')?;
overlap: NUMBER '%';
windowFunction: ('Hanning' | 'Hamming' | 'Blackman' | 'Triangular');
pitch: NUMBER;
modelType: ('String' | 'Membrane' | 'Tube');
modelParam: ('Stiffness:' stiffness | 'Damping:' damping | 'Tension:' tension | 'Length:' length);
excitation: ('Impulse' | 'Noise' | 'Sine');
amplitude: NUMBER;
phase: NUMBER;
index: NUMBER;
filterType: ('Low-pass' | 'High-pass' | 'Band-pass' | 'Notch');
cutoff: NUMBER;
resonance: NUMBER;
decay: NUMBER;
bandwidth: NUMBER;
gain: NUMBER;

delayTime: NUMBER ('s' | 'ms')?;
feedback: NUMBER;
rate: NUMBER 'Hz'? ;
depth: NUMBER;
drive: NUMBER;
tone: NUMBER;
stiffness: NUMBER;
damping: NUMBER;
tension: NUMBER;
length: NUMBER;

COMMA: ',';
NUMBER: '-'? ('0'..'9')+ ('.' ('0'..'9')+)?;
WS: [ \t\r\n]+ -> skip;
