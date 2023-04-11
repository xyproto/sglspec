// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // SampleGenerationLanguage

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SampleGenerationLanguageParser struct {
	*antlr.BaseParser
}

var samplegenerationlanguageParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func samplegenerationlanguageParserInit() {
	staticData := &samplegenerationlanguageParserStaticData
	staticData.literalNames = []string{
		"", "'#'", "'SampleGenerationLanguage'", "'v'", "'Oscillator:'", "'Mix:'",
		"'Envelope:'", "'Attack:'", "'Decay:'", "'Sustain:'", "'Release:'",
		"'Effect:'", "'FM:'", "'Granular:'", "'Size:'", "'Overlap:'", "'Window:'",
		"'Pitch:'", "'PhysicalModel:'", "'Excitation:'", "'Additive:'", "'Frequency:'",
		"'Amplitude:'", "'Phase:'", "'Wavetable:'", "'Waveform:'", "'Index:'",
		"'Subtractive:'", "'Filter:'", "'Cutoff:'", "'Resonance:'", "'KarplusStrong:'",
		"'Formant:'", "'Bandwidth:'", "'Gain:'", "'Sine'", "'Square'", "'Triangle'",
		"'Sawtooth'", "'Noise'", "'Hz'", "'Amplitude'", "'Filter'", "'ms'",
		"'Reverb'", "'Delay'", "'Chorus'", "'Distortion'", "'High-pass'", "'Low-pass'",
		"'Time:'", "'Feedback:'", "'Rate:'", "'Depth:'", "'Drive:'", "'Tone:'",
		"'Carrier'", "'Modulator'", "'%'", "'Hanning'", "'Hamming'", "'Blackman'",
		"'Triangular'", "'String'", "'Membrane'", "'Tube'", "'Stiffness:'",
		"'Damping:'", "'Tension:'", "'Length:'", "'Impulse'", "'Band-pass'",
		"'Notch'", "','",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "COMMA", "NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"prog", "header", "oscillator", "envelope", "effect", "fm", "granular",
		"physicalModel", "additive", "wavetable", "subtractive", "karplusStrong",
		"formant", "waveform", "frequency", "mix", "envType", "attackTime",
		"decayTime", "sustainLevel", "releaseTime", "effectType", "effectParam",
		"effectReverb", "effectDelay", "effectChorus", "effectDistortion", "effectFilter",
		"fmType", "size", "overlap", "windowFunction", "pitch", "modelType",
		"modelParam", "excitation", "amplitude", "phase", "index", "filterType",
		"cutoff", "resonance", "decay", "bandwidth", "gain", "delayTime", "feedback",
		"rate", "depth", "drive", "tone", "stiffness", "damping", "tension",
		"length",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 75, 402, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 4, 0, 123, 8, 0, 11, 0, 12, 0, 124, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4, 159, 8, 4, 11, 4, 12,
		4, 160, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 172,
		8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 191, 8, 7, 10, 7, 12, 7, 194,
		9, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 3, 9, 222, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 3, 14, 270,
		8, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 3, 22, 294, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 300,
		8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 3, 34, 357, 8, 34, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1,
		42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46,
		1, 47, 1, 47, 3, 47, 386, 8, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1,
		50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 0, 0,
		55, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
		106, 108, 0, 8, 1, 0, 35, 39, 1, 0, 41, 42, 2, 0, 39, 39, 44, 49, 1, 0,
		56, 57, 1, 0, 59, 62, 1, 0, 63, 65, 3, 0, 35, 35, 39, 39, 70, 70, 2, 0,
		48, 49, 71, 72, 371, 0, 110, 1, 0, 0, 0, 2, 126, 1, 0, 0, 0, 4, 131, 1,
		0, 0, 0, 6, 139, 1, 0, 0, 0, 8, 154, 1, 0, 0, 0, 10, 162, 1, 0, 0, 0, 12,
		173, 1, 0, 0, 0, 14, 186, 1, 0, 0, 0, 16, 202, 1, 0, 0, 0, 18, 212, 1,
		0, 0, 0, 20, 223, 1, 0, 0, 0, 22, 239, 1, 0, 0, 0, 24, 252, 1, 0, 0, 0,
		26, 265, 1, 0, 0, 0, 28, 267, 1, 0, 0, 0, 30, 271, 1, 0, 0, 0, 32, 273,
		1, 0, 0, 0, 34, 275, 1, 0, 0, 0, 36, 278, 1, 0, 0, 0, 38, 281, 1, 0, 0,
		0, 40, 283, 1, 0, 0, 0, 42, 286, 1, 0, 0, 0, 44, 293, 1, 0, 0, 0, 46, 299,
		1, 0, 0, 0, 48, 301, 1, 0, 0, 0, 50, 310, 1, 0, 0, 0, 52, 319, 1, 0, 0,
		0, 54, 328, 1, 0, 0, 0, 56, 334, 1, 0, 0, 0, 58, 336, 1, 0, 0, 0, 60, 339,
		1, 0, 0, 0, 62, 342, 1, 0, 0, 0, 64, 344, 1, 0, 0, 0, 66, 346, 1, 0, 0,
		0, 68, 356, 1, 0, 0, 0, 70, 358, 1, 0, 0, 0, 72, 360, 1, 0, 0, 0, 74, 362,
		1, 0, 0, 0, 76, 364, 1, 0, 0, 0, 78, 366, 1, 0, 0, 0, 80, 368, 1, 0, 0,
		0, 82, 370, 1, 0, 0, 0, 84, 372, 1, 0, 0, 0, 86, 374, 1, 0, 0, 0, 88, 376,
		1, 0, 0, 0, 90, 378, 1, 0, 0, 0, 92, 381, 1, 0, 0, 0, 94, 383, 1, 0, 0,
		0, 96, 387, 1, 0, 0, 0, 98, 389, 1, 0, 0, 0, 100, 391, 1, 0, 0, 0, 102,
		393, 1, 0, 0, 0, 104, 395, 1, 0, 0, 0, 106, 397, 1, 0, 0, 0, 108, 399,
		1, 0, 0, 0, 110, 122, 3, 2, 1, 0, 111, 123, 3, 4, 2, 0, 112, 123, 3, 6,
		3, 0, 113, 123, 3, 8, 4, 0, 114, 123, 3, 10, 5, 0, 115, 123, 3, 12, 6,
		0, 116, 123, 3, 14, 7, 0, 117, 123, 3, 16, 8, 0, 118, 123, 3, 18, 9, 0,
		119, 123, 3, 20, 10, 0, 120, 123, 3, 22, 11, 0, 121, 123, 3, 24, 12, 0,
		122, 111, 1, 0, 0, 0, 122, 112, 1, 0, 0, 0, 122, 113, 1, 0, 0, 0, 122,
		114, 1, 0, 0, 0, 122, 115, 1, 0, 0, 0, 122, 116, 1, 0, 0, 0, 122, 117,
		1, 0, 0, 0, 122, 118, 1, 0, 0, 0, 122, 119, 1, 0, 0, 0, 122, 120, 1, 0,
		0, 0, 122, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0,
		124, 125, 1, 0, 0, 0, 125, 1, 1, 0, 0, 0, 126, 127, 5, 1, 0, 0, 127, 128,
		5, 2, 0, 0, 128, 129, 5, 3, 0, 0, 129, 130, 5, 74, 0, 0, 130, 3, 1, 0,
		0, 0, 131, 132, 5, 4, 0, 0, 132, 133, 3, 26, 13, 0, 133, 134, 5, 73, 0,
		0, 134, 135, 3, 28, 14, 0, 135, 136, 5, 73, 0, 0, 136, 137, 5, 5, 0, 0,
		137, 138, 3, 30, 15, 0, 138, 5, 1, 0, 0, 0, 139, 140, 5, 6, 0, 0, 140,
		141, 3, 32, 16, 0, 141, 142, 5, 73, 0, 0, 142, 143, 5, 7, 0, 0, 143, 144,
		3, 34, 17, 0, 144, 145, 5, 73, 0, 0, 145, 146, 5, 8, 0, 0, 146, 147, 3,
		36, 18, 0, 147, 148, 5, 73, 0, 0, 148, 149, 5, 9, 0, 0, 149, 150, 3, 38,
		19, 0, 150, 151, 5, 73, 0, 0, 151, 152, 5, 10, 0, 0, 152, 153, 3, 40, 20,
		0, 153, 7, 1, 0, 0, 0, 154, 155, 5, 11, 0, 0, 155, 158, 3, 42, 21, 0, 156,
		157, 5, 73, 0, 0, 157, 159, 3, 44, 22, 0, 158, 156, 1, 0, 0, 0, 159, 160,
		1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 9, 1, 0, 0,
		0, 162, 163, 5, 12, 0, 0, 163, 164, 3, 56, 28, 0, 164, 165, 5, 73, 0, 0,
		165, 166, 3, 26, 13, 0, 166, 167, 5, 73, 0, 0, 167, 171, 3, 28, 14, 0,
		168, 169, 5, 73, 0, 0, 169, 170, 5, 5, 0, 0, 170, 172, 3, 30, 15, 0, 171,
		168, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 11, 1, 0, 0, 0, 173, 174, 5,
		13, 0, 0, 174, 175, 5, 14, 0, 0, 175, 176, 3, 58, 29, 0, 176, 177, 5, 73,
		0, 0, 177, 178, 5, 15, 0, 0, 178, 179, 3, 60, 30, 0, 179, 180, 5, 73, 0,
		0, 180, 181, 5, 16, 0, 0, 181, 182, 3, 62, 31, 0, 182, 183, 5, 73, 0, 0,
		183, 184, 5, 17, 0, 0, 184, 185, 3, 64, 32, 0, 185, 13, 1, 0, 0, 0, 186,
		187, 5, 18, 0, 0, 187, 188, 3, 66, 33, 0, 188, 192, 5, 73, 0, 0, 189, 191,
		3, 68, 34, 0, 190, 189, 1, 0, 0, 0, 191, 194, 1, 0, 0, 0, 192, 190, 1,
		0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 195, 1, 0, 0, 0, 194, 192, 1, 0, 0,
		0, 195, 196, 5, 73, 0, 0, 196, 197, 5, 19, 0, 0, 197, 198, 3, 70, 35, 0,
		198, 199, 5, 73, 0, 0, 199, 200, 5, 5, 0, 0, 200, 201, 3, 30, 15, 0, 201,
		15, 1, 0, 0, 0, 202, 203, 5, 20, 0, 0, 203, 204, 5, 21, 0, 0, 204, 205,
		3, 28, 14, 0, 205, 206, 5, 73, 0, 0, 206, 207, 5, 22, 0, 0, 207, 208, 3,
		72, 36, 0, 208, 209, 5, 73, 0, 0, 209, 210, 5, 23, 0, 0, 210, 211, 3, 74,
		37, 0, 211, 17, 1, 0, 0, 0, 212, 213, 5, 24, 0, 0, 213, 214, 5, 25, 0,
		0, 214, 215, 3, 26, 13, 0, 215, 216, 5, 73, 0, 0, 216, 217, 5, 26, 0, 0,
		217, 221, 3, 76, 38, 0, 218, 219, 5, 73, 0, 0, 219, 220, 5, 5, 0, 0, 220,
		222, 3, 30, 15, 0, 221, 218, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 19,
		1, 0, 0, 0, 223, 224, 5, 27, 0, 0, 224, 225, 5, 25, 0, 0, 225, 226, 3,
		26, 13, 0, 226, 227, 5, 73, 0, 0, 227, 228, 5, 21, 0, 0, 228, 229, 3, 28,
		14, 0, 229, 230, 5, 73, 0, 0, 230, 231, 5, 28, 0, 0, 231, 232, 3, 78, 39,
		0, 232, 233, 5, 73, 0, 0, 233, 234, 5, 29, 0, 0, 234, 235, 3, 80, 40, 0,
		235, 236, 5, 73, 0, 0, 236, 237, 5, 30, 0, 0, 237, 238, 3, 82, 41, 0, 238,
		21, 1, 0, 0, 0, 239, 240, 5, 31, 0, 0, 240, 241, 5, 21, 0, 0, 241, 242,
		3, 28, 14, 0, 242, 243, 5, 73, 0, 0, 243, 244, 5, 8, 0, 0, 244, 245, 3,
		84, 42, 0, 245, 246, 5, 73, 0, 0, 246, 247, 5, 28, 0, 0, 247, 248, 3, 78,
		39, 0, 248, 249, 5, 73, 0, 0, 249, 250, 5, 29, 0, 0, 250, 251, 3, 80, 40,
		0, 251, 23, 1, 0, 0, 0, 252, 253, 5, 32, 0, 0, 253, 254, 5, 28, 0, 0, 254,
		255, 3, 78, 39, 0, 255, 256, 5, 73, 0, 0, 256, 257, 5, 21, 0, 0, 257, 258,
		3, 28, 14, 0, 258, 259, 5, 73, 0, 0, 259, 260, 5, 33, 0, 0, 260, 261, 3,
		86, 43, 0, 261, 262, 5, 73, 0, 0, 262, 263, 5, 34, 0, 0, 263, 264, 3, 88,
		44, 0, 264, 25, 1, 0, 0, 0, 265, 266, 7, 0, 0, 0, 266, 27, 1, 0, 0, 0,
		267, 269, 5, 74, 0, 0, 268, 270, 5, 40, 0, 0, 269, 268, 1, 0, 0, 0, 269,
		270, 1, 0, 0, 0, 270, 29, 1, 0, 0, 0, 271, 272, 5, 74, 0, 0, 272, 31, 1,
		0, 0, 0, 273, 274, 7, 1, 0, 0, 274, 33, 1, 0, 0, 0, 275, 276, 5, 74, 0,
		0, 276, 277, 5, 43, 0, 0, 277, 35, 1, 0, 0, 0, 278, 279, 5, 74, 0, 0, 279,
		280, 5, 43, 0, 0, 280, 37, 1, 0, 0, 0, 281, 282, 5, 74, 0, 0, 282, 39,
		1, 0, 0, 0, 283, 284, 5, 74, 0, 0, 284, 285, 5, 43, 0, 0, 285, 41, 1, 0,
		0, 0, 286, 287, 7, 2, 0, 0, 287, 43, 1, 0, 0, 0, 288, 294, 3, 46, 23, 0,
		289, 294, 3, 48, 24, 0, 290, 294, 3, 50, 25, 0, 291, 294, 3, 52, 26, 0,
		292, 294, 3, 54, 27, 0, 293, 288, 1, 0, 0, 0, 293, 289, 1, 0, 0, 0, 293,
		290, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293, 292, 1, 0, 0, 0, 294, 45, 1,
		0, 0, 0, 295, 296, 5, 8, 0, 0, 296, 300, 3, 36, 18, 0, 297, 298, 5, 5,
		0, 0, 298, 300, 3, 30, 15, 0, 299, 295, 1, 0, 0, 0, 299, 297, 1, 0, 0,
		0, 300, 47, 1, 0, 0, 0, 301, 302, 5, 50, 0, 0, 302, 303, 3, 90, 45, 0,
		303, 304, 5, 73, 0, 0, 304, 305, 5, 51, 0, 0, 305, 306, 3, 92, 46, 0, 306,
		307, 5, 73, 0, 0, 307, 308, 5, 5, 0, 0, 308, 309, 3, 30, 15, 0, 309, 49,
		1, 0, 0, 0, 310, 311, 5, 52, 0, 0, 311, 312, 3, 94, 47, 0, 312, 313, 5,
		73, 0, 0, 313, 314, 5, 53, 0, 0, 314, 315, 3, 96, 48, 0, 315, 316, 5, 73,
		0, 0, 316, 317, 5, 5, 0, 0, 317, 318, 3, 30, 15, 0, 318, 51, 1, 0, 0, 0,
		319, 320, 5, 54, 0, 0, 320, 321, 3, 98, 49, 0, 321, 322, 5, 73, 0, 0, 322,
		323, 5, 55, 0, 0, 323, 324, 3, 100, 50, 0, 324, 325, 5, 73, 0, 0, 325,
		326, 5, 5, 0, 0, 326, 327, 3, 30, 15, 0, 327, 53, 1, 0, 0, 0, 328, 329,
		5, 29, 0, 0, 329, 330, 3, 80, 40, 0, 330, 331, 5, 73, 0, 0, 331, 332, 5,
		5, 0, 0, 332, 333, 3, 30, 15, 0, 333, 55, 1, 0, 0, 0, 334, 335, 7, 3, 0,
		0, 335, 57, 1, 0, 0, 0, 336, 337, 5, 74, 0, 0, 337, 338, 5, 43, 0, 0, 338,
		59, 1, 0, 0, 0, 339, 340, 5, 74, 0, 0, 340, 341, 5, 58, 0, 0, 341, 61,
		1, 0, 0, 0, 342, 343, 7, 4, 0, 0, 343, 63, 1, 0, 0, 0, 344, 345, 5, 74,
		0, 0, 345, 65, 1, 0, 0, 0, 346, 347, 7, 5, 0, 0, 347, 67, 1, 0, 0, 0, 348,
		349, 5, 66, 0, 0, 349, 357, 3, 102, 51, 0, 350, 351, 5, 67, 0, 0, 351,
		357, 3, 104, 52, 0, 352, 353, 5, 68, 0, 0, 353, 357, 3, 106, 53, 0, 354,
		355, 5, 69, 0, 0, 355, 357, 3, 108, 54, 0, 356, 348, 1, 0, 0, 0, 356, 350,
		1, 0, 0, 0, 356, 352, 1, 0, 0, 0, 356, 354, 1, 0, 0, 0, 357, 69, 1, 0,
		0, 0, 358, 359, 7, 6, 0, 0, 359, 71, 1, 0, 0, 0, 360, 361, 5, 74, 0, 0,
		361, 73, 1, 0, 0, 0, 362, 363, 5, 74, 0, 0, 363, 75, 1, 0, 0, 0, 364, 365,
		5, 74, 0, 0, 365, 77, 1, 0, 0, 0, 366, 367, 7, 7, 0, 0, 367, 79, 1, 0,
		0, 0, 368, 369, 5, 74, 0, 0, 369, 81, 1, 0, 0, 0, 370, 371, 5, 74, 0, 0,
		371, 83, 1, 0, 0, 0, 372, 373, 5, 74, 0, 0, 373, 85, 1, 0, 0, 0, 374, 375,
		5, 74, 0, 0, 375, 87, 1, 0, 0, 0, 376, 377, 5, 74, 0, 0, 377, 89, 1, 0,
		0, 0, 378, 379, 5, 74, 0, 0, 379, 380, 5, 43, 0, 0, 380, 91, 1, 0, 0, 0,
		381, 382, 5, 74, 0, 0, 382, 93, 1, 0, 0, 0, 383, 385, 5, 74, 0, 0, 384,
		386, 5, 40, 0, 0, 385, 384, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 95,
		1, 0, 0, 0, 387, 388, 5, 74, 0, 0, 388, 97, 1, 0, 0, 0, 389, 390, 5, 74,
		0, 0, 390, 99, 1, 0, 0, 0, 391, 392, 5, 74, 0, 0, 392, 101, 1, 0, 0, 0,
		393, 394, 5, 74, 0, 0, 394, 103, 1, 0, 0, 0, 395, 396, 5, 74, 0, 0, 396,
		105, 1, 0, 0, 0, 397, 398, 5, 74, 0, 0, 398, 107, 1, 0, 0, 0, 399, 400,
		5, 74, 0, 0, 400, 109, 1, 0, 0, 0, 11, 122, 124, 160, 171, 192, 221, 269,
		293, 299, 356, 385,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SampleGenerationLanguageParserInit initializes any static state used to implement SampleGenerationLanguageParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSampleGenerationLanguageParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SampleGenerationLanguageParserInit() {
	staticData := &samplegenerationlanguageParserStaticData
	staticData.once.Do(samplegenerationlanguageParserInit)
}

// NewSampleGenerationLanguageParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSampleGenerationLanguageParser(input antlr.TokenStream) *SampleGenerationLanguageParser {
	SampleGenerationLanguageParserInit()
	this := new(SampleGenerationLanguageParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &samplegenerationlanguageParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// SampleGenerationLanguageParser tokens.
const (
	SampleGenerationLanguageParserEOF    = antlr.TokenEOF
	SampleGenerationLanguageParserT__0   = 1
	SampleGenerationLanguageParserT__1   = 2
	SampleGenerationLanguageParserT__2   = 3
	SampleGenerationLanguageParserT__3   = 4
	SampleGenerationLanguageParserT__4   = 5
	SampleGenerationLanguageParserT__5   = 6
	SampleGenerationLanguageParserT__6   = 7
	SampleGenerationLanguageParserT__7   = 8
	SampleGenerationLanguageParserT__8   = 9
	SampleGenerationLanguageParserT__9   = 10
	SampleGenerationLanguageParserT__10  = 11
	SampleGenerationLanguageParserT__11  = 12
	SampleGenerationLanguageParserT__12  = 13
	SampleGenerationLanguageParserT__13  = 14
	SampleGenerationLanguageParserT__14  = 15
	SampleGenerationLanguageParserT__15  = 16
	SampleGenerationLanguageParserT__16  = 17
	SampleGenerationLanguageParserT__17  = 18
	SampleGenerationLanguageParserT__18  = 19
	SampleGenerationLanguageParserT__19  = 20
	SampleGenerationLanguageParserT__20  = 21
	SampleGenerationLanguageParserT__21  = 22
	SampleGenerationLanguageParserT__22  = 23
	SampleGenerationLanguageParserT__23  = 24
	SampleGenerationLanguageParserT__24  = 25
	SampleGenerationLanguageParserT__25  = 26
	SampleGenerationLanguageParserT__26  = 27
	SampleGenerationLanguageParserT__27  = 28
	SampleGenerationLanguageParserT__28  = 29
	SampleGenerationLanguageParserT__29  = 30
	SampleGenerationLanguageParserT__30  = 31
	SampleGenerationLanguageParserT__31  = 32
	SampleGenerationLanguageParserT__32  = 33
	SampleGenerationLanguageParserT__33  = 34
	SampleGenerationLanguageParserT__34  = 35
	SampleGenerationLanguageParserT__35  = 36
	SampleGenerationLanguageParserT__36  = 37
	SampleGenerationLanguageParserT__37  = 38
	SampleGenerationLanguageParserT__38  = 39
	SampleGenerationLanguageParserT__39  = 40
	SampleGenerationLanguageParserT__40  = 41
	SampleGenerationLanguageParserT__41  = 42
	SampleGenerationLanguageParserT__42  = 43
	SampleGenerationLanguageParserT__43  = 44
	SampleGenerationLanguageParserT__44  = 45
	SampleGenerationLanguageParserT__45  = 46
	SampleGenerationLanguageParserT__46  = 47
	SampleGenerationLanguageParserT__47  = 48
	SampleGenerationLanguageParserT__48  = 49
	SampleGenerationLanguageParserT__49  = 50
	SampleGenerationLanguageParserT__50  = 51
	SampleGenerationLanguageParserT__51  = 52
	SampleGenerationLanguageParserT__52  = 53
	SampleGenerationLanguageParserT__53  = 54
	SampleGenerationLanguageParserT__54  = 55
	SampleGenerationLanguageParserT__55  = 56
	SampleGenerationLanguageParserT__56  = 57
	SampleGenerationLanguageParserT__57  = 58
	SampleGenerationLanguageParserT__58  = 59
	SampleGenerationLanguageParserT__59  = 60
	SampleGenerationLanguageParserT__60  = 61
	SampleGenerationLanguageParserT__61  = 62
	SampleGenerationLanguageParserT__62  = 63
	SampleGenerationLanguageParserT__63  = 64
	SampleGenerationLanguageParserT__64  = 65
	SampleGenerationLanguageParserT__65  = 66
	SampleGenerationLanguageParserT__66  = 67
	SampleGenerationLanguageParserT__67  = 68
	SampleGenerationLanguageParserT__68  = 69
	SampleGenerationLanguageParserT__69  = 70
	SampleGenerationLanguageParserT__70  = 71
	SampleGenerationLanguageParserT__71  = 72
	SampleGenerationLanguageParserCOMMA  = 73
	SampleGenerationLanguageParserNUMBER = 74
	SampleGenerationLanguageParserWS     = 75
)

// SampleGenerationLanguageParser rules.
const (
	SampleGenerationLanguageParserRULE_prog             = 0
	SampleGenerationLanguageParserRULE_header           = 1
	SampleGenerationLanguageParserRULE_oscillator       = 2
	SampleGenerationLanguageParserRULE_envelope         = 3
	SampleGenerationLanguageParserRULE_effect           = 4
	SampleGenerationLanguageParserRULE_fm               = 5
	SampleGenerationLanguageParserRULE_granular         = 6
	SampleGenerationLanguageParserRULE_physicalModel    = 7
	SampleGenerationLanguageParserRULE_additive         = 8
	SampleGenerationLanguageParserRULE_wavetable        = 9
	SampleGenerationLanguageParserRULE_subtractive      = 10
	SampleGenerationLanguageParserRULE_karplusStrong    = 11
	SampleGenerationLanguageParserRULE_formant          = 12
	SampleGenerationLanguageParserRULE_waveform         = 13
	SampleGenerationLanguageParserRULE_frequency        = 14
	SampleGenerationLanguageParserRULE_mix              = 15
	SampleGenerationLanguageParserRULE_envType          = 16
	SampleGenerationLanguageParserRULE_attackTime       = 17
	SampleGenerationLanguageParserRULE_decayTime        = 18
	SampleGenerationLanguageParserRULE_sustainLevel     = 19
	SampleGenerationLanguageParserRULE_releaseTime      = 20
	SampleGenerationLanguageParserRULE_effectType       = 21
	SampleGenerationLanguageParserRULE_effectParam      = 22
	SampleGenerationLanguageParserRULE_effectReverb     = 23
	SampleGenerationLanguageParserRULE_effectDelay      = 24
	SampleGenerationLanguageParserRULE_effectChorus     = 25
	SampleGenerationLanguageParserRULE_effectDistortion = 26
	SampleGenerationLanguageParserRULE_effectFilter     = 27
	SampleGenerationLanguageParserRULE_fmType           = 28
	SampleGenerationLanguageParserRULE_size             = 29
	SampleGenerationLanguageParserRULE_overlap          = 30
	SampleGenerationLanguageParserRULE_windowFunction   = 31
	SampleGenerationLanguageParserRULE_pitch            = 32
	SampleGenerationLanguageParserRULE_modelType        = 33
	SampleGenerationLanguageParserRULE_modelParam       = 34
	SampleGenerationLanguageParserRULE_excitation       = 35
	SampleGenerationLanguageParserRULE_amplitude        = 36
	SampleGenerationLanguageParserRULE_phase            = 37
	SampleGenerationLanguageParserRULE_index            = 38
	SampleGenerationLanguageParserRULE_filterType       = 39
	SampleGenerationLanguageParserRULE_cutoff           = 40
	SampleGenerationLanguageParserRULE_resonance        = 41
	SampleGenerationLanguageParserRULE_decay            = 42
	SampleGenerationLanguageParserRULE_bandwidth        = 43
	SampleGenerationLanguageParserRULE_gain             = 44
	SampleGenerationLanguageParserRULE_delayTime        = 45
	SampleGenerationLanguageParserRULE_feedback         = 46
	SampleGenerationLanguageParserRULE_rate             = 47
	SampleGenerationLanguageParserRULE_depth            = 48
	SampleGenerationLanguageParserRULE_drive            = 49
	SampleGenerationLanguageParserRULE_tone             = 50
	SampleGenerationLanguageParserRULE_stiffness        = 51
	SampleGenerationLanguageParserRULE_damping          = 52
	SampleGenerationLanguageParserRULE_tension          = 53
	SampleGenerationLanguageParserRULE_length           = 54
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Header() IHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *ProgContext) AllOscillator() []IOscillatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOscillatorContext); ok {
			len++
		}
	}

	tst := make([]IOscillatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOscillatorContext); ok {
			tst[i] = t.(IOscillatorContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Oscillator(i int) IOscillatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOscillatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOscillatorContext)
}

func (s *ProgContext) AllEnvelope() []IEnvelopeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnvelopeContext); ok {
			len++
		}
	}

	tst := make([]IEnvelopeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnvelopeContext); ok {
			tst[i] = t.(IEnvelopeContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Envelope(i int) IEnvelopeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnvelopeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnvelopeContext)
}

