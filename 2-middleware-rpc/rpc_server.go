package middleware_rpc

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// type Args struct{}

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	return nil
}

func RPCServer() {
	timeserver := new(TimeServer)
	rpc.Register(timeserver)
	rpc.HandleHTTP()

	connection, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error: ", err)
	}

	http.Serve(connection, nil)
}
