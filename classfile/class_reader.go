package classfile

import "encoding/binary"

/**
	用于辅助读取class文件字节数据

	Go标准库 encoding/binary包中定义一个
	变量BigEndian(大端字节)

	修改对象中的数据切片
 */


type ClassReader struct {
	data	[]byte
}

func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}
