package utils

import (
    "time"
    "strconv"
    "crypto/sha1"
    "encoding/hex"
)

func GenerateMeetingId(roomId string) string {
    h := sha1.New()
    h.Write([]byte(roomId))
    roomIdSha := hex.EncodeToString(h.Sum(nil))

    timestamp := time.Now().UnixNano() / 1000 // ms instead of nano
    timestampStr := strconv.FormatInt(timestamp, 10)

    return roomIdSha + "-" + timestampStr
}
