package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var tmp int16

	num := uint16(0xffff)
	byteArray := []byte{byte(num >> 8), byte(num)}
	fmt.Printf("byteArray[%v]\n", byteArray)

	bytesBuffer := bytes.NewBuffer(byteArray)
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)

	fmt.Printf("%v\n", tmp)
}
