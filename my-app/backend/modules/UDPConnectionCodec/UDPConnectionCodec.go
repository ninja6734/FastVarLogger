package UDPConnectionCodec

import (
	"fmt"
)

const (
	MagicNumber = "FVLR"
	DataSize    = 2042
	EndCodec    = "\x00\x00"
)

func ReadUDPData(data []byte) {
	for i, b := range data {
		fmt.Printf("Byte %d: %x\n", i, b)
	}
	// Implementation for reading buffer
}

func CreateUDPDataPacket(data []byte) []byte {
	// magic number: FVLR as bytes so 46 56 4c 52 :   4 bytes
	// data: 2042 bytes

	packet := make([]byte, 0)
	packet = append(packet, []byte{0x46, 0x56, 0x4c, 0x52}...) // magic number
	packet = append(packet, data...)                           // data
	return packet
}

func DecodeUDPData(data []byte) ([]byte, error) {
	if len(data) < DataSize+4 || string(data[:4]) != MagicNumber {
		return nil, fmt.Errorf("Invalid UDP data")
	}
	return data[4:], nil
}

func EncodeUDPData(data []byte) ([]byte, error) {
	packet := make([]byte, 0)
	packet = append(packet, []byte{0x46, 0x56, 0x4c, 0x52}...) // magic number
	packet = append(packet, data...)
	return packet, nil
}
