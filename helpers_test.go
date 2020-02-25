package goqradar

import (
	"testing"
)

func TestParseContentRange(t *testing.T) {
	cr1 := "0-10/40"
	_, _, _, err := parseContentRange(cr1)
	if err != nil {
		t.Fatalf("should not error but error is: %s", err)
	}

	cr2 := "0_10/40"
	_, _, _, err = parseContentRange(cr2)
	if err == nil {
		t.Fatal("should error")
	}

	cr3 := "0-10_40"
	_, _, _, err = parseContentRange(cr3)
	if err == nil {
		t.Fatal("should error")
	}

	cr4 := "0-0/1"
	_, _, _, err = parseContentRange(cr4)
	if err != nil {
		t.Fatalf("should not error but error is: %s", err)
	}

	cr5 := "*/3"
	_, _, _, err = parseContentRange(cr5)
	if err != nil {
		t.Fatalf("should not error but error is: %s", err)
	}

}
