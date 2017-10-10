package decoder

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestGetRecord(t *testing.T) {
	dec := NewDecoder(bytes.NewReader(ExampleMessage))

	found, _, record := GetRecord(dec)
	if found != 0 {
		t.Error("Expected to find a record")
	}

	timestamp := time.Now()
	fmt.Printf("[%d] %s: [%s, {", 0, "empty", timestamp.String())
	for k, v := range record {
		fmt.Printf("\"%s\": %v, ", k, v)
	}
	fmt.Printf("}]\n")
}
