package icndb

import "testing"

func TestGetRandomQuote(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "Test if random quote is obtained",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRandomQuote()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRandomQuote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) == 0 {
				t.Errorf("GetRandomQuote() got = %v, Wanted a Quote", got)
			}
		})
	}
}