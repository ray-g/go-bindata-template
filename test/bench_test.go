package main

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	bin "github.com/ray-g/go-bindata-template/test/bindata"
)

var index = "templates/index.tmpl"

func BenchmarkAsset(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bin.Asset(index)
	}
}

func BenchmarkAssetString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bin.AssetString(index)
	}
}

func BenchmarkReadFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ioutil.ReadFile(index)
	}
}

func BenchmarkConcatString(b *testing.B) {
	var str string
	idx, err := bin.AssetString(index)

	if err == nil {
		length := len(idx)
		for n := 0; n < b.N; n++ {
			str += idx
		}
		b.StopTimer()

		if len(str) != b.N*length {
			b.Errorf("error length not same, expect %d, got %d", b.N*length, len(str))
		}
	} else {
		b.Errorf("error loading asset. %v", err)
	}
}

func BenchmarkBufferString(b *testing.B) {
	var buffer bytes.Buffer
	idx, err := bin.AssetString(index)

	if err == nil {
		length := len(idx)
		for n := 0; n < b.N; n++ {
			buffer.WriteString(idx)
		}
		str := buffer.String()
		b.StopTimer()
		if len(str) != b.N*length {
			b.Errorf("error length not same, expect %d, got %d", b.N*length, len(str))
		}
	} else {
		b.Errorf("error loading asset. %v", err)
	}
}

func BenchmarkCopyString(b *testing.B) {
	idx, err := bin.AssetString(index)

	if err == nil {
		length := len(idx)
		bs := make([]byte, b.N*length)
		bl := 0
		for n := 0; n < b.N; n++ {
			bl += copy(bs[bl:], idx)
		}
		str := string(bs)
		b.StopTimer()
		if len(str) != b.N*length {
			b.Errorf("error length not same, expect %d, got %d", b.N*length, len(str))
		}
	} else {
		b.Errorf("error loading asset. %v", err)
	}
}

// Go 1.10
func BenchmarkStringBuilder(b *testing.B) {
	var strBuilder strings.Builder
	idx, err := bin.AssetString(index)

	if err == nil {
		length := len(idx)
		for n := 0; n < b.N; n++ {
			strBuilder.WriteString(idx)
		}
		str := strBuilder.String()
		b.StopTimer()
		if len(str) != b.N*length {
			b.Errorf("error length not same, expect %d, got %d", b.N*length, len(str))
		}
	} else {
		b.Errorf("error loading asset. %v", err)

	}
}

func BenchmarkAppendBytes(b *testing.B) {
	var bytes []byte
	idx, err := bin.Asset(index)

	if err == nil {
		for n := 0; n < b.N; n++ {
			bytes = append(bytes, idx...)
		}
		b.StopTimer()
	} else {
		b.Errorf("error loading asset. %v", err)
	}
}

func BenchmarkBufferBytes(b *testing.B) {
	var buffer bytes.Buffer
	idx, err := bin.Asset(index)

	if err == nil {
		length := len(idx)
		for n := 0; n < b.N; n++ {
			buffer.Write(idx)
		}
		str := buffer.String()
		b.StopTimer()
		if len(str) != b.N*length {
			b.Errorf("error length not same, expect %d, got %d", b.N*length, len(str))
		}
	} else {
		b.Errorf("error loading asset. %v", err)
	}
}

func BenchmarkCopyBytes(b *testing.B) {
	idx, err := bin.Asset(index)

	if err == nil {
		length := len(idx)
		bs := make([]byte, b.N*length)
		bl := 0
		for n := 0; n < b.N; n++ {
			bl += copy(bs[bl:], idx)
		}
		str := string(bs)
		b.StopTimer()
		if len(str) != b.N*length {
			b.Errorf("error length not same, expect %d, got %d", b.N*length, len(str))
		}
	} else {
		b.Errorf("error loading asset. %v", err)
	}
}
