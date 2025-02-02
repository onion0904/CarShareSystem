package scalar

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalDateTime は time.Time を GraphQL のスカラーに変換
func MarshalDateTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf(`"%s"`, t.Format(time.RFC3339)))
	})
}

// UnmarshalDateTime は GraphQL のスカラーを time.Time に変換
func UnmarshalDateTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return time.Parse(time.RFC3339, v)
	default:
		return time.Time{}, fmt.Errorf("invalid DateTime format")
	}
}
