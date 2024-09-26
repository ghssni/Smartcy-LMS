package utils

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

// ExtractSQLErrorCode extracts SQLSTATE code from an error message
func ExtractSQLErrorCode(err error) (string, error) {
	if err == nil {
		return "", nil
	}

	// Regex to extract SQLSTATE code (5 digits) from the error message
	re := regexp.MustCompile(`SQLSTATE (\d{5})`)
	matches := re.FindStringSubmatch(err.Error())

	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", fmt.Errorf("SQLSTATE code not found in error message")
}

// HandlePostgresError checks for specific SQLSTATE codes and returns appropriate gRPC errors
func HandlePostgresError(err error) error {
	if err == nil {
		return nil
	}

	sqlState, extractErr := ExtractSQLErrorCode(err)
	if extractErr != nil {
		return status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	switch sqlState {
	//Not Found
	case "02000":
		return status.Errorf(codes.NotFound, "Record not found: %v", err)
	case "23505": // Unique violation
		return status.Errorf(codes.AlreadyExists, "Record already exists: %v", err)
	case "23503": // Foreign key violation
		return status.Errorf(codes.NotFound, "Referenced record not found: %v", err)
	case "23502": // Not null violation
		return status.Errorf(codes.InvalidArgument, "Required field is empty: %v", err)
	default:
		return status.Errorf(codes.Internal, "Unhandled database error: %v", err)
	}
}
