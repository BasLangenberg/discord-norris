package mcsrvstat_us

import (
	"testing"
)

func TestGetOnlinePlayers(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "Test if function returns no error",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetOnlinePlayers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOnlinePlayers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}