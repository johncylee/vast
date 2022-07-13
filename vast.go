package vast

// A minimum skeleton to parse VAST version in order to decide real struct to use
type VAST struct {
	XMLName string `xml:"VAST"`
	Version string `xml:"version,attr"`
}
