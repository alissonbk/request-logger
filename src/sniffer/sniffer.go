package sniffer

import (
	"fmt"
	"requestlogger/src/request"
	"requestlogger/src/utils"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type Sniffer struct {
	ifName              string
	maxSizeOfEachPacket int32
	promisc             bool
	onlyASCII           bool
}

func NewSniffer() *Sniffer {
	return &Sniffer{ifName: "enp6s0", maxSizeOfEachPacket: 2048, promisc: true}
}

func (sniffer *Sniffer) handlePacket(packet gopacket.Packet) {
	strPacket := utils.BArrayToString(packet.Data())
	if request.HasReadableContent(strPacket) {
		content := request.BuildRequestContent(strPacket)
		//var content = strconv.Quote(strPacket)
		fmt.Println(content)
	} else {
		fmt.Println("-----------------Received a unknown packet-----------------")
	}

}

func StartSniffer(sniffer *Sniffer) {
	fmt.Println("Start sniffing...")
	if handle, err := pcap.OpenLive(sniffer.ifName, sniffer.maxSizeOfEachPacket, sniffer.promisc, pcap.BlockForever); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter("tcp and port 80"); err != nil { // optional
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			sniffer.handlePacket(packet) // Do something with a packet here.
		}
	}
}
