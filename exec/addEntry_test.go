package exec

import (
	"testing"
	"time"
)

func TestCreateEntryNLP(t *testing.T) {
	msg, err := CreateEntry("Test", "10 minutes ago", "")
	if err != nil {
		t.Fatalf("CreateEntry() returned error: %v", err)
	}

	if msg.Name != "Test" {
		t.Fatal("Wrong name")
	}
}

func TestCreateEntry(t *testing.T) {
	msg, err := CreateEntry("Test", "2021-05-24T16:25:56.558Z", "")
	if err != nil {
		t.Fatalf("CreateEntry() returned error: %v", err)
	}

	tm, _ := time.Parse(time.RFC3339, msg.Start)

	d := time.Date(2021, 5, 24, 16, 25, 56, 558, time.UTC).Local().Format(time.RFC3339)

	if d != tm.Local().Format(time.RFC3339) {
		t.Fatalf("Times are not equal: %s", tm.String())
	}
}
