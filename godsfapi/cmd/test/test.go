package main

import (
	"log"
	"os"

	"github.com/Duet3D/DSF-APIs/godsfapi/commands"
	"github.com/Duet3D/DSF-APIs/godsfapi/connection"
	"github.com/Duet3D/DSF-APIs/godsfapi/connection/initmessages"
	"github.com/Duet3D/DSF-APIs/godsfapi/types"
)

const (
	LocalSock = "/home/manuel/tmp/duet.sock"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}
	switch os.Args[1] {
	case "subscribe":
		subscribe()
	case "intercept":
		intercept()
	case "command":
		if len(os.Args) > 2 {
			for _, c := range os.Args[2:] {
				command(c)
			}
		} else {
			command("")
		}
	}
}

func command(code string) {
	cc := connection.CommandConnection{}
	err := cc.Connect(LocalSock)
	if err != nil {
		panic(err)
	}
	defer cc.Close()
	if code != "" {
		r, err := cc.PerformSimpleCode(code, types.SPI)
		if err != nil {
			log.Panic(err)
		}
		log.Println(r)
	} else {
		mm, err := cc.GetSerializedMachineModel()
		if err != nil {
			log.Panic(err)
		}
		log.Println(string(mm))
	}
}

func subscribe() {
	sc := connection.SubscribeConnection{}
	sc.Debug = true
	err := sc.Connect(initmessages.SubscriptionModePatch, "heat/**", LocalSock)
	if err != nil {
		log.Panic(err)
	}
	defer sc.Close()
	m, err := sc.GetMachineModelPatch()
	if err != nil {
		log.Panic(err)
	}
	log.Println(m)
}

func intercept() {
	ic := connection.InterceptConnection{}
	ic.Debug = true
	err := ic.Connect(initmessages.InterceptionModePre, LocalSock)
	if err != nil {
		log.Panic(err)
	}
	defer ic.Close()
	for {
		c, err := ic.ReceiveCode()
		if err != nil {
			log.Panic(err)
		}
		cc := c.Clone()
		cc.Flags |= commands.Asynchronous
		ic.PerformCode(cc)
		// log.Println(c)
		err = ic.IgnoreCode()
		if err != nil {
			log.Panic(err)
		}
	}
}
