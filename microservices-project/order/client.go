package order

import (
	"context"
	"microservices-project/order/pb"
	"time"

	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.OrderServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewOrderServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostOrder(ctx context.Context, accountID string, products []OrderedProduct) (*Order, error) {
	protoProducts := []*pb.PostOrderRequest_OrderProduct{}
	for _, p := range products {
		protoProducts = append(protoProducts, &pb.PostOrderRequest_OrderProduct{
			ProductId: p.ID,
			Quantity:  p.Quantity,
		})
	}
	r, err := c.service.PostOrder(ctx, &pb.PostOrderRequest{
		AccountId: accountID,
		Products:  protoProducts,
	})
	if err != nil {
		return nil, err
	}
	o := r.Order
	newOrder := &Order{
		ID:         o.Id,
		CreatedAt:  time.Time{},
		TotalPrice: o.TotalPrice,
		AccountID:  accountID,
	}
	for _, p := range o.Products {
		newOrder.Products = append(newOrder.Products, OrderedProduct{
			ID:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Quantity:    p.Quantity,
		})
	}
	return newOrder, nil
}

func (c *Client) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {
	r, err := c.service.GetOrdersForAccount(ctx, &pb.GetOrdersForAccountRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}
	var orders []Order
	for _, o := range r.Orders {
		newOrder := Order{
			ID:         o.Id,
			TotalPrice: o.TotalPrice,
			AccountID:  accountID,
		}
		for _, p := range o.Products {
			newOrder.Products = append(newOrder.Products, OrderedProduct{
				ID:          p.Id,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Quantity:    p.Quantity,
			})
		}
		orders = append(orders, newOrder)
	}
	return orders, nil
}
