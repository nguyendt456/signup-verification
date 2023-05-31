package deploy

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/mail.v2"

	"github.com/nguyendt456/signup-with-verification/api/signup"
	"github.com/nguyendt456/signup-with-verification/database"
	pb "github.com/nguyendt456/signup-with-verification/pb"
)

func StartGW() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	endpoint := os.Getenv("GRPC")
	addr := os.Getenv("GRPC_GW")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = pb.RegisterSignupServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Starting gRPC gateway with hostname: %s\n", addr)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func StartgRPCservice() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	addr := os.Getenv("GRPC")

	g, err := database.InitPostgresFromDotenv()
	if err != nil {
		log.Println(err)
		return
	}
	r, err := database.InitRedisFromDotenv()
	if err != nil {
		log.Println(err)
		return
	}

	e := signup.EmailSender{Mail: mail.NewMessage()}
	err = e.InitFromDotenv()
	if err != nil {
		log.Println(err)
		return
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSignupServiceServer(grpcServer, &signup.SignupAPI{
		Email:    &e,
		Database: g,
		Redis:    r,
	})

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Starting gRPC service with hostname: %s\n", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Println(err)
		return
	}
}
