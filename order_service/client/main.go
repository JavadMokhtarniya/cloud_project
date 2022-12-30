package main

import (
//    "context"
//    "fmt"
//    "server"

//    "log"
//    "net"
)
func main(){
r, err = c.GetProduct(ctx, &pb.PID{id: *id})
if err != nil {
        log.Fatalf("could not greet: %v", err)
}
log.Printf("Greeting: %s", r.GetMessage())
}