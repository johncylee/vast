package vast

import (
	"testing"
)

func testVAST2File(t *testing.T, fn string) {
	var v VAST2
	testXML(t, fn, &v)
}

func TestVAST2(t *testing.T) {
	for _, fn := range []string{
		"v2test/empty.xml",
		"v2test/hivestack.xml",
		"v2test/cf.xml",
		"v2test/Inline_LinearRegular_VAST2.0.xml",
		"v2test/Inline_LinearVAST2vpaid.xml",
		"v2test/Inline_NonLinear_VAST2.0.xml",
		"v2test/Inline_NonLinear_Verification_VAST2.0.xml",
	} {
		t.Log("Testing ", fn)
		testVAST2File(t, fn)
	}
}
