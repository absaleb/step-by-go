package impl

import (
	"context"

	"gitlab.okta-solutions.com/mashroom/backend/verification"
	"gitlab.okta-solutions.com/mashroom/backend/common/errs"
)


func (server *serverImpl) VerifyPostcode(ctx context.Context, field *verification.VerifyPostcodeRequest) (*verification.VerifyPostcodeResult, error) {
	if field == nil {
		return nil, errs.NilRequest
	}
	return VerifyPostcodeImpl(field)
}

func (server *serverImpl) VerifyPostcodeQuery(ctx context.Context, field *verification.VerifyPostcodeQueryRequest) (*verification.VerifyPostcodeResult, error) {
	if field == nil {
		return nil, errs.NilRequest
	}
	return VerifyPostcodeQueryImpl(field)
}
