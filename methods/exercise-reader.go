package main

import (
	"golang.org/x/tour/reader" 
	//ascii "encoding/ascii85"
	//"fmt"
)

type MyReader struct{
	i int//index
}


// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte)(int, error){
	b[m.i] = 'A'
	m.i++
	return m.i, nil//ascii.CorruptInputError(float64('A'))
	
}
func main() {
	reader.Validate(MyReader{})
}

