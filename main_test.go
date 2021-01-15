package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
