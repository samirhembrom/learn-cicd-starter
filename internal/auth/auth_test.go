package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		want    string
		wantErr bool
	}{
		{
			name:    "Valid ApiKey header",
			header:  "ApiKey hello.world",
			want:    "hello.world",
			wantErr: false,
		},
		{
			name:    "Empty header",
			header:  "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Wrong scheme",
			header:  "Bearer hello.world",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/example", nil)
			if tt.header != "" {
				req.Header.Set("Authorization", tt.header)
			}

			got, err := GetAPIKey(req.Header)
			if err != nil {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
