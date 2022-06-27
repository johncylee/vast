package vast

import (
	"encoding/xml"
	"os"
	"testing"
)

func testVAST2File(t *testing.T, fn string) {
	var v VAST2
	b, err := os.ReadFile(fn)
	if err != nil {
		t.Fatal(err)
	}
	err = xml.Unmarshal(b, &v)
	if err != nil {
		t.Fatal(err)
	}
	b, err = xml.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
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
