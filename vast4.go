package vast

type VAST4 struct {
	VAST
	Ad    []V4Ad     `xml:",omitempty"`
	Error []CDataURI `xml:",omitempty"`
}

type V4Ad struct {
	InLine        *V4InLine  `xml:",omitempty"`
	Wrapper       *V4Wrapper `xml:",omitempty"`
	Id            string     `xml:"id,attr,omitempty"`
	Sequence      int        `xml:"sequence,attr,omitempty"`
	ConditionalAd *bool      `xml:"conditionalAd,attr,omitempty"` // deprecated >=4.1
	AdType        string     `xml:"adType,attr,omitempty"`        // 4.1, 4.2
}

type V4AdDefBase struct {
	AdSystem           V2AdSystem
	Error              []CDataURI     `xml:",omitempty"`
	Extensions         *[]V2Extension `xml:">Extension,omitempty"`
	Impression         []IdURI
	Pricing            *V3Pricing            `xml:",omitempty"`
	ViewableImpression *V4ViewableImpression `xml:",omitempty"`
}

type V4InLine struct {
	V4AdDefBase
	AdServingId     string `xml:",omitempty"` // 4.1, 4.2 required
	AdTitle         string
	AdVerifications *[]V4AdVerification `xml:">Verification,omitempty"`
	Advertiser      string              `xml:",omitempty"`
	Category        []AuthorityElement  `xml:",omitempty"`
	Creatives       []V4Creative        `xml:">Creative"`
	Description     string              `xml:",omitempty"`
	Expires         int                 `xml:",omitempty"` // 4.1, 4.2
	Survey          []V4Survey          `xml:",omitempty"`
}

type V4AdVerification struct {
	ExecutableResource     []V4ExecutableResource `xml:",omitempty"` // 4.1, 4.2
	FlashResource          []V4FlashResource      `xml:",omitempty"` // 4.0
	JavaScriptResource     []V4JavaScriptResource `xml:",omitempty"`
	TrackingEvents         *[]V3Tracking          `xml:">Tracking,omitempty"` // 4.1, 4.2
	VerificationParameters *CDataElement          `xml:",omitempty"`          // 4.1, 4.2
	ViewableImpression     *IdURI                 `xml:",omitempty"`          // 4.0
	Vendor                 string                 `xml:"vendor,attr,omitempty"`
}

type V4ExecutableResource struct {
	Value        AnyURI `xml:",cdata"`
	ApiFramework string `xml:"apiFramework,attr,omitempty"`
	Type         string `xml:"type,attr,omitempty"`
}

type V4FlashResource struct {
	Value        AnyURI `xml:",cdata"`
	ApiFramework string `xml:"apiFramework,attr,omitempty"`
}

type V4JavaScriptResource struct {
	Value           AnyURI `xml:",cdata"`
	ApiFramework    string `xml:"apiFramework,attr,omitempty"`
	BrowserOptional *bool  `xml:"browserOptional,attr,omitempty"` // 4.1, 4.2
}

type V4ViewableImpression struct {
	Viewable         []CDataURI `xml:",omitempty"`
	NotViewable      []CDataURI `xml:",omitempty"`
	ViewUndetermined []CDataURI `xml:",omitempty"`
	Id               string     `xml:"id,attr,omitempty"`
}

type V4CreativeBase struct {
	Sequence     *int   `xml:"sequence,attr,omitempty"`
	ApiFramework string `xml:"apiFramework,attr,omitempty"`
	Id           string `xml:"id,attr,omitempty"`
	AdId         string `xml:"adId,attr,omitempty"`
}

type V4Creative struct {
	V4CreativeBase
	CompanionAds       *V4CompanionAds   `xml:",omitempty"`
	CreativeExtensions *[]V2Extension    `xml:">CreativeExtension,omitempty"`
	Linear             *V4Linear         `xml:",omitempty"`
	NonLinearAds       *V4NonLinearAds   `xml:",omitempty"`
	UniversalAdId      []V4UniversalAdId // 1 in 4.0 and 4.1, 1..* in 4.2
}

type V4CompanionAds struct {
	Companion []V4Companion `xml:",omitempty"`
	Required  string        `xml:"required,attr,omitempty"`
}

type V4CreativeResource struct {
	HTMLResource   []CDataElement     `xml:",omitempty"`
	IFrameResource []CDataURI         `xml:",omitempty"`
	StaticResource []V2StaticResource `xml:",omitempty"`
}

