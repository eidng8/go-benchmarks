package benchmarks

// see also https://dev.to/pmalhaire/concatenate-strings-in-golang-a-quick-benchmark-4ahh

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var (
	emptyStr = ""
	shortStr = "short"
	shortLen = len(shortStr)
	longStr  = strings.Repeat("long ", 100)
	longLen  = len(longStr)
	shortArr = []string{"short", "short", "short"}
	longArr  = []string{
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
		"long", "long", "long", "long", "long", "long", "long", "long", "long",
	}
)

func Benchmark_string_conditional_assign_short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//goland:noinspection GoBoolExpressions
		if "" != emptyStr {
			_ = shortStr
		}
	}
}

func Benchmark_string_unconditional_assign_short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = shortStr
	}
}

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

func Benchmark_string_sprintf_float_short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%.10f", shortStr, 1.2)
	}
}

func Benchmark_string_concat_Itoa_short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = shortStr + strconv.Itoa(1)
	}
}

func Benchmark_string_concat_FormatInt_short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = shortStr + strconv.FormatInt(1, 10)
	}
}

func Benchmark_string_concat_FormatFloat_short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = shortStr + strconv.FormatFloat(1, 'f', 10, 64)
	}
}

func Benchmark_string_builder_short(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(shortStr)
		_ = sb.String()
	}
}

func Benchmark_string_builder_fprintf_short(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		_, _ = fmt.Fprintf(&sb, "%s%d", shortStr, 1)
		_ = sb.String()
	}
}

func Benchmark_string_builder_fprintf_float_short(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		_, _ = fmt.Fprintf(&sb, "%s%.10f", shortStr, 1.2)
		_ = sb.String()
	}
}

func Benchmark_string_builder_Itoa_short(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(shortStr)
		sb.WriteString(strconv.Itoa(1))
		_ = sb.String()
	}
}

func Benchmark_string_builder_FormatInt_short(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(shortStr)
		sb.WriteString(strconv.FormatInt(1, 10))
		_ = sb.String()
	}
}

func Benchmark_string_builder_FormatFloat_short(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(shortStr)
		sb.WriteString(strconv.FormatFloat(1, 'f', 10, 64))
		_ = sb.String()
	}
}

func Benchmark_string_builder_grown_short(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow(shortLen * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString(shortStr)
		_ = sb.String()
	}
}

func Benchmark_string_builder_grown_fprintf_short(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow(shortLen * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = fmt.Fprintf(&sb, "%s%d", shortStr, 1)
		_ = sb.String()
	}
}

func Benchmark_string_builder_grown_fprintf_float_short(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow(shortLen * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = fmt.Fprintf(&sb, "%s%.10f", shortStr, 1.2)
		_ = sb.String()
	}
}

func Benchmark_string_builder_grown_Itoa_short(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow((shortLen + 1) * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString(shortStr)
		sb.WriteString(strconv.Itoa(1))
		_ = sb.String()
	}
}

func Benchmark_string_builder_grown_FormatInt_short(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow((shortLen + 1) * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString(shortStr)
		sb.WriteString(strconv.FormatInt(1, 10))
		_ = sb.String()
	}
}

func Benchmark_string_builder_grown_FormatFloat_short(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow((shortLen + 1) * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString(shortStr)
		sb.WriteString(strconv.FormatFloat(1, 'f', 10, 64))
		_ = sb.String()
	}
}

func Benchmark_string_conditional_assign_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//goland:noinspection GoBoolExpressions
		if "" != emptyStr {
			_ = longStr
		}
	}
}

func Benchmark_string_unconditional_assign_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = longStr
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

func Benchmark_string_sprintf_float_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%.10f", longStr, 1.2)
	}
}

func Benchmark_string_concat_Itoa_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = longStr + strconv.Itoa(1)
	}
}

func Benchmark_string_concat_FormatInt_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = longStr + strconv.FormatInt(1, 10)
	}
}

func Benchmark_string_concat_FormatFloat_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = longStr + strconv.FormatFloat(1, 'f', 10, 64)
	}
}

func Benchmark_string_builder_long(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(longStr)
		_ = sb.String()
	}
}

