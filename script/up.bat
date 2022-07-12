build-lambda-zip -output main.zip main
aws lambda update-function-code --function-name unzip --zip-file fileb://main.zip
