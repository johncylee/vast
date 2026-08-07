package vast

import "encoding/xml"

// A minimum skeleton to parse VAST version in order to decide real struct to use
type VAST struct {
	XMLName string `xml:"VAST"`
	XmlnsXs Xs     `xml:"xs,attr,omitempty"`
	Xmlns   string `xml:"xmlns,attr,omitempty"`
	Version string `xml:"version,attr"`
}

type Xs string

func (t *Xs) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  xml.Name{Local: "xmlns:xs"},
		Value: string(*t),
	}, nil
}

func (t *Xs) UnmarshalXMLAttr(attr xml.Attr) error {
	*t = Xs(attr.Value)
	return nil
}
