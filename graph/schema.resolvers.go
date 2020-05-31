package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oryono/banking-go/graph/generated"
	"github.com/oryono/banking-go/models"
)

func (r *queryResolver) Clients(ctx context.Context) ([]*models.Client, error) {
	var clients []*models.Client
	r.DB.Find(&clients)
	return clients, nil
}

func (r *queryResolver) Customers(ctx context.Context) ([]*models.Customer, error) {
	var customers []*models.Customer
	r.DB.Preload("Client").Find(&customers)
	return customers, nil
}

func (r *queryResolver) Entries(ctx context.Context) ([]*models.Entry, error) {
	var entries []*models.Entry
	r.DB.Find(&entries)
	return entries, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
