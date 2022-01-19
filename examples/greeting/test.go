package greeting

import "fmt"

// Test 一个包下的文件重名民没有任何影响，同一目录下所有文件都属于同一个包(main包除外)
// 这样设计的好处在于同一个目录下文件可以任意重命名，包为逻辑集合体，可以做到文件无关性，可以任意拆分重组
func Test() {
	fmt.Println("package can contains any files,and file names")
}
