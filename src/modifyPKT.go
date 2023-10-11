package main

/**
 * 修改数据包中的单个字段值, 或整层字段值的修改
 */

import (
	"common"
	"fmt"
	"net"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
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
	newPcapFile = "C:/Users/17444/Downloads/test1.pcap"
	pktHandle *pcap.Handle
	options gopacket.SerializeOptions

	srcPort layers.TCPPort
	mac1 = net.HardwareAddr{0x00, 0x50, 0x56, 0xc0, 0x05, 0x01}
	mac2 = net.HardwareAddr{0x00, 0x50, 0x56, 0xc0, 0x04, 0x08}
)

func test1() {
	pktHandle, err := pcap.OpenOffline(pcapFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer pktHandle.Close()

	fp, err := os.Create(newPcapFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer fp.Close()

	w := pcapgo.NewWriter(fp)
	w.WriteFileHeader(65535, layers.LinkTypeEthernet)

	pktSource := gopacket.NewPacketSource(pktHandle, pktHandle.LinkType())
	i := 0

	for packet := range pktSource.Packets() {
		//fmt.Println(packet)
		//showPacket(packet)
		if i == 0 {
			setC2S(packet)
		}
		modifyPKT(packet, w)
		i++
	}
}

func setC2S(pkt gopacket.Packet) {
	var tcpL3 *layers.TCP
	
	L3 := pkt.Layer(layers.LayerTypeTCP)
	if L3 != nil {
		tcpL3, _ = L3.(*layers.TCP)
		srcPort = tcpL3.SrcPort
	}
}

func modifyPKT(pkt gopacket.Packet, w *pcapgo.Writer) {
	var ethernetL1 *layers.Ethernet
	var ipL2 *layers.IPv4
	var tcpL3 *layers.TCP
	var appL4 []byte

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

	L1 := pkt.Layer(layers.LayerTypeEthernet)
	if L1 != nil {
		ethernetL1, _ = L1.(*layers.Ethernet)
	} else {
		var smac net.HardwareAddr
		var dmac net.HardwareAddr
		
		if srcPort == tcpL3.SrcPort {
			smac, dmac = mac1, mac2
		} else {
			smac, dmac = mac2, mac1
		}

		ethernetL1 = &layers.Ethernet{
			EthernetType: 0x0800,
			SrcMAC:       smac,
			DstMAC:       dmac,
		}
	}

	newPkt := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(
		newPkt, options,
		ethernetL1,
		ipL2,
		tcpL3,
		gopacket.Payload(appL4),
	)

	outData := newPkt.Bytes()
	//fmt.Printf("len(%d) %v \n\n", len(outData), outData)

	ci := gopacket.CaptureInfo{
		Timestamp:     pkt.Metadata().CaptureInfo.Timestamp,
		Length:        65535,
		CaptureLength: len(outData),
	}

	err := w.WritePacket(ci, outData)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
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

