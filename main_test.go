package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestOpenAPIv2(t *testing.T) {
	data, err := readOpenAPI("test/openapiv2.yaml")
	assert.IsType(t, OpenAPIv2{}, data)
	assert.NoError(t, err)
}

func TestOpenAPIv3(t *testing.T) {
	data, err := readOpenAPI("test/openapiv3.yaml")
	assert.IsType(t, OpenAPIv3{}, data)
	assert.NoError(t, err)
}

func TestKinV3(t *testing.T) {

	fileBytes, err := ioutil.ReadFile("test/openapiv3.yaml")
	assert.NoError(t, err)

	loader := openapi3.NewSwaggerLoader()
	loader.IsExternalRefsAllowed = true
	swagger, err := loader.LoadSwaggerFromData(fileBytes)
	assert.NoError(t, err)
	assert.IsType(t, &openapi3.Swagger{}, swagger)
	b, err := swagger.MarshalJSON()
	assert.NoError(t, err)

	data := make(map[string]interface{})
	err = json.Unmarshal(b, &data)
	assert.NoError(t, err)

	b, err = yaml.Marshal(data)
	assert.NoError(t, err)

	err = ioutil.WriteFile("test_v3_out.yaml", b, 0777)
	assert.NoError(t, err)
}
