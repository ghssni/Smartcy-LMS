package utils

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

// StringToUint converts a string to a uint.
// If the conversion fails, it returns 0.
func StringToUint(s string) uint32 {
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		// Return 0 if there's an error
		return 0
	}
	return uint32(val)
}

func ProtoTimeToDate(protoTimestamp *timestamppb.Timestamp) string {
	if protoTimestamp == nil {
		return ""
	}

	// Convert google.protobuf.Timestamp to time.Time
	goTime := protoTimestamp.AsTime()

	// Format time.Time to "dd-mm-yyyy"
	return goTime.Format("02-01-2006")
}
