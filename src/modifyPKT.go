package main

import (
	"common"
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
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
	options gopacket.SerializeOptions
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
		//fmt.Println(packet)
		//showPacket(packet)
		modifyPKT(packet)
	}
}

func showPacket(pkt gopacket.Packet) {
	L1 := pkt.Layer(layers.LayerTypeEthernet)
	if L1 != nil {
		ethernetL1, _ := L1.(*layers.Ethernet)
		fmt.Printf("%v \n\n", ethernetL1)
	}

	L2 := pkt.Layer(layers.LayerTypeIPv4)
	if L2 != nil {
		ipL2, _ := L2.(*layers.IPv4)
		fmt.Printf("%v \n\n", ipL2)
	}

	L3 := pkt.Layer(layers.LayerTypeTCP)
	if L3 != nil {
		tcpL3, _ := L3.(*layers.TCP)
		fmt.Printf("%v \n\n", tcpL3)
	}

	L4 := pkt.ApplicationLayer()
	if L4 != nil {
		fmt.Printf("%v \n\n", L4.Payload())
	} else {
		fmt.Printf("L4 == NULL \n\n")
	}
}

func modifyPKT(pkt gopacket.Packet) {
	var ethernetL1 *layers.Ethernet
	var ipL2 *layers.IPv4
	var tcpL3 *layers.TCP

	L1 := pkt.Layer(layers.LayerTypeEthernet)
	if L1 != nil {
		ethernetL1, _ = L1.(*layers.Ethernet)
	} else {
		ethernetL1 = nil
	}

	L2 := pkt.Layer(layers.LayerTypeIPv4)
	if L2 != nil {
		ipL2, _ = L2.(*layers.IPv4)
	} else {
		ipL2 = nil
	}

	L3 := pkt.Layer(layers.LayerTypeTCP)
	if L3 != nil {
		tcpL3, _ = L3.(*layers.TCP)
	} else {
		tcpL3 = nil
	}

	newPkt := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(
		newPkt, options,
		ethernetL1,
		ipL2,
		tcpL3,
		nil,
	)

}



