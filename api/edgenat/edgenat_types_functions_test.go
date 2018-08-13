package edgenat

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup() (edgeNatsList *EdgeNats) {
	edgeNatsList = &EdgeNats{}
	firstNat := EdgeNat{
		Description:        "firstNat",
		InterfaceID:        0,
		Type:               "snat",
		Index:              1,
		SourceAddress:      "10.0.0.0/8",
		DestinationAddress: "1.1.1.1/32",
	}
	secondNat := EdgeNat{
		Description:        "firstNat",
		InterfaceID:        0,
		Type:               "dnat",
		Index:              2,
		SourceAddress:      "1.1.1.1/32",
		DestinationAddress: "10.0.1.10/32",
		SourcePort:         "80",
		DestinationPort:    "8080",
	}
	edgeNatsList.Nats = []EdgeNat{firstNat, secondNat}
	return edgeNatsList
}

func TestFilterByName(t *testing.T) {
	edgeNats := setup()

	firstFiltered := edgeNats.FilterByIndex(1)
	assert.Equal(t, 1, firstFiltered.Index)

	secondFiltered := edgeNats.FilterByIndex(2)
	assert.Equal(t, 2, secondFiltered.Index)
}

func TestStringImplementation(t *testing.T) {
	edgeNat := setup()
	expectedOutput := "[{{ } 1 0 snat 10.0.0.0/8 1.1.1.1/32 false firstNat   } {{ } 2 0 dnat 1.1.1.1/32 10.0.1.10/32 false firstNat  80 8080}]"
	assert.Equal(t, expectedOutput, fmt.Sprint(edgeNat))
}
