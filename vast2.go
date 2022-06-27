package vast

type VAST2 struct {
	XMLName string `xml:"VAST"`
	Ad      []V2Ad `xml:",omitempty"`
	Version string `xml:"version,attr"`
}

type V2Ad struct {
	InLine  *V2InLine  `xml:",omitempty"`
	Wrapper *V2Wrapper `xml:",omitempty"`
	Id      string     `xml:"id,attr"`
}

type V2InLine struct {
	AdSystem    V2AdSystem
	AdTitle     string
	Description string    `xml:",omitempty"`
	Survey      *CDATAURI `xml:",omitempty"`
	Error       *CDATAURI `xml:",omitempty"`
	Impression  []IdURI
	Creatives   []V2Creative  `xml:">Creative"`
	Extensions  *V2Extensions `xml:",omitempty"`
}

type V2AdSystem struct {
	Version string `xml:"version,attr,omitempty"`
	Value   string `xml:",chardata"`
}

type V2Creative struct {
	Linear       *V2Linear       `xml:",omitempty"`
	CompanionAds *V2Companions   `xml:",omitempty"`
	NonLinearAds *V2NonLinearAds `xml:",omitempty"`
	Id           string          `xml:"id,attr,omitempty"`
	Sequence     *int            `xml:"sequence,attr,omitempty"`
	AdID         string          `xml:",attr,omitempty"`
}

type V2Linear struct {
	Duration       XsTime
	TrackingEvents *V2Trackings   `xml:",omitempty"`
	AdParameters   string         `xml:",omitempty"`
	VideoClicks    *V2VideoClicks `xml:",omitempty"`
	MediaFiles     *struct {
		MediaFile []V2MediaFile
	} `xml:",omitempty"`
}

type V2Trackings struct {
	Tracking []V2Tracking `xml:",omitempty"`
}

type V2Tracking struct {
	Event string `xml:"event,attr"`
	Value AnyURI `xml:",cdata"`
}

type V2VideoClicks struct {
	ClickThrough  *IdURI  `xml:",omitempty"`
	ClickTracking []IdURI `xml:",omitempty"`
	CustomClick   []IdURI `xml:",omitempty"`
}

type V2MediaFile struct {
	Id                  string `xml:"id,attr,omitempty"`
	Delivery            string `xml:"delivery,attr"`
	Type                string `xml:"type,attr"`
	Bitrate             int    `xml:"bitrate,attr,omitempty"`
	Width               int    `xml:"width,attr"`
	Height              int    `xml:"height,attr"`
	Scalable            bool   `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio bool   `xml:"maintainAspectRatio,attr,omitempty"`
	ApiFramework        string `xml:"apiFramework,attr,omitempty"`
	Value               AnyURI `xml:",cdata"`
}

type V2Companions struct {
	Companion []V2Companion `xml:",omitempty"`
}

type V2Companion struct {
	StaticResource        *V2StaticResource `xml:",omitempty"`
	IFrameResource        *CDATAURI         `xml:",omitempty"`
	HTMLResource          *V2HTMLResource   `xml:",omitempty"`
	TrackingEvents        *V2Trackings      `xml:",omitempty"`
	CompanionClickThrough *CDATAURI         `xml:",omitempty"`
	AltText               string            `xml:",omitempty"`
	AdParameters          string            `xml:",omitempty"`
	Id                    string            `xml:"id,attr,omitempty"`
	Width                 int               `xml:"width,attr"`
	Height                int               `xml:"height,attr"`
	ExpandedWidth         int               `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight        int               `xml:"expanededHeight,attr,omitempty"`
	ApiFramework          string            `xml:"apiFramework,attr,omitempty"`
}

type V2StaticResource struct {
	CreativeType string `xml:"creativeType,attr"`
	Value        AnyURI `xml:",cdata"`
}

type V2HTMLResource struct {
	Value string `xml:",cdata"`
}

type V2NonLinearAds struct {
	TrackingEvents *V2Trackings `xml:",omitempty"`
	NonLinear      []V2NonLinear
}

type V2NonLinear struct {
	StaticResource        *V2StaticResource `xml:",omitempty"`
	IFrameResource        *CDATAURI         `xml:",omitempty"`
	HTMLResource          *V2HTMLResource   `xml:",omitempty"`
	NonLinearClickThrough *CDATAURI         `xml:",omitempty"`
	AdParameters          string            `xml:",omitempty"`
	Id                    string            `xml:"id,attr,omitempty"`
	Width                 int               `xml:"width,attr"`
	Height                int               `xml:"height,attr"`
	ExpandedWidth         int               `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight        int               `xml:"expanededHeight,attr,omitempty"`
	Scalable              bool              `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio   bool              `xml:"maintainAspectRatio,attr,omitempty"`
	MinSuggestedDuration  *XsTime           `xml:"minSuggestedDuration,attr,omitempty"`
	ApiFramework          string            `xml:"apiFramework,attr,omitempty"`
}

type V2Extensions struct {
	Extension []V2Extension `xml:",omitempty"`
}

type V2Extension struct {
	Type  string `xml:"type,attr,omitempty"`
	Value []byte `xml:",innerxml"`
}

type V2Wrapper struct {
	AdSystem     V2AdSystem
	VASTAdTagURI CDATAURI
	Error        *CDATAURI `xml:",omitempty"`
	Impression   []CDATAURI
	Creatives    []V2WrappedCreative `xml:">Creative,omitempty"`
	Extensions   *V2Extensions       `xml:",omitempty"`
}

type V2WrappedCreative struct {
	Linear       *V2WrappedLinear       `xml:",omitempty"`
	CompanionAds *V2Companions          `xml:",omitempty"`
	NonLinearAds *V2WrappedNonLinearAds `xml:",omitempty"`
	Id           string                 `xml:"id,attr,omitempty"`
	Sequence     *int                   `xml:"sequence,attr,omitempty"`
	AdID         string                 `xml:",omitempty"`
}

type V2WrappedLinear struct {
	TrackingEvents *V2Trackings   `xml:",omitempty"`
	VideoClicks    *V2VideoClicks `xml:",omitempty"`
	ClickTracking  []IdURI        `xml:",omitempty"`
}

type V2WrappedNonLinearAds struct {
	TrackingEvents *V2Trackings  `xml:",omitempty"`
	NonLinear      []V2NonLinear `xml:",omitempty"`
}
