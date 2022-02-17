watch:
	find . -name '*.go' | entr -r go run *.go *.html