func (s *ProgContext) AllEffect() []IEffectContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEffectContext); ok {
			len++
		}
	}

	tst := make([]IEffectContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEffectContext); ok {
			tst[i] = t.(IEffectContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Effect(i int) IEffectContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEffectContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEffectContext)
}

func (s *ProgContext) AllFm() []IFmContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFmContext); ok {
			len++
		}
	}

	tst := make([]IFmContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFmContext); ok {
			tst[i] = t.(IFmContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Fm(i int) IFmContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFmContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFmContext)
}

func (s *ProgContext) AllGranular() []IGranularContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGranularContext); ok {
			len++
		}
	}

	tst := make([]IGranularContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGranularContext); ok {
			tst[i] = t.(IGranularContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Granular(i int) IGranularContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGranularContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGranularContext)
}

func (s *ProgContext) AllPhysicalModel() []IPhysicalModelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPhysicalModelContext); ok {
			len++
		}
	}

	tst := make([]IPhysicalModelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPhysicalModelContext); ok {
			tst[i] = t.(IPhysicalModelContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) PhysicalModel(i int) IPhysicalModelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPhysicalModelContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPhysicalModelContext)
}

func (s *ProgContext) AllAdditive() []IAdditiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveContext); ok {
			tst[i] = t.(IAdditiveContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Additive(i int) IAdditiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveContext)
}

