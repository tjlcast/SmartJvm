package classfile

/**
	和类一样，字段和方法也有自己的访问标志。访问标志之后是一个常量池索引，给出字段名或方法名。
	然后又是一个常量池索引，给出字段或方法的描述符，最后是属性表。

	这里使用一个结构统一表示字段和方法。
 */

type MemberInfo struct {
	cp				ConstantPool
	accessFlags		uint16			//	--->	idx of cp
	nameIndex		uint16			// 	---> 	idx of cp
	descriptorIndex	uint16			// 	---> 	idx of cp
	attributes		[]AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:					cp,
		accessFlags:		reader.readUint16(),
		nameIndex:			reader.readUint16(),
		descriptorIndex:	reader.readUint16(),
		attributes:			readerAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	// todo
}

/**
	从常量池查找字段或方法名
 */
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

/**
	从常量池查找字段或方法描述符号
 */
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

