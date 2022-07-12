package vast

type VAST3 struct {
	XMLName string `xml:"VAST"`
	Ad      []V3Ad `xml:",omitempty"`
	Version string `xml:"version,attr"`
}

type V3Ad struct {
	InLine   *V3InLine  `xml:",omitempty"`
	Wrapper  *V3Wrapper `xml:",omitempty"`
	Id       string     `xml:"id,attr,omitempty"`
	Sequence *int       `xml:"sequence,attr,omitempty"`
}

// The "pointer to array" is a workaround. Without it, unmarshaling may creates an empty
// array that leads to "<Extensions></Extensions>" after marshaling.
type V3InLine struct {
	AdSystem    V3AdSystem
	AdTitle     string
	Description string     `xml:",omitempty"`
	Advertiser  string     `xml:",omitempty"`
	Pricing     *V3Pricing `xml:",omitempty"`
	Survey      *CDATAURI  `xml:",omitempty"`
	Error       *CDATAURI  `xml:",omitempty"`
	Impression  []IdURI
	Creatives   []V3Creative   `xml:">Creative"`
	Extensions  *[]V2Extension `xml:">Extension,omitempty"`
}

type V3AdSystem struct {
	Version string `xml:"version,attr,omitempty"`
	Value   string `xml:",chardata"`
}

type V3Pricing struct {
	Model    string  `xml:"model,attr"`
	Currency string  `xml:"currency,attr"`
	Value    float32 `xml:",cdata"`
}

type V3Creative struct {
	Linear       *V3Linear       `xml:",omitempty"`
	CompanionAds *V3CompanionAds `xml:",omitempty"`
	NonLinearAds *V3NonLinearAds `xml:",omitempty"`
	Id           string          `xml:"id,attr,omitempty"`
	Sequence     *int            `xml:"sequence,attr,omitempty"`
	AdID         string          `xml:",omitempty"`
}

type V3Linear struct {
	Icons              *[]V3Icon      `xml:">Icon,omitempty"`
	CreativeExtensions *[]V2Extension `xml:">CreativeExtension,omitempty"`
	Duration           XsTime
	TrackingEvents     *[]V3Tracking   `xml:">Tracking,omitempty"`
	AdParameters       *V3AdParameters `xml:",omitempty"`
	VideoClicks        *V2VideoClicks  `xml:",omitempty"`
	MediaFiles         *[]V3MediaFile  `xml:">MediaFile,omitempty"`
	SkipOffset         string          `xml:"skipoffset,attr,omitempty"`
}

type V3Icon struct {
	StaticResource   *V2StaticResource `xml:",omitempty"`
	IFrameResource   *CDATAURI         `xml:",omitempty"`
	HTMLResource     *V3HTMLResource   `xml:",omitempty"`
	IconClicks       *V3IconClicks     `xml:",omitempty"`
	IconViewTracking []CDATAURI        `xml:",omitempty"`
	Program          string            `xml:"program,attr"`
	Width            int               `xml:"width,attr"`
	Height           int               `xml:"height,attr"`
	XPosition        string            `xml:"xPosition,attr"`
	YPosition        string            `xml:"yPosition,attr"`
	Offset           *XsTime           `xml:"offset,attr,omitempty"`
	Duration         *XsTime           `xml:"duration,attr,omitempty"`
	ApiFramework     string            `xml:"apiFramework,attr,omitempty"`
}

type V3HTMLResource struct {
	XmlEncoded bool   `xml:"xmlEncoded,attr,omitempty"`
	Value      string `xml:",cdata"`
}

type V3IconClicks struct {
	IconClickTracking []CDATAURI `xml:",omitempty"`
	IconClickThrough  *CDATAURI  `xml:",omitempty"`
}

type V3Tracking struct {
	Event  string `xml:"event,attr"`
	Offset string `xml:"offset,attr,omitempty"`
	Value  AnyURI `xml:",cdata"`
}

type V3AdParameters struct {
	XmlEncoded bool   `xml:"xmlEncoded,attr,omitempty"`
	Value      string `xml:",chardata"`
}

type V3MediaFile struct {
	Id                  string `xml:"id,attr,omitempty"`
	Delivery            string `xml:"delivery,attr"`
	Type                string `xml:"type,attr"`
	Bitrate             int    `xml:"bitrate,attr,omitempty"`
	MinBitrate          int    `xml:"minBitrate,attr,omitempty"`
	MaxBitrate          int    `xml:"maxBitrate,attr,omitempty"`
	Width               int    `xml:"width,attr"`
	Height              int    `xml:"height,attr"`
	Scalable            bool   `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio bool   `xml:"maintainAspectRatio,attr,omitempty"`
	ApiFramework        string `xml:"apiFramework,attr,omitempty"`
	Codec               string `xml:"codec,attr,omitempty"`
	Value               AnyURI `xml:",cdata"`
}

type V3CompanionAds struct {
	Companion []V3Companion `xml:",omitempty"`
	Required  string        `xml:"required,attr,omitempty"`
}

type V3Companion struct {
	StaticResource        *V2StaticResource `xml:",omitempty"`
	IFrameResource        *CDATAURI         `xml:",omitempty"`
	HTMLResource          *V3HTMLResource   `xml:",omitempty"`
	CreativeExtensions    *[]V2Extension    `xml:">CreativeExtension,omitempty"`
	TrackingEvents        *[]V3Tracking     `xml:">Tracking,omitempty"`
	CompanionClickThrough *CDATAURI         `xml:",omitempty"`
	AltText               string            `xml:",omitempty"`
	AdParameters          *V3AdParameters   `xml:",omitempty"`
	Id                    string            `xml:"id,attr,omitempty"`
	Width                 int               `xml:"width,attr"`
	Height                int               `xml:"height,attr"`
	AssetWidth            int               `xml:"assetWidth,attr,omitempty"`
	AssetHeight           int               `xml:"assetHeight,attr,omitempty"`
	ExpandedWidth         int               `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight        int               `xml:"expandedHeight,attr,omitempty"`
	ApiFramework          string            `xml:"apiFramework,attr,omitempty"`
	AdSlotId              string            `xml:"adSlotId,attr,omitempty"`
}

