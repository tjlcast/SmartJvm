package classfile

type ConstantPool []ConstantInfo

/**
	读取常量池信息
	1、表头给出的常量池的大小比实际大1
	2、有效索引是1～n-1
	3、ConstantLongInfo 和 ConstantDoubleInfo 各占两个位置，1～N-1的某些数会变成无效索引
 */
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	for i:=1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

/**
	根据索引查找常量
 */
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

/**
	根据索引查找字段或方法的名字和描述符
 */
func (self ConstantPool) getNamedAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

/**
	 根据索引查找类名
 */
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

