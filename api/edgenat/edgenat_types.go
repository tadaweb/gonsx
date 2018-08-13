package edgenat

import "encoding/xml"

// EdgeNats top level list object
type EdgeNats struct {
	XMLName xml.Name  `xml:"nat"`
	Version string    `xml:"version"`
	Enabled bool      `xml:"enabled"`
	Nats    []EdgeNat `xml:"natRules>natRule"`
}

// EdgeNat object within EdgeNats list.
type EdgeNat struct {
	XMLName            xml.Name `xml:"natRule"`
	Index              int      `xml:"ruleTag,omitempty"`
	InterfaceID        int      `xml:"vnic"`
	Type               string   `xml:"action"`
	SourceAddress      string   `xml:"originalAddress"`
	DestinationAddress string   `xml:"translatedAddress"`
	Logging            bool     `xml:"loggingEnabled,omitempty"`
	Description        string   `xml:"description,omitempty"`
	Protocol           string   `xml:"protocol,omitempty"`
	SourcePort         string   `xml:"originalPort,omitempty"`
	DestinationPort    string   `xml:"translatedPort,omitempty"`
}
