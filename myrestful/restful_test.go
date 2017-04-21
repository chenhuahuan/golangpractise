package main

import (
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_dispatchHandler(t *testing.T) {
	type args struct {
		response http.ResponseWriter
		request  *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		dispatchHandler(tt.args.response, tt.args.request)
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		main()
	}
}
