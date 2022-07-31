package gateway_test

import (
	"bytes"
	"encoding/json"
	"log"
	"mkstacks/adapter/gateway"
	"os"
	"testing"
)

type DummyStruct struct {
	Dummy string
}

func TestLocalFS(t *testing.T) {
	testText1 := "Hello LocalFS test"
	testFile1 := "testdata.txt"
	testStruct := DummyStruct{
		Dummy: "dummy_text",
	}
	s := gateway.LocalFS{}
	err := os.WriteFile(testFile1, bytes.NewBufferString(testText1).Bytes(), 0755)
	b, err := os.ReadFile(testFile1)
	log.Printf("%s", string(b))
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Read", func(t *testing.T) {
		t.Cleanup(func() {
			if err := os.Remove(testFile1); err != nil {
				t.Fatal(err)
			}
		})
		buf, err := s.Read(testFile1)
		if err != nil {
			t.Failed()
		}
		str := string(buf)
		if str != testText1 {
			t.Fatalf("expected is %s, but actual %s", testText1, str)
		}
	})
	t.Run("Write", func(t *testing.T) {
		t.Cleanup(func() {
			if err := os.Remove(testFile1); err != nil {
				t.Fatal(err)
			}
		})
		expected, err := json.Marshal(testStruct)
		if err != nil {
			t.Fatal(err)
		}

		if err := s.Write(testStruct, testFile1); err != nil {
			t.Fatal(err)
		}

		actual, err := s.Read(testFile1)
		if bytes.Compare(actual, expected) != 0 {
			t.Fatalf("expected is %s, but actual %s", string(expected), string(actual))
		}
	})
}
