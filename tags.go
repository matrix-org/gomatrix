package gomatrix

// Tag contains the data for an m.tag message type
type Tag struct {
	Tags map[string]TagProperties `json:"tags"`
}

// TagProperties contains the properties of an MTag
type TagProperties struct {
	Order float32 `json:"order,omitempty"` // Empty values must be neglected
}