func (s *ProgContext) AllWavetable() []IWavetableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWavetableContext); ok {
			len++
		}
	}

	tst := make([]IWavetableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWavetableContext); ok {
			tst[i] = t.(IWavetableContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Wavetable(i int) IWavetableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWavetableContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWavetableContext)
}

func (s *ProgContext) AllSubtractive() []ISubtractiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubtractiveContext); ok {
			len++
		}
	}

	tst := make([]ISubtractiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubtractiveContext); ok {
			tst[i] = t.(ISubtractiveContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Subtractive(i int) ISubtractiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubtractiveContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubtractiveContext)
}

func (s *ProgContext) AllKarplusStrong() []IKarplusStrongContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKarplusStrongContext); ok {
			len++
		}
	}

	tst := make([]IKarplusStrongContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKarplusStrongContext); ok {
			tst[i] = t.(IKarplusStrongContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) KarplusStrong(i int) IKarplusStrongContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKarplusStrongContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKarplusStrongContext)
}

func (s *ProgContext) AllFormant() []IFormantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFormantContext); ok {
			len++
		}
	}

	tst := make([]IFormantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFormantContext); ok {
			tst[i] = t.(IFormantContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Formant(i int) IFormantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormantContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SampleGenerationLanguageParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Header()
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6594771024) != 0 {
		p.SetState(122)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SampleGenerationLanguageParserT__3:
			{
				p.SetState(111)
				p.Oscillator()
			}

		case SampleGenerationLanguageParserT__5:
			{
				p.SetState(112)
				p.Envelope()
			}

		case SampleGenerationLanguageParserT__10:
			{
				p.SetState(113)
				p.Effect()
			}

		case SampleGenerationLanguageParserT__11:
			{
				p.SetState(114)
				p.Fm()
			}

		case SampleGenerationLanguageParserT__12:
			{
				p.SetState(115)
				p.Granular()
			}

		case SampleGenerationLanguageParserT__17:
			{
				p.SetState(116)
				p.PhysicalModel()
			}

		case SampleGenerationLanguageParserT__19:
			{
				p.SetState(117)
				p.Additive()
			}

		case SampleGenerationLanguageParserT__23:
			{
				p.SetState(118)
				p.Wavetable()
			}

		case SampleGenerationLanguageParserT__26:
			{
				p.SetState(119)
				p.Subtractive()
			}

		case SampleGenerationLanguageParserT__30:
			{
				p.SetState(120)
				p.KarplusStrong()
			}

		case SampleGenerationLanguageParserT__31:
			{
				p.SetState(121)
				p.Formant()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_header
	return p
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (s *HeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Header() (localctx IHeaderContext) {
	this := p
	_ = this

	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SampleGenerationLanguageParserRULE_header)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(SampleGenerationLanguageParserT__0)
	}
	{
		p.SetState(127)
		p.Match(SampleGenerationLanguageParserT__1)
	}
	{
		p.SetState(128)
		p.Match(SampleGenerationLanguageParserT__2)
	}
	{
		p.SetState(129)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IOscillatorContext is an interface to support dynamic dispatch.
type IOscillatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOscillatorContext differentiates from other interfaces.
	IsOscillatorContext()
}

type OscillatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOscillatorContext() *OscillatorContext {
	var p = new(OscillatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_oscillator
	return p
}

func (*OscillatorContext) IsOscillatorContext() {}

func NewOscillatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OscillatorContext {
	var p = new(OscillatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_oscillator

	return p
}

func (s *OscillatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OscillatorContext) Waveform() IWaveformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWaveformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWaveformContext)
}

func (s *OscillatorContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *OscillatorContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *OscillatorContext) Frequency() IFrequencyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrequencyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrequencyContext)
}

func (s *OscillatorContext) Mix() IMixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMixContext)
}

func (s *OscillatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OscillatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OscillatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterOscillator(s)
	}
}

func (s *OscillatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitOscillator(s)
	}
}

func (s *OscillatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitOscillator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Oscillator() (localctx IOscillatorContext) {
	this := p
	_ = this

	localctx = NewOscillatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SampleGenerationLanguageParserRULE_oscillator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(SampleGenerationLanguageParserT__3)
	}
	{
		p.SetState(132)
		p.Waveform()
	}
	{
		p.SetState(133)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(134)
		p.Frequency()
	}
	{
		p.SetState(135)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(136)
		p.Match(SampleGenerationLanguageParserT__4)
	}
	{
		p.SetState(137)
		p.Mix()
	}

	return localctx
}

// IEnvelopeContext is an interface to support dynamic dispatch.
type IEnvelopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnvelopeContext differentiates from other interfaces.
	IsEnvelopeContext()
}

type EnvelopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnvelopeContext() *EnvelopeContext {
	var p = new(EnvelopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_envelope
	return p
}

func (*EnvelopeContext) IsEnvelopeContext() {}

func NewEnvelopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvelopeContext {
	var p = new(EnvelopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_envelope

	return p
}

func (s *EnvelopeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnvelopeContext) EnvType() IEnvTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnvTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnvTypeContext)
}

func (s *EnvelopeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *EnvelopeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *EnvelopeContext) AttackTime() IAttackTimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttackTimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttackTimeContext)
}

func (s *EnvelopeContext) DecayTime() IDecayTimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecayTimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecayTimeContext)
}

func (s *EnvelopeContext) SustainLevel() ISustainLevelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISustainLevelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISustainLevelContext)
}

func (s *EnvelopeContext) ReleaseTime() IReleaseTimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReleaseTimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReleaseTimeContext)
}

func (s *EnvelopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnvelopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnvelopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterEnvelope(s)
	}
}

func (s *EnvelopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitEnvelope(s)
	}
}

