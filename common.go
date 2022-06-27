package vast

import (
	"net/url"
	"strings"
	"time"
)

type AnyURI url.URL

func (p *AnyURI) MarshalText() ([]byte, error) {
	return ([]byte)((*url.URL)(p).String()), nil
}

func (p *AnyURI) UnmarshalText(text []byte) (err error) {
	u, err := url.ParseRequestURI(strings.TrimSpace(string(text)))
	if err != nil {
		return
	}
	*p = AnyURI(*u)
	return
}

type CDATAURI struct {
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
