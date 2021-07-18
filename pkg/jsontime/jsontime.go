package jsontime

import "time"

// Mon Jan 2 15:04:05 -0700 MST 2006
// 2016-06-03T10:50:56 -10:00

// ZDTimeFormat (Zendesk Time Format) is the time format for Zendesk files
const ZDTimeFormat = "2006-01-02T15:04:05 -07:00"

type Time struct {
	time.Time
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	t.Time, err = time.Parse(`"`+ZDTimeFormat+`"`, string(data))
	return err
}