func (s *EnvelopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitEnvelope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Envelope() (localctx IEnvelopeContext) {
	this := p
	_ = this

	localctx = NewEnvelopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SampleGenerationLanguageParserRULE_envelope)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(SampleGenerationLanguageParserT__5)
	}
	{
		p.SetState(140)
		p.EnvType()
	}
	{
		p.SetState(141)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(142)
		p.Match(SampleGenerationLanguageParserT__6)
	}
	{
		p.SetState(143)
		p.AttackTime()
	}
	{
		p.SetState(144)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(145)
		p.Match(SampleGenerationLanguageParserT__7)
	}
	{
		p.SetState(146)
		p.DecayTime()
	}
	{
		p.SetState(147)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(148)
		p.Match(SampleGenerationLanguageParserT__8)
	}
	{
		p.SetState(149)
		p.SustainLevel()
	}
	{
		p.SetState(150)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(151)
		p.Match(SampleGenerationLanguageParserT__9)
	}
	{
		p.SetState(152)
		p.ReleaseTime()
	}

	return localctx
}

// IEffectContext is an interface to support dynamic dispatch.
type IEffectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEffectContext differentiates from other interfaces.
	IsEffectContext()
}

type EffectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEffectContext() *EffectContext {
	var p = new(EffectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_effect
	return p
}

func (*EffectContext) IsEffectContext() {}

func NewEffectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EffectContext {
	var p = new(EffectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_effect

	return p
}

func (s *EffectContext) GetParser() antlr.Parser { return s.parser }

func (s *EffectContext) EffectType() IEffectTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEffectTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEffectTypeContext)
}

func (s *EffectContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *EffectContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *EffectContext) AllEffectParam() []IEffectParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEffectParamContext); ok {
			len++
		}
	}

	tst := make([]IEffectParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEffectParamContext); ok {
			tst[i] = t.(IEffectParamContext)
			i++
		}
	}

	return tst
}

func (s *EffectContext) EffectParam(i int) IEffectParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEffectParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEffectParamContext)
}

func (s *EffectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EffectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EffectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterEffect(s)
	}
}

func (s *EffectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitEffect(s)
	}
}

func (s *EffectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitEffect(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Effect() (localctx IEffectContext) {
	this := p
	_ = this

	localctx = NewEffectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SampleGenerationLanguageParserRULE_effect)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(SampleGenerationLanguageParserT__10)
	}
	{
		p.SetState(155)
		p.EffectType()
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SampleGenerationLanguageParserCOMMA {
		{
			p.SetState(156)
			p.Match(SampleGenerationLanguageParserCOMMA)
		}
		{
			p.SetState(157)
			p.EffectParam()
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFmContext is an interface to support dynamic dispatch.
type IFmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFmContext differentiates from other interfaces.
	IsFmContext()
}

type FmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFmContext() *FmContext {
	var p = new(FmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_fm
	return p
}

func (*FmContext) IsFmContext() {}

func NewFmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FmContext {
	var p = new(FmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_fm

	return p
}

func (s *FmContext) GetParser() antlr.Parser { return s.parser }

func (s *FmContext) FmType() IFmTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFmTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFmTypeContext)
}

func (s *FmContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *FmContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *FmContext) Waveform() IWaveformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWaveformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWaveformContext)
}

func (s *FmContext) Frequency() IFrequencyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrequencyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrequencyContext)
}

func (s *FmContext) Mix() IMixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMixContext)
}

func (s *FmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterFm(s)
	}
}

func (s *FmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitFm(s)
	}
}

func (s *FmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitFm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Fm() (localctx IFmContext) {
	this := p
	_ = this

	localctx = NewFmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SampleGenerationLanguageParserRULE_fm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(SampleGenerationLanguageParserT__11)
	}
	{
		p.SetState(163)
		p.FmType()
	}
	{
		p.SetState(164)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(165)
		p.Waveform()
	}
	{
		p.SetState(166)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(167)
		p.Frequency()
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SampleGenerationLanguageParserCOMMA {
		{
			p.SetState(168)
			p.Match(SampleGenerationLanguageParserCOMMA)
		}
		{
			p.SetState(169)
			p.Match(SampleGenerationLanguageParserT__4)
		}
		{
			p.SetState(170)
			p.Mix()
		}

	}

	return localctx
}

// IGranularContext is an interface to support dynamic dispatch.
type IGranularContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGranularContext differentiates from other interfaces.
	IsGranularContext()
}

type GranularContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGranularContext() *GranularContext {
	var p = new(GranularContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_granular
	return p
}

func (*GranularContext) IsGranularContext() {}

func NewGranularContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GranularContext {
	var p = new(GranularContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_granular

	return p
}

func (s *GranularContext) GetParser() antlr.Parser { return s.parser }

func (s *GranularContext) Size() ISizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISizeContext)
}

func (s *GranularContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *GranularContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *GranularContext) Overlap() IOverlapContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOverlapContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOverlapContext)
}

func (s *GranularContext) WindowFunction() IWindowFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWindowFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWindowFunctionContext)
}

func (s *GranularContext) Pitch() IPitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPitchContext)
}

func (s *GranularContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GranularContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GranularContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterGranular(s)
	}
}

func (s *GranularContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitGranular(s)
	}
}

func (s *GranularContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitGranular(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Granular() (localctx IGranularContext) {
	this := p
	_ = this

	localctx = NewGranularContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SampleGenerationLanguageParserRULE_granular)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(SampleGenerationLanguageParserT__12)
	}
	{
		p.SetState(174)
		p.Match(SampleGenerationLanguageParserT__13)
	}
	{
		p.SetState(175)
		p.Size()
	}
	{
		p.SetState(176)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(177)
		p.Match(SampleGenerationLanguageParserT__14)
	}
	{
		p.SetState(178)
		p.Overlap()
	}
	{
		p.SetState(179)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(180)
		p.Match(SampleGenerationLanguageParserT__15)
	}
	{
		p.SetState(181)
		p.WindowFunction()
	}
	{
		p.SetState(182)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(183)
		p.Match(SampleGenerationLanguageParserT__16)
	}
	{
		p.SetState(184)
		p.Pitch()
	}

	return localctx
}

// IPhysicalModelContext is an interface to support dynamic dispatch.
type IPhysicalModelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPhysicalModelContext differentiates from other interfaces.
	IsPhysicalModelContext()
}

type PhysicalModelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPhysicalModelContext() *PhysicalModelContext {
	var p = new(PhysicalModelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_physicalModel
	return p
}

func (*PhysicalModelContext) IsPhysicalModelContext() {}

func NewPhysicalModelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PhysicalModelContext {
	var p = new(PhysicalModelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_physicalModel

	return p
}

func (s *PhysicalModelContext) GetParser() antlr.Parser { return s.parser }

func (s *PhysicalModelContext) ModelType() IModelTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModelTypeContext)
}

func (s *PhysicalModelContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *PhysicalModelContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *PhysicalModelContext) Excitation() IExcitationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExcitationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExcitationContext)
}

func (s *PhysicalModelContext) Mix() IMixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMixContext)
}

func (s *PhysicalModelContext) AllModelParam() []IModelParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IModelParamContext); ok {
			len++
		}
	}

	tst := make([]IModelParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IModelParamContext); ok {
			tst[i] = t.(IModelParamContext)
			i++
		}
	}

	return tst
}

func (s *PhysicalModelContext) ModelParam(i int) IModelParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModelParamContext)
}

func (s *PhysicalModelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhysicalModelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PhysicalModelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterPhysicalModel(s)
	}
}

func (s *PhysicalModelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitPhysicalModel(s)
	}
}

func (s *PhysicalModelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitPhysicalModel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) PhysicalModel() (localctx IPhysicalModelContext) {
	this := p
	_ = this

	localctx = NewPhysicalModelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SampleGenerationLanguageParserRULE_physicalModel)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(SampleGenerationLanguageParserT__17)
	}
	{
		p.SetState(187)
		p.ModelType()
	}
	{
		p.SetState(188)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&15) != 0 {
		{
			p.SetState(189)
			p.ModelParam()
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(195)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(196)
		p.Match(SampleGenerationLanguageParserT__18)
	}
	{
		p.SetState(197)
		p.Excitation()
	}
	{
		p.SetState(198)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(199)
		p.Match(SampleGenerationLanguageParserT__4)
	}
	{
		p.SetState(200)
		p.Mix()
	}

	return localctx
}

// IAdditiveContext is an interface to support dynamic dispatch.
type IAdditiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveContext differentiates from other interfaces.
	IsAdditiveContext()
}

type AdditiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveContext() *AdditiveContext {
	var p = new(AdditiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_additive
	return p
}

func (*AdditiveContext) IsAdditiveContext() {}

func NewAdditiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveContext {
	var p = new(AdditiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_additive

	return p
}

func (s *AdditiveContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveContext) Frequency() IFrequencyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrequencyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrequencyContext)
}

func (s *AdditiveContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *AdditiveContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *AdditiveContext) Amplitude() IAmplitudeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAmplitudeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAmplitudeContext)
}

func (s *AdditiveContext) Phase() IPhaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPhaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPhaseContext)
}

func (s *AdditiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterAdditive(s)
	}
}

func (s *AdditiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitAdditive(s)
	}
}

func (s *AdditiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitAdditive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Additive() (localctx IAdditiveContext) {
	this := p
	_ = this

	localctx = NewAdditiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SampleGenerationLanguageParserRULE_additive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(SampleGenerationLanguageParserT__19)
	}
	{
		p.SetState(203)
		p.Match(SampleGenerationLanguageParserT__20)
	}
	{
		p.SetState(204)
		p.Frequency()
	}
	{
		p.SetState(205)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(206)
		p.Match(SampleGenerationLanguageParserT__21)
	}
	{
		p.SetState(207)
		p.Amplitude()
	}
	{
		p.SetState(208)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(209)
		p.Match(SampleGenerationLanguageParserT__22)
	}
	{
		p.SetState(210)
		p.Phase()
	}

	return localctx
}

// IWavetableContext is an interface to support dynamic dispatch.
type IWavetableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWavetableContext differentiates from other interfaces.
	IsWavetableContext()
}

type WavetableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWavetableContext() *WavetableContext {
	var p = new(WavetableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_wavetable
	return p
}

func (*WavetableContext) IsWavetableContext() {}

func NewWavetableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WavetableContext {
	var p = new(WavetableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_wavetable

	return p
}

func (s *WavetableContext) GetParser() antlr.Parser { return s.parser }

func (s *WavetableContext) Waveform() IWaveformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWaveformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWaveformContext)
}

func (s *WavetableContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *WavetableContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *WavetableContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *WavetableContext) Mix() IMixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMixContext)
}

func (s *WavetableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WavetableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WavetableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterWavetable(s)
	}
}

func (s *WavetableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitWavetable(s)
	}
}

func (s *WavetableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitWavetable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Wavetable() (localctx IWavetableContext) {
	this := p
	_ = this

	localctx = NewWavetableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SampleGenerationLanguageParserRULE_wavetable)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(SampleGenerationLanguageParserT__23)
	}
	{
		p.SetState(213)
		p.Match(SampleGenerationLanguageParserT__24)
	}
	{
		p.SetState(214)
		p.Waveform()
	}
	{
		p.SetState(215)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(216)
		p.Match(SampleGenerationLanguageParserT__25)
	}
	{
		p.SetState(217)
		p.Index()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SampleGenerationLanguageParserCOMMA {
		{
			p.SetState(218)
			p.Match(SampleGenerationLanguageParserCOMMA)
		}
		{
			p.SetState(219)
			p.Match(SampleGenerationLanguageParserT__4)
		}
		{
			p.SetState(220)
			p.Mix()
		}

	}

	return localctx
}

// ISubtractiveContext is an interface to support dynamic dispatch.
type ISubtractiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubtractiveContext differentiates from other interfaces.
	IsSubtractiveContext()
}

type SubtractiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubtractiveContext() *SubtractiveContext {
	var p = new(SubtractiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_subtractive
	return p
}

func (*SubtractiveContext) IsSubtractiveContext() {}

func NewSubtractiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubtractiveContext {
	var p = new(SubtractiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_subtractive

	return p
}

func (s *SubtractiveContext) GetParser() antlr.Parser { return s.parser }

func (s *SubtractiveContext) Waveform() IWaveformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWaveformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWaveformContext)
}

func (s *SubtractiveContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *SubtractiveContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *SubtractiveContext) Frequency() IFrequencyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrequencyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrequencyContext)
}

func (s *SubtractiveContext) FilterType() IFilterTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterTypeContext)
}

func (s *SubtractiveContext) Cutoff() ICutoffContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICutoffContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICutoffContext)
}

func (s *SubtractiveContext) Resonance() IResonanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResonanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResonanceContext)
}

func (s *SubtractiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtractiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubtractiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterSubtractive(s)
	}
}

func (s *SubtractiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitSubtractive(s)
	}
}

func (s *SubtractiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitSubtractive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Subtractive() (localctx ISubtractiveContext) {
	this := p
	_ = this

	localctx = NewSubtractiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SampleGenerationLanguageParserRULE_subtractive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(SampleGenerationLanguageParserT__26)
	}
	{
		p.SetState(224)
		p.Match(SampleGenerationLanguageParserT__24)
	}
	{
		p.SetState(225)
		p.Waveform()
	}
	{
		p.SetState(226)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(227)
		p.Match(SampleGenerationLanguageParserT__20)
	}
	{
		p.SetState(228)
		p.Frequency()
	}
	{
		p.SetState(229)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(230)
		p.Match(SampleGenerationLanguageParserT__27)
	}
	{
		p.SetState(231)
		p.FilterType()
	}
	{
		p.SetState(232)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(233)
		p.Match(SampleGenerationLanguageParserT__28)
	}
	{
		p.SetState(234)
		p.Cutoff()
	}
	{
		p.SetState(235)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(236)
		p.Match(SampleGenerationLanguageParserT__29)
	}
	{
		p.SetState(237)
		p.Resonance()
	}

	return localctx
}

// IKarplusStrongContext is an interface to support dynamic dispatch.
type IKarplusStrongContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKarplusStrongContext differentiates from other interfaces.
	IsKarplusStrongContext()
}

type KarplusStrongContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKarplusStrongContext() *KarplusStrongContext {
	var p = new(KarplusStrongContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_karplusStrong
	return p
}

func (*KarplusStrongContext) IsKarplusStrongContext() {}

func NewKarplusStrongContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KarplusStrongContext {
	var p = new(KarplusStrongContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_karplusStrong

	return p
}

func (s *KarplusStrongContext) GetParser() antlr.Parser { return s.parser }

func (s *KarplusStrongContext) Frequency() IFrequencyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrequencyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrequencyContext)
}

func (s *KarplusStrongContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *KarplusStrongContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *KarplusStrongContext) Decay() IDecayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecayContext)
}

func (s *KarplusStrongContext) FilterType() IFilterTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterTypeContext)
}

func (s *KarplusStrongContext) Cutoff() ICutoffContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICutoffContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICutoffContext)
}

func (s *KarplusStrongContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KarplusStrongContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KarplusStrongContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterKarplusStrong(s)
	}
}

func (s *KarplusStrongContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitKarplusStrong(s)
	}
}

func (s *KarplusStrongContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitKarplusStrong(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) KarplusStrong() (localctx IKarplusStrongContext) {
	this := p
	_ = this

	localctx = NewKarplusStrongContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SampleGenerationLanguageParserRULE_karplusStrong)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(SampleGenerationLanguageParserT__30)
	}
	{
		p.SetState(240)
		p.Match(SampleGenerationLanguageParserT__20)
	}
	{
		p.SetState(241)
		p.Frequency()
	}
	{
		p.SetState(242)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(243)
		p.Match(SampleGenerationLanguageParserT__7)
	}
	{
		p.SetState(244)
		p.Decay()
	}
	{
		p.SetState(245)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(246)
		p.Match(SampleGenerationLanguageParserT__27)
	}
	{
		p.SetState(247)
		p.FilterType()
	}
	{
		p.SetState(248)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(249)
		p.Match(SampleGenerationLanguageParserT__28)
	}
	{
		p.SetState(250)
		p.Cutoff()
	}

	return localctx
}

// IFormantContext is an interface to support dynamic dispatch.
type IFormantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormantContext differentiates from other interfaces.
	IsFormantContext()
}

type FormantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormantContext() *FormantContext {
	var p = new(FormantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_formant
	return p
}

func (*FormantContext) IsFormantContext() {}

func NewFormantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormantContext {
	var p = new(FormantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_formant

	return p
}

func (s *FormantContext) GetParser() antlr.Parser { return s.parser }

func (s *FormantContext) FilterType() IFilterTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterTypeContext)
}

func (s *FormantContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *FormantContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *FormantContext) Frequency() IFrequencyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrequencyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrequencyContext)
}

func (s *FormantContext) Bandwidth() IBandwidthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBandwidthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBandwidthContext)
}

func (s *FormantContext) Gain() IGainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGainContext)
}

func (s *FormantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterFormant(s)
	}
}

func (s *FormantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitFormant(s)
	}
}

func (s *FormantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitFormant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Formant() (localctx IFormantContext) {
	this := p
	_ = this

	localctx = NewFormantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SampleGenerationLanguageParserRULE_formant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(SampleGenerationLanguageParserT__31)
	}
	{
		p.SetState(253)
		p.Match(SampleGenerationLanguageParserT__27)
	}
	{
		p.SetState(254)
		p.FilterType()
	}
	{
		p.SetState(255)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(256)
		p.Match(SampleGenerationLanguageParserT__20)
	}
	{
		p.SetState(257)
		p.Frequency()
	}
	{
		p.SetState(258)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(259)
		p.Match(SampleGenerationLanguageParserT__32)
	}
	{
		p.SetState(260)
		p.Bandwidth()
	}
	{
		p.SetState(261)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(262)
		p.Match(SampleGenerationLanguageParserT__33)
	}
	{
		p.SetState(263)
		p.Gain()
	}

	return localctx
}

// IWaveformContext is an interface to support dynamic dispatch.
type IWaveformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWaveformContext differentiates from other interfaces.
	IsWaveformContext()
}

type WaveformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWaveformContext() *WaveformContext {
	var p = new(WaveformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_waveform
	return p
}

func (*WaveformContext) IsWaveformContext() {}

func NewWaveformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WaveformContext {
	var p = new(WaveformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_waveform

	return p
}

func (s *WaveformContext) GetParser() antlr.Parser { return s.parser }
func (s *WaveformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WaveformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WaveformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterWaveform(s)
	}
}

func (s *WaveformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitWaveform(s)
	}
}

func (s *WaveformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitWaveform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Waveform() (localctx IWaveformContext) {
	this := p
	_ = this

	localctx = NewWaveformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SampleGenerationLanguageParserRULE_waveform)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1065151889408) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFrequencyContext is an interface to support dynamic dispatch.
type IFrequencyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFrequencyContext differentiates from other interfaces.
	IsFrequencyContext()
}

type FrequencyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrequencyContext() *FrequencyContext {
	var p = new(FrequencyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_frequency
	return p
}

func (*FrequencyContext) IsFrequencyContext() {}

func NewFrequencyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FrequencyContext {
	var p = new(FrequencyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_frequency

	return p
}

func (s *FrequencyContext) GetParser() antlr.Parser { return s.parser }

func (s *FrequencyContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *FrequencyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FrequencyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FrequencyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterFrequency(s)
	}
}

func (s *FrequencyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitFrequency(s)
	}
}

func (s *FrequencyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitFrequency(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Frequency() (localctx IFrequencyContext) {
	this := p
	_ = this

	localctx = NewFrequencyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SampleGenerationLanguageParserRULE_frequency)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SampleGenerationLanguageParserT__39 {
		{
			p.SetState(268)
			p.Match(SampleGenerationLanguageParserT__39)
		}

	}

	return localctx
}

// IMixContext is an interface to support dynamic dispatch.
type IMixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMixContext differentiates from other interfaces.
	IsMixContext()
}

type MixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMixContext() *MixContext {
	var p = new(MixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_mix
	return p
}

func (*MixContext) IsMixContext() {}

func NewMixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MixContext {
	var p = new(MixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_mix

	return p
}

func (s *MixContext) GetParser() antlr.Parser { return s.parser }

func (s *MixContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *MixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterMix(s)
	}
}

func (s *MixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitMix(s)
	}
}

func (s *MixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitMix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Mix() (localctx IMixContext) {
	this := p
	_ = this

	localctx = NewMixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SampleGenerationLanguageParserRULE_mix)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IEnvTypeContext is an interface to support dynamic dispatch.
type IEnvTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnvTypeContext differentiates from other interfaces.
	IsEnvTypeContext()
}

type EnvTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnvTypeContext() *EnvTypeContext {
	var p = new(EnvTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_envType
	return p
}

func (*EnvTypeContext) IsEnvTypeContext() {}

func NewEnvTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvTypeContext {
	var p = new(EnvTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_envType

	return p
}

func (s *EnvTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *EnvTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnvTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnvTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterEnvType(s)
	}
}

func (s *EnvTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitEnvType(s)
	}
}

func (s *EnvTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitEnvType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) EnvType() (localctx IEnvTypeContext) {
	this := p
	_ = this

	localctx = NewEnvTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SampleGenerationLanguageParserRULE_envType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SampleGenerationLanguageParserT__40 || _la == SampleGenerationLanguageParserT__41) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAttackTimeContext is an interface to support dynamic dispatch.
type IAttackTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttackTimeContext differentiates from other interfaces.
	IsAttackTimeContext()
}

type AttackTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttackTimeContext() *AttackTimeContext {
	var p = new(AttackTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_attackTime
	return p
}

func (*AttackTimeContext) IsAttackTimeContext() {}

func NewAttackTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttackTimeContext {
	var p = new(AttackTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_attackTime

	return p
}

func (s *AttackTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttackTimeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *AttackTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttackTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttackTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterAttackTime(s)
	}
}

func (s *AttackTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitAttackTime(s)
	}
}

func (s *AttackTimeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitAttackTime(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) AttackTime() (localctx IAttackTimeContext) {
	this := p
	_ = this

	localctx = NewAttackTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SampleGenerationLanguageParserRULE_attackTime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}
	{
		p.SetState(276)
		p.Match(SampleGenerationLanguageParserT__42)
	}

	return localctx
}

// IDecayTimeContext is an interface to support dynamic dispatch.
type IDecayTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecayTimeContext differentiates from other interfaces.
	IsDecayTimeContext()
}

type DecayTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecayTimeContext() *DecayTimeContext {
	var p = new(DecayTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_decayTime
	return p
}

func (*DecayTimeContext) IsDecayTimeContext() {}

func NewDecayTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecayTimeContext {
	var p = new(DecayTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_decayTime

	return p
}

func (s *DecayTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *DecayTimeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *DecayTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecayTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecayTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterDecayTime(s)
	}
}

func (s *DecayTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitDecayTime(s)
	}
}

func (s *DecayTimeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitDecayTime(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) DecayTime() (localctx IDecayTimeContext) {
	this := p
	_ = this

	localctx = NewDecayTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SampleGenerationLanguageParserRULE_decayTime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}
	{
		p.SetState(279)
		p.Match(SampleGenerationLanguageParserT__42)
	}

	return localctx
}

// ISustainLevelContext is an interface to support dynamic dispatch.
type ISustainLevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSustainLevelContext differentiates from other interfaces.
	IsSustainLevelContext()
}

type SustainLevelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySustainLevelContext() *SustainLevelContext {
	var p = new(SustainLevelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_sustainLevel
	return p
}

func (*SustainLevelContext) IsSustainLevelContext() {}

func NewSustainLevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SustainLevelContext {
	var p = new(SustainLevelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_sustainLevel

	return p
}

func (s *SustainLevelContext) GetParser() antlr.Parser { return s.parser }

func (s *SustainLevelContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *SustainLevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SustainLevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SustainLevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterSustainLevel(s)
	}
}

func (s *SustainLevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitSustainLevel(s)
	}
}

