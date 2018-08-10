package edgenat

import "fmt"

func (s EdgeNats) String() string {
	return fmt.Sprintf("%v", s.Nats)
}

//func (s EdgeNat) String() string {
//return fmt.Sprintf("index: %d, vnic: %v, type: %s, src: %s, dst: %s", s.Index, s.InterfaceID, s.Type, s.SourceAddress, s.DestinationAddress)
//}

// FilterByIndex filters EdgeNats list by given name of nat and returns
// the found nat object.
func (s EdgeNats) FilterByIndex(index int) *EdgeNat {
	var edgeNatFound EdgeNat
	for _, edgeNat := range s.Nats {
		if edgeNat.Index == index {
			edgeNatFound = edgeNat
			break
		}
	}
	return &edgeNatFound
}
