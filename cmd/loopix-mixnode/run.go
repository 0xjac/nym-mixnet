package main

import (
	"fmt"
	"os"

	"github.com/nymtech/loopix-messaging/helpers"
	"github.com/nymtech/loopix-messaging/pki"
	"github.com/nymtech/loopix-messaging/server"
	"github.com/nymtech/loopix-messaging/sphinx"
	"github.com/tav/golly/optparse"
)

const (
	// PkiDb is the location of the database file, relative to the project root. TODO: move this to homedir.
	PkiDb        = "pki/database.db"
	defaultHost  = "localhost"
	defaultID    = "Mix1"
	defaultPort  = "6666"
	defaultLayer = -1
)

func cmdRun(args []string, usage string) {
	opts := newOpts("run [OPTIONS]", usage)
	id := opts.Flags("--id").Label("ID").String("Id of the loopix-client we want to run", defaultID)
	host := opts.Flags("--host").Label("HOST").String("The host on which the loopix-client is running", defaultHost)
	port := opts.Flags("--port").Label("PORT").String("Port on which loopix-client listens", defaultPort)
	layer := opts.Flags("--layer").Label("Layer").Int("Mixnet layer of this particular node", defaultLayer)

	params := opts.Parse(args)
	if len(params) != 0 {
		opts.PrintUsage()
		os.Exit(1)
	}

	err := pki.EnsurePkiDb(PkiDb)
	if err != nil {
		fmt.Println("PkiDb problem ")
		panic(err)
	}

	ip, err := helpers.GetLocalIP()
	if err != nil {
		panic(err)
	}

	if host != &ip {
		host = &ip
	}

	pubM, privM, err := sphinx.GenerateKeyPair()
	if err != nil {
		panic(err)
	}

	mixServer, err := server.NewMixServer(*id, *host, *port, pubM, privM, PkiDb, *layer)
	if err != nil {
		panic(err)
	}

	err = mixServer.Start()
	if err != nil {
		panic(err)
	}

	wait := make(chan struct{})
	<-wait
}

func newOpts(command string, usage string) *optparse.Parser {
	return optparse.New("Usage: loopix-mixnode " + command + "\n\n  " + usage + "\n")
}
