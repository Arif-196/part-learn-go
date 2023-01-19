package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Table Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{name: "Arif", request: "Arif"}, {name: "Fadillah", request: "Fadillah"},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Arif", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Arif")
		}
	})
	b.Run("Fadillah", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fadillah")
		}
	})

}

// Benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Arif")
	}
}
func BenchmarkHelloWorldDua(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fadillah")
	}
}

func TestTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Arif)",
			request:  "Arif",
			expected: "Hello Arif",
		},
		{
			name:     "HelloWorld(Fadillah)",
			request:  "Fadillah",
			expected: "Hello Fadillah",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Arif", func(t *testing.T) {
		result := HelloWorld("Arif")
		require.Equal(t, "Hello Arif", result, "Result must be 'Hello Arif'")
	})

	t.Run("Fadillah", func(t *testing.T) {
		result := HelloWorld("Fadillah")
		require.Equal(t, "Hello Fadillah", result, "Result must be 'Hello Fadillah'")
	})

}

// Main Test
func TestMain(m *testing.M) {
	fmt.Println("Before unit test")
	m.Run()
	fmt.Println("After unit test")
}

func TestSkip(t *testing.T) {
	// checking operating system
	if runtime.GOOS == "windows" {
		t.Skip("Cannot Run on Windows OS")
	}
}

func TestHelloWorldRequired(t *testing.T) {
	result := HelloWorld("Arif")
	require.Equal(t, "Hello Arif", result, "Result must be 'Hello Arif'")
	fmt.Println("Function TestHelloWorldRequired Executed")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Arif")
	assert.Equal(t, "Hello Arif", result, "Result must be 'Hello Arif'")
	fmt.Println("Function TestHelloWorldAssert Executed")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Arif")
	if result != "Hello Arif" {
		// error
		// t.Fail() // if use Fail() Program will continue to proccess finish this function
		t.Error("Result must be 'Hello Arif'")
		// panic("Result is not Hello Arif")
	}
	fmt.Println("Function TestHelloWorld Executed")
}

func TestHelloWorldFadillah(t *testing.T) {
	result := HelloWorld("Fadillah")
	if result != "Hello Fadillah" {
		// error
		// t.FailNow() // if use FailNow() Program will be stopped in here  and cannot continue to fmt.Println()
		t.Fatal("Result must be 'Hello Fadillah'")
		// panic("Result is not Hello Fadillah")
	}
	fmt.Println("Function TestHelloWorldFadillah Executed")
}
