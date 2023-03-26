package advanced

import (
	"bytes"
	"fmt"
)

func Interfaces_Concept(){
	var w writer = consoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type writer interface {
	Write([]byte) (int, error)
}

type consoleWriter struct {}

func (cw consoleWriter) Write(data []byte) (int, error){
	n, err := fmt.Println(string(data))
	return n, err
}

func Interfaces_WithPrimitiveTypes(){
	myInt := intCounter(0)
	var inc incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type incrementer interface {
	Increment() (int)
}

type intCounter int

func (ic *intCounter) Increment() (int){
	*ic++
	return int(*ic)
}

func Interfaces_EmbeddedInterfaces(){
	var wc writerCloser2 = newBufferedWriterCloser()
	wc.Write([]byte("Hello Youtube listeners, this is a test"))
	wc.Close()
}

type writer2 interface {
	Write([]byte) (int, error)
}

type closer2 interface {
	Close() (error)
}

type writerCloser2 interface {
	writer2
	closer2
}

type bufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *bufferedWriterCloser) Write(data []byte) (int, error){
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)

	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

func (bwc *bufferedWriterCloser) Close() (error) {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)

		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}

	return nil
}

func newBufferedWriterCloser() (*bufferedWriterCloser) {
	return &bufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

