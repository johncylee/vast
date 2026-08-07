package vast

import (
	"testing"
)

func testVAST4File(t *testing.T, fn string) {
	var v VAST4
	testXML(t, fn, &v)
}

func TestVAST4(t *testing.T) {
	for _, fn := range []string{
		"v4.0_test/Ad_Verification-test.xml",
		"v4.0_test/Category-test.xml",
		"v4.0_test/Conditional_Ad-test.xml",
		"v4.0_test/Event_Tracking-test.xml",
		"v4.0_test/Inline_Companion_Tag-test.xml",
		"v4.0_test/Inline_Linear_Tag-test.xml",
		"v4.0_test/Inline_Non-Linear_Tag-test.xml",
		"v4.0_test/Inline_Simple.xml",
		"v4.0_test/No_Wrapper_Tag-test.xml",
		"v4.0_test/Ready_to_serve_Media_Files_check-test.xml",
		"v4.0_test/SSAI_stitching_VPAID_separation-test.xml",
		"v4.0_test/SSAI_stitching_mezzanine_file_support-test.xml",
		"v4.0_test/Universal_Ad_ID-test.xml",
		"v4.0_test/Video_Clicks_and_click_tracking-Inline-test.xml",
		"v4.0_test/Viewable_Impression-test.xml",
		"v4.0_test/Wrapper_Tag-test.xml",
		"v4.1_test/Ad_Verification-test.xml",
		"v4.1_test/Audio_DAAST_Sample.xml",
		"v4.1_test/Category-test.xml",
		"v4.1_test/Closed_Caption_Test.xml",
		"v4.1_test/Conditional_Ad-test.xml",
		"v4.1_test/Event_Tracking-test.xml",
		"v4.1_test/Inline_Companion_Tag-test.xml",
		"v4.1_test/Inline_Linear_Tag-test.xml",
		"v4.1_test/Inline_Non-Linear_Tag-test.xml",
		"v4.1_test/Inline_Simple.xml",
		"v4.1_test/No_Wrapper_Tag-test.xml",
		"v4.1_test/Ready_to_serve_Media_Files_check-test.xml",
		"v4.1_test/SSAI_stitching_VPAID_separation-test.xml",
		"v4.1_test/SSAI_stitching_mezzanine_file_support-test.xml",
		"v4.1_test/Universal_Ad_ID-test.xml",
		"v4.1_test/Video_Clicks_and_click_tracking-Inline-test.xml",
		"v4.1_test/Viewable_Impression-test.xml",
		"v4.1_test/Wrapper_Tag-test.xml",
		"v4.2_test/Ad_Verification-test.xml",
		"v4.2_test/Category-test.xml",
		"v4.2_test/Closed_Caption_Test.xml",
		"v4.2_test/Event_Tracking-test.xml",
		"v4.2_test/IconClickFallbacks.xml",
		"v4.2_test/Inline_Companion_Tag-test.xml",
		"v4.2_test/Inline_Linear_Tag-test.xml",
		"v4.2_test/Inline_Non-Linear_Tag-test.xml",
		"v4.2_test/Inline_Simple.xml",
		"v4.2_test/No_Wrapper_Tag-test.xml",
		"v4.2_test/Ready_to_serve_Media_Files_check-test.xml",
		"v4.2_test/Universal_Ad_ID-multi-test.xml",
		"v4.2_test/Video_Clicks_and_click_tracking-Inline-test.xml",
		"v4.2_test/Viewable_Impression-test.xml",
		"v4.2_test/Wrapper_Tag-test.xml",
	} {
		testVAST4File(t, fn)
	}
}
