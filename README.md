Add swagger json:

`swagger generate spec -o ./swagger-ui/swagger-base.json --scan-models`

Build:

`GOOS=linux GOARCH=amd64 go build`