func (s *SustainLevelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitSustainLevel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) SustainLevel() (localctx ISustainLevelContext) {
	this := p
	_ = this

	localctx = NewSustainLevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SampleGenerationLanguageParserRULE_sustainLevel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IReleaseTimeContext is an interface to support dynamic dispatch.
type IReleaseTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReleaseTimeContext differentiates from other interfaces.
	IsReleaseTimeContext()
}

type ReleaseTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReleaseTimeContext() *ReleaseTimeContext {
	var p = new(ReleaseTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_releaseTime
	return p
}

func (*ReleaseTimeContext) IsReleaseTimeContext() {}

func NewReleaseTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReleaseTimeContext {
	var p = new(ReleaseTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_releaseTime

	return p
}

func (s *ReleaseTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReleaseTimeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *ReleaseTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReleaseTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReleaseTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterReleaseTime(s)
	}
}

func (s *ReleaseTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitReleaseTime(s)
	}
}

func (s *ReleaseTimeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitReleaseTime(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) ReleaseTime() (localctx IReleaseTimeContext) {
	this := p
	_ = this

	localctx = NewReleaseTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SampleGenerationLanguageParserRULE_releaseTime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}
	{
		p.SetState(284)
		p.Match(SampleGenerationLanguageParserT__42)
	}

	return localctx
}

// IEffectTypeContext is an interface to support dynamic dispatch.
type IEffectTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEffectTypeContext differentiates from other interfaces.
	IsEffectTypeContext()
}

type EffectTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEffectTypeContext() *EffectTypeContext {
	var p = new(EffectTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectType
	return p
}

func (*EffectTypeContext) IsEffectTypeContext() {}

func NewEffectTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EffectTypeContext {
	var p = new(EffectTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectType

	return p
}

func (s *EffectTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *EffectTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EffectTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EffectTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterEffectType(s)
	}
}

func (s *EffectTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitEffectType(s)
	}
}

func (s *EffectTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitEffectType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) EffectType() (localctx IEffectTypeContext) {
	this := p
	_ = this

	localctx = NewEffectTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SampleGenerationLanguageParserRULE_effectType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1108857476612096) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEffectParamContext is an interface to support dynamic dispatch.
type IEffectParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEffectParamContext differentiates from other interfaces.
	IsEffectParamContext()
}

type EffectParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEffectParamContext() *EffectParamContext {
	var p = new(EffectParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectParam
	return p
}

func (*EffectParamContext) IsEffectParamContext() {}

func NewEffectParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EffectParamContext {
	var p = new(EffectParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectParam

	return p
}

func (s *EffectParamContext) GetParser() antlr.Parser { return s.parser }

func (s *EffectParamContext) EffectReverb() IEffectReverbContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEffectReverbContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEffectReverbContext)
}

func (s *EffectParamContext) EffectDelay() IEffectDelayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEffectDelayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEffectDelayContext)
}

func (s *EffectParamContext) EffectChorus() IEffectChorusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEffectChorusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEffectChorusContext)
}

func (s *EffectParamContext) EffectDistortion() IEffectDistortionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEffectDistortionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEffectDistortionContext)
}

func (s *EffectParamContext) EffectFilter() IEffectFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEffectFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEffectFilterContext)
}

func (s *EffectParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EffectParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EffectParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterEffectParam(s)
	}
}

func (s *EffectParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitEffectParam(s)
	}
}

func (s *EffectParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitEffectParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) EffectParam() (localctx IEffectParamContext) {
	this := p
	_ = this

	localctx = NewEffectParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SampleGenerationLanguageParserRULE_effectParam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(293)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SampleGenerationLanguageParserT__4, SampleGenerationLanguageParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(288)
			p.EffectReverb()
		}

	case SampleGenerationLanguageParserT__49:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(289)
			p.EffectDelay()
		}

	case SampleGenerationLanguageParserT__51:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(290)
			p.EffectChorus()
		}

	case SampleGenerationLanguageParserT__53:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(291)
			p.EffectDistortion()
		}

	case SampleGenerationLanguageParserT__28:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(292)
			p.EffectFilter()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEffectReverbContext is an interface to support dynamic dispatch.
type IEffectReverbContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEffectReverbContext differentiates from other interfaces.
	IsEffectReverbContext()
}

type EffectReverbContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEffectReverbContext() *EffectReverbContext {
	var p = new(EffectReverbContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectReverb
	return p
}

func (*EffectReverbContext) IsEffectReverbContext() {}

func NewEffectReverbContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EffectReverbContext {
	var p = new(EffectReverbContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectReverb

	return p
}

func (s *EffectReverbContext) GetParser() antlr.Parser { return s.parser }

func (s *EffectReverbContext) DecayTime() IDecayTimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecayTimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecayTimeContext)
}

func (s *EffectReverbContext) Mix() IMixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMixContext)
}

func (s *EffectReverbContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EffectReverbContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EffectReverbContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterEffectReverb(s)
	}
}

func (s *EffectReverbContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitEffectReverb(s)
	}
}

func (s *EffectReverbContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitEffectReverb(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) EffectReverb() (localctx IEffectReverbContext) {
	this := p
	_ = this

	localctx = NewEffectReverbContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SampleGenerationLanguageParserRULE_effectReverb)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(299)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SampleGenerationLanguageParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(295)
			p.Match(SampleGenerationLanguageParserT__7)
		}
		{
			p.SetState(296)
			p.DecayTime()
		}

	case SampleGenerationLanguageParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(297)
			p.Match(SampleGenerationLanguageParserT__4)
		}
		{
			p.SetState(298)
			p.Mix()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEffectDelayContext is an interface to support dynamic dispatch.
type IEffectDelayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEffectDelayContext differentiates from other interfaces.
	IsEffectDelayContext()
}

type EffectDelayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEffectDelayContext() *EffectDelayContext {
	var p = new(EffectDelayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectDelay
	return p
}

func (*EffectDelayContext) IsEffectDelayContext() {}

func NewEffectDelayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EffectDelayContext {
	var p = new(EffectDelayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectDelay

	return p
}

func (s *EffectDelayContext) GetParser() antlr.Parser { return s.parser }

func (s *EffectDelayContext) DelayTime() IDelayTimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDelayTimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDelayTimeContext)
}

func (s *EffectDelayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *EffectDelayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *EffectDelayContext) Feedback() IFeedbackContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFeedbackContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFeedbackContext)
}

func (s *EffectDelayContext) Mix() IMixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMixContext)
}

func (s *EffectDelayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EffectDelayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EffectDelayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterEffectDelay(s)
	}
}

func (s *EffectDelayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitEffectDelay(s)
	}
}

func (s *EffectDelayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitEffectDelay(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) EffectDelay() (localctx IEffectDelayContext) {
	this := p
	_ = this

	localctx = NewEffectDelayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SampleGenerationLanguageParserRULE_effectDelay)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(SampleGenerationLanguageParserT__49)
	}
	{
		p.SetState(302)
		p.DelayTime()
	}
	{
		p.SetState(303)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(304)
		p.Match(SampleGenerationLanguageParserT__50)
	}
	{
		p.SetState(305)
		p.Feedback()
	}
	{
		p.SetState(306)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(307)
		p.Match(SampleGenerationLanguageParserT__4)
	}
	{
		p.SetState(308)
		p.Mix()
	}

	return localctx
}

// IEffectChorusContext is an interface to support dynamic dispatch.
type IEffectChorusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEffectChorusContext differentiates from other interfaces.
	IsEffectChorusContext()
}

type EffectChorusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEffectChorusContext() *EffectChorusContext {
	var p = new(EffectChorusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectChorus
	return p
}

func (*EffectChorusContext) IsEffectChorusContext() {}

func NewEffectChorusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EffectChorusContext {
	var p = new(EffectChorusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectChorus

	return p
}

func (s *EffectChorusContext) GetParser() antlr.Parser { return s.parser }

func (s *EffectChorusContext) Rate() IRateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRateContext)
}

func (s *EffectChorusContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *EffectChorusContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *EffectChorusContext) Depth() IDepthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDepthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDepthContext)
}

func (s *EffectChorusContext) Mix() IMixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMixContext)
}

func (s *EffectChorusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EffectChorusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EffectChorusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterEffectChorus(s)
	}
}

func (s *EffectChorusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitEffectChorus(s)
	}
}

func (s *EffectChorusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitEffectChorus(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) EffectChorus() (localctx IEffectChorusContext) {
	this := p
	_ = this

	localctx = NewEffectChorusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SampleGenerationLanguageParserRULE_effectChorus)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(SampleGenerationLanguageParserT__51)
	}
	{
		p.SetState(311)
		p.Rate()
	}
	{
		p.SetState(312)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(313)
		p.Match(SampleGenerationLanguageParserT__52)
	}
	{
		p.SetState(314)
		p.Depth()
	}
	{
		p.SetState(315)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(316)
		p.Match(SampleGenerationLanguageParserT__4)
	}
	{
		p.SetState(317)
		p.Mix()
	}

	return localctx
}

// IEffectDistortionContext is an interface to support dynamic dispatch.
type IEffectDistortionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEffectDistortionContext differentiates from other interfaces.
	IsEffectDistortionContext()
}

type EffectDistortionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEffectDistortionContext() *EffectDistortionContext {
	var p = new(EffectDistortionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectDistortion
	return p
}

func (*EffectDistortionContext) IsEffectDistortionContext() {}

func NewEffectDistortionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EffectDistortionContext {
	var p = new(EffectDistortionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectDistortion

	return p
}

func (s *EffectDistortionContext) GetParser() antlr.Parser { return s.parser }

func (s *EffectDistortionContext) Drive() IDriveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDriveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDriveContext)
}

func (s *EffectDistortionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SampleGenerationLanguageParserCOMMA)
}

func (s *EffectDistortionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, i)
}

func (s *EffectDistortionContext) Tone() IToneContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IToneContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IToneContext)
}

func (s *EffectDistortionContext) Mix() IMixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMixContext)
}

func (s *EffectDistortionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EffectDistortionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EffectDistortionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterEffectDistortion(s)
	}
}

func (s *EffectDistortionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitEffectDistortion(s)
	}
}

func (s *EffectDistortionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitEffectDistortion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) EffectDistortion() (localctx IEffectDistortionContext) {
	this := p
	_ = this

	localctx = NewEffectDistortionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SampleGenerationLanguageParserRULE_effectDistortion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)
		p.Match(SampleGenerationLanguageParserT__53)
	}
	{
		p.SetState(320)
		p.Drive()
	}
	{
		p.SetState(321)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(322)
		p.Match(SampleGenerationLanguageParserT__54)
	}
	{
		p.SetState(323)
		p.Tone()
	}
	{
		p.SetState(324)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(325)
		p.Match(SampleGenerationLanguageParserT__4)
	}
	{
		p.SetState(326)
		p.Mix()
	}

	return localctx
}

// IEffectFilterContext is an interface to support dynamic dispatch.
type IEffectFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEffectFilterContext differentiates from other interfaces.
	IsEffectFilterContext()
}

type EffectFilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEffectFilterContext() *EffectFilterContext {
	var p = new(EffectFilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectFilter
	return p
}

func (*EffectFilterContext) IsEffectFilterContext() {}

func NewEffectFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EffectFilterContext {
	var p = new(EffectFilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_effectFilter

	return p
}

func (s *EffectFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *EffectFilterContext) Cutoff() ICutoffContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICutoffContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICutoffContext)
}

func (s *EffectFilterContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserCOMMA, 0)
}

func (s *EffectFilterContext) Mix() IMixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMixContext)
}

func (s *EffectFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EffectFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EffectFilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterEffectFilter(s)
	}
}

func (s *EffectFilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitEffectFilter(s)
	}
}

