package models

type Request struct {
	DocumentID string `json:"document_id"`
	PageID     int    `json:"page_id"`
}

type Segment struct {
	DocumentID    string `json:"document_id"`
	PageID        int    `json:"page_id"`
	SegmentID     int    `json:"segment_id"`
	TotalSegments int    `json:"total_segments"`
	Payload       string `json:"payload"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
