package models

import (
    _ "encoding/json"
)

// Description of an error
type APIError struct {
    Code string        `json:"code"`
    Description string `json:"description"`
    Source *string     `json:"source"`
}

// A generic API error response
type APIErrorResponse struct {
    Errors []APIError `json:"errors"`
}

// // Default response for meeting routes
// type RecordingResponse struct {
//     Data []*Recording `json:"data"`
//     Errors []APIError `json:"errors"`
// }

// Parameters accepted when filtering meetings (used in multiple routes)
type MeetingFilters struct {
    // MeetingIds []string `json:"meetingId"`
    RoomIds []string    `json:"roomId"`
}

// Parameters accepted when receiving requests for meetings
type MeetingParams struct {
    Filters MeetingFilters `json:"filters"`
    Attributes Meeting     `json:"attributes"`
}
