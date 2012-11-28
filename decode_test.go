// Copyright (c) 2012 Graeme Connell. All rights reserved.
// Copyright (c) 2009-2012 Andreas Krennmair. All rights reserved.

package gopacket

import (
	"fmt"
	"testing"
)

var testSimpleTcpPacket []byte = []byte{
	0x00, 0x00, 0x0c, 0x9f, 0xf0, 0x20, 0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49,
	0x08, 0x00, 0x45, 0x00, 0x01, 0xa4, 0x39, 0xdf, 0x40, 0x00, 0x40, 0x06,
	0x55, 0x5a, 0xac, 0x11, 0x51, 0x49, 0xad, 0xde, 0xfe, 0xe1, 0xc5, 0xf7,
	0x00, 0x50, 0xc5, 0x7e, 0x0e, 0x48, 0x49, 0x07, 0x42, 0x32, 0x80, 0x18,
	0x00, 0x73, 0xab, 0xb1, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x03, 0x77,
	0x37, 0x9c, 0x42, 0x77, 0x5e, 0x3a, 0x47, 0x45, 0x54, 0x20, 0x2f, 0x20,
	0x48, 0x54, 0x54, 0x50, 0x2f, 0x31, 0x2e, 0x31, 0x0d, 0x0a, 0x48, 0x6f,
	0x73, 0x74, 0x3a, 0x20, 0x77, 0x77, 0x77, 0x2e, 0x66, 0x69, 0x73, 0x68,
	0x2e, 0x63, 0x6f, 0x6d, 0x0d, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x6b, 0x65, 0x65, 0x70, 0x2d, 0x61,
	0x6c, 0x69, 0x76, 0x65, 0x0d, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x2d, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x3a, 0x20, 0x4d, 0x6f, 0x7a, 0x69, 0x6c, 0x6c,
	0x61, 0x2f, 0x35, 0x2e, 0x30, 0x20, 0x28, 0x58, 0x31, 0x31, 0x3b, 0x20,
	0x4c, 0x69, 0x6e, 0x75, 0x78, 0x20, 0x78, 0x38, 0x36, 0x5f, 0x36, 0x34,
	0x29, 0x20, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x57, 0x65, 0x62, 0x4b, 0x69,
	0x74, 0x2f, 0x35, 0x33, 0x35, 0x2e, 0x32, 0x20, 0x28, 0x4b, 0x48, 0x54,
	0x4d, 0x4c, 0x2c, 0x20, 0x6c, 0x69, 0x6b, 0x65, 0x20, 0x47, 0x65, 0x63,
	0x6b, 0x6f, 0x29, 0x20, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2f, 0x31,
	0x35, 0x2e, 0x30, 0x2e, 0x38, 0x37, 0x34, 0x2e, 0x31, 0x32, 0x31, 0x20,
	0x53, 0x61, 0x66, 0x61, 0x72, 0x69, 0x2f, 0x35, 0x33, 0x35, 0x2e, 0x32,
	0x0d, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x3a, 0x20, 0x74, 0x65,
	0x78, 0x74, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x2c, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x68, 0x74, 0x6d,
	0x6c, 0x2b, 0x78, 0x6d, 0x6c, 0x2c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x6d, 0x6c, 0x3b, 0x71, 0x3d,
	0x30, 0x2e, 0x39, 0x2c, 0x2a, 0x2f, 0x2a, 0x3b, 0x71, 0x3d, 0x30, 0x2e,
	0x38, 0x0d, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x45, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x67, 0x7a, 0x69, 0x70,
	0x2c, 0x64, 0x65, 0x66, 0x6c, 0x61, 0x74, 0x65, 0x2c, 0x73, 0x64, 0x63,
	0x68, 0x0d, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x3a, 0x20, 0x65, 0x6e, 0x2d, 0x55,
	0x53, 0x2c, 0x65, 0x6e, 0x3b, 0x71, 0x3d, 0x30, 0x2e, 0x38, 0x0d, 0x0a,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2d, 0x43, 0x68, 0x61, 0x72, 0x73,
	0x65, 0x74, 0x3a, 0x20, 0x49, 0x53, 0x4f, 0x2d, 0x38, 0x38, 0x35, 0x39,
	0x2d, 0x31, 0x2c, 0x75, 0x74, 0x66, 0x2d, 0x38, 0x3b, 0x71, 0x3d, 0x30,
	0x2e, 0x37, 0x2c, 0x2a, 0x3b, 0x71, 0x3d, 0x30, 0x2e, 0x33, 0x0d, 0x0a,
	0x0d, 0x0a,
}

