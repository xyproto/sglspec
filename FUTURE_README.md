# Sample Generation Language (SGL) Specification

The Sample Generation Language (SGL) is a text-based format for describing and generating audio samples using various synthesis techniques. The format is designed to be versatile and can be used to create samples for various music genres, including chiptune, acid trance, drum & bass, trance, EDM, demoscene music, and new experimental sounds.

Note that the included sgl2wav utility is a work in progress, and currently does not work!

[Antlr frammar file](antlr/SampleGenerationLanguage.g4)

## Header

The header consists of a single line that starts with a hash symbol (#) and indicates the version of the format.

Example:

    # SampleGenerationLanguage v2.0

## Oscillator

The oscillator section describes the basic waveforms and parameters for sound generation. Each oscillator line starts with the keyword "Oscillator" followed by the type, frequency (in Hz), and an optional mix level.

Supported waveforms:

    Sine
    Square
    Triangle
    Sawtooth
    Noise

Example:

    Oscillator: Sine, 440, Mix: 0.5
    Oscillator: Square, 440, Mix: 0.5

## Envelopes

The envelopes section describes the amplitude and filter envelopes of the sound. It starts with the keyword "Envelope" followed by the type (Amplitude or Filter), attack time, decay time, sustain level, and release time.

Example:

    Envelope: Amplitude, Attack: 10ms, Decay: 100ms, Sustain: 0.5, Release: 300ms
    Envelope: Filter, Attack: 20ms, Decay: 50ms, Sustain: 0.8, Release: 100ms

## Effects

The effects section describes the audio processing effects applied to the sound. Each line starts with the keyword "Effect" followed by the type and parameters.

Supported effects:

    Reverb: Decay time, Mix level
    Delay: Delay time, Feedback, Mix level
    Chorus: Rate, Depth, Mix level
    Distortion: Drive, Tone, Mix level

Example:

    Effect: Reverb, Decay: 2s, Mix: 0.3
    Effect: Delay, Time: 500ms, Feedback: 0.4, Mix: 0.25

## FM Synthesis

The FM synthesis section describes the carrier and modulator oscillators. Each line starts with the keyword "FM" followed by the oscillator type (Carrier or Modulator), waveform, frequency, and optional mix level.

Example:

    FM: Carrier, Sine, 440, Mix: 0.8
    FM: Modulator, Sine, 880, Mix: 0.2

## Granular Synthesis

The granular synthesis section describes the granular parameters. It starts with the keyword "Granular" followed by the grain size, overlap, window function, and pitch scaling factor.

Supported window functions:

    Hanning
    Hamming
    Blackman
    Triangular

Example:

    Granular: Size: 50ms, Overlap: 75%, Window: Hanning, Pitch: 1.0

## Physical Modeling

The physical modeling section describes the type of physical model, its parameters, and any excitation methods. Each line starts with the keyword "PhysicalModel" followed by the model type, parameters, and excitation method.

Supported physical models:

    String: Stiffness, Damping
    Membrane: Tension, Damping
    Tube: Length, Damping

Supported excitation methods:

    Impulse
    Noise
    Sine

Example:

    PhysicalModel: String, Stiffness: 0.5, Damping: 0.2, Excitation: Impulse, Mix: 1.0
    PhysicalModel: Membrane, Tension: 0.8, Damping: 0.1, Excitation: Noise, Mix: 0.6

## Additive Synthesis

The additive synthesis section describes the combination of multiple sinusoidal waveforms with different frequencies and amplitudes. Each line starts with the keyword "Additive" followed by the frequency, amplitude, and phase offset.

Example:

    Additive: Frequency: 440, Amplitude: 0.5, Phase: 0
    Additive: Frequency: 880, Amplitude: 0.25, Phase: 0

## Wavetable Synthesis

The wavetable synthesis section describes the use of a wavetable for generating sound. Each line starts with the keyword "Wavetable" followed by the waveform, index, and optional mix level.

Example:

    Wavetable: Waveform: Sine, Index: 0, Mix: 0.5
    Wavetable: Waveform: Square, Index: 1, Mix: 0.5

## Subtractive Synthesis

The subtractive synthesis section describes the use of filters to process harmonically rich waveforms. Each line starts with the keyword "Subtractive" followed by the waveform, frequency, filter type, and filter parameters.

Supported filter types:

    Low-pass
    High-pass
    Band-pass
    Notch

Example:

    Subtractive: Waveform: Sawtooth, Frequency: 440, Filter: Low-pass, Cutoff: 1000, Resonance: 0.5

## Karplus-Strong Synthesis

The Karplus-Strong synthesis section describes the simulation of plucked strings using a delay line and a low-pass filter. Each line starts with the keyword "KarplusStrong" followed by the frequency, decay, and filter parameters.

Example:

    KarplusStrong: Frequency: 440, Decay: 0.99, Filter: Low-pass, Cutoff: 1000

## Formant Synthesis

The formant synthesis section describes the use of multiple resonant filters to mimic the human voice. Each line starts with the keyword "Formant" followed by the filter type, frequency, bandwidth, and gain.

Example:

    Formant: Filter: Band-pass, Frequency: 500, Bandwidth: 100, Gain: 0.8
    Formant: Filter: Band-pass, Frequency: 1500, Bandwidth: 200, Gain: 0.6

---

The Sample Generation Language (SGL) is versatile and capable of generating a wide variety of sounds for various music genres and styles. Remember that a dedicated tool or software is required to interpret the SGL format and generate the corresponding audio samples.

## General info

* Version: 2.0.0
* License: BSD-3
* The content has been modified and curated by a human, but the samples needs more testing.
