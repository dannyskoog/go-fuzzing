setup:
	go install golang.org/dl/gotip@latest
	gotip download

test:
	gotip test ./internal/pkg/calculation --fuzz=Fuzz