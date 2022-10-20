package main

import (
	"testing"
	"reflect"
)

// Define a function named toJSONBytes that takes two string arguments,
// and returns a bytes[].
// In other words, define a function with the following signature:
// toJSONBytes(string, string) bytes[]

func TesttoJSONBytes(t *testing.T) {
	var message = "Hello Slack!"
	expected := []byte(message)
	if observed := toJSONBytes("text", message); reflect.DeepEqual(observed, expected) {
		t.Fatalf("toJSONBytes() = %v, want %v", observed, expected)
	}
}

// BenchmarktoJSONBytes() is a benchmarking function. These functions follow the
// form `func BenchmarkXxx(*testing.B)` and can be used to test the performance
// of your implementation. They may not be present in every exercise, but when
// they are you can run them by including the `-bench` flag with the `go test`
// command, like so: `go test -v --bench . --benchmem`
//
// You will see output similar to the following:
//
// BenchmarktoJSONBytes   	2000000000	         0.46 ns/op
//
// This means that the loop ran 2000000000 times at a speed of 0.46 ns per loop.
//
// While benchmarking can be useful to compare different iterations of the same
// exercise, keep in mind that others will run the same benchmarks on different
// machines, with different specs, so the results from these benchmark tests may
// vary.
func BenchmarktoJSONBytes(b *testing.B) {
	var message = "Hello Slack!"	
	for i := 0; i < b.N; i++ {
		toJSONBytes("text", message)
	}
}
