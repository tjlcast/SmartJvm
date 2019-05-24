package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

/**
	使用Entry表示类路径项，分别有4种实现：
	DirEntry、ZipEntry、CompositeEntry、WildcardEntry
 */

type Entry interface {
	// 查找和加载class文件，className是class文件的相对路径
	readClass(className string) ([]byte, Entry, error)
	String() string
}

/**
	newEntry根据 文件路径 创建不同类型的entry实例
 */
func newEntry(path string) Entry {
	// 有 目录分隔符 时
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	// 有 星号 时
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	// .jar 或则 .zip
	if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") ||
		strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
		return newZipEntry(path)
	}

	// 具体路径 时
	return newDirEntry(path)
}
