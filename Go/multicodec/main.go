package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"reflect"

	ma "github.com/jbenet/go-multiaddr"
	mc "github.com/jbenet/go-multicodec"
	mux "github.com/jbenet/go-multicodec/mux"
)

func doStdin() {
	c := mux.StandardMux()
	dec := c.Decoder(os.Stdin)

	var v interface{}
	if err := dec.Decode(&v); err != nil {
		panic(err)
	}

	fmt.Println(v)
}

type MyStruct struct {
	Hello     string
	greetings string // won't be encoded, because it's not exported
}

func doStruct() {
	s := MyStruct{"world", "earth"}
	fmt.Printf("Encoding: %+v\n", s)

	var buf bytes.Buffer
	mx := mux.StandardMux()
	mx.Encoder(&buf).Encode(s)

	fmt.Printf("Encoded buffer: %s\n", buf.String())

	var v MyStruct
	mx.Decoder(&buf).Decode(&v)

	fmt.Printf("Decoded: %+v\n", v)
}

type IPeer interface {
	// Adding a method for a more real life example.
	// Now Peer implements IPeer, but interface{} doesn't.
	GetId() string
}

type Peer struct {
	Id string
	// Can't use type []ma.Multiaddr for Addrs, because the underlying type
	// is not exported and can't be copied or set using reflection.
	Addrs  [][]byte
	Params map[string]interface{}
}

func (p Peer) GetId() string {
	return p.Id
}

func doInterface() {
	m, _ := ma.NewMultiaddr("/ip4/127.0.0.1")
	ps := []IPeer{
		Peer{"me", [][]byte{m.Bytes()}, nil},
		Peer{"they", [][]byte{m.Bytes()}, map[string]interface{}{"age": 3}},
	}
	fmt.Printf("Encoding: %+v\n", ps)

	var buf bytes.Buffer
	mx := mux.StandardMux()
	mx.Encoder(&buf).Encode(ps)

	fmt.Printf("Encoded buffer: %s\n", buf.String())

	// Now we'll decode the []Peer structure into []IPeer slice
	// without mentioning the type Peer.
	dec := mx.Decoder(&buf)
	typ := reflect.SliceOf(reflect.TypeOf(ps[0]))
	val, err := reflectDecode(dec, typ)
	if err != nil {
		panic(err)
	}

	v := make([]IPeer, val.Len())
	for i := 0; i < val.Len(); i++ {
		v[i] = val.Index(i).Interface().(IPeer)
	}

	fmt.Printf("Decoded: %+v\n", v)
}

func reflectDecode(dec mc.Decoder, typ reflect.Type) (reflect.Value, error) {
	x := reflect.New(typ).Interface()
	if err := dec.Decode(x); err != nil {
		return reflect.Value{}, err
	}

	return reflect.Indirect(reflect.ValueOf(x)), nil
}

func main() {
	stdin := flag.Bool("stdin", false, "")
	flag.Parse()

	if *stdin {
		doStdin()
	} else {
		doStruct()
		fmt.Println("---")
		doInterface()
	}
}
