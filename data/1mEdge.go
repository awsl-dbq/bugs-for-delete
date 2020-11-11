package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: ./1mEdge start num \n\t ./1mEdge 1 100000")
		return
	}
	s := os.Args[1]
	n := os.Args[2]
	su, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	nu, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	name := fmt.Sprintf("%d__%d.rdf", su, nu)
	fmt.Println("output to", name)
	fh, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModePerm)
	rdf(fh, uint64(su), uint64(nu))
	fmt.Println("Rdf generated!")
	// fmt.Println("run dgo")
	// conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer conn.Close()
	// dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	// txn := dgraphClient.NewTxn()
	// ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	// rdf, _ := ioutil.ReadFile(name)
	// mu := &api.Mutation{
	// 	SetNquads: rdf,
	// }
	// res, err := txn.Mutate(ctx, mu)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res)
	// defer cancel()
	// defer txn.Discard(ctx)

}
func rdf(fh *os.File, start uint64, num uint64) {
	s := fmt.Sprintf("_:s%x", start)
	for i := uint64(1); i < num+1; i++ {
		o := fmt.Sprintf("_:s%x", i+start)
		fh.WriteString(fmt.Sprintf("%s <follow> %s . \n", s, o))
	}
	fh.Close()
}
