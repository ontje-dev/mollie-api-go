package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type balancesServiceSuite struct{ suite.Suite }

func (bs *balancesServiceSuite) SetupSuite() { setEnv() }

func (bs *balancesServiceSuite) TearDownSuite() { unsetEnv() }

func (bs *balancesServiceSuite) TestBalancesService_Get() {
	type args struct {
		ctx     context.Context
		balance string
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		handler http.HandlerFunc
		pre     func()
	}{
		{
			"get balance by id works as expected",
			args{
				context.Background(),
				"bal_gVMhHKqSSRYJyPsuoPNFH",
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(bs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(bs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.GetBalanceResponse))
			},
			noPre,
		},
		{
			"get balance by id returns an error from the server",
			args{
				context.Background(),
				"bal_gVMhHKqSSRYJyPsuoPNFH",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			errorHandler,
			noPre,
		},
		{
			"get balance by id returns an error when creating the request",
			args{
				context.Background(),
				"bal_gVMhHKqSSRYJyPsuoPNFH",
			},
			true,
			errBadBaseURL,
			errorHandler,
			crashSrv,
		},
		{
			"get balance by id returns an error when trying to parse the json response",
			args{
				context.Background(),
				"bal_gVMhHKqSSRYJyPsuoPNFH",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			encodingHandler,
			noPre,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		bs.T().Run(c.name, func(t *testing.T) {
			c.pre()

			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/balances/%s",
					c.args.balance,
				),
				c.handler,
			)

			res, capture, err := tClient.Balances.Get(c.args.ctx, c.args.balance)
			if c.wantErr {
				bs.NotNil(err)
				bs.EqualError(err, c.err.Error())
			} else {
				bs.Nil(err)
				bs.IsType(&Balance{}, capture)
				bs.Same(c.args.ctx, res.Request.Context())
				bs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (bs *balancesServiceSuite) TestBalancesService_Primary() {
	type args struct {
		ctx context.Context
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		handler http.HandlerFunc
		pre     func()
	}{
		{
			"get primary balance works as expected",
			args{
				context.Background(),
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(bs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(bs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.GetBalanceResponse))
			},
			noPre,
		},
		{
			"get primary balance returns an error from the server",
			args{
				context.Background(),
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			errorHandler,
			noPre,
		},
		{
			"get primary balance returns an error when creating the request",
			args{
				context.Background(),
			},
			true,
			errBadBaseURL,
			errorHandler,
			crashSrv,
		},
		{
			"get primary balance returns an error when trying to parse the json response",
			args{
				context.Background(),
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			encodingHandler,
			noPre,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		bs.T().Run(c.name, func(t *testing.T) {
			c.pre()

			tMux.HandleFunc(
				"/v2/balances/primary",
				c.handler,
			)

			res, capture, err := tClient.Balances.Primary(c.args.ctx)
			if c.wantErr {
				bs.NotNil(err)
				bs.EqualError(err, c.err.Error())
			} else {
				bs.Nil(err)
				bs.IsType(&Balance{}, capture)
				bs.Same(c.args.ctx, res.Request.Context())
				bs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (bs *balancesServiceSuite) TestBalancesService_List() {
	type args struct {
		ctx     context.Context
		options *BalanceListOptions
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		handler http.HandlerFunc
		pre     func()
	}{
		{
			"list balances works as expected",
			args{
				context.Background(),
				nil,
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(bs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(bs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.ListBalancesResponse))
			},
			noPre,
		},
		{
			"list balances returns an error from the server",
			args{
				context.Background(),
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			errorHandler,
			noPre,
		},
		{
			"list balances returns an error when creating the request",
			args{
				context.Background(),
				nil,
			},
			true,
			errBadBaseURL,
			errorHandler,
			crashSrv,
		},
		{
			"list balances returns an error when trying to parse the json response",
			args{
				context.Background(),
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			encodingHandler,
			noPre,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		bs.T().Run(c.name, func(t *testing.T) {
			c.pre()

			tMux.HandleFunc(
				"/v2/balances",
				c.handler,
			)

			res, capture, err := tClient.Balances.List(c.args.ctx, c.args.options)
			if c.wantErr {
				bs.NotNil(err)
				bs.EqualError(err, c.err.Error())
			} else {
				bs.Nil(err)
				bs.IsType(&BalancesList{}, capture)
				bs.Same(c.args.ctx, res.Request.Context())
				bs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (bs *balancesServiceSuite) TestBalancesService_GetReport() {
	type args struct {
		ctx     context.Context
		balance string
		options *BalanceReportOptions
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		handler http.HandlerFunc
		pre     func()
	}{
		{
			"get balance report works as expected",
			args{
				context.Background(),
				"bal_gVMhHKqSSRYJyPsuoPNFH",
				nil,
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(bs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(bs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.GetBalanceReportResponse))
			},
			noPre,
		},
		{
			"get balance report returns an error from the server",
			args{
				context.Background(),
				"bal_gVMhHKqSSRYJyPsuoPNFH",
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			errorHandler,
			noPre,
		},
		{
			"get balance report returns an error when creating the request",
			args{
				context.Background(),
				"bal_gVMhHKqSSRYJyPsuoPNFH",
				nil,
			},
			true,
			errBadBaseURL,
			errorHandler,
			crashSrv,
		},
		{
			"get balance report returns an error when trying to parse the json response",
			args{
				context.Background(),
				"bal_gVMhHKqSSRYJyPsuoPNFH",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			encodingHandler,
			noPre,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		bs.T().Run(c.name, func(t *testing.T) {
			c.pre()

			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/balances/%s/report",
					"bal_gVMhHKqSSRYJyPsuoPNFH",
				),
				c.handler,
			)

			res, balanceReport, err := tClient.Balances.GetReport(c.args.ctx, c.args.balance, c.args.options)
			if c.wantErr {
				bs.NotNil(err)
				bs.EqualError(err, c.err.Error())
			} else {
				bs.Nil(err)
				bs.IsType(&BalanceReport{}, balanceReport)
				bs.Same(c.args.ctx, res.Request.Context())
				bs.IsType(&http.Response{}, res.Response)
				fmt.Printf("\n%+v\n", *balanceReport)
			}
		})
	}
}

func (bs *balancesServiceSuite) TestBalancesService_GetPrimaryReport() {
	type args struct {
		ctx     context.Context
		options *BalanceReportOptions
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		handler http.HandlerFunc
		pre     func()
	}{
		{
			"get primary balance report works as expected",
			args{
				context.Background(),

				nil,
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(bs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(bs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.GetBalanceReportResponse))
			},
			noPre,
		},
		{
			"get primary balance report returns an error from the server",
			args{
				context.Background(),

				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			errorHandler,
			noPre,
		},
		{
			"get primary balance report returns an error when creating the request",
			args{
				context.Background(),

				nil,
			},
			true,
			errBadBaseURL,
			errorHandler,
			crashSrv,
		},
		{
			"get primary balance report returns an error when trying to parse the json response",
			args{
				context.Background(),

				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			encodingHandler,
			noPre,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		bs.T().Run(c.name, func(t *testing.T) {
			c.pre()

			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/balances/%s/report",
					"primary",
				),
				c.handler,
			)

			res, balanceReport, err := tClient.Balances.GetPrimaryReport(c.args.ctx, c.args.options)
			if c.wantErr {
				bs.NotNil(err)
				bs.EqualError(err, c.err.Error())
			} else {
				bs.Nil(err)
				bs.IsType(&BalanceReport{}, balanceReport)
				bs.Same(c.args.ctx, res.Request.Context())
				bs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (bs *balancesServiceSuite) TestBalancesService_ListBalanceTransactions() {
	type args struct {
		ctx     context.Context
		balance string
		options *BalanceTransactionsListOptions
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		handler http.HandlerFunc
		pre     func()
	}{
		{
			"list balance transactions works as expected",
			args{
				context.Background(),
				"bal_gVMhHKqSSRYJyPsuoPNFH",
				nil,
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(bs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(bs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.ListBalanceTransactionsResponse))
			},
			noPre,
		},
		{
			"list balance transactions returns an error from the server",
			args{
				context.Background(),
				"bal_gVMhHKqSSRYJyPsuoPNFH",
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			errorHandler,
			noPre,
		},
		{
			"list balance transactions returns an error when creating the request",
			args{
				context.Background(),
				"bal_gVMhHKqSSRYJyPsuoPNFH",
				nil,
			},
			true,
			errBadBaseURL,
			errorHandler,
			crashSrv,
		},
		{
			"list balance transactions returns an error when trying to parse the json response",
			args{
				context.Background(),
				"bal_gVMhHKqSSRYJyPsuoPNFH",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			encodingHandler,
			noPre,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		bs.T().Run(c.name, func(t *testing.T) {
			c.pre()

			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/balances/%s/transactions",
					c.args.balance,
				),
				c.handler,
			)

			res, balanceReport, err := tClient.Balances.GetTransactionsList(c.args.ctx, c.args.balance, c.args.options)
			if c.wantErr {
				bs.NotNil(err)
				bs.EqualError(err, c.err.Error())
			} else {
				bs.Nil(err)
				bs.IsType(&BalanceTransactionsList{}, balanceReport)
				bs.Same(c.args.ctx, res.Request.Context())
				bs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (bs *balancesServiceSuite) TestBalancesService_ListPrimaryBalanceTransactions() {
	type args struct {
		ctx     context.Context
		options *BalanceTransactionsListOptions
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		handler http.HandlerFunc
		pre     func()
	}{
		{
			"list primary balance transactions works as expected",
			args{
				context.Background(),
				nil,
			},
			false,
			nil,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(bs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(bs.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				_, _ = w.Write([]byte(testdata.ListBalanceTransactionsResponse))
			},
			noPre,
		},
		{
			"list primary balance transactions returns an error from the server",
			args{
				context.Background(),
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			errorHandler,
			noPre,
		},
		{
			"list primary balance transactions returns an error when creating the request",
			args{
				context.Background(),
				nil,
			},
			true,
			errBadBaseURL,
			errorHandler,
			crashSrv,
		},
		{
			"list primary balance transactions returns an error when trying to parse the json response",
			args{
				context.Background(),
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			encodingHandler,
			noPre,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		bs.T().Run(c.name, func(t *testing.T) {
			c.pre()

			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/balances/%s/transactions",
					"primary",
				),
				c.handler,
			)

			res, balanceReport, err := tClient.Balances.GetPrimaryTransactionsList(c.args.ctx, c.args.options)
			if c.wantErr {
				bs.NotNil(err)
				bs.EqualError(err, c.err.Error())
			} else {
				bs.Nil(err)
				bs.IsType(&BalanceTransactionsList{}, balanceReport)
				bs.Same(c.args.ctx, res.Request.Context())
				bs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestBalancesService(t *testing.T) {
	suite.Run(t, new(balancesServiceSuite))
}
