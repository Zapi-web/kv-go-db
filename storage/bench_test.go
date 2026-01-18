package storage

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"go.uber.org/zap"
)

func BenchmarkSet(b *testing.B) {
	path := filepath.Join(b.TempDir(), "bench.db")
	d, _ := Init(path, zap.NewNop())
	defer d.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		value := "value"

		err := d.Set(key, value)
		if err != nil {
			log.Fatal(err)
		}
	}
}
