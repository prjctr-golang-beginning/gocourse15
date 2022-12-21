package main

import (
	"context"
	"github.com/maximorov/auditor"
	"pubsub/pubsub"
)

type myRepository struct {
}

func (r *myRepository) CreateMany(ctx context.Context, enities []auditor.Valuable) (int, error) {
	return 3, nil
}

func main() {
	ctx := context.Background()
	p := pubsub.NewPublisher(ctx)
}
