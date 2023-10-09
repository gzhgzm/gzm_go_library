package main

import (
	"os"
	"common"
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	var cts common.TimeStat
	fmt.Printf("----------- begin ------------\n")
	cts.TimeStatInit()

	test1()

	cts.TimeStatShow()
	fmt.Printf("------------ end -------------\n")
	common.PressKeyExit()
}

var (
	pcapFile = "C:/Users/17444/Downloads/test.pcap"
	pktHandle *pcap.Handle
)

func test1() {
	pktHandle, err := pcap.OpenOffline(pcapFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer pktHandle.Close()

	pktSource := gopacket.NewPacketSource(pktHandle, pktHandle.LinkType())

	for packet := range pktSource.Packets() {
		fmt.Println(packet)
	}
}


