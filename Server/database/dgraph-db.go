package database

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

type CancelFunc func()

func GetDgraphClient() (*dgo.Dgraph, CancelFunc) {

	direccion := "localhost:9080"
	conn, err := grpc.Dial(direccion, grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	dc := api.NewDgraphClient(conn)
	dgraphClient := dgo.NewDgraphClient(dc)

	if err != nil {
		log.Fatalf("While trying to login %v", err.Error())
	}

	return dgraphClient, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}

func MutateDatabase(object []byte) {
	log.Println(object)
	dg, cancel := GetDgraphClient()
	defer cancel()

	ctx := context.Background()

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	mu := &api.Mutation{
		SetJson: object,
	}

	res, err := txn.Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	err2 := txn.Commit(context.Background())
	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println(string(res.Json))
}

func GetDBConsult(query string) []byte {
	dg, cancel := GetDgraphClient()
	defer cancel()

	ctx := context.Background()

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	res, err := txn.Query(ctx, query)

	if err == nil {

		log.Println(string(res.Json))
		return (res.Json)
	} else {
		return nil
	}
}
