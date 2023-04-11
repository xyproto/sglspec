all:
	antlr4 -Dlanguage=Go -visitor antlr/SampleGenerationLanguage.g4
	(cd sgl2wav; go build -mod=vendor)

clean:
	(cd sgl2wav; go clean)
