package vast

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"
)

func testXML(t *testing.T, fn string, v interface{}) {
	b, err := os.ReadFile(fn)
	if err != nil {
		t.Fatal("os.ReadFile:", err)
	}
	if err = xml.Unmarshal(b, v); err != nil {
		t.Fatal("xml.Unmarshal:", err)
	}
	marshaled, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatal("xml.MarshalIndent:", err)
	}
	expected := fn[:len(fn)-4] + "-marshaled.xml"
	if b, err = os.ReadFile(expected); err != nil {
		t.Fatal("os.ReadFile:", err)
	}
	if bytes.Compare(marshaled, b) != 0 {
		t.Fatal("Unexpected marshel output:", fn)
	}
}

func testVAST2File(t *testing.T, fn string) {
	var v VAST2
	testXML(t, fn, &v)
}

func TestVAST2(t *testing.T) {
	for _, fn := range []string{
		"v2test/empty.xml",
		"v2test/hivestack.xml",
		"v2test/Inline_LinearRegular_VAST2.0.xml",
		"v2test/Inline_LinearVAST2vpaid.xml",
		"v2test/Inline_NonLinear_VAST2.0.xml",
		"v2test/Inline_NonLinear_Verification_VAST2.0.xml",
	} {
		testVAST2File(t, fn)
	}
}
