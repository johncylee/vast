package vast

import (
	"net/url"
	"regexp"
	"strings"
	"time"
)

var reV4 = regexp.MustCompile(`^4\.[0-2]$`)

// return: 2, 3, or 4 (4.0 ~ 4.2). -1 means unsupported.
func IsSupported(ver string) (v int) {
	if strings.HasPrefix(ver, "2.") {
		v = 2
	} else if strings.HasPrefix(ver, "3.") {
		v = 3
	} else if reV4.Match([]byte(ver)) {
		v = 4
	} else {
		v = -1
	}
	return
}

type AnyURI url.URL

func (p *AnyURI) MarshalText() ([]byte, error) {
	return ([]byte)((*url.URL)(p).String()), nil
}

func (p *AnyURI) UnmarshalText(text []byte) (err error) {
	s := strings.TrimSpace(string(text))
	if s == "" {
		*p = AnyURI(url.URL{})
		return
	}
	u, err := url.Parse(s)
	if err != nil {
		return
	}
	*p = AnyURI(*u)
	return
}

type CDataURI struct {
	Value AnyURI `xml:",cdata"`
}

type IdURI struct {
	Id    string `xml:"id,attr,omitempty"`
	Value AnyURI `xml:",cdata"`
}

const durationPattern string = "15:04:05"

type XsTime time.Duration

func (d *XsTime) MarshalText() ([]byte, error) {
	t, _ := time.Parse(durationPattern, "00:00:00")
	t = t.Add(*(*time.Duration)(d))
	return ([]byte)(t.Format(durationPattern)), nil
}

func (d *XsTime) UnmarshalText(text []byte) (err error) {
	start, _ := time.Parse(durationPattern, "00:00:00")
	t, err := time.Parse(durationPattern, string(text))
	if err != nil {
		return
	}
	*d = XsTime(t.Sub(start))
	return
}

type CDataElement struct {
	Value string `xml:",cdata"`
}

type AuthorityElement struct {
	Value     string  `xml:",chardata"`
	Authority *AnyURI `xml:"authority,attr,omitempty"`
}

type XmlEncodedStringElement struct {
	Value      string `xml:",chardata"`
	XmlEncoded *bool  `xml:"xmlEncoded,attr,omitempty"`
}
