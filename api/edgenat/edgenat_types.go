package edgenat

import "encoding/xml"

// EdgeNats top level list object
type EdgeNats struct {
	XMLName xml.Name  `xml:"nat"`
	Nats    []EdgeNat `xml:"natRules>natRule"`
}

// EdgeNat object within EdgeNats list.
type EdgeNat struct {
	XMLName            xml.Name `xml:"natRule"`
	Index              int      `xml:"ruleTag,omitempty"`
	InterfaceID        int      `xml:"vnic"`
	Type               string   `xml:"type"`
	SourceAddress      string   `xml:"originalAddress"`
	DestinationAddress string   `xml:"translatedAddress"`
	Logging            bool     `xml:"loggingEnabled,omitempty"`
	Description        string   `xml:"description,omitempty"`
	Protocol           string   `xml:"protocol,omitempty"`
	SourcePort         int      `xml:"originalPort,omitempty"`
	DestinationPort    int      `xml:"translatedPort,omitempty"`
}
