package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// GetBasicToken fetches a basic authorization token from the http request
func GetBasicToken(r *http.Request) (username, password string, err error) {
	if r == nil {
		return "", "", fmt.Errorf("invalid HTTP request")
	}
	value := r.Header.Get("Authorization")
	if strings.HasPrefix(value, "Basic ") || strings.HasPrefix(value, "basic ") {
		split := strings.Split(value, " ")
		if len(split) == 2 {
			decoded, err := base64.StdEncoding.DecodeString(split[1])
			if err != nil {
				return "", "", fmt.Errorf("basic Authorization token is malformed")
			}
			split = strings.Split(string(decoded), ":")
			if len(split) == 2 {
				return split[0], split[1], nil
			}
		}
	}
	return "", "", fmt.Errorf("basic Authorization token is missing")
}

// GetBearerToken fetches a bearer authorization token from the http request
func GetBearerToken(r *http.Request) (string, error) {
	if r == nil {
		return "", fmt.Errorf("invalid HTTP request")
	}
	value := r.Header.Get("Authorization")
	if strings.HasPrefix(value, "Bearer ") || strings.HasPrefix(value, "bearer ") {
		split := strings.Split(value, " ")
		if len(split) == 2 {
			return split[1], nil
		}
	}
	return "", fmt.Errorf("bearer Authorization token is missing")
}

// ValidateIPv6 validates a v6 IP address
func ValidateIPv6(ip string, withPrefix bool) bool {
	ipv6Regex := `\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])(\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])){3}))|:)))(%.+)?\s*`
	if withPrefix {
		ipv6Regex = ipv6Regex + "/[0-9]{1,3}"
	}
	return regexp.MustCompile(ipv6Regex).MatchString(ip)
}

// ValidateIPv4 validates a v4 IP address
func ValidateIPv4(ip string, withPrefix bool) bool {
	ipv4Regex := `^(?:[0-9]{1,3}[.]){3}[0-9]{1,3}$`
	if withPrefix {
		ipv4Regex = `^(?:[0-9]{1,3}[.]){3}[0-9]{1,3}/[0-9]{1,2}$`
	}
	return regexp.MustCompile(ipv4Regex).MatchString(ip)
}

// define a key for our context value
type contextKey string

const key contextKey = "context-values"

// EnrichContextWithValues enriches the current context with a map ok key value pairs
func EnrichContextWithValues(ctx context.Context, values map[string]interface{}) context.Context {
	return context.WithValue(ctx, key, values)
}

// FetchContextValues returns the map of "context-values" values added in the current context
func FetchContextValues(ctx context.Context) map[string]interface{} {
	values := ctx.Value(key)
	if values == nil {
		return map[string]interface{}{}
	}
	result, ok := values.(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}
	return result
}

// EnrichContextWithValue enriches the context with a key value pair
func EnrichContextWithValue[T any](ctx context.Context, key string, value T) context.Context {
	values := FetchContextValues(ctx)
	values[key] = value
	return EnrichContextWithValues(ctx, values)
}

// FetchContextValue return a pointer to the specified value got from the context
func FetchContextValue[T any](ctx context.Context, key string) *T {
	values := FetchContextValues(ctx)
	value, ok := values[key]
	if !ok {
		return nil
	}
	v, ok := value.(T)
	if !ok {
		return nil
	}
	return &v
}
