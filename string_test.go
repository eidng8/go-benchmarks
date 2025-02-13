package benchmarks

// see also https://dev.to/pmalhaire/concatenate-strings-in-golang-a-quick-benchmark-4ahh

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var (
	shortStr = "short"
	shortLen = len(shortStr)
	longStr  = strings.Repeat("long ", 100)
	longLen  = len(longStr)
)

func Benchmark_string_concat_short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = shortStr + shortStr
	}
}

func Benchmark_string_sprintf_short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%d", shortStr, 1)
	}
}

func Benchmark_string_concat_conv_short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = shortStr + strconv.Itoa(1)
	}
}

func Benchmark_string_builder_short(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(shortStr)
	}
}

func Benchmark_string_builder_fprintf_short(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		_, _ = fmt.Fprintf(&sb, "%s%d", shortStr, 1)
	}
}

func Benchmark_string_builder_conv_short(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(shortStr)
		sb.WriteString(strconv.Itoa(1))
	}
}

func Benchmark_string_builder_grown_short(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow(shortLen * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString(shortStr)
	}
}

func Benchmark_string_builder_grown_fprintf_short(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow(shortLen * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = fmt.Fprintf(&sb, "%s%d", shortStr, 1)
	}
}

func Benchmark_string_builder_grown_conv_short(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow((shortLen + 1) * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString(shortStr)
		sb.WriteString(strconv.Itoa(1))
	}
}

func Benchmark_string_concat_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = longStr + longStr
	}
}

func Benchmark_string_sprintf_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%d", longStr, 1)
	}
}

func Benchmark_string_concat_conv_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = longStr + strconv.Itoa(1)
	}
}

func Benchmark_string_builder_long(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(longStr)
	}
}

func Benchmark_string_builder_fprintf_long(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		_, _ = fmt.Fprintf(&sb, "%s%d", longStr, 1)
	}
}

func Benchmark_string_builder_conv_long(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(longStr)
		sb.WriteString(strconv.Itoa(1))
	}
}

func Benchmark_string_builder_grown_long(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow(longLen * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString(longStr)
	}
}

func Benchmark_string_builder_grown_fprintf_long(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow(longLen * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = fmt.Fprintf(&sb, "%s%d", longStr, 1)
	}
}

func Benchmark_string_builder_grown_conv_long(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow((longLen + 1) * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString(longStr)
		sb.WriteString(strconv.Itoa(1))
	}
}
