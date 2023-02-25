build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main .
deploy_prod: build
	serverless deploy --stage prod --verbose