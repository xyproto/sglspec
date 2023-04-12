#!/bin/sh
# Suggested by: https://github.com/antlr/antlr4/blob/master/doc/go-target.md
alias antlr4='java -Xmx500M -cp "./antlr-4.12.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
#antlr4 -Dlanguage=Go -no-visitor -package parser *.g4
antlr4 -Dlanguage=Go -visitor -package parser *.g4
