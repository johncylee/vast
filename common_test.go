package vast

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"
)

func testXML(t *testing.T, fn string, v any) {
	b, err := os.ReadFile(fn)
	if err != nil {
		t.Fatalf("%s os.ReadFile: %s", fn, err)
	}
	if err = xml.Unmarshal(b, v); err != nil {
		t.Fatalf("%s xml.Unmarshal: %s", fn, err)
	}
	marshaled, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatal("xml.MarshalIndent:", err)
	}
	expected := fn[:len(fn)-4] + "-marshaled.xml"
	if b, err = os.ReadFile(expected); err != nil {
		t.Fatalf("%s os.ReadFile: %s", expected, err)
	}
	if !bytes.Equal(marshaled, b) {
		f, err := os.CreateTemp("", "vast-*.xml")
		if err == nil {
			if _, err = f.Write(marshaled); err == nil {
				t.Logf("diff %s %s", f.Name(), expected)
			}
		}
		t.Fatal("Unexpected marshel output:", expected)
	}
}

func TestEmpty(t *testing.T) {
	var v2 VAST2
	v2.Version = "2.0"
	marshaled, err := xml.MarshalIndent(v2, "", "  ")
	if err != nil {
		t.Fatal("xml.MarshalIndent:", err)
	}
	fn := "v2test/empty-marshaled.xml"
	b, err := os.ReadFile(fn)
	if err != nil {
		t.Fatal("os.ReadFile:", err)
	}
	if !bytes.Equal(marshaled, b) {
		t.Fatal("Unexpected marshel output:", string(marshaled))
	}
	var v3 VAST3
	v3.Version = "3.0"
	if marshaled, err = xml.MarshalIndent(v3, "", "  "); err != nil {
		t.Fatal("xml.MarshalIndent:", err)
	}
	fn = "v3test/empty-marshaled.xml"
	if b, err = os.ReadFile(fn); err != nil {
		t.Fatal("os.ReadFile:", err)
	}
	if !bytes.Equal(marshaled, b) {
		t.Fatal("Unexpected marshel output:", string(marshaled))
	}
}

func TestIsSupported(t *testing.T) {
	for in, expect := range map[string]int{
		"2.0": 2,
		"3.0": 3,
		"4.0": 4,
		"4.1": 4,
		"4.2": 4,
		"4.4": -1,
	} {
		if out := IsSupported(in); out != expect {
			t.Errorf("IsSupported(%s) = %d, expect %d", in, out, expect)
		}
	}

}