type V4Companion struct {
	V4CreativeResource
	AdParameters           *XmlEncodedStringElement `xml:",omitempty"`
	AltText                string                   `xml:",omitempty"`
	CompanionClickThrough  *CDataURI                `xml:",omitempty"`
	CompanionClickTracking []IdURI                  `xml:",omitempty"`
	CreativeExtensions     *[]V2Extension           `xml:">CreativeExtension,omitempty"`
	TrackingEvents         *[]V3Tracking            `xml:">Tracking,omitempty"`
	Id                     string                   `xml:"id,attr,omitempty"`
	Width                  int                      `xml:"width,attr"`
	Height                 int                      `xml:"height,attr"`
	AssetWidth             int                      `xml:"assetWidth,attr,omitempty"`
	AssetHeight            int                      `xml:"assetHeight,attr,omitempty"`
	ExpandedWidth          int                      `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight         int                      `xml:"expandedHeight,attr,omitempty"`
	ApiFramework           string                   `xml:"apiFramework,attr,omitempty"`
	AdSlotID               string                   `xml:"adSlotID,attr,omitempty"` // 4.0
	AdSlotId               string                   `xml:"adSlotId,attr,omitempty"` // 4.1, 4.2
	PxRatio                float32                  `xml:"pxratio,attr,omitempty"`
	RenderingMode          string                   `xml:"renderingMode,attr,omitempty"` // 4.1, 4.2
}

type V4LinearBase struct {
	Icons          *[]V4Icon     `xml:">Icon,omitempty"`
	TrackingEvents *[]V3Tracking `xml:">Tracking,omitempty"`
	SkipOffset     string        `xml:"skipoffset,attr,omitempty"`
}

type V4Linear struct {
	V4LinearBase
	AdParameters *XmlEncodedStringElement `xml:",omitempty"`
	Duration     XsTime
	MediaFiles   V4MediaFiles
	VideoClicks  *V2VideoClicks `xml:",omitempty"`
}

type V4Icon struct {
	V4CreativeResource
	IconClicks       *V4IconClicks `xml:",omitempty"`
	IconViewTracking []CDataURI    `xml:",omitempty"`
	Program          string        `xml:"program,attr,omitempty"`
	Width            int           `xml:"width,attr,omitempty"`
	Height           int           `xml:"height,attr,omitempty"`
	XPosition        string        `xml:"xPosition,attr,omitempty"`
	YPosition        string        `xml:"yPosition,attr,omitempty"`
	Duration         *XsTime       `xml:"duration,attr,omitempty"`
	Offset           *XsTime       `xml:"offset,attr,omitempty"`
	ApiFramework     string        `xml:"apiFramework,attr,omitempty"`
	PxRatio          float32       `xml:"pxratio,attr,omitempty"`
}

type V4IconClicks struct {
	IconClickFallbackImages *[]V4IconClickFallbackImage `xml:">IconClickFallbackImage,omitempty"` // 4.2
	IconClickThrough        *CDataURI                   `xml:",omitempty"`
	IconClickTracking       []IdURI                     `xml:",omitempty"`
}

type V4IconClickFallbackImage struct {
	AltText        string    `xml:",omitempty"`
	StaticResource *CDataURI `xml:",omitempty"`
	Height         int       `xml:"height,attr,omitempty"`
	Width          int       `xml:"width,attr,omitempty"`
}

type V4MediaFiles struct {
	ClosedCaptionFiles      *[]V4ClosedCaptionFile `xml:">ClosedCaptionFile,omitempty"` // 4.1, 4.2
	MediaFile               []V4MediaFile
	Mezzanine               *[]V4Mezzanine              `xml:",omitempty"`
	InteractiveCreativeFile []V4InteractiveCreativeFile `xml:",omitempty"`
}

type V4ClosedCaptionFile struct {
	Value    AnyURI `xml:",cdata"`
	Type     string `xml:"type,attr,omitempty"`
	Language string `xml:"language,attr,omitempty"`
}

type V4MediaFile struct {
	V3MediaFile
	FileSize  int    `xml:"fileSize,attr,omitempty"`  // 4.1, 4.2
	MediaType string `xml:"mediaType,attr,omitempty"` // 4.1, 4.2
}

// 4.0 -> 4.1, 4.2 incompatible
type V4Mezzanine struct {
	Value     AnyURI `xml:",cdata"`
	Id        string `xml:"id,attr,omitempty"`
	Delivery  string `xml:"delivery,attr,omitempty"` // 4.1, 4.2 required
	Type      string `xml:"type,attr,omitempty"`     // 4.1, 4.2 required
	Width     int    `xml:"width,attr,omitempty"`    // 4.1, 4.2 required
	Height    int    `xml:"height,attr,omitempty"`   // 4.1, 4.2 required
	Codec     string `xml:"codec,attr,omitempty"`
	FileSize  int    `xml:"fileSize,attr,omitempty"`
	MediaType string `xml:"mediaType,attr,omitempty"`
}

type V4InteractiveCreativeFile struct {
	Value            AnyURI `xml:",cdata"`
	Type             string `xml:"type,attr,omitempty"`
	ApiFramework     string `xml:"apiFramework,attr,omitempty"`
	VariableDuration *bool  `xml:"variableDuration,attr,omitempty"` // 4.1, 4.2
}

type V4NonLinearAds struct {
	TrackingEvents *[]V3Tracking `xml:">Tracking,omitempty"`
	NonLinear      []V4NonLinear
}

type V4NonLinearBase struct {
	NonLinearClickTracking []IdURI `xml:",omitempty"`
}

type V4NonLinear struct {
	V4CreativeResource
	AdParameters           *XmlEncodedStringElement `xml:",omitempty"`
	NonLinearClickThrough  *CDataURI                `xml:",omitempty"`
	NonLinearClickTracking []IdURI                  `xml:",omitempty"`
	Id                     string                   `xml:"id,attr,omitempty"`                   // 4.1, 4.2
	Width                  int                      `xml:"width,attr,omitempty"`                // 4.1, 4.2 required
	Height                 int                      `xml:"height,attr,omitempty"`               // 4.1, 4.2 required
	ExpandedWidth          int                      `xml:"expandedWidth,attr,omitempty"`        // 4.1, 4.2
	ExpandedHeight         int                      `xml:"expandedHeight,attr,omitempty"`       // 4.1, 4.2
	Scalable               *bool                    `xml:"scalable,attr,omitempty"`             // 4.1, 4.2
	MaintainAspectRatio    *bool                    `xml:"maintainAspectRatio,attr,omitempty"`  // 4.1, 4.2
	MinSuggestedDuration   *XsTime                  `xml:"minSuggestedDuration,attr,omitempty"` // 4.1, 4.2
	ApiFramework           string                   `xml:"apiFramework,attr,omitempty"`         // 4.1, 4.2
}

type V4UniversalAdId struct {
	Value      string `xml:",chardata"`
	IdRegistry string `xml:"idRegistry,attr"`
	IdValue    string `xml:"idValue,attr,omitempty"` // 4.0 required
}

type V4Survey struct {
	Value AnyURI `xml:",cdata"`
	Type  string `xml:"type,attr,omitempty"`
}

type V4Wrapper struct {
	V4AdDefBase
	AdVerifications          *[]V4AdVerification `xml:">Verification,omitempty"`
	BlockedAdCategories      []AuthorityElement  `xml:",omitempty"` // 4.1, 4.2
	Creatives                []V4CreativeWrapper `xml:">Creative"`
	VASTAdTagURI             CDataURI
	FollowAdditionalWrappers *bool `xml:"followAdditionalWrappers,attr,omitempty"`
	AllowMultipleAds         *bool `xml:"allowMultipleAds,attr,omitempty"`
	FallbackOnNoAd           *bool `xml:"fallbackOnNoAd,attr,omitempty"`
}

type V4CreativeWrapper struct {
	V4CreativeBase
	CompanionAds *V4CompanionAds        `xml:",omitempty"`
	Linear       *V4LinearWrapper       `xml:",omitempty"`
	NonLinearAds *V4NonLinearAdsWrapper `xml:",omitempty"`
}

type V4LinearWrapper struct {
	V4LinearBase
	VideoClicks *V2VideoClicks `xml:",omitempty"`
}

type V4NonLinearAdsWrapper struct {
	TrackingEvents *[]V3Tracking `xml:">Tracking,omitempty"`
	NonLinear      []V4NonLinearBase
}
