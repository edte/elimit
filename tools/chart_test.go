package tools

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestGenerateLineGraph(t *testing.T) {
	type args struct {
		data []Item
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateLineGraph(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateLineGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListen(t *testing.T) {
	type args struct {
		port int
		data []Item
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				port: 8899,
				data: []Item{
					{
						Time: time.Now(),
						Data: false,
					},
					{
						Time: time.Now(),
						Data: true,
					},
					{
						Time: time.Now(),
						Data: false,
					},
					{
						Time: time.Now(),
						Data: false,
					},
					{
						Time: time.Now(),
						Data: true,
					},
					{
						Time: time.Now(),
						Data: false,
					},
					{
						Time: time.Now(),
						Data: true,
					},
					{
						Time: time.Now(),
						Data: false,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Listen(tt.args.port, tt.args.data)
		})
	}
}
