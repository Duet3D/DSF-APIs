package main

import (
	"log"
	"os"

	"github.com/Duet3D/DSF-APIs/godsfapi/v2/commands"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/connection"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/connection/initmessages"
	"github.com/Duet3D/DSF-APIs/godsfapi/v2/types"
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

	// Make sure we close the connection
	defer cc.Close()
	connection.CloseOnSignals(&cc)

	if code != "" {
		r, err := cc.PerformSimpleCode(code, types.SBC)
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
	err := sc.Connect(initmessages.SubscriptionModePatch, "", LocalSock)
	if err != nil {
		log.Panic(err)
	}

	// Make sure we close the connection
	defer sc.Close()
	connection.CloseOnSignals(&sc)

	m, err := sc.GetMachineModel()
	if err != nil {
		log.Panic(err)
	}
	log.Println(m)
	k := m.Move.Kinematics
	log.Println(k)
	ck, err := k.AsCoreKinematics()
	if err != nil {
		log.Panic(err)
	}
	log.Println(ck)
	log.Println(ck.AsKinematics())
	_, err = sc.GetMachineModelPatch()
	if err != nil {
		log.Panic(err)
	}
	// log.Println(ms)
}

func intercept() {
	ic := connection.InterceptConnection{}
	ic.Debug = true
	err := ic.Connect(initmessages.InterceptionModePre, LocalSock)
	if err != nil {
		log.Panic(err)
	}

	// Make sure we close the connection
	defer ic.Close()
	connection.CloseOnSignals(&ic)

	for {
		c, err := ic.ReceiveCode()
		if err != nil {
			log.Panic(err)
		}
		success, err := ic.Flush(c.Channel)
		if err != nil {
			log.Panic(err)
		}
		if !success {
			ic.CancelCode()
			continue
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
