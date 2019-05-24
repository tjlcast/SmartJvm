package classpath

import (
	"path/filepath"
	"os"
)

type Classpath struct {
	bootClasspath		Entry
	extClasspath		Entry
	userClasspath		Entry
}

/**
	Classpath对象实例的创建接口
 */
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

/**
	1、首先根据参数Xjre
	2、如果没有Xjre，在当前路径下寻找
	3、根据系统的JAVA_HOME路径加载
 */
func getJreDir(jreOption string) string {
	if jreOption != "" && exits(jreOption) {
		return jreOption
	}

	if exits("./jre") {
		return "./jre'"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return jh
	}

	panic("Can not find jre folder!")
}

/**
	doc: 	判断路径是否存在
	input: 	一个路径字符串
 */
func exits(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

/**
	配置boot和ext类路径
	分别为：/jre/lib/*  /jre/lib/ext/*
 */
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

/**
	配置用户类路径
	默认为当前路径下
 */
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" || cpOption == " " {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

/**
	doc:	加载class
			依次从boot、ext、user中查找并加载类数据
	input:	完整的类名
 */
func (self *Classpath) ReadClass(classname string) ([]byte, Entry, error) {
	// boot_cp
	classname = classname + ".class"
	// 这里的class文件命名： java/lang/Object.class
	if data, entry, err := self.bootClasspath.readClass(classname); err == nil {
		return data, entry, err
	}

	// ext_cp
	if data, entry, err := self.extClasspath.readClass(classname); err == nil {
		return data, entry, err
	}

	// user_cp
	return self.userClasspath.readClass(classname)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}


