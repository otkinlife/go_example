package example_pprof
// 使用runtime的pprof包记录内存使用信息
// 可使用top，trace查看文件
import (
	"go_example/test_func"
	"log"
	"os"
	"runtime/pprof"
)

func RuntimePprofMem () {
	f, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}

	test_func.Fibonacci(40)
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}

	defer f.Close()
}
