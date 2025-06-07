package loan

import (
	"database/sql"

	"github.com/davyc98/billing-engine/internal/billing/internal/gateway"
	"github.com/davyc98/billing-engine/internal/billing/internal/interactor"
	"github.com/davyc98/billing-engine/internal/pkg/pkgsql"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type Exposed struct {
}

type Dependencies struct {
	DB           *sql.DB
	Logger       *zap.SugaredLogger
	QueryBuilder pkgsql.GoquBuilder
	HttpRouter   *httprouter.Router
	Validator    *validator.Validate
}

func New(deps Dependencies) *Exposed {
	loanSQLstore := gateway.NewLoanSQLGateway(deps.DB, deps.Logger, deps.QueryBuilder)
	getOutstandingUsecase := interactor.NewGetOustandingLoan(
		loanSQLstore,
		deps.Logger,
	)

	loanHTTPEndpoint := gateway.NewLoanHTTPEndpoint(
		getOutstandingUsecase,
		deps.Logger,
		deps.Validator,
	)

	gateway.NewLoanHTTPGateway(deps.HttpRouter, deps.Logger, loanHTTPEndpoint, deps.Validator)

	return &Exposed{}
}
