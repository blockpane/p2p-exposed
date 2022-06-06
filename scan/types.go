package scan

// TODO: actually implement advanced fingerprinting

// Signature holds the information needed to match against an HTTP response and attribute a server type
type Signature struct {
	// Chain will be the most likely blockchain the service is associated with.
	Chain string `yaml:"chain"`
	// Service will be the service, RPC, Prometheus, etc.
	Service string `yaml:"service"`
	// Value is what to search for, this is a regular expression
	Value string `yaml:"value"`
	// Location is where to search, can be 'header' or 'body'
	Location string `yaml:"location"`
	// Key is only used if Location is header: what http header to match against
	Key string `yaml:"key"`
}
