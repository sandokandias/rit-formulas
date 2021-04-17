// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"io"
	"strconv"

	"context"
	"log"
	"time"

	"github.com/sandokandias/card-vault-service/pkg/grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:5050"
)

type Formula struct {
	CardNumber string
	CardHolder string
	ExpMonth   string
	ExpYear    string
	UserID     string
}

func (f Formula) Run(writer io.Writer) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := api.NewCardServiceClient(conn)

	ctx := metadata.AppendToOutgoingContext(context.Background(), "tenant_id", "org")
	clientDeadline := time.Now().Add(time.Duration(10) * time.Second)
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()

	expM, _ := strconv.Atoi(f.ExpMonth)
	expY, _ := strconv.Atoi(f.ExpYear)

	r, err := c.AddCard(ctx, &api.AddCardRequest{
		CardNumber: f.CardNumber,
		CardHolder: f.CardHolder,
		ExpMonth:   uint32(expM),
		ExpYear:    uint32(expY),
		UserId:     f.UserID,
	})
	if err != nil {
		log.Fatalf("failed add card: %v", err)
	}

	log.Printf("Card token: %s", r.CardToken)
}
