package models

// DocumentMeta contains all meta data used to identifier a document.
type DocumentMeta struct {
	Key string `json:"_key,omitempty"`
	ID  string `json:"_id,omitempty"`
	Rev string `json:"_rev,omitempty"`
}
