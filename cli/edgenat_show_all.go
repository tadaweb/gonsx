package main

import (
	"flag"
	"fmt"
	"github.com/tadaweb/gonsx"
	"github.com/tadaweb/gonsx/api/edgenat"
	"net/http"
	"os"
)

func readAllEdgeNat(client *gonsx.NSXClient, flagSet *flag.FlagSet) {

	if edgeid == "" {
		fmt.Println("edgeid must be set - usage: -edgeid <a valid edge id>")
		os.Exit(1)
	}

	readAllEdgeNat := edgenat.NewGetAll(edgeid)
	err := client.Do(readAllEdgeNat)
	if err != nil {
		fmt.Println("error retrieving a list of all edge nat rules" + err.Error())
		os.Exit(1)
	}

	if readAllEdgeNat.StatusCode() == http.StatusOK {
		edgenatsShowAllResponse := readAllEdgeNat.GetResponse()
		rows := []map[string]interface{}{}
		headers := []string{"EdgeNatID", "vNIC", "Type", "Source", "Destination"}
		for _, edgenat := range edgenatsShowAllResponse.Nats {
			fmt.Println(edgenat)
			row := map[string]interface{}{}
			row["EdgeNatID"] = edgenat.Index
			row["vNIC"] = edgenat.InterfaceID
			row["Type"] = edgenat.Type
			row["Source"] = edgenat.SourceAddress + ":" + edgenat.SourcePort
			row["Destination"] = edgenat.DestinationAddress + ":" + edgenat.DestinationPort
			rows = append(rows, row)
		}
		PrettyPrintMany(headers, rows)
	} else {
		fmt.Printf("\nError nat-show-all HTTP return code != 200: %d", readAllEdgeNat.StatusCode())
		os.Exit(2)
	}
}

func init() {
	readAllEdgeNatFlags := flag.NewFlagSet("edgenat-show-all", flag.ExitOnError)
	readAllEdgeNatFlags.StringVar(&edgeid, "edgeid", "", "usage: -edgeid <a valid edge id to attach interfaces to>")
	RegisterCliCommand("edgenat-show-all", readAllEdgeNatFlags, readAllEdgeNat)
}
