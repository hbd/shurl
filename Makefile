all: install
	@shurl

debug:
	@go install .
	@shurl -debug

install:
	@go install .
