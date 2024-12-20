package nex

import (
	"fmt"
	"strings"
	"time"
)

type DateTime struct {
	value uint64
}

// FromComponents creates a DateTime instance from individual date and time components
func (dt *DateTime) FromComponents(year, month, day, hour, minute, second int) *DateTime {
	dt.value = uint64(second) |
		(uint64(minute) << 6) |
		(uint64(hour) << 12) |
		(uint64(day) << 17) |
		(uint64(month) << 22) |
		(uint64(year) << 26)

	return dt
}

// FromTimestamp converts a time.Time object into a DateTime
func (dt *DateTime) FromTimestamp(timestamp time.Time) *DateTime {
	return dt.FromComponents(
		timestamp.Year(),
		int(timestamp.Month()),
		timestamp.Day(),
		timestamp.Hour(),
		timestamp.Minute(),
		timestamp.Second(),
	)
}

// Now returns a DateTime set to the current UTC time
func (dt *DateTime) Now() *DateTime {
	return dt.FromTimestamp(time.Now().UTC())
}

// Getters for DateTime components
func (dt *DateTime) Second() int { return int(dt.value & 0x3F) }
func (dt *DateTime) Minute() int { return int((dt.value >> 6) & 0x3F) }
func (dt *DateTime) Hour() int   { return int((dt.value >> 12) & 0x1F) }
func (dt *DateTime) Day() int    { return int((dt.value >> 17) & 0x1F) }
func (dt *DateTime) Month() int  { return int((dt.value >> 22) & 0x0F) }
func (dt *DateTime) Year() int   { return int(dt.value >> 26) }

// Value returns the raw uint64 value stored in DateTime
func (dt *DateTime) Value() uint64 {
	return dt.value
}

// ToTime converts the DateTime to a standard time.Time object
func (dt *DateTime) ToTime() time.Time {
	return time.Date(
		dt.Year(),
		time.Month(dt.Month()),
		dt.Day(),
		dt.Hour(),
		dt.Minute(),
		dt.Second(),
		0,
		time.UTC,
	)
}

// String returns a human-readable representation of the DateTime
func (dt *DateTime) String() string {
	return fmt.Sprintf("DateTime{value: %d, time: %s}", dt.value, dt.ToTime().Format("2006-01-02 15:04:05"))
}

// DebugString returns a formatted string with indentation for debugging
func (dt *DateTime) DebugString(indentLevel int) string {
	indent := strings.Repeat("\t", indentLevel)
	return fmt.Sprintf("%sDateTime{\n%s\tValue: %d,\n%s\tTime: %s\n%s}",
		indent,
		indent,
		dt.value,
		indent,
		dt.ToTime().Format("2006-01-02 15:04:05"),
		indent,
	)
}

// NewDateTime initializes a DateTime with the given value
func NewDateTime(value uint64) *DateTime {
	return &DateTime{value: value}
}
