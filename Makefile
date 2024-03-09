build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o ./bin/bootstrap

deploy: build
	cd bin && zip -r bootstrap.zip bootstrap
	cd bin && aws lambda update-function-code --function-name "lamb" --zip-file fileb://bootstrap.zip --region="us-east-1"