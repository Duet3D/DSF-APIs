package main

import (
	"log"
	"os"

	"github.com/Duet3D/DSF-APIs/godsfapi/v3/commands"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/connection"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/connection/initmessages"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/machine/messages"
	"github.com/Duet3D/DSF-APIs/godsfapi/v3/types"
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
	err := sc.Connect(initmessages.SubscriptionModePatch, nil, LocalSock)
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
	err := ic.Connect(initmessages.InterceptionModePre, nil, nil, false, LocalSock)
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
		if c.Type == commands.MCode && c.IsMajorNumber(1234) {

			success, err := ic.Flush()
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
			ic.ResolveCode(messages.Success, "")
			// log.Println(c)
		} else {
			err = ic.IgnoreCode()
		}
		if err != nil {
			log.Panic(err)
		}
	}
}
