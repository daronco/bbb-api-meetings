package models

import (
    "time"
    _ "errors"
    _ "../utils"
    "strconv"
    "crypto/sha1"
    "encoding/hex"
    "text/template"
    "fmt"
    "bytes"
)

var (
    MeetingList map[string]*Meeting
)

func init() {
    MeetingList = make(map[string]*Meeting)
    r1 := Meeting{MeetingId: "meeting_1", RoomId: "room_1", Name: "A Nice Meeting", Record: true}
    MeetingList["meeting_1"] = &r1
    // r2 := Meeting{"meeting_2", "room_1", "Another Nice Meeting", false}
    // MeetingList["meeting_2"] = &r2
    // r3 := Meeting{"meeting_3", "room_2", "Yet Another", true}
    // MeetingList["meeting_3"] = &r3
}

type Meeting struct {
    MeetingId         string    `json:"meetingId"`
    RoomId            string    `json:"roomId"`
    Name              string    `json:"name"`
    Running           bool      `json:"running"`
    Record            bool      `json:"record"`
    Duration          int       `json:"duration"`
    CreateTime        time.Time `json:"createTime"`
    ModeratorPassword string    `json:"moderatorPassword"`
    AttendeePassword  string    `json:"attendeePassword"`
    VoiceBridge       int       `json:"voiceBridge"`
}

func (this Meeting) CreateTimeAsString() string {
    return "Testing"
}

func (this Meeting) CreateTimeAsTimestamp() int {
    return 123123123
}

func GenerateMeetingId(roomId string) string {
    h := sha1.New()
    h.Write([]byte(roomId))
    roomIdSha := hex.EncodeToString(h.Sum(nil))

    timestamp := time.Now().UnixNano() / 1000 // ms instead of nano
    timestampStr := strconv.FormatInt(timestamp, 10)

    return roomIdSha + "-" + timestampStr
}

func GetAllMeetings(filters *MeetingFilters) []*Meeting {
    ret := []*Meeting{}
    for _, meeting := range MeetingList {
        ret = append(ret, meeting)
    }
    return ret
}

func CreateMeeting(params *Meeting) (a *Meeting, err *APIError) {
    var msg bytes.Buffer
    t, _ := template.ParseFiles("templates/create-meeting-request.json.tmpl")
    t.Execute(&msg, params)
    fmt.Println("---", msg.String())

    MeetingList[params.MeetingId] = params
    return MeetingList[params.MeetingId], nil
}
