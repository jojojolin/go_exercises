package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
)

type rot13Reader struct {
	r io.Reader
}



func (rot rot13Reader) Read(b []byte) (int, error){
	r := rot.r
	n, e := r.Read(b)
	if e!= nil{
		return 0, e
	}
	for i:=0; i<n;i++{
		if b[i]>=65 && b[i]<=90{
			b[i]= (b[i]-64+13)%26+64
		} else if b[i]>=97 && b[i]<=122{
			b[i]= (b[i]-96+13)%26+96
		}
		//fmt.Println(b[i])
	}
	return n, nil
	
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

