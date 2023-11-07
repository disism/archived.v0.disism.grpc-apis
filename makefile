.PHONY: version git

VERSION_FILE := BUF_VERSION

run:
	buf lint && buf generate

version:
	@buf_version=$$(buf --version); \
	expected_version=$$(head -n 1 $(VERSION_FILE)); \
	[ "$$buf_version" = "$$expected_version" ] && \
		echo "buf version: $$buf_version" || \
		{ echo "😅 Error: buf version mismatch. Expected version: $$expected_version, actual version: $$buf_version"; exit 1;}

git:
	chmod +x .git/hooks/pre-commit