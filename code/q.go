package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	dgo "github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

var (
	dgraph = flag.String("d", "127.0.0.1:9080", "Dgraph Alpha address")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	log.Println("Before Delete Query:")
	ret := query(dg)
	deleteFollow(dg, ret.Uid)
	log.Println("After Delete Query:")
	ret = query(dg)
	if ret.C != 0 {
		log.Printf("should be zero ,but got %v", ret.C)
	}
}

type c struct {
	Uid string
	C   int
}
type q struct {
	Q []c
}

func query(dg *dgo.Dgraph) c {
	resp, err := dg.NewTxn().Query(context.Background(), `{
		  q(func: has(follow)) { 	
		uid
 		c: count(follow)
	
}}`)

	if err != nil {
		log.Fatal(err)
	}
	var str bytes.Buffer
	json.Indent(&str, resp.Json, "", "\t")

	log.Printf("Response: %s\n", str.String())
	var r q
	json.Unmarshal(str.Bytes(), &r)
	return r.Q[0]
}

func deleteFollow(dg *dgo.Dgraph, uid string) {
	del := fmt.Sprintf("<%s> <follow>  * .", uid)
	log.Printf("%s", del)
	mu := &api.Mutation{
		DelNquads: []byte(del),
	}
	req := &api.Request{
		//Query: query,
		Mutations: []*api.Mutation{mu},
		CommitNow: true,
	}
	resp, err := dg.NewTxn().Do(context.Background(), req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var str bytes.Buffer
	json.Indent(&str, resp.Json, "", "\t")

	fmt.Printf("Response: %s\n", str.String())
}
