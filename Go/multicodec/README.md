## Run ##

### Decode standard input ###

```bash
$ echo -e '\x0c/multicodec\n\x06/json\n{"hello":"world"}' | ./multicodec -stdin
map[hello:world]
```

### Internal encode/decode examples. ###

The upper one demonstrates that unexported fields do not get encoded.

The lower one shows how to decode an array of structs into a
an array of interfaces that the struct implements.

```bash
$ ./multicodec
Encoding: {Hello:world greetings:earth}
Encoded buffer: /multicodec
/cbor
¡eHelloeworld
Decoded: {Hello:world greetings:}
---
Encoding: [{Id:me Addrs:[[4 127 0 0 1]] Params:map[]} {Id:they Addrs:[[4 127 0 0 1]] Params:map[age:3]}]
Encoded buffer: /multicodec
/cbor
£bIdbmeeAddrsE  fParams £bIddtheyeAddrsE  fParams¡cage
Decoded: [{Id:me Addrs:[[4 127 0 0 1]] Params:map[]} {Id:they Addrs:[[4 127 0 0 1]] Params:map[age:3]}]
```

