package main

/**
 * 修改数据包中的单个字段值, 或整层字段值的修改
 */

import (
	"common"
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"

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
	fPathName = flag.String("f", "", "需要修改的pcap文件") 
	pcapFile string
	newPcapFile string
	
	pktHandle *pcap.Handle
	options gopacket.SerializeOptions

	srcPort layers.TCPPort
	mac1 = net.HardwareAddr{0x00, 0x50, 0x56, 0xc0, 0x05, 0x01}
	mac2 = net.HardwareAddr{0x00, 0x50, 0x56, 0xc0, 0x04, 0x08}
)

func test1() {
	flag.Parse()
	if *fPathName == "" {
		os.Exit(1)
	}

	pcapFile = *fPathName
	dirpath, filename := filepath.Split(pcapFile)
	filesuff := filepath.Ext(filename)
	filebase := filename[0 : len(filename) - len(filesuff)]
	newPcapFile = dirpath + filebase + "_new" + filesuff

	modifyPKTProcess()
}

func modifyPKTProcess() {
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
		modifyPKT(packet, w, i)
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

func modifyPKT(pkt gopacket.Packet, w *pcapgo.Writer, i int) {
	var eL1 *layers.Ethernet
	var iL2 *layers.IPv4
	var tL3 *layers.TCP
	
	var ethernetL1 []byte
	var ipL2 []byte
	var tcpL3 []byte
	var appL4 []byte
	var outData []byte

	fmt.Printf("packet(%d) \n", i)

	L4 := pkt.ApplicationLayer()
	if L4 != nil {
		appL4 = L4.Payload()
	}

	L3 := pkt.Layer(layers.LayerTypeTCP)
	if L3 != nil {
		tL3, _ = L3.(*layers.TCP)
		tcpL3 = tL3.BaseLayer.Contents
	}

	L2 := pkt.Layer(layers.LayerTypeIPv4)
	if L2 != nil {
		iL2, _ = L2.(*layers.IPv4)
		ipL2 = iL2.BaseLayer.Contents
	}

	L1 := pkt.Layer(layers.LayerTypeEthernet)
	if L1 != nil {
		eL1 = L1.(*layers.Ethernet)
		ethernetL1 = eL1.BaseLayer.Contents
	} else {
		var smac net.HardwareAddr
		var dmac net.HardwareAddr
		
		if srcPort == tL3.SrcPort {
			smac, dmac = mac1, mac2
		} else {
			smac, dmac = mac2, mac1
		}

		ethernetL1 = append(ethernetL1, dmac...)
		ethernetL1 = append(ethernetL1, smac...)
		ethernetL1 = append(ethernetL1, []byte{0x08, 0x00}...)
		//fmt.Printf("ethernetL1(%d) %v \n", len(ethernetL1), ethernetL1)
	}
	
	outData = append(outData, ethernetL1...)
	outData = append(outData, ipL2...)
	outData = append(outData, tcpL3...)
	outData = append(outData, appL4...)
	//fmt.Printf("len(%d) %v \n\n", len(outData), outData)

	ci := gopacket.CaptureInfo{
		Timestamp:     pkt.Metadata().CaptureInfo.Timestamp,
		Length:        len(outData),
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

