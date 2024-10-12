package ctags

import (
	"os"
	"testing"
)

func readContent() []byte {
	b, err := os.ReadFile("./../gitindex/index.go")
	if err != nil {
		panic(err)
	}
	return b
}

func BenchmarkParserUniversal(b *testing.B) {
	content := readContent()

	parser := NewCTagsParser(map[CTagsParserType]string{UniversalCTags: "/Users/hiromu.nakamura/intel/homebrew/bin/ctags"})
	defer parser.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := parser.Parse("test.go", content, UniversalCTags)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParserScip(b *testing.B) {
	content := readContent()

	parser := NewCTagsParser(map[CTagsParserType]string{ScipCTags: "scip-ctags"})
	defer parser.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := parser.Parse("test.go", content, ScipCTags)
		if err != nil {
			b.Fatal(err)
		}
	}
}
