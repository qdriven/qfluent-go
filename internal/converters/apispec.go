package converters

import (
	"github.com/grokify/spectrum/openapi2/openapi2postman2"
	"github.com/grokify/spectrum/openapi3/openapi3postman2"
	"github.com/grokify/spectrum/postman2"
	"qfluent-go/internal"
)

//TODO: use https://github.com/grokify/spectrum/blob/master/examples/ringcentral-engage-digital/convert.go
//Convert Different OpenAPI Format

func SwaggerToPostman(swaggerFilePath string, postmanFilePath string) {

	cfg := openapi2postman2.Configuration{
		PostmanURLBase: "{{BASE_URL}}",
		PostmanHeaders: []postman2.Header{{
			Key:   "Authorization",
			Value: "Bearer {{my_access_token}}"}}}
	cov := openapi2postman2.NewConverter(cfg)
	err := cov.Convert(swaggerFilePath, postmanFilePath)
	internal.HandlerError(err)
}

func OpenApi3ToPostman(swaggerFilePath string, postmanFilePath string) {

	cfg := openapi3postman2.Configuration{
		UseXTagGroups:            true,
		PostmanServerURLBasePath: "1.0",
		PostmanServerURL:         "{{BASE_URL}}",
		PostmanHeaders: []postman2.Header{{
			Key:   "Authorization",
			Value: "Bearer {{my_access_token}}"}}}
	cov := openapi3postman2.NewConverter(cfg)
	err := cov.ConvertFile(swaggerFilePath, postmanFilePath)
	internal.HandlerError(err)
}
