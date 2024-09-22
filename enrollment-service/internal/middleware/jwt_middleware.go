package middleware

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"os"
	"strings"
)

// JWTInterceptor is a middleware that checks if the request has a valid JWT token
func JWTInterceptor(secretKey string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if info.FullMethod == "/payments.PaymentsService/UpdateExpiredPaymentStatus" {
			return handler(ctx, req)
		}

		// Retrieve metadata from context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("missing metadata")
		}

		// Retrieve the authorization header
		authHeader, ok := md["authorization"]
		if !ok || len(authHeader) == 0 {
			return nil, errors.New("authorization token is not provided")
		}

		// Extract token string from the Bearer token
		tokenString := strings.TrimSpace(strings.Replace(authHeader[0], "Bearer", "", 1))

		// Parse and verify the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method")
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			return nil, errors.New("invalid or expired token")
		}

		// Proceed to the actual handler
		return handler(ctx, req)
	}
}

// GetUserIDFromToken extracts the user ID from the JWT token
func GetUserIDFromToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	tokenString := strings.TrimSpace(strings.Replace(authHeader[0], "Bearer", "", 1))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return "", status.Errorf(codes.Unauthenticated, "invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", status.Errorf(codes.Unauthenticated, "invalid token claims")
	}

	userID, ok := claims["id"].(string)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "user ID not found in token")
	}

	return userID, nil
}

// GetEmailFromToken extracts the email from the JWT token
func GetEmailFromToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	tokenString := strings.TrimSpace(strings.Replace(authHeader[0], "Bearer", "", 1))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return "", status.Errorf(codes.Unauthenticated, "invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", status.Errorf(codes.Unauthenticated, "invalid token claims")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "email not found in token")
	}

	return email, nil
}

// GetTokenFromContext extracts the JWT token from the context
func GetTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	tokenString := strings.TrimSpace(strings.Replace(authHeader[0], "Bearer", "", 1))

	return tokenString, nil
}

// AccessKeyInterceptor is a middleware that checks if the request has a valid access key
func AccessKeyInterceptor(expectedKey string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if info.FullMethod == "/payments.PaymentsService/UpdateExpiredPaymentStatus" {
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, status.Errorf(codes.Unauthenticated, "Missing metadata")
			}

			accessKeys := md.Get("X-ACCESS-KEY")
			if len(accessKeys) == 0 || accessKeys[0] != expectedKey {
				return nil, status.Errorf(codes.Unauthenticated, "Invalid access key")
			}
		}
		return handler(ctx, req)
	}
}