type V3NonLinearAds struct {
	TrackingEvents *[]V3Tracking `xml:">Tracking,omitempty"`
	NonLinear      []V3NonLinear
}

type V3NonLinear struct {
	StaticResource         *V2StaticResource `xml:",omitempty"`
	IFrameResource         *CDATAURI         `xml:",omitempty"`
	HTMLResource           *V3HTMLResource   `xml:",omitempty"`
	CreativeExtensions     *[]V2Extension    `xml:">CreativeExtension,omitempty"`
	NonLinearClickTracking []CDATAURI        `xml:",omitempty"`
	NonLinearClickThrough  *CDATAURI         `xml:",omitempty"`
	AdParameters           *V3AdParameters   `xml:",omitempty"`
	Id                     string            `xml:"id,attr,omitempty"`
	Width                  int               `xml:"width,attr"`
	Height                 int               `xml:"height,attr"`
	ExpandedWidth          int               `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight         int               `xml:"expandedHeight,attr,omitempty"`
	Scalable               bool              `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio    bool              `xml:"maintainAspectRatio,attr,omitempty"`
	MinSuggestedDuration   *XsTime           `xml:"minSuggestedDuration,attr,omitempty"`
	ApiFramework           string            `xml:"apiFramework,attr,omitempty"`
}

type V3Wrapper struct {
	AdSystem     V3AdSystem
	VASTAdTagURI CDATAURI
	Error        *CDATAURI `xml:",omitempty"`
	Impression   []CDATAURI
	Creatives    *[]V3WrappedCreative `xml:">Creative,omitempty"`
	Extensions   *[]V2Extension       `xml:">Extension,omitempty"`
}

type V3WrappedCreative struct {
	Linear       *V3WrappedLinear       `xml:",omitempty"`
	CompanionAds *[]V3CompanionWrapper  `xml:">Companion,omitempty"`
	NonLinearAds *V3WrappedNonLinearAds `xml:",omitempty"`
	Id           string                 `xml:"id,attr,omitempty"`
	Sequence     *int                   `xml:"sequence,attr,omitempty"`
	AdID         string                 `xml:",omitempty"`
}

type V3WrappedLinear struct {
	CreativeExtensions *[]V2Extension        `xml:">CreativeExtension,omitempty"`
	Icons              *[]V3Icon             `xml:">Icon,omitempty"`
	TrackingEvents     *[]V3Tracking         `xml:">Tracking,omitempty"`
	VideoClicks        *V3WrappedVideoClicks `xml:",omitempty"`
}

type V3WrappedVideoClicks struct {
	ClickTracking []IdURI `xml:",omitempty"`
	CustomClick   []IdURI `xml:",omitempty"`
}

type V3CompanionWrapper struct {
	StaticResource         *V2StaticResource `xml:",omitempty"`
	IFrameResource         *CDATAURI         `xml:",omitempty"`
	HTMLResource           *V3HTMLResource   `xml:",omitempty"`
	CreativeExtensions     *[]V2Extension    `xml:">CreativeExtension,omitempty"`
	TrackingEvents         *[]V3Tracking     `xml:">Tracking,omitempty"`
	CompanionClickThrough  *CDATAURI         `xml:",omitempty"`
	CompanionClickTracking []CDATAURI        `xml:",omitempty"`
	AltText                string            `xml:",omitempty"`
	AdParameters           *V3AdParameters   `xml:",omitempty"`
	Id                     string            `xml:"id,attr,omitempty"`
	Width                  int               `xml:"width,attr"`
	Height                 int               `xml:"height,attr"`
	AssetWidth             int               `xml:"assetWidth,attr,omitempty"`
	AssetHeight            int               `xml:"assetHeight,attr,omitempty"`
	ExpandedWidth          int               `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight         int               `xml:"expandedHeight,attr,omitempty"`
	ApiFramework           string            `xml:"apiFramework,attr,omitempty"`
	AdSlotId               string            `xml:"adSlotId,attr,omitempty"`
}

type V3WrappedNonLinearAds struct {
	TrackingEvents *[]V3Tracking        `xml:">Tracking,omitempty"`
	NonLinear      []V3NonLinearWrapper `xml:",omitempty"`
}

type V3NonLinearWrapper struct {
	CreativeExtensions     *[]V2Extension `xml:">CreativeExtension,omitempty"`
	NonLinearClickTracking []CDATAURI     `xml:",omitempty"`
	Id                     string         `xml:"id,attr,omitempty"`
	Width                  int            `xml:"width,attr"`
	Height                 int            `xml:"height,attr"`
	ExpandedWidth          int            `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight         int            `xml:"expandedHeight,attr,omitempty"`
	Scalable               bool           `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio    bool           `xml:"maintainAspectRatio,attr,omitempty"`
	MinSuggestedDuration   *XsTime        `xml:"minSuggestedDuration,attr,omitempty"`
	ApiFramework           string         `xml:"apiFramework,attr,omitempty"`
}
