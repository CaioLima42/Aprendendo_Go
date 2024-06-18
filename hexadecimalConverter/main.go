package main

import (
	"fmt"
)

const hextable = "0123456789abcdef"
func Encode(dst, src []byte) int {
	j := 0
	for _, v := range src {
		dst[j] = hextable[v>>4]
		dst[j+1] = hextable[v&0x0f]
		j += 2
	}
	return len(src) * 2
}

//1010 1111
func main(){
	fmt.Println(Encode([]byte{0b01111001}))
}