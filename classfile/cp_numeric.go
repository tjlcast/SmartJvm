package classfile

import "math"

/**
	cp:	整形
 */
type ConstantIntegerInfo struct {
	val		int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

/**
	cp:	浮点型
 */
type ConstantFloatInfo struct {
	val		float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

/**
	cp:	长整型
 */
type ConstantLongInfo struct {
	val	int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

/**
	cp:	长浮点型
 */
type ConstantDoubleInfo struct {
	val	float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}