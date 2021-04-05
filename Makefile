.PHONY: test coverage

test:
	@if [[ -z "${n}" ]]; then \
	 	go test; \
	else \
		go test -run ^${n}$$; \
	fi

coverage:
	@go test -coverprofile='coverage.out' \
		&& go tool cover -html='coverage.out' -o='coverage.html' \
		&& rm 'coverage.out'
