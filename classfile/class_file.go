package classfile

import "fmt"

type ClassFile struct {
	// magic				uint32
	minorVersion 			uint16
	majorVersion			uint16
	constantPool			ConstantPool
	accessFlags				uint16
	thisClass				uint16
	superClass				uint16
	interfaces				[]uint16
	fields					[]*MemberInfo
	methods					[]*MemberInfo
	attributes				[]AttributeInfo
}

/**
	Parse() 函数把[]byte解析成ClassFile结构体
 */
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	} ()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)

	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)

	self.readAndCheckVersion(reader)

	self.constantPool = readConstantPool(reader)

	self.accessFlags = reader.readUint16()

	self.thisClass = reader.readUint16()

	self.superClass = reader.readUint16()

	self.interfaces = reader.readUint16s()

	self.fields = readMembers(reader, self.constantPool)

	self.methods = readMembers(reader, self.constantPool)

	self.attributes = readAttributes(reader, self.constantPool)
}



/**
	魔术
 */
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/**
	版本号
 */
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()

	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

/**
	class 文件的最小版本
 */
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

/**
	class 文件的最大版本
 */
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

/**
	class 文件的常量池
 */
func (self *ClassFile) ConstantPool() ConstantPool {
	// todo
	return nil
}

/**
	类访问标志
	指出 class 文件定义的是类还是接口，以及访问级别是public还是private
 */
func (self *ClassFile) AccessFlag() uint16 {
	// todo
	return 12
}

/**
	className by idx
 */
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

/**
	superClassName by idx
 */
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" // 只有 java.lang.Object 没有超类
}

/**
	接口索引表
 */
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

func (self *ClassFile) Fields() []*MemberInfo {
	// todo
}

func (self *ClassFile) Methods() []*MemberInfo {
	// todo
}
