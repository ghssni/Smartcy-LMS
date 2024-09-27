package utils

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
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

func ProtoTimetoTime(protoTimestamp *timestamppb.Timestamp) string {
	if protoTimestamp == nil {
		return ""
	}

	// Convert google.protobuf.Timestamp to time.Time
	goTime := protoTimestamp.AsTime()

	// Format time.Time to "dd-mm-yyyy hh:mm:ss"
	return goTime.Format("02-01-2006 15:04:05")
}

func ProtoTimetoTimeWithSeconds(protoTimestamp *timestamppb.Timestamp) string {
	if protoTimestamp == nil {
		return ""
	}

	// Convert google.protobuf.Timestamp to time.Time
	goTime := protoTimestamp.AsTime()

	// Format time.Time to "dd-mm-yyyy hh:mm:ss"
	return goTime.Format("02-01-2006 15:04:05")
}

func TimeToProtoTime(goTime string) *timestamppb.Timestamp {
	// Parse time string to time.Time
	t, err := time.Parse("02-01-2006", goTime)
	if err != nil {
		return nil
	}

	// Convert time.Time to google.protobuf.Timestamp
	return timestamppb.New(t)
}

// TimeToProtoTimestamp time.Time to google.protobuf.Timestamp
func TimeToProtoTimestamp(goTime time.Time) *timestamppb.Timestamp {
	return timestamppb.New(goTime)
}

func FormatTimestamp(ts *timestamp.Timestamp) string {
	if ts == nil {
		return ""
	}
	return ts.AsTime().Format(time.RFC3339)
}
