package benchmarks

import (
	"encoding/json"
	"testing"

	"github.com/json-iterator/go"
)

var (
	jsonIter = jsoniter.ConfigFastest

	sampleJSON = []byte(`{"name":"John","age":10,"boy":true,"physical":{"height":123.4,"weight":32.1},"hobbies":["swimming","jogging","reading"],"scores":[5,4,3,2,1]}`)
	sampleMap  = map[string]interface{}{
		"name": "John",
		"age":  10,
		"boy":  true,
		"physical": map[string]interface{}{
			"height": 123.4,
			"weight": 32.1,
		},
		"hobbies": []interface{}{"swimming", "jogging", "reading"},
		"scores":  []interface{}{5, 4, 3, 2, 1},
	}
	sampleStruct = person{
		Name: "John",
		Age:  10,
		Boy:  true,
		Physical: struct {
			Height float64 `json:"height,omitempty"`
			Weight float64 `json:"weight,omitempty"`
		}{
			Height: 123.4,
			Weight: 32.1,
		},
		Hobbies: []string{"swimming", "jogging", "reading"},
		Scores:  []int{5, 4, 3, 2, 1},
	}
)

type person struct {
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age,omitempty"`
	Boy      bool   `json:"boy,omitempty"`
	Physical struct {
		Height float64 `json:"height,omitempty"`
		Weight float64 `json:"weight,omitempty"`
	}
	Hobbies []string `json:"hobbies,omitempty"`
	Scores  []int    `json:"scores,omitempty"`
}

func Benchmark_json_unmarshal_map(b *testing.B) {
	sut := map[string]interface{}{}
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(sampleJSON, &sut)
	}
}

func Benchmark_json_unmarshal_struct(b *testing.B) {
	sut := person{}
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(sampleJSON, &sut)
	}
}

func Benchmark_json_marshal_map(b *testing.B) {
	sut := sampleMap
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(sut)
	}
}

func Benchmark_json_marshal_struct(b *testing.B) {
	sut := sampleStruct
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(sut)
	}
}

func Benchmark_jsoniter_unmarshal_map(b *testing.B) {
	sut := map[string]interface{}{}
	for i := 0; i < b.N; i++ {
		_ = jsonIter.Unmarshal(sampleJSON, &sut)
	}
}

func Benchmark_jsoniter_unmarshal_struct(b *testing.B) {
	sut := person{}
	for i := 0; i < b.N; i++ {
		_ = jsonIter.Unmarshal(sampleJSON, &sut)
	}
}

func Benchmark_jsoniter_marshal_map(b *testing.B) {
	sut := sampleMap
	for i := 0; i < b.N; i++ {
		_, _ = jsonIter.Marshal(sut)
	}
}

func Benchmark_jsoniter_marshal_struct(b *testing.B) {
	sut := sampleStruct
	for i := 0; i < b.N; i++ {
		_, _ = jsonIter.Marshal(sut)
	}
}
