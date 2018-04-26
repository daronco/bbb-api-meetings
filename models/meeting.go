package models

import (
    _ "errors"
    _ "../utils"
    _ "fmt"
    _ "strconv"
    _ "time"
)

var (
    MeetingList map[string]*Meeting
)

func init() {
    MeetingList = make(map[string]*Meeting)
    r1 := Meeting{"meeting_1", "A Nice Meeting", "room_1", true}
    MeetingList["meeting_1"] = &r1
    // r2 := Meeting{"meeting_2", "Another Nice Meeting", "room_1", false}
    // MeetingList["meeting_2"] = &r2
    // r3 := Meeting{"meeting_3", "Yet Another", "room_2", true}
    // MeetingList["meeting_3"] = &r3
}

type Meeting struct {
    MeetingId string `json:"meetingId"`
    Name      string `json:"name"`
    RoomId    string `json:"roomId"`
    Running   bool   `json:"running"`
}

func GetAllMeetings(filters *MeetingFilters) []*Meeting {
    ret := []*Meeting{}
    for _, meeting := range MeetingList {
        ret = append(ret, meeting)
    }
    return ret
}

func CreateMeeting(params *Meeting) (a *Meeting, err *APIError) {
    MeetingList[params.MeetingId] = params
    return MeetingList[params.MeetingId], nil
}