func BenchmarkCreatePacket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LINKTYPE_ETHERNET.Decode(testSimpleTcpPacket, Lazy)
	}
}

func BenchmarkGetEthLayer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LINKTYPE_ETHERNET.Decode(testSimpleTcpPacket, Lazy).Layer(TYPE_ETHERNET)
	}
}

func BenchmarkTypeAssertion(b *testing.B) {
	var eth LinkLayer = &Ethernet{}
	c := 0
	for i := 0; i < b.N; i++ {
		if _, ok := eth.(*Ethernet); ok {
			c++
		}
	}
}

func BenchmarkGetIpLayer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LINKTYPE_ETHERNET.Decode(testSimpleTcpPacket, Lazy).Layer(TYPE_IP4)
	}
}

func BenchmarkGetTcpLayer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LINKTYPE_ETHERNET.Decode(testSimpleTcpPacket, Lazy).Layer(TYPE_TCP)
	}
}

func BenchmarkGetAllLayers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LINKTYPE_ETHERNET.Decode(testSimpleTcpPacket, Lazy).Layers()
	}
}

func BenchmarkDecodeEager(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LINKTYPE_ETHERNET.Decode(testSimpleTcpPacket, Eager)
	}
}

func BenchmarkAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &TCP{}
	}
}

func BenchmarkFlowKey(b *testing.B) {
	p := LINKTYPE_ETHERNET.Decode(testSimpleTcpPacket, Lazy)
	net := p.NetworkLayer()
	trans := p.TransportLayer()
	for i := 0; i < b.N; i++ {
		NewFlowKey(net, trans)
	}
}

func BenchmarkPacketFlowKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LINKTYPE_ETHERNET.Decode(testSimpleTcpPacket, Lazy).FlowKey()
	}
}

func BenchmarkFlowMapKey(b *testing.B) {
	p := LINKTYPE_ETHERNET.Decode(testSimpleTcpPacket, Lazy)
	net := p.NetworkLayer()
	trans := p.TransportLayer()
	m := map[FlowKey]bool{}
	for i := 0; i < b.N; i++ {
		m[NewFlowKey(net, trans)] = true
	}
}

func TestDecodeSimpleTcpPacket(t *testing.T) {
	Equal := func(desc, expected string, actual fmt.Stringer) {
		if expected != actual.String() {
			t.Errorf("%s: Expected %s Got %s", desc, expected, actual)
		}
	}
	p := LINKTYPE_ETHERNET.Decode(testSimpleTcpPacket, Lazy)
	if eth := p.LinkLayer(); eth == nil {
		t.Error("No ethernet layer found")
	} else {
		Equal("Eth Src", "bc:30:5b:e8:d3:49", eth.SrcLinkAddr())
		Equal("Eth Dst", "00:00:0c:9f:f0:20", eth.DstLinkAddr())
	}
	if net := p.NetworkLayer(); net == nil {
		t.Error("No net layer found")
	} else if ip, ok := net.(*IPv4); !ok {
		t.Error("Net layer is not IP layer")
	} else {
		Equal("IP Src", "172.17.81.73", net.SrcNetAddr())
		Equal("IP Dst", "173.222.254.225", net.DstNetAddr())
		if ip.Version != 4 {
			t.Error("ip Version", ip.Version)
		}
		if ip.Ihl != 5 {
			t.Error("ip header length", ip.Ihl)
		}
		if ip.Tos != 0 {
			t.Error("ip TOS", ip.Tos)
		}
		if ip.Length != 420 {
			t.Error("ip Length", ip.Length)
		}
		if ip.Id != 14815 {
			t.Error("ip ID", ip.Id)
		}
		if ip.Flags != 0x02 {
			t.Error("ip Flags", ip.Flags)
		}
		if ip.FragOffset != 0 {
			t.Error("ip Fragoffset", ip.FragOffset)
		}
		if ip.Ttl != 64 {
			t.Error("ip TTL", ip.Ttl)
		}
		if ip.Protocol != 6 {
			t.Error("ip Protocol", ip.Protocol)
		}
		if ip.Checksum != 0x555A {
			t.Error("ip Checksum", ip.Checksum)
		}
	}
	if trans := p.TransportLayer(); trans == nil {
		t.Error("No transport layer found")
	} else if tcp, ok := trans.(*TCP); !ok {
		t.Error("Transport layer is not TCP layer")
	} else {
		Equal("TCP Src", "50679", trans.SrcAppAddr())
		Equal("TCP Dst", "80", trans.DstAppAddr())
		if tcp.SrcPort != 50679 {
			t.Error("tcp srcport", tcp.SrcPort)
		}
		if tcp.DstPort != 80 {
			t.Error("tcp destport", tcp.DstPort)
		}
		if tcp.Seq != 0xc57e0e48 {
			t.Error("tcp seq", tcp.Seq)
		}
		if tcp.Ack != 0x49074232 {
			t.Error("tcp ack", tcp.Ack)
		}
		if tcp.DataOffset != 8 {
			t.Error("tcp dataoffset", tcp.DataOffset)
		}
		if tcp.Flags != 0x18 {
			t.Error("tcp flags", tcp.Flags)
		}
		if tcp.Window != 0x73 {
			t.Error("tcp window", tcp.Window)
		}
		if tcp.Checksum != 0xabb1 {
			t.Error("tcp checksum", tcp.Checksum)
		}
		if tcp.Urgent != 0 {
			t.Error("tcp urgent", tcp.Urgent)
		}
	}
	if payload, ok := p.Layer(TYPE_PAYLOAD).(*Payload); payload == nil || !ok {
		t.Error("No payload layer found")
	} else {
		if string(payload.Data) != "GET / HTTP/1.1\r\nHost: www.fish.com\r\nConnection: keep-alive\r\nUser-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.2 (KHTML, like Gecko) Chrome/15.0.874.121 Safari/535.2\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: gzip,deflate,sdch\r\nAccept-Language: en-US,en;q=0.8\r\nAccept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.3\r\n\r\n" {
			t.Error("--- PAYLOAD STRING ---\n", string(payload.Data), "\n--- PAYLOAD BYTES ---\n", payload.Data)
		}
	}
}

