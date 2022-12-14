## How to make .zip file

### MacOS

1. Go to root directory of project
2. Do `GOARCH=amd64 GOOS=linux go build main.go`
3. Do `zip main.zip main`

## How to deployment to Lambda AWS
1. Open AWS Web 
2. Select us-west-2 Oregon Region 
3. Open Lambda section
4. Select lambda with 'test' name
5. Open Code tab
6. Upload from (.zip file)