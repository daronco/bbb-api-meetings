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
    WelcomeMessage    string    `json:"welcomeMessage"`
    DialNumber        string    `json:"dialNumber"`
}

// Returns CreateTime as a string in the format:
// "Tue Apr 24 21:27:53 UTC 2018"
func (this Meeting) CreateTimeAsString() string {
    return this.CreateTime.UTC().Format("Mon Jan 2 15:04:05 MST 2006")
}

func (this Meeting) CreateTimeAsTimestamp() int {
    return int(this.CreateTime.UnixNano() / 1000000) // ms instead of nano
}

func GenerateMeetingId(roomId string) string {
    h := sha1.New()
    h.Write([]byte(roomId))
    roomIdSha := hex.EncodeToString(h.Sum(nil))

    timestamp := time.Now().UnixNano() / 1000000 // ms instead of nano
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
    params.CreateTime = time.Now()

    // TODO: tmp
    var msg bytes.Buffer
    t, _ := template.ParseFiles("templates/create-meeting-request.json.tmpl")
    t.Execute(&msg, params)
    fmt.Println("---", msg.String())
    // redis.CreateMeeting(params)

    MeetingList[params.MeetingId] = params
    return MeetingList[params.MeetingId], nil
}
