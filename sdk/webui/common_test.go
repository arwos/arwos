package webui_test

import (
	"encoding/json"
	"testing"

	"go.osspkg.com/casecheck"
)

func TestAPI(t *testing.T) {
	data := AreaFluid(
		Row(
			Col(
				ID("id"),
				Class("md", "text"),
				Row(
					ID("id"),
				),
			),
			Col(
				Row(),
			),
		),
	)

	b, err := json.Marshal(data)
	casecheck.NoError(t, err)
	casecheck.Equal(t, ``, string(b))
}

func BenchmarkAPI(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data := Row(
				Col(
					ID("id"),
					Row(
						ID("id"),
					),
				),
				Col(
					Row(),
				),
			)
			json.Marshal(data)
		}
	})
}
