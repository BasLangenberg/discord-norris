package giphy

import "testing"

func TestGetRandomChuckGifDownSizedLarge(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test if gif is returned",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetRandomChuckGifDownSizedLarge()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRandomChuckGifDownSizedLarge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