// Makes sure packet payload doesn't display the 6 trailing null of this packet
// as part of the payload.  They're actually the ethernet trailer.
func TestDecodeSmallTcpPacketHasEmptyPayload(t *testing.T) {
	p := LINKTYPE_ETHERNET.Decode(
		[]byte{
			0xbc, 0x30, 0x5b, 0xe8, 0xd3, 0x49, 0xb8, 0xac, 0x6f, 0x92, 0xd5, 0xbf,
			0x08, 0x00, 0x45, 0x00, 0x00, 0x28, 0x00, 0x00, 0x40, 0x00, 0x40, 0x06,
			0x3f, 0x9f, 0xac, 0x11, 0x51, 0xc5, 0xac, 0x11, 0x51, 0x49, 0x00, 0x63,
			0x9a, 0xef, 0x00, 0x00, 0x00, 0x00, 0x2e, 0xc1, 0x27, 0x83, 0x50, 0x14,
			0x00, 0x00, 0xc3, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}, Eager)

	if payload := p.Layer(TYPE_PAYLOAD); payload != nil {
		t.Error("Payload found for empty TCP packet")
	}
}

func TestDecodeVlanPacket(t *testing.T) {
	p := LINKTYPE_ETHERNET.Decode(
		[]byte{
			0x00, 0x10, 0xdb, 0xff, 0x10, 0x00, 0x00, 0x15, 0x2c, 0x9d, 0xcc, 0x00,
			0x81, 0x00, 0x01, 0xf7, 0x08, 0x00, 0x45, 0x00, 0x00, 0x28, 0x29, 0x8d,
			0x40, 0x00, 0x7d, 0x06, 0x83, 0xa0, 0xac, 0x1b, 0xca, 0x8e, 0x45, 0x16,
			0x94, 0xe2, 0xd4, 0x0a, 0x00, 0x50, 0xdf, 0xab, 0x9c, 0xc6, 0xcd, 0x1e,
			0xe5, 0xd1, 0x50, 0x10, 0x01, 0x00, 0x5a, 0x74, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
		}, Eager)
	if err := p.ErrorLayer(); err != nil {
		t.Error("Error while parsing vlan packet:", err.Error())
	}
	if vlan := p.Layer(TYPE_DOT1Q); vlan == nil {
		t.Error("Didn't detect vlan")
	} else if _, ok := vlan.(*Dot1Q); !ok {
		t.Error("TYPE_DOT1Q layer is not a Dot1Q object")
	}
	for i, l := range p.Layers() {
		t.Logf("Layer %d: %#v", i, l)
	}
	expected := []LayerType{TYPE_ETHERNET, TYPE_DOT1Q, TYPE_IP4, TYPE_TCP}
	if len(p.Layers()) != len(expected) {
		t.Error("Incorrect number of headers:", len(p.Layers()))
		t.FailNow()
	}
	for i, l := range p.Layers() {
		if l.LayerType() != expected[i] {
			t.Error("At index", i, "expected layer type", expected[i], ", received", l.LayerType())
		}
	}
}