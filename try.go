package main
import(
	"fmt"
	"hash/fnv"
	"encoding/binary"
)
func main(){
	in := 99999999
	fmt.Println(in)
	h := fnv.New64a()
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(in))
	h.Write(bs)
	fmt.Print(h.Sum64())
}
