package controllers

import (
    "../models"
    _ "encoding/json"
    _ "fmt"

    "github.com/astaxie/beego"
)

type MeetingController struct {
    beego.Controller
}

func (c *MeetingController) ParseParams() *models.MeetingParams {
    var params models.MeetingParams

    // // parse accepted URL parameters
    // meetingIds := c.GetStrings("meetingId")
    // roomIds := c.GetStrings("roomId")

    // // parse request body
    // err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
    // if err != nil {
    //     fmt.Println("Error parsing request body", err)
    // }

    // // give priority to parameters set in the URL
    // if len(meetingIds) > 0 {
    //     params.Filters.MeetingIds = meetingIds
    // }
    // if len(roomIds) > 0 {
    //     params.Filters.RoomIds = roomIds
    // }

    return &params
}

// @Title Create Meetings
// @Description creates a Meeting
// @Success 200 {object} models.Meeting
// @router / [post]
func (c *MeetingController) Create() {
    // params := c.ParseParams()
    // filters := &params.Filters
    // var response models.RecordingResponse

    // if filters == nil || (len(filters.RoomIds) == 0 && len(filters.MeetingIds) == 0) {
    //     err := models.APIError{"NoFilters", "Request aborted because no filters were provided", nil}
    //     errs := []models.APIError{err}

    //     // TODO: set another status code?
    //     response = models.RecordingResponse{nil, errs}
    // } else {
    //     recs, errs := models.DeleteAllRecordings(filters)
    //     if len(recs) == 0 { recs = nil }
    //     if len(errs) == 0 { errs = nil }

    //     response = models.RecordingResponse{recs, errs}
    // }
    // c.Data["json"] = response
    // c.ServeJson()
}
