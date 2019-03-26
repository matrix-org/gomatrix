package gomatrix

// Tag contains the data for a Tag which can be referenced by the Tag name
type Tag struct {
	Tags map[string]TagProperties `json:"tags"`
}

// TagProperties contains the properties of an MTag
type TagProperties struct {
	Order float32 `json:"order,omitempty"` // Empty values must be neglected
}
