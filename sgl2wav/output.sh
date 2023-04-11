#!/bin/sh
grep -n -r "GenerateWhiteNoise" *.go
grep -n -r "GeneratePinkNoise" *.go
grep -n -r "GenerateBrownianNoise" *.go
grep -n -r "EffectData" *.go
grep -n -r "listener.data" *.go
grep -n -r "ctx.GetChild" *.go
