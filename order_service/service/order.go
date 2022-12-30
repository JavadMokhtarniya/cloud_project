package order

import (
   "context"
   "fmt"
//    "https://github.com/cloud-1401-project/auth-service.git"
   "google.golang.org/grpc"
   "log"
   "net"
)

type server struct{}

func main() {
   fmt.Println("Server is running...")

   // Make a listener
   lis, err := net.Listen("tcp", "0.0.0.0:50051")
   if err != nil {
      log.Fatalf("Failed to listen: %v", err)
   }

   // Make a gRPC server
   grpcServer := grpc.NewServer()
   product.RegisterCalculatorServiceServer(grpcServer, &server{})

   // Run the gRPC server
   if err := grpcServer.Serve(lis); err != nil {
      log.Fatalf("Failed to serve: %v", err)
   }
}

func (*server) product(ctx context.Context, req *product.PID) (*product.PInfo, error) {
   fmt.Printf("Received product RPC: %v", req)

//    id := req.id()
   

//    sum := firstNumber + secondNumber

//    res := &product.PInfo{
//       title: sum,
//    }

//    return res, nil
}