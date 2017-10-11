package decoder

import (
	"bytes"
	"io"
	"testing"
)

func TestGetRecord(t *testing.T) {
	dec := NewDecoder(bytes.NewReader(ExampleMessage))

	count := 0
	for {
		record, err := dec.GetRecord()
		if err != nil && err != io.EOF {
			t.Fatal(err)
		}
		if err == io.EOF {
			break // No More Records
		}

		if record == nil {
			t.Fatal("Expected Record")
		}
		count++

		_, err = ExtractTime(record)
		if err != nil {
			t.Fatal(err)
		}

		_, err = ExtractData(record)
		if err != nil {
			t.Fatal(err)
		}
	}

	if count <= 0 {
		t.Fatalf("processed %d records; wanted at least 1", count)
	}
}

func BenchmarkGetRecord(b *testing.B) {
	b.SetBytes(int64(len(ExampleMessage)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec := NewDecoder(bytes.NewReader(ExampleMessage))
		count := 0
		for {
			record, err := dec.GetRecord()
			if err != nil && err != io.EOF {
				b.Fatal(err)
			}
			if err == io.EOF {
				break // No More Records
			}

			if record == nil {
				b.Fatal("Expected Record")
			}
			count++

			_, err = ExtractTime(record)
			if err != nil {
				b.Fatal(err)
			}

			_, err = ExtractData(record)
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}
