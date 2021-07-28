package example_pprof
// 使用runtime的pprof包记录内存使用信息
// 可使用top，trace查看文件
import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func RuntimePprofCpu() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
	for i := 0; i < 102400; i++ {
		s := make([]byte, 3)
		fmt.Sprint(s, "hello world\n")
	}
}
