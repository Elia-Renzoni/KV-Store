package main

import (
	"log"
	"net/rpc"
)

const (
	MAX_CONN int = 30
	WRITER   int = 14
	READER   int = 6
	DELETER  int = 10
)

type serverResponse []byte

type Pair struct {
	key   int
	value []byte
}

func main() {
	var reply serverResponse
	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < MAX_CONN; i++ {
		if i < WRITER {
			go func() {
				toSend := &Pair{
					key: ,
					value: ,
				}
				client.Call("DistributedCache.Set")
			}()
		}

	}

}
