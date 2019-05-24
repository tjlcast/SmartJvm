package classfile

/**
	Constant_Fieldref_info				表示字段符号引用
	Constant_Methodref_info				表示普通方法符号引用
	Constant_interfaceMethodref_info	表示接口方法符号引用
 */
type ConstantMemberrefInfo struct {
	cp					ConstantPool
	classIndex			uint16
	nameAndTypeIndex	uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex 		= reader.readUint16()
	self.nameAndTypeIndex 	= reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNamedAndType(self.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}