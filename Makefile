build:
	env GOOS=linux go build -ldflags="-s -w" -o app/main cmd/main.go
deploy_prod: build
	serverless deploy --stage prod --verbose