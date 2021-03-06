package main

import (
	"fmt"
)

func printUnknown(uh []byte) {
	// print unknown header and data
}

func printEther(eh *EtherHeader) {
	// print ether header
	fmt.Printf("-------------Ether-------------\n")
	fmt.Println("Dst: ", eh.DstMacAddr)
	fmt.Println("Src: ", eh.SrcMacAddr)
	t := uint16ToByte(eh.ProtoType)
	switch eh.ProtoType {
	case EthTypeArp:
		fmt.Printf("Type: ARP(%X)\n", t)

	case EthTypeIpv4:
		fmt.Printf("Type: IP(%X)\n", t)

	case EthTypeIpv6:
		fmt.Printf("Type: IPv6(%X)\n", t)

	default:
		fmt.Printf("Type: Unknown(%X)\n", t)
	}
	fmt.Printf("-------------------------------\n")
}

func printArp(ah *ArpHeader, pd []byte) {
	// print arp header and padding data
	t := uint16ToByte(ah.ProtoType)
	fmt.Printf("--------------Arp--------------\n")
	fmt.Printf("HardwareType: %X\n", ah.HardwareType)
	fmt.Printf("ProtoType: IP(%X)\n", t)
	fmt.Printf("MacAddrLen: %X\n", ah.MacAddrLen)
	fmt.Printf("IpAddrLen: %X\n", ah.IpAddrLen)
	switch ah.OperationCode {
	case OpCodeRequest:
		fmt.Printf("OprationCode: %X(request)\n", ah.OperationCode)

	case OpCodeReply:
		fmt.Printf("OprationCode: %X(reply)\n", ah.OperationCode)

	case OpCodeReqRev:
		fmt.Printf("OprationCode: %X(request reverse)\n", ah.OperationCode)

	case OpCodeRepRev:
		fmt.Printf("OprationCode: %X(reply reverse)\n", ah.OperationCode)

	default:
		fmt.Printf("OprationCode: %X(Unknown)\n", ah.OperationCode)
	}
	fmt.Println("SenderMacAddr: ", ah.SenderMacAddr)
	fmt.Println("SenderIpAddr: ", ah.SenderIpAddr)
	fmt.Println("TargetMacAddr: ", ah.TargetMacAddr)
	fmt.Println("TargetIpAddr: ", ah.TargetIpAddr)
	fmt.Printf("PaddingData: %X\n", pd)
	fmt.Printf("-------------------------------\n")
}

func printIpv4(ih *IpHeader) {
	fmt.Printf("--------------IP---------------\n")
	fmt.Printf("Version: %X\n", ih.IpVersion)
	fmt.Printf("IHL: %X\n", ih.HeaderLen)
	fmt.Printf("ServiceType: %X\n", ih.ServiceType)
	fmt.Println("TotalLen:", ih.TotalLen)
	fmt.Println("Identification:", ih.Identification)
	fmt.Printf("Flags: %X\n", ih.Flags)
	fmt.Printf("Offset: %X\n", ih.FragmentOffset)
	fmt.Println("TTL:", ih.TTL)
	fmt.Println("Protocol:", ih.NextProto)
	fmt.Printf("CheckSum: 0x%x\n", ih.CheckSum)
	fmt.Println("SrcIpAddr:", ih.SrcIpAddr)
	fmt.Println("DstIpAddr:", ih.DstIpAddr)
	fmt.Printf("-------------------------------\n")
}

func printIpv6(ih6 *Ipv6Header) {
	fmt.Printf("-------------IPv6--------------\n")
	fmt.Printf("Version: %X\n", ih6.Ipv6Version)
	fmt.Printf("TrafficClass: %X\n", ih6.TrafficClass)
	fmt.Printf("FlowLabel: 0x%x\n", ih6.FlowLabel)
	fmt.Println("PayloadLen:", ih6.PayloadLen)
	fmt.Println("NextHeader:", ih6.NextHeader)
	fmt.Println("HopLimit:", ih6.HopLimit)
	fmt.Println("SrcIpv6Addr:", ih6.SrcIpv6Addr)
	fmt.Println("DstIpv6Addr:", ih6.DstIpv6Addr)
	fmt.Printf("-------------------------------\n")
}

func printIcmp(icmph *ICMPHeader, data []byte) {
	fmt.Printf("-------------ICMP--------------\n")
	switch icmph.ICMPType {
	case IcmpEchoReply | Icmp6Echoreply:
		fmt.Printf("Type: Echo Reply(0x%x)\n", icmph.ICMPType)

	case IcmpDstUnreach | Icmp6DstUnreath:
		fmt.Printf("Type: Dst Unreach(0x%x)\n", icmph.ICMPType)

	case IcmpRedirect:
		fmt.Printf("Type: Redirect(0x%x)\n", icmph.ICMPType)

	case IcmpEchoReq | Icmp6Echoreq:
		fmt.Printf("Type: Echo Request(0x%x)\n", icmph.ICMPType)

	case IcmpExceeded:
		fmt.Printf("Type: Time Exceeded(0x%x)\n", icmph.ICMPType)

	default:
		fmt.Printf("Type: Unknown(%X)\n", icmph.ICMPType)
	}
	fmt.Printf("ICMPCode: %X\n", icmph.ICMPCode)
	fmt.Printf("CheckSum: %X\n", icmph.CheckSum)
	fmt.Printf("ICMPData: %X\n", data)
	fmt.Printf("-------------------------------\n")
}

func printTcp(tcph *TCPHeader, data []byte) {
	fmt.Printf("--------------TCP--------------\n")
	fmt.Println("SrcPort:", tcph.SrcPortNum)
	fmt.Println("DstPort:", tcph.DstPortNum)
	fmt.Println("SequenceNum:", tcph.SequenceNum)
	fmt.Println("AckNwNum:", tcph.AckNwlNum)
	fmt.Println("HeaderLen:", tcph.HeaderLen)
	fmt.Printf("Reservation: %x\n", tcph.Reservation)
	fmt.Printf("CtrlFlag: %x\n", tcph.CtrlFlag)
	fmt.Println("WindowSize:", tcph.WindowSize)
	fmt.Printf("CheckSum: 0x%x\n", tcph.CheckSum)
	fmt.Printf("UrgPointer: %x\n", tcph.UrgPointer)
	fmt.Printf("Data: %x\n", data)
	fmt.Printf("-------------------------------\n")
}

func printUdp(udph *UDPHeader, data []byte) {
	fmt.Printf("--------------UDP--------------\n")
	fmt.Println("SrcPort:", udph.SrcPortNum)
	fmt.Println("DstPort:", udph.DstPortNum)
	fmt.Println("PacketLen:", udph.PacketLen)
	fmt.Printf("CheckSum: 0x%x\n", udph.CheckSum)
	fmt.Printf("-------------------------------\n")
}
