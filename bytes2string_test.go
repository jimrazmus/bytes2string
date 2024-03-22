package bytes2string

import "testing"

func TestByteCountSI(t *testing.T) {
	tests := []struct {
		name  string
		bytes int64
		want  string
	}{
		{"Test 0 kB", 0, "0 B"},
		{"Test bytes less than 1000", 500, "500 B"},
		{"Test 1 kB", 1000, "1.0 kB"},
		{"Test 1 MB", 1000 * 1000, "1.0 MB"},
		{"Test 1 GB", 1000 * 1000 * 1000, "1.0 GB"},
		{"Test max int64", 9223372036854775807, "9.2 EB"},
		{"Test negative bytes", -1, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteCountSI(tt.bytes); got != tt.want {
				t.Errorf("ByteCountSI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteCountIEC(t *testing.T) {
	tests := []struct {
		name  string
		bytes int64
		want  string
	}{
		{"Test 0 kB", 0, "0 B"},
		{"Test bytes less than 1024", 500, "500 B"},
		{"Test 1 KiB", 1024, "1.0 KiB"},
		{"Test 1 MiB", 1024 * 1024, "1.0 MiB"},
		{"Test 1 GiB", 1024 * 1024 * 1024, "1.0 GiB"},
		{"Test max int64", 9223372036854775807, "8.0 EiB"},
		{"Test negative bytes", -1, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteCountIEC(tt.bytes); got != tt.want {
				t.Errorf("ByteCountIEC() = %v, want %v", got, tt.want)
			}
		})
	}
}