func (s *EffectFilterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitEffectFilter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) EffectFilter() (localctx IEffectFilterContext) {
	this := p
	_ = this

	localctx = NewEffectFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SampleGenerationLanguageParserRULE_effectFilter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(SampleGenerationLanguageParserT__28)
	}
	{
		p.SetState(329)
		p.Cutoff()
	}
	{
		p.SetState(330)
		p.Match(SampleGenerationLanguageParserCOMMA)
	}
	{
		p.SetState(331)
		p.Match(SampleGenerationLanguageParserT__4)
	}
	{
		p.SetState(332)
		p.Mix()
	}

	return localctx
}

// IFmTypeContext is an interface to support dynamic dispatch.
type IFmTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFmTypeContext differentiates from other interfaces.
	IsFmTypeContext()
}

type FmTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFmTypeContext() *FmTypeContext {
	var p = new(FmTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_fmType
	return p
}

func (*FmTypeContext) IsFmTypeContext() {}

func NewFmTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FmTypeContext {
	var p = new(FmTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_fmType

	return p
}

func (s *FmTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *FmTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FmTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FmTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterFmType(s)
	}
}

func (s *FmTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitFmType(s)
	}
}

func (s *FmTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitFmType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) FmType() (localctx IFmTypeContext) {
	this := p
	_ = this

	localctx = NewFmTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SampleGenerationLanguageParserRULE_fmType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SampleGenerationLanguageParserT__55 || _la == SampleGenerationLanguageParserT__56) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISizeContext is an interface to support dynamic dispatch.
type ISizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSizeContext differentiates from other interfaces.
	IsSizeContext()
}

type SizeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySizeContext() *SizeContext {
	var p = new(SizeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_size
	return p
}

func (*SizeContext) IsSizeContext() {}

func NewSizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SizeContext {
	var p = new(SizeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_size

	return p
}

func (s *SizeContext) GetParser() antlr.Parser { return s.parser }

func (s *SizeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *SizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterSize(s)
	}
}

func (s *SizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitSize(s)
	}
}

func (s *SizeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitSize(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Size() (localctx ISizeContext) {
	this := p
	_ = this

	localctx = NewSizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SampleGenerationLanguageParserRULE_size)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}
	{
		p.SetState(337)
		p.Match(SampleGenerationLanguageParserT__42)
	}

	return localctx
}

// IOverlapContext is an interface to support dynamic dispatch.
type IOverlapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOverlapContext differentiates from other interfaces.
	IsOverlapContext()
}

type OverlapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOverlapContext() *OverlapContext {
	var p = new(OverlapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_overlap
	return p
}

func (*OverlapContext) IsOverlapContext() {}

func NewOverlapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OverlapContext {
	var p = new(OverlapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_overlap

	return p
}

func (s *OverlapContext) GetParser() antlr.Parser { return s.parser }

func (s *OverlapContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *OverlapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OverlapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OverlapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterOverlap(s)
	}
}

func (s *OverlapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitOverlap(s)
	}
}

func (s *OverlapContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitOverlap(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Overlap() (localctx IOverlapContext) {
	this := p
	_ = this

	localctx = NewOverlapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SampleGenerationLanguageParserRULE_overlap)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}
	{
		p.SetState(340)
		p.Match(SampleGenerationLanguageParserT__57)
	}

	return localctx
}

// IWindowFunctionContext is an interface to support dynamic dispatch.
type IWindowFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWindowFunctionContext differentiates from other interfaces.
	IsWindowFunctionContext()
}

type WindowFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindowFunctionContext() *WindowFunctionContext {
	var p = new(WindowFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_windowFunction
	return p
}

func (*WindowFunctionContext) IsWindowFunctionContext() {}

func NewWindowFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowFunctionContext {
	var p = new(WindowFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_windowFunction

	return p
}

func (s *WindowFunctionContext) GetParser() antlr.Parser { return s.parser }
func (s *WindowFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindowFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterWindowFunction(s)
	}
}

func (s *WindowFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitWindowFunction(s)
	}
}

func (s *WindowFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitWindowFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) WindowFunction() (localctx IWindowFunctionContext) {
	this := p
	_ = this

	localctx = NewWindowFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SampleGenerationLanguageParserRULE_windowFunction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8646911284551352320) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPitchContext is an interface to support dynamic dispatch.
type IPitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPitchContext differentiates from other interfaces.
	IsPitchContext()
}

type PitchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPitchContext() *PitchContext {
	var p = new(PitchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_pitch
	return p
}

func (*PitchContext) IsPitchContext() {}

func NewPitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PitchContext {
	var p = new(PitchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_pitch

	return p
}

func (s *PitchContext) GetParser() antlr.Parser { return s.parser }

func (s *PitchContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *PitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterPitch(s)
	}
}

func (s *PitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitPitch(s)
	}
}

func (s *PitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitPitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Pitch() (localctx IPitchContext) {
	this := p
	_ = this

	localctx = NewPitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SampleGenerationLanguageParserRULE_pitch)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IModelTypeContext is an interface to support dynamic dispatch.
type IModelTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModelTypeContext differentiates from other interfaces.
	IsModelTypeContext()
}

type ModelTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelTypeContext() *ModelTypeContext {
	var p = new(ModelTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_modelType
	return p
}

func (*ModelTypeContext) IsModelTypeContext() {}

func NewModelTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelTypeContext {
	var p = new(ModelTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_modelType

	return p
}

func (s *ModelTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *ModelTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterModelType(s)
	}
}

func (s *ModelTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitModelType(s)
	}
}

func (s *ModelTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitModelType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) ModelType() (localctx IModelTypeContext) {
	this := p
	_ = this

	localctx = NewModelTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SampleGenerationLanguageParserRULE_modelType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-63)) & ^0x3f) == 0 && ((int64(1)<<(_la-63))&7) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IModelParamContext is an interface to support dynamic dispatch.
type IModelParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModelParamContext differentiates from other interfaces.
	IsModelParamContext()
}

type ModelParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelParamContext() *ModelParamContext {
	var p = new(ModelParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_modelParam
	return p
}

func (*ModelParamContext) IsModelParamContext() {}

func NewModelParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelParamContext {
	var p = new(ModelParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_modelParam

	return p
}

func (s *ModelParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelParamContext) Stiffness() IStiffnessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStiffnessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStiffnessContext)
}

func (s *ModelParamContext) Damping() IDampingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDampingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDampingContext)
}

func (s *ModelParamContext) Tension() ITensionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITensionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITensionContext)
}

func (s *ModelParamContext) Length() ILengthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthContext)
}

func (s *ModelParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterModelParam(s)
	}
}

func (s *ModelParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitModelParam(s)
	}
}

func (s *ModelParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitModelParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) ModelParam() (localctx IModelParamContext) {
	this := p
	_ = this

	localctx = NewModelParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SampleGenerationLanguageParserRULE_modelParam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(356)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SampleGenerationLanguageParserT__65:
		{
			p.SetState(348)
			p.Match(SampleGenerationLanguageParserT__65)
		}
		{
			p.SetState(349)
			p.Stiffness()
		}

	case SampleGenerationLanguageParserT__66:
		{
			p.SetState(350)
			p.Match(SampleGenerationLanguageParserT__66)
		}
		{
			p.SetState(351)
			p.Damping()
		}

	case SampleGenerationLanguageParserT__67:
		{
			p.SetState(352)
			p.Match(SampleGenerationLanguageParserT__67)
		}
		{
			p.SetState(353)
			p.Tension()
		}

	case SampleGenerationLanguageParserT__68:
		{
			p.SetState(354)
			p.Match(SampleGenerationLanguageParserT__68)
		}
		{
			p.SetState(355)
			p.Length()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExcitationContext is an interface to support dynamic dispatch.
type IExcitationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExcitationContext differentiates from other interfaces.
	IsExcitationContext()
}

type ExcitationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExcitationContext() *ExcitationContext {
	var p = new(ExcitationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_excitation
	return p
}

func (*ExcitationContext) IsExcitationContext() {}

func NewExcitationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExcitationContext {
	var p = new(ExcitationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_excitation

	return p
}

func (s *ExcitationContext) GetParser() antlr.Parser { return s.parser }
func (s *ExcitationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExcitationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExcitationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterExcitation(s)
	}
}

func (s *ExcitationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitExcitation(s)
	}
}

func (s *ExcitationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitExcitation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Excitation() (localctx IExcitationContext) {
	this := p
	_ = this

	localctx = NewExcitationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SampleGenerationLanguageParserRULE_excitation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-35)) & ^0x3f) == 0 && ((int64(1)<<(_la-35))&34359738385) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAmplitudeContext is an interface to support dynamic dispatch.
type IAmplitudeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAmplitudeContext differentiates from other interfaces.
	IsAmplitudeContext()
}

type AmplitudeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAmplitudeContext() *AmplitudeContext {
	var p = new(AmplitudeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_amplitude
	return p
}

func (*AmplitudeContext) IsAmplitudeContext() {}

func NewAmplitudeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AmplitudeContext {
	var p = new(AmplitudeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_amplitude

	return p
}

func (s *AmplitudeContext) GetParser() antlr.Parser { return s.parser }

func (s *AmplitudeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *AmplitudeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AmplitudeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AmplitudeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterAmplitude(s)
	}
}

func (s *AmplitudeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitAmplitude(s)
	}
}

func (s *AmplitudeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitAmplitude(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Amplitude() (localctx IAmplitudeContext) {
	this := p
	_ = this

	localctx = NewAmplitudeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SampleGenerationLanguageParserRULE_amplitude)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IPhaseContext is an interface to support dynamic dispatch.
type IPhaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPhaseContext differentiates from other interfaces.
	IsPhaseContext()
}

type PhaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPhaseContext() *PhaseContext {
	var p = new(PhaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_phase
	return p
}

func (*PhaseContext) IsPhaseContext() {}

func NewPhaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PhaseContext {
	var p = new(PhaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_phase

	return p
}

func (s *PhaseContext) GetParser() antlr.Parser { return s.parser }

func (s *PhaseContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *PhaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PhaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterPhase(s)
	}
}

func (s *PhaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitPhase(s)
	}
}

func (s *PhaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitPhase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Phase() (localctx IPhaseContext) {
	this := p
	_ = this

	localctx = NewPhaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SampleGenerationLanguageParserRULE_phase)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_index
	return p
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Index() (localctx IIndexContext) {
	this := p
	_ = this

	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SampleGenerationLanguageParserRULE_index)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IFilterTypeContext is an interface to support dynamic dispatch.
type IFilterTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterTypeContext differentiates from other interfaces.
	IsFilterTypeContext()
}

type FilterTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterTypeContext() *FilterTypeContext {
	var p = new(FilterTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_filterType
	return p
}

func (*FilterTypeContext) IsFilterTypeContext() {}

func NewFilterTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterTypeContext {
	var p = new(FilterTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_filterType

	return p
}

func (s *FilterTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *FilterTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterFilterType(s)
	}
}

func (s *FilterTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitFilterType(s)
	}
}

func (s *FilterTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitFilterType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) FilterType() (localctx IFilterTypeContext) {
	this := p
	_ = this

	localctx = NewFilterTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SampleGenerationLanguageParserRULE_filterType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-48)) & ^0x3f) == 0 && ((int64(1)<<(_la-48))&25165827) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICutoffContext is an interface to support dynamic dispatch.
type ICutoffContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCutoffContext differentiates from other interfaces.
	IsCutoffContext()
}

type CutoffContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCutoffContext() *CutoffContext {
	var p = new(CutoffContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_cutoff
	return p
}

func (*CutoffContext) IsCutoffContext() {}

func NewCutoffContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CutoffContext {
	var p = new(CutoffContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_cutoff

	return p
}

func (s *CutoffContext) GetParser() antlr.Parser { return s.parser }

func (s *CutoffContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *CutoffContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CutoffContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CutoffContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterCutoff(s)
	}
}

func (s *CutoffContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitCutoff(s)
	}
}

