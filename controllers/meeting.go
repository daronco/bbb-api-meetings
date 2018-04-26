package controllers

import (
    "../models"
    "../utils"
    "encoding/json"
    "fmt"
    "strings"

    "github.com/astaxie/beego"
)

type MeetingController struct {
    beego.Controller
}

// @Title Get Meetings
// @Description get a list of Meetings
// @Success 200 {object} models.Meeting
// @router / [get]
func (c *MeetingController) Index() {
    params := c.ParseParams()
    meetings := models.GetAllMeetings(&params.Filters)
    if len(meetings) == 0 { meetings = nil }

    response := models.MeetingResponse{meetings, nil}
    c.Data["json"] = response
    c.ServeJson()
}

// @Title Create Meetings
// @Description creates a Meeting
// @Success 200 {object} models.Meeting
// @router / [post]
func (c *MeetingController) Create() {
    var response models.MeetingResponse
    params := c.ParseParams()

    valid, errs := c.ValidateParamsForCreate(params)
    if !valid {
        // TODO: set another status code?
        response = models.MeetingResponse{nil, errs}
    } else {
        params = c.PrepareParamsForCreate(params)

        var errs []models.APIError
        beego.Info("Creating meeting with the attributes", params.Attributes)
        meeting, err := models.CreateMeeting(&params.Attributes)
        if err != nil {
            errs = []models.APIError{*err}
        } else {
            errs = nil
        }
        meetings := []*models.Meeting{}
        if meeting != nil { meetings = append(meetings, meeting) }

        response = models.MeetingResponse{meetings, errs}
    }

    c.Data["json"] = response
    c.ServeJson()
}

func (c *MeetingController) ParseParams() *models.MeetingParams {
    var params models.MeetingParams

    // parse request body
    err := json.Unmarshal(c.Ctx.Input.RequestBody, &params)
    if err != nil {
        fmt.Println("Error parsing request body", err)
    }

    return &params
}

func (c *MeetingController) ValidateParamsForCreate(params *models.MeetingParams) (a bool, errs []models.APIError) {
    errors := []models.APIError{}

    if strings.TrimSpace(params.Attributes.RoomId) == "" {
        err := models.APIError{"InvalidParameter", "Invalid parameter roomId", nil}
        errors = append(errors, err)
    }

    return (len(errors) <= 0), errors
}

func (c *MeetingController) PrepareParamsForCreate(params *models.MeetingParams) *models.MeetingParams {
    params.Attributes.MeetingId = models.GenerateMeetingId(params.Attributes.RoomId)
    params.Attributes.AttendeePassword, _ = utils.RandomHex(8)
    params.Attributes.ModeratorPassword, _ = utils.RandomHex(8)
    return params
}
