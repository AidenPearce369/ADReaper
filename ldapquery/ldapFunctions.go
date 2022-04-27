package ldapquery

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	"github.com/audibleblink/bamflags"
	"github.com/bwmarrin/go-objectsid"
)

func ParseLDAPFlag(attr int64, flagMap map[int]string) (flags []string) {
	values, err := bamflags.ParseInt(attr)
	if err != nil {
		return
	}
	for _, value := range values {
		if propName := flagMap[value]; propName != "" {
			flags = append(flags, propName)
		}
	}
	return
}

func ParseGUID(b []byte) (string, error) {
	if len(b) != 16 {
		return "", fmt.Errorf("[+]GUID must be 16 bytes")
	}
	return fmt.Sprintf(
		"%08x-%04x-%04x-%04x-%012x",
		binary.LittleEndian.Uint32(b[:4]),
		binary.LittleEndian.Uint16(b[4:6]),
		binary.LittleEndian.Uint16(b[6:8]),
		b[8:10],
		b[10:]), nil
}

func ParseSID(b []byte) (string, error) {
	if len(b) < 12 {
		return "", fmt.Errorf("[+]Invalid Windows SID")
	}
	sid := objectsid.Decode(b)
	return sid.String(), nil
}

func ParseTimeStamp(s string) (timestamp time.Time) {
	ticks, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	secs := (int64)((ticks / 10000000) - 11644473600)
	nsecs := (int64)((ticks % 10000000) * 100)

	return time.Unix(secs, nsecs).UTC()
}

func ParseSystemObject(s string) string {
	if len(s) > 0 {
		return "a"
	}
	return "a"
}
