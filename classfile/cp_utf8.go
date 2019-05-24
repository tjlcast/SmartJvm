package classfile

/**
	存放的是MUTF-8编码的字符串
 */
type ConstantUtf8Info struct {
	str		string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

/**
	简单版
 */
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
