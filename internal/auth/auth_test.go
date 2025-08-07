package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := make(map[string]http.Header)
	headers["correct"] = http.Header{}
	headers["empty"] = http.Header{}
	headers["malformed_not_enough_args"] = http.Header{}
	headers["malformed_too_many_args"] = http.Header{}

	const keyType = "ApiKey"
	const token = "TOKEN"
	headers["correct"].Add("Authorization", fmt.Sprintf("%s %s", keyType, token))
	headers["empty"].Add("Authorization", "")
	headers["malformed_not_enough_args"].Add("Authorization", "TOKEN")
	headers["malformed_too_many_args"].Add("Authorization", "TOKEN")

	tests := map[string]struct {
		input      http.Header
		want       string
		shouldFail bool
	}{
		"correct": {
			input:      headers["correct"],
			want:       token,
			shouldFail: false,
		},
		"empty": {
			input:      headers["empty"],
			want:       "",
			shouldFail: true,
		},
		"malformed_not_enough_args": {
			input:      headers["malformed_not_enough_args"],
			want:       "",
			shouldFail: true,
		},
		"malformed_too_many_args": {
			input:      headers["malformed_too_many_args"],
			want:       "",
			shouldFail: true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if err == nil && tc.shouldFail {
				t.Fatalf("got: %v, shouldFail: %v", got, tc.shouldFail)
			}
			if err != nil && !tc.shouldFail {
				t.Fatalf("error: %v, shouldFail: %v", err, tc.shouldFail)
			}
			if got != tc.want {
				t.Fatalf("expected: %v, got: %v, shouldFail: %v", tc.want, got, tc.shouldFail)
			}
		})
	}

}
