package gateway

import (
	"context"
	"net/http"
	"strconv"

	"github.com/davyc98/billing-engine/internal/billing/internal/usecase"
	"github.com/davyc98/billing-engine/internal/pkg/pkgerror"
	"github.com/davyc98/billing-engine/internal/pkg/pkghttp/v1"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func NewLoanHTTPGateway(
	httpRouter *httprouter.Router,
	logger *zap.SugaredLogger,
	loanHTTPEndpoint *LoanHTTPEndpoint,
	validator *validator.Validate,
) {
	server := pkghttp.NewServer(
		pkghttp.WithResponseEncoder(pkghttp.CodeMessageResponseEncoder),
		pkghttp.WithErrorResponseEncoder(pkghttp.CodeMessageErrorEncoder),
	)

	httpRouter.Handler(
		http.MethodGet,
		"/billing/:loan_id",
		server.Serve(loanHTTPEndpoint.GetOutstanding),
	)

	httpRouter.Handler(
		http.MethodPost,
		"/billing",
		server.Serve(loanHTTPEndpoint.MakePayment),
	)
}

type LoanHTTPEndpoint struct {
	getOutstandingUsecase usecase.GetOutstanding
	paymentUsecase        usecase.MakePayment
	validator             *validator.Validate
	logger                *zap.SugaredLogger
}

func NewLoanHTTPEndpoint(
	getOutstandingUsecase usecase.GetOutstanding,
	paymentUsecase usecase.MakePayment,
	logger *zap.SugaredLogger,
	validator *validator.Validate,

) *LoanHTTPEndpoint {
	return &LoanHTTPEndpoint{
		getOutstandingUsecase: getOutstandingUsecase,
		paymentUsecase:        paymentUsecase,
		logger:                logger,
		validator:             validator,
	}
}

func (l *LoanHTTPEndpoint) GetOutstanding(
	ctx context.Context,
	request pkghttp.Request,
) (resp any, err error) {
	var input usecase.GetOutstandingInput
	params := httprouter.ParamsFromContext(ctx)

	loanID := params.ByName("loan_id")

	input.LoanID, err = strconv.ParseUint(loanID, 10, 64)
	if err != nil {
		l.logger.Errorw("failed to parse loan id", "error", err)

		return nil, pkgerror.ValidationErrorFrom(err)
	}

	if err := l.validator.Struct(input); err != nil {
		l.logger.Errorw("failed to validate request", "error", err)

		return nil, pkgerror.ValidationErrorFrom(err)
	}

	var res *usecase.GetOustandingOutput
	if res, err = l.getOutstandingUsecase.Execute(ctx, input); err != nil {
		l.logger.Errorw("failed to get oustanding", "error", err)

		return nil, err
	}

	return res, nil
}

func (l *LoanHTTPEndpoint) MakePayment(
	ctx context.Context,
	request pkghttp.Request,
) (resp any, err error) {
	var input usecase.PaymentInput
	if err := request.Decode(&input); err != nil {
		l.logger.Errorw("failed to decode request", "error", err)

		return nil, pkgerror.ServerErrorFrom(err)
	}

	if err := l.validator.Struct(input); err != nil {
		l.logger.Errorw("failed to validate request", "error", err)

		return nil, pkgerror.ValidationErrorFrom(err)
	}

	if err = l.paymentUsecase.Execute(ctx, input); err != nil {
		l.logger.Errorw("failed to get oustanding", "error", err)

		return nil, err
	}

	return "ok", nil
}
