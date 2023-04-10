all:
	antlr4 -Dlanguage=Go -visitor antlr/SampleGenerationLanguage.g4
	go build -mod=vendor

clean:
	go clean
