all:
	(cd sgl2wav; go build -mod=vendor)

clean:
	(cd sgl2wav; go clean)
