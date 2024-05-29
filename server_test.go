package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {

	testserver := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testserver.Close()

	testClient := testserver.Client()

	fmt.Println(testserver.URL)
	response, err := testClient.Get(testserver.URL)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, 200, response.StatusCode)
}

func TestHelloWorldFail(t *testing.T) {
	testserver := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testserver.Close()

	testClient := testserver.Client()

	fmt.Println(testserver.URL)

	body := strings.NewReader("test body")

	response, err := testClient.Post(testserver.URL, "application/json", body)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, 405, response.StatusCode)
}