func (s *CutoffContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitCutoff(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Cutoff() (localctx ICutoffContext) {
	this := p
	_ = this

	localctx = NewCutoffContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SampleGenerationLanguageParserRULE_cutoff)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IResonanceContext is an interface to support dynamic dispatch.
type IResonanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResonanceContext differentiates from other interfaces.
	IsResonanceContext()
}

type ResonanceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResonanceContext() *ResonanceContext {
	var p = new(ResonanceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_resonance
	return p
}

func (*ResonanceContext) IsResonanceContext() {}

func NewResonanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResonanceContext {
	var p = new(ResonanceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_resonance

	return p
}

func (s *ResonanceContext) GetParser() antlr.Parser { return s.parser }

func (s *ResonanceContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *ResonanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResonanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResonanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterResonance(s)
	}
}

func (s *ResonanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitResonance(s)
	}
}

func (s *ResonanceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitResonance(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Resonance() (localctx IResonanceContext) {
	this := p
	_ = this

	localctx = NewResonanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SampleGenerationLanguageParserRULE_resonance)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IDecayContext is an interface to support dynamic dispatch.
type IDecayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecayContext differentiates from other interfaces.
	IsDecayContext()
}

type DecayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecayContext() *DecayContext {
	var p = new(DecayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_decay
	return p
}

func (*DecayContext) IsDecayContext() {}

func NewDecayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecayContext {
	var p = new(DecayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_decay

	return p
}

func (s *DecayContext) GetParser() antlr.Parser { return s.parser }

func (s *DecayContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *DecayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterDecay(s)
	}
}

func (s *DecayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitDecay(s)
	}
}

func (s *DecayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitDecay(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Decay() (localctx IDecayContext) {
	this := p
	_ = this

	localctx = NewDecayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SampleGenerationLanguageParserRULE_decay)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IBandwidthContext is an interface to support dynamic dispatch.
type IBandwidthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBandwidthContext differentiates from other interfaces.
	IsBandwidthContext()
}

type BandwidthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBandwidthContext() *BandwidthContext {
	var p = new(BandwidthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_bandwidth
	return p
}

func (*BandwidthContext) IsBandwidthContext() {}

func NewBandwidthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BandwidthContext {
	var p = new(BandwidthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_bandwidth

	return p
}

func (s *BandwidthContext) GetParser() antlr.Parser { return s.parser }

func (s *BandwidthContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *BandwidthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BandwidthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BandwidthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterBandwidth(s)
	}
}

func (s *BandwidthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitBandwidth(s)
	}
}

func (s *BandwidthContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitBandwidth(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Bandwidth() (localctx IBandwidthContext) {
	this := p
	_ = this

	localctx = NewBandwidthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SampleGenerationLanguageParserRULE_bandwidth)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IGainContext is an interface to support dynamic dispatch.
type IGainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGainContext differentiates from other interfaces.
	IsGainContext()
}

type GainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGainContext() *GainContext {
	var p = new(GainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_gain
	return p
}

func (*GainContext) IsGainContext() {}

func NewGainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GainContext {
	var p = new(GainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_gain

	return p
}

func (s *GainContext) GetParser() antlr.Parser { return s.parser }

func (s *GainContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *GainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterGain(s)
	}
}

func (s *GainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitGain(s)
	}
}

func (s *GainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitGain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Gain() (localctx IGainContext) {
	this := p
	_ = this

	localctx = NewGainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SampleGenerationLanguageParserRULE_gain)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IDelayTimeContext is an interface to support dynamic dispatch.
type IDelayTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelayTimeContext differentiates from other interfaces.
	IsDelayTimeContext()
}

type DelayTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelayTimeContext() *DelayTimeContext {
	var p = new(DelayTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_delayTime
	return p
}

func (*DelayTimeContext) IsDelayTimeContext() {}

func NewDelayTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelayTimeContext {
	var p = new(DelayTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_delayTime

	return p
}

func (s *DelayTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *DelayTimeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *DelayTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelayTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelayTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterDelayTime(s)
	}
}

func (s *DelayTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitDelayTime(s)
	}
}

func (s *DelayTimeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitDelayTime(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) DelayTime() (localctx IDelayTimeContext) {
	this := p
	_ = this

	localctx = NewDelayTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SampleGenerationLanguageParserRULE_delayTime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}
	{
		p.SetState(379)
		p.Match(SampleGenerationLanguageParserT__42)
	}

	return localctx
}

// IFeedbackContext is an interface to support dynamic dispatch.
type IFeedbackContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFeedbackContext differentiates from other interfaces.
	IsFeedbackContext()
}

type FeedbackContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFeedbackContext() *FeedbackContext {
	var p = new(FeedbackContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_feedback
	return p
}

func (*FeedbackContext) IsFeedbackContext() {}

func NewFeedbackContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FeedbackContext {
	var p = new(FeedbackContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_feedback

	return p
}

func (s *FeedbackContext) GetParser() antlr.Parser { return s.parser }

func (s *FeedbackContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *FeedbackContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FeedbackContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FeedbackContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterFeedback(s)
	}
}

func (s *FeedbackContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitFeedback(s)
	}
}

func (s *FeedbackContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitFeedback(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Feedback() (localctx IFeedbackContext) {
	this := p
	_ = this

	localctx = NewFeedbackContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SampleGenerationLanguageParserRULE_feedback)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IRateContext is an interface to support dynamic dispatch.
type IRateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRateContext differentiates from other interfaces.
	IsRateContext()
}

type RateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRateContext() *RateContext {
	var p = new(RateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_rate
	return p
}

func (*RateContext) IsRateContext() {}

func NewRateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RateContext {
	var p = new(RateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_rate

	return p
}

func (s *RateContext) GetParser() antlr.Parser { return s.parser }

func (s *RateContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *RateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterRate(s)
	}
}

func (s *RateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitRate(s)
	}
}

func (s *RateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitRate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Rate() (localctx IRateContext) {
	this := p
	_ = this

	localctx = NewRateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SampleGenerationLanguageParserRULE_rate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SampleGenerationLanguageParserT__39 {
		{
			p.SetState(384)
			p.Match(SampleGenerationLanguageParserT__39)
		}

	}

	return localctx
}

// IDepthContext is an interface to support dynamic dispatch.
type IDepthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDepthContext differentiates from other interfaces.
	IsDepthContext()
}

type DepthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDepthContext() *DepthContext {
	var p = new(DepthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_depth
	return p
}

func (*DepthContext) IsDepthContext() {}

func NewDepthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DepthContext {
	var p = new(DepthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_depth

	return p
}

func (s *DepthContext) GetParser() antlr.Parser { return s.parser }

func (s *DepthContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *DepthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DepthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DepthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterDepth(s)
	}
}

func (s *DepthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitDepth(s)
	}
}

func (s *DepthContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitDepth(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Depth() (localctx IDepthContext) {
	this := p
	_ = this

	localctx = NewDepthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, SampleGenerationLanguageParserRULE_depth)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IDriveContext is an interface to support dynamic dispatch.
type IDriveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDriveContext differentiates from other interfaces.
	IsDriveContext()
}

type DriveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDriveContext() *DriveContext {
	var p = new(DriveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_drive
	return p
}

func (*DriveContext) IsDriveContext() {}

func NewDriveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DriveContext {
	var p = new(DriveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_drive

	return p
}

func (s *DriveContext) GetParser() antlr.Parser { return s.parser }

func (s *DriveContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *DriveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DriveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DriveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterDrive(s)
	}
}

func (s *DriveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitDrive(s)
	}
}

func (s *DriveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitDrive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Drive() (localctx IDriveContext) {
	this := p
	_ = this

	localctx = NewDriveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, SampleGenerationLanguageParserRULE_drive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IToneContext is an interface to support dynamic dispatch.
type IToneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsToneContext differentiates from other interfaces.
	IsToneContext()
}

type ToneContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyToneContext() *ToneContext {
	var p = new(ToneContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_tone
	return p
}

func (*ToneContext) IsToneContext() {}

func NewToneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ToneContext {
	var p = new(ToneContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_tone

	return p
}

func (s *ToneContext) GetParser() antlr.Parser { return s.parser }

func (s *ToneContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *ToneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ToneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterTone(s)
	}
}

func (s *ToneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitTone(s)
	}
}

func (s *ToneContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitTone(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Tone() (localctx IToneContext) {
	this := p
	_ = this

	localctx = NewToneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, SampleGenerationLanguageParserRULE_tone)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IStiffnessContext is an interface to support dynamic dispatch.
type IStiffnessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStiffnessContext differentiates from other interfaces.
	IsStiffnessContext()
}

type StiffnessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStiffnessContext() *StiffnessContext {
	var p = new(StiffnessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_stiffness
	return p
}

func (*StiffnessContext) IsStiffnessContext() {}

func NewStiffnessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StiffnessContext {
	var p = new(StiffnessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_stiffness

	return p
}

func (s *StiffnessContext) GetParser() antlr.Parser { return s.parser }

func (s *StiffnessContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *StiffnessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StiffnessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StiffnessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterStiffness(s)
	}
}

func (s *StiffnessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitStiffness(s)
	}
}

func (s *StiffnessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitStiffness(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Stiffness() (localctx IStiffnessContext) {
	this := p
	_ = this

	localctx = NewStiffnessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, SampleGenerationLanguageParserRULE_stiffness)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// IDampingContext is an interface to support dynamic dispatch.
type IDampingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDampingContext differentiates from other interfaces.
	IsDampingContext()
}

type DampingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDampingContext() *DampingContext {
	var p = new(DampingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_damping
	return p
}

func (*DampingContext) IsDampingContext() {}

func NewDampingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DampingContext {
	var p = new(DampingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_damping

	return p
}

func (s *DampingContext) GetParser() antlr.Parser { return s.parser }

func (s *DampingContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *DampingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DampingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DampingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterDamping(s)
	}
}

func (s *DampingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitDamping(s)
	}
}

func (s *DampingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitDamping(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Damping() (localctx IDampingContext) {
	this := p
	_ = this

	localctx = NewDampingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, SampleGenerationLanguageParserRULE_damping)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// ITensionContext is an interface to support dynamic dispatch.
type ITensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTensionContext differentiates from other interfaces.
	IsTensionContext()
}

type TensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTensionContext() *TensionContext {
	var p = new(TensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_tension
	return p
}

func (*TensionContext) IsTensionContext() {}

func NewTensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TensionContext {
	var p = new(TensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_tension

	return p
}

func (s *TensionContext) GetParser() antlr.Parser { return s.parser }

func (s *TensionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *TensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterTension(s)
	}
}

func (s *TensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitTension(s)
	}
}

func (s *TensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitTension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Tension() (localctx ITensionContext) {
	this := p
	_ = this

	localctx = NewTensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, SampleGenerationLanguageParserRULE_tension)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}

// ILengthContext is an interface to support dynamic dispatch.
type ILengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLengthContext differentiates from other interfaces.
	IsLengthContext()
}

type LengthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthContext() *LengthContext {
	var p = new(LengthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SampleGenerationLanguageParserRULE_length
	return p
}

func (*LengthContext) IsLengthContext() {}

func NewLengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthContext {
	var p = new(LengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SampleGenerationLanguageParserRULE_length

	return p
}

func (s *LengthContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SampleGenerationLanguageParserNUMBER, 0)
}

func (s *LengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.EnterLength(s)
	}
}

func (s *LengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SampleGenerationLanguageListener); ok {
		listenerT.ExitLength(s)
	}
}

func (s *LengthContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SampleGenerationLanguageVisitor:
		return t.VisitLength(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SampleGenerationLanguageParser) Length() (localctx ILengthContext) {
	this := p
	_ = this

	localctx = NewLengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, SampleGenerationLanguageParserRULE_length)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.Match(SampleGenerationLanguageParserNUMBER)
	}

	return localctx
}
