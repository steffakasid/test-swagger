package main

import (
	"errors"
	"io/ioutil"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v2"
)

type v3orv2 struct {
	OpenAPI *string `json:"openapi" yaml:"openapi"`
	Swagger *string `json:"swagger" yaml:"swagger"`
}

type OpenAPIv3 struct {
	openapi3.ExtensionProps
	OpenAPI      string                        `json:"openapi" yaml:"openapi"` // Required
	Components   openapi3.Components           `json:"components,omitempty" yaml:"components,omitempty"`
	Info         *openapi3.Info                `json:"info" yaml:"info"`   // Required
	Paths        openapi3.Paths                `json:"paths" yaml:"paths"` // Required
	Security     openapi3.SecurityRequirements `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      openapi3.Servers              `json:"servers,omitempty" yaml:"servers,omitempty"`
	Tags         openapi3.Tags                 `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs *openapi3.ExternalDocs        `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

type SecurityScheme struct {
	openapi3.ExtensionProps
	Ref              string            `json:"$ref,omitempty" yaml:"$ref"`
	Description      string            `json:"description,omitempty" yaml:"description"`
	Type             string            `json:"type,omitempty" yaml:"type"`
	In               string            `json:"in,omitempty" yaml:"in"`
	Name             string            `json:"name,omitempty" yaml:"name"`
	Flow             string            `json:"flow,omitempty" yaml:"flow"`
	AuthorizationURL string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl"`
	TokenURL         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl"`
	Scopes           map[string]string `json:"scopes,omitempty" yaml:"scopes"`
	Tags             openapi3.Tags     `json:"tags,omitempty" yaml:"tags"`
}

type OpenAPIv2 struct {
	openapi3.ExtensionProps
	Swagger             string                         `json:"swagger" yaml:"swagger"`
	Info                openapi3.Info                  `json:"info" yaml:"info"`
	ExternalDocs        *openapi3.ExternalDocs         `json:"externalDocs,omitempty" yaml:"externalDocs"`
	Schemes             []string                       `json:"schemes,omitempty" yaml:"schemes"`
	Consumes            []string                       `json:"consumes,omitempty" yaml:"consumes"`
	Produces            []string                       `json:"produces,omitempty" yaml:"produces"`
	Host                string                         `json:"host,omitempty" yaml:"host"`
	BasePath            string                         `json:"basePath,omitempty" yaml:"basePath"`
	Paths               map[string]*interface{}        `json:"paths,omitempty" yaml:"paths"`                   // FIXME: Overwritten with map[string]*interface{}
	Definitions         map[string]*interface{}        `json:"definitions,omitempty,noref" yaml:"definitions"` // FIXME: Overwritten with map[string]*interface{}
	Parameters          map[string]*openapi2.Parameter `json:"parameters,omitempty,noref" yaml:"parameters"`
	Responses           map[string]*openapi2.Response  `json:"responses,omitempty,noref" yaml:"responses"`
	SecurityDefinitions map[string]*SecurityScheme     `json:"securityDefinitions,omitempty" yaml:"securityDefinitions"` // FIXME: Overwritten with local type to get yaml mapping
	Security            openapi2.SecurityRequirements  `json:"security,omitempty" yaml:"security"`
	Tags                openapi3.Tags                  `json:"tags,omitempty" yaml:"tags"`
}

func main() {
	readOpenAPI("test/openapiv3.yaml")
	readOpenAPI("test/swagger.yaml")

}

func readOpenAPI(filepath string) (interface{}, error) {
	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	isV3orV2 := v3orv2{}
	err = yaml.Unmarshal(fileContent, &isV3orV2)
	if err != nil {
		return nil, err
	}

	if isV3orV2.OpenAPI != nil {
		openapi3Doc := OpenAPIv3{}
		err = yaml.UnmarshalStrict(fileContent, &openapi3Doc)
		if err != nil {
			return nil, err
		}
		return openapi3Doc, err
	}

	if isV3orV2.Swagger != nil {
		swaggerDoc := OpenAPIv2{}
		err = yaml.UnmarshalStrict(fileContent, &swaggerDoc)
		if err != nil {
			return nil, err
		}
		return swaggerDoc, err
	}
	return nil, errors.New("blub")
}
