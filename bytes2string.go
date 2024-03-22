// Package bytes2string provides functions for converting int64 byte counts to
// a shortened SI, or IEC, string format.
package bytes2string

import "fmt"

// ByteCountSI returns a formatted string using the appropriate International
// System of Units prefix for any positive byte count or an empty string
// otherwise.
func ByteCountSI(bytes int64) string {
	if bytes < 0 {
		return ""
	}

	const unit = 1000

	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	div, exp := int64(unit), 0

	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "kMGTPE"[exp])
}

// ByteCountIEC returns a formatted string using the appropriate International
// Electrotechnical Commission prefix for any positive byte count or an empty
// string otherwise.
func ByteCountIEC(bytes int64) string {
	if bytes < 0 {
		return ""
	}

	const unit = 1024

	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	div, exp := int64(unit), 0

	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %ciB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
