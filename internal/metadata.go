package internal

type Metadata struct {
	PageSize     int `json:"page_size,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
	LastSeenID   int `json:"last_seen_id,omitempty"`
}

type MetadataCursor struct {
	NextCursor int64 `json:"next_cursor"`
} // @name MetadataCursor