func Benchmark_string_builder_fprintf_long(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		_, _ = fmt.Fprintf(&sb, "%s%d", longStr, 1)
		_ = sb.String()
	}
}

func Benchmark_string_builder_fprintf_float_long(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		_, _ = fmt.Fprintf(&sb, "%s%.10f", longStr, 1.2)
		_ = sb.String()
	}
}

func Benchmark_string_builder_Itoa_long(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(longStr)
		sb.WriteString(strconv.Itoa(1))
		_ = sb.String()
	}
}

func Benchmark_string_builder_FormatInt_long(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(longStr)
		sb.WriteString(strconv.FormatInt(1, 10))
		_ = sb.String()
	}
}

func Benchmark_string_builder_FormatFloat_long(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		sb.WriteString(longStr)
		sb.WriteString(strconv.FormatFloat(1, 'f', 10, 64))
		_ = sb.String()
	}
}

func Benchmark_string_builder_grown_long(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow(longLen * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString(longStr)
		_ = sb.String()
	}
}

func Benchmark_string_builder_grown_fprintf_long(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow(longLen * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = fmt.Fprintf(&sb, "%s%d", longStr, 1)
		_ = sb.String()
	}
}

func Benchmark_string_builder_grown_fprintf_float_long(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow(longLen * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = fmt.Fprintf(&sb, "%s%.10f", longStr, 1.2)
		_ = sb.String()
	}
}

func Benchmark_string_builder_grown_Itoa_long(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow((longLen + 1) * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString(longStr)
		sb.WriteString(strconv.Itoa(1))
		_ = sb.String()
	}
}

func Benchmark_string_builder_grown_FormatInt_long(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow((longLen + 1) * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString(longStr)
		sb.WriteString(strconv.FormatInt(1, 10))
		_ = sb.String()
	}
}

func Benchmark_string_builder_grown_FormatFloat_long(b *testing.B) {
	b.StopTimer()
	var sb strings.Builder
	sb.Grow((longLen + 1) * b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString(longStr)
		sb.WriteString(strconv.FormatFloat(1, 'f', 10, 64))
		_ = sb.String()
	}
}

func Benchmark_string_compare_short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//goland:noinspection GoBoolExpressions
		_ = shortStr == shortStr
	}
}

func Benchmark_string_compare_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//goland:noinspection GoDfaConstantCondition
		_ = longStr == longStr
	}
}

func Benchmark_string_HasPrefix_short(b *testing.B) {
	prefix := shortStr[:3]
	for i := 0; i < b.N; i++ {
		strings.HasPrefix(shortStr, prefix)
	}
}

func Benchmark_string_HasPrefix_long(b *testing.B) {
	prefix := longStr[:100]
	for i := 0; i < b.N; i++ {
		strings.HasPrefix(longStr, prefix)
	}
}

func Benchmark_string_HasSuffix_short(b *testing.B) {
	prefix := shortStr[:3]
	for i := 0; i < b.N; i++ {
		strings.HasSuffix(shortStr, prefix)
	}
}

func Benchmark_string_HasSuffix_long(b *testing.B) {
	prefix := longStr[:100]
	for i := 0; i < b.N; i++ {
		strings.HasSuffix(longStr, prefix)
	}
}

func Benchmark_strings_join_short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(shortArr, ",")
	}
}

func Benchmark_strings_join_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(longArr, ",")
	}
}

func Benchmark_strings_join_short_by_string_build(b *testing.B) {
	var sb strings.Builder
	sb.Grow(65536)
	b.ResetTimer()
	sb.WriteString(shortArr[0])
	for i := 0; i < b.N; i++ {
		sb.Reset()
		for _, s := range shortArr[1:] {
			sb.WriteString(",")
			sb.WriteString(s)
		}
		_ = sb.String()
	}
}

func Benchmark_strings_join_long_by_string_build(b *testing.B) {
	var sb strings.Builder
	sb.Grow(65536)
	b.ResetTimer()
	sb.WriteString(longArr[0])
	for i := 0; i < b.N; i++ {
		sb.Reset()
		for _, s := range longArr[1:] {
			sb.WriteString(",")
			sb.WriteString(s)
		}
		_ = sb.String()
	}
}
