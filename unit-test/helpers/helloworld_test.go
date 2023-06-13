package helpers

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Dicki")
	if result != "Hello Dicki" {
		// error
		// jangan pake panic
		panic("result is not hello dicki")
	}
}

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("Dicki")
	if result != "Hello Dicki" {
		// error
		t.Fail()
	}

	fmt.Println("test helloworldfail done")
}

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("Dicki")
	if result != "Hello Dicki" {
		// error
		t.FailNow()
	}

	fmt.Println("test helloworldfailnow done")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Dicki")
	if result != "Hello Dicki" {
		// error
		t.Error("result mus be 'hello dicki'")
	}

	fmt.Println("test helloworldfailnow done")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Dicki")
	if result != "Hello Dicki" {
		// error
		t.Fatal("result mus be 'hello dicki'")
	}

	fmt.Println("test helloworldfailnow done")
}

// cek di dokumentasi testify
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Dicki")
	assert.Equal(t, "Hello Dicki", result, "result must be 'Hello Dicki'")

	fmt.Println("test TestHelloWorldAssertion done")
}

// cek di dokumentasi testify
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Dicki")
	require.Equal(t, "Hello Dicki", result, "result must be 'Hello Dicki'")

	fmt.Println("test TestHelloWorldRequire done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("can not run on windows os")
	}

	result := HelloWorld("Dicki")
	require.Equal(t, "Hello Dicki", result, "result must be 'Hello Dicki'")
}

// tidak peduli testnya gagal before & after tetap diesksekusi dalam satu package
func TestMain(m *testing.M) {
	// before
	fmt.Println("before unit test")

	m.Run()

	// after
	fmt.Println("after unit test")
}

func TestSubTest(t *testing.T) {
	t.Run("Dicki", func(t *testing.T) {
		result := HelloWorld("Dicki")
		require.Equal(t, "Hello Dicki", result, "result must be 'Hello Dicki'")
	})
	t.Run("Darmawan", func(t *testing.T) {
		result := HelloWorld("Darmawan")
		require.Equal(t, "Hello Darmawan", result, "result must be 'Hello Darmawan'")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Dicki",
			request:  "Dicki",
			expected: "Hello Dicki",
		},
		{
			name:     "Darmawan",
			request:  "Darmawan",
			expected: "Hello Darmawan",
		},
		{
			name:     "Saputra",
			request:  "Saputra",
			expected: "Hello Saputra",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Dicki")
	}
}

// mirip sub unit test
func BenchmarkSubBenchmark(b *testing.B) {
	b.Run("Dicki", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Dicki")
		}
	})
	b.Run("Darmawan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Darmawan")
		}
	})
	b.Run("Saputra", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Saputra")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:    "Dicki",
			request: "Dicki",
		},
		{
			name:    "Darmawan",
			request: "Darmawan",
		},
		{
			name:    "Saputra",
			request: "Saputra",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
