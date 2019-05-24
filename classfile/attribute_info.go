package classfile

/**
	jvm规范没有使用tag，而是使用属性名来区别不同的属性
	属性表中存放的的属性名实际上并不是编码后的字符串，而是常量池索引，指向常量池中的constant_utf8_info常量
 */
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i, _ := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttirbuteInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return nil
}

func newAttirbuteInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	// todo
	return nil
}
