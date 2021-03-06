package vast

import (
	"testing"
)

func testVAST3File(t *testing.T, fn string) {
	var v VAST3
	testXML(t, fn, &v)
}

func TestVAST3(t *testing.T) {
	for _, fn := range []string{
		"v3test/Event_Tracking-test.xml",
		"v3test/Inline_Companion_Tag-test.xml",
		"v3test/Inline_Linear_Tag-test.xml",
		"v3test/Inline_Non-Linear_Tag-test.xml",
		"v3test/No_Wrapper_Tag-test.xml",
		"v3test/tenmax-vast.xml",
		"v3test/Video_Clicks_and_click_tracking-Inline-test.xml",
		"v3test/Wrapper_Tag-test.xml",
	} {
		testVAST3File(t, fn)
	}
}
