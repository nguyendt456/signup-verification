package test

import (
	"context"
	"log"
	"testing"

	pb "github.com/nguyendt456/signup-with-verification/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type gRPCTestSuite struct {
	suite.Suite
}

func (s *gRPCTestSuite) Test1_ClientDial() {
	conn, err := grpc.Dial("0.0.0.0:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	new_client := pb.NewSignupServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := new_client.Signup(ctx, &pb.User{
		Email:    "nguyendt456@gmail.com",
		Password: "12345112",
		Name:     "Nguyen",
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	assert.Equal(s.T(), "verif", res.State)

}

func TestMain(t *testing.T) {
	suite.Run(t, new(gRPCTestSuite))
}
