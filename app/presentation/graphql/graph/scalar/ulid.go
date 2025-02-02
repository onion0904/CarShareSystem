package scalar

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/onion0904/go-pkg/ulid"
)

// MarshalID は ULID を文字列に変換
func MarshalID(u string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf(`"%s"`, u))
	})
}

// UnmarshalID は文字列を ULID に変換
func UnmarshalID(v interface{}) (string, error) {
	switch v := v.(type) {
	case string:
		if ulid.IsValid(v) {
			return v, nil
		}
		return "", fmt.Errorf("invalid ULID format")
	default:
		return "", fmt.Errorf("invalid ULID format")
	}
}
