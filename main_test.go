package main

import (
	"reflect"
	"testing"

	"github.com/jessevdk/go-flags"
)

func TestNamespacesFlag(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{"comma-separated", []string{"--namespaces=ns-1,ns-2"}, []string{"ns-1", "ns-2"}},
		{"repeated flag", []string{"--namespaces=ns-1", "--namespaces=ns-2"}, []string{"ns-1", "ns-2"}},
		{"mixed", []string{"-n", "ns-1, ns-2", "-n", "ns-3"}, []string{"ns-1", "ns-2", "ns-3"}},
		{"empty entries", []string{"--namespaces=ns-1,,ns-2,"}, []string{"ns-1", "ns-2"}},
		{"not set", nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var parsed struct {
				Namespaces []string `long:"namespaces" short:"n"`
			}
			if _, err := flags.ParseArgs(&parsed, tt.args); err != nil {
				t.Fatalf("parsing %v: %v", tt.args, err)
			}
			if got := splitCommaSeparated(parsed.Namespaces); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("namespaces for %v = %q, want %q", tt.args, got, tt.want)
			}
		})
	}
}
