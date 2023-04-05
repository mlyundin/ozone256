package checkout

import (
	"context"
	//"database/sql"
	checkoutDomain "route256/checkout/internal/domain"
	checkoutDomainMock "route256/checkout/internal/domain/mocks"
	desc "route256/checkout/pkg/checkout"
	"route256/libs/logger"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestListCart(t *testing.T) {
	logger.Init(true)

	type cartHandlerMockFunc func(mc *minimock.Controller) checkoutDomain.CartHandler
	type productCheckerMockFunc func(mc *minimock.Controller) checkoutDomain.ProductChecker

	type args struct {
		ctx context.Context
		req *desc.ListCartRequest
	}

	var (
		mc           = minimock.NewController(t)
		ctx          = context.Background()
		userId       = gofakeit.Int64()
		n            = 10
		skus         = make([]uint32, 0, n)
		productsDesc = checkoutDomain.ProductsDesc{}
		req          = &desc.ListCartRequest{User: userId}
		cart         = &checkoutDomain.Cart{}
		resp         = &desc.ListCartResponse{}
		cartErr      = errors.New("cart error")
	)
	t.Cleanup(mc.Finish)

	for len(productsDesc) < n {
		productsDesc[gofakeit.Uint32()] = checkoutDomain.ProductDesc{
			Name:  gofakeit.FarmAnimal(),
			Price: uint32(gofakeit.Uint16()),
			Err:   nil}
	}

	for sku, item := range productsDesc {
		skus = append(skus, sku)
		count := uint16(gofakeit.Uint8())
		cart.Items = append(cart.Items, checkoutDomain.Product{
			ProductDesc: item,
			Count:       count,
			Sku:         sku,
		})
		cart.Total += item.Price * uint32(count)
	}

	for _, item := range cart.Items {
		resp.Items = append(resp.Items, &desc.Item{Sku: item.Sku, Count: uint32(item.Count), Name: item.Name, Price: item.Price})
	}
	resp.TotalPrice = cart.Total

	tests := []struct {
		name               string
		args               args
		want               *desc.ListCartResponse
		err                error
		cartHandlerMock    cartHandlerMockFunc
		productCheckerMock productCheckerMockFunc
	}{
		{
			name: "positive case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: resp,
			err:  nil,
			cartHandlerMock: func(mc *minimock.Controller) checkoutDomain.CartHandler {
				mock := checkoutDomainMock.NewCartHandlerMock(mc)
				mock.ListCartMock.Expect(ctx, userId).Return(cart, nil)
				return mock
			},

			productCheckerMock: func(mc *minimock.Controller) checkoutDomain.ProductChecker {
				mock := checkoutDomainMock.NewProductCheckerMock(mc)
				mock.ProductsMock.Expect(ctx, skus).Return(productsDesc)
				return mock
			},
		},
		{
			name: "negative case - cart error",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  cartErr,
			cartHandlerMock: func(mc *minimock.Controller) checkoutDomain.CartHandler {
				mock := checkoutDomainMock.NewCartHandlerMock(mc)
				mock.ListCartMock.Expect(ctx, userId).Return(nil, cartErr)
				return mock
			},

			productCheckerMock: func(mc *minimock.Controller) checkoutDomain.ProductChecker {
				return nil
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			api := New(checkoutDomain.NewMock(
				tt.cartHandlerMock(mc),
				tt.productCheckerMock(mc),
			))

			res, err := api.ListCart(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, res)
			if tt.err != nil {
				require.ErrorContains(t, err, tt.err.Error())
			} else {
				require.Equal(t, tt.err, err)
			}
		})
	}
}

func TestPurchase(t *testing.T) {
	type cartHandlerMockFunc func(mc *minimock.Controller) checkoutDomain.CartHandler
	type stocksCheckerMockFunc func(mc *minimock.Controller) checkoutDomain.StocksChecker

	logger.Init(true)

	type args struct {
		ctx context.Context
		req *desc.PurchaseRequest
	}

	var (
		mc           = minimock.NewController(t)
		ctx          = context.Background()
		userId       = gofakeit.Int64()
		n            = 10
		productsDesc = checkoutDomain.ProductsDesc{}
		req          = &desc.PurchaseRequest{User: userId}
		cart         = &checkoutDomain.Cart{}
		cartErr      = errors.New("cart error")
		stockErr     = errors.New("stocks error")
	)
	t.Cleanup(mc.Finish)

	for len(productsDesc) < n {
		productsDesc[gofakeit.Uint32()] = checkoutDomain.ProductDesc{
			Name:  gofakeit.FarmAnimal(),
			Price: uint32(gofakeit.Uint16()),
			Err:   nil}
	}

	for sku, item := range productsDesc {
		count := uint16(gofakeit.Uint8())
		cart.Items = append(cart.Items, checkoutDomain.Product{
			ProductDesc: item,
			Count:       count,
			Sku:         sku,
		})
		cart.Total += item.Price * uint32(count)
	}

	tests := []struct {
		name              string
		args              args
		want              *emptypb.Empty
		err               error
		cartHandlerMock   cartHandlerMockFunc
		stocksCheckerMock stocksCheckerMockFunc
	}{
		{
			name: "positive case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  nil,
			cartHandlerMock: func(mc *minimock.Controller) checkoutDomain.CartHandler {
				mock := checkoutDomainMock.NewCartHandlerMock(mc)
				mock.ListCartMock.Expect(ctx, userId).Return(cart, nil)
				return mock
			},

			stocksCheckerMock: func(mc *minimock.Controller) checkoutDomain.StocksChecker {
				mock := checkoutDomainMock.NewStocksCheckerMock(mc)
				mock.CreateOrderMock.Expect(ctx, userId, cart).Return(gofakeit.Int64(), nil)
				return mock
			},
		},
		{
			name: "negative case - cart error",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  cartErr,
			cartHandlerMock: func(mc *minimock.Controller) checkoutDomain.CartHandler {
				mock := checkoutDomainMock.NewCartHandlerMock(mc)
				mock.ListCartMock.Expect(ctx, userId).Return(nil, cartErr)
				return mock
			},
			stocksCheckerMock: func(mc *minimock.Controller) checkoutDomain.StocksChecker {
				return nil
			},
		},

		{
			name: "negative case - stocks error",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  stockErr,
			cartHandlerMock: func(mc *minimock.Controller) checkoutDomain.CartHandler {
				mock := checkoutDomainMock.NewCartHandlerMock(mc)
				mock.ListCartMock.Expect(ctx, userId).Return(cart, nil)
				return mock
			},

			stocksCheckerMock: func(mc *minimock.Controller) checkoutDomain.StocksChecker {
				mock := checkoutDomainMock.NewStocksCheckerMock(mc)
				mock.CreateOrderMock.Expect(ctx, userId, cart).Return(gofakeit.Int64(), stockErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			api := New(checkoutDomain.NewMock(
				tt.cartHandlerMock(mc),
				tt.stocksCheckerMock(mc),
			))

			res, err := api.Purchase(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, res)
			if tt.err != nil {
				require.ErrorContains(t, err, tt.err.Error())
			} else {
				require.Equal(t, tt.err, err)
			}
		})
	}
}
