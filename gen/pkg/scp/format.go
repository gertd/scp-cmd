package scp

// Format -- enum
type Format int

// Format -- OpenAPI spec serialization format enum
const (
	FormatUnknown Format = 0 + iota
	FormatJSON
	FormatYAML
)

// Format enum string representations
const (
	formatUnknown = "unknown"
	formatJSON    = "json"
	formatYAML    = "yaml"
)

// String -- Format
func (t Format) String() string {
	return [...]string{formatUnknown, formatJSON, formatYAML}[t]
}
