package model

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	for _, c := range []struct {
		name       string
		marshalled []byte
	}{
		{
			name:       "Empty",
			marshalled: []byte("{}"),
		},
		{
			name:       "EmptyBaz",
			marshalled: []byte(`{
				"baz": {}
			}`),
		},
		{
			name:       "BazTypeFoo",
			marshalled: []byte(`{
				"baz": {
					"foo": "some-data"
				}
			}`),
		},
		{
			name:       "BazTypeFooWithTypename",
			marshalled: []byte(`{
				"baz": {
					"__typename": "Foo",
					"foo": "some-data"
				}
			}`),
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			d := Data{}
			if err := json.Unmarshal(c.marshalled, &d); err != nil {
				t.Fatal(err)
			}
		})
	}
}
