
this is a project about jvm.

to be continue....

测试：
    - 启动时进行类加载
    命令：
        order -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_101.jdk/Contents/Home/jre" java.lang.Object


https://github.com/zxh0/jvm.go


jvm 解析器大致逻辑

do {
    atomically calculate pc and fetch opcode at pc
    if (operands) fetch operands;
    execute the action for the opcode;
} while (there is more to do);