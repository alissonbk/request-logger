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
	port                string
	maxSizeOfEachPacket int32
	promisc             bool
	onlyASCII           bool
}

func NewSniffer(iface string, port string) *Sniffer {
	return &Sniffer{ifName: iface, port: port, maxSizeOfEachPacket: 2048, promisc: true}
}

func (sniffer *Sniffer) handlePacket(packet gopacket.Packet) {
	strPacket := utils.BArrayToString(packet.Data())
	if request.HasReadableContent(strPacket) {
		content := request.BuildRequestContent(strPacket)
		//var content = strconv.Quote(strPacket)
		if content != "" {
			fmt.Println("------------------------------------------------------------------------------------------------------------------------")
			fmt.Println(content)
		}
	} else {
		// fmt.Println("-----------------Received a unknown packet-----------------")
		// content := request.BuildRequestContent(strPacket)
		// //var content = strconv.Quote(strPacket)
		// fmt.Println(content)
		// fmt.Println("-----------------End of unknown packet-----------------")
	}

}

func StartSniffer(sniffer *Sniffer) {
	fmt.Println("Start sniffing... iface: " + sniffer.ifName + " port: " + sniffer.port)
	if handle, err := pcap.OpenLive(sniffer.ifName, sniffer.maxSizeOfEachPacket, sniffer.promisc, pcap.BlockForever); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter("tcp and port " + sniffer.port); err != nil { // optional
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			sniffer.handlePacket(packet) // Do something with a packet here.
		}
	}
}
