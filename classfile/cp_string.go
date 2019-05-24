package classfile

/**
	表示java.lang.String字面量
 */
type ConstantStringInfo struct {
	cp				ConstantPool
	stringIndex		uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
