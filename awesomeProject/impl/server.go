package impl

import (
	"context"
	"os"
	"strconv"
	"time"

	"gitlab.okta-solutions.com/mashroom/backend/verification"

	"gitlab.okta-solutions.com/mashroom/backend/common/errs"
	"gitlab.okta-solutions.com/mashroom/backend/verification/verifiers"
	_ "gitlab.okta-solutions.com/mashroom/backend/verification/verifiers/fields"

	"gitlab.okta-solutions.com/mashroom/backend/account"
	"gitlab.okta-solutions.com/mashroom/backend/auth"
	"gitlab.okta-solutions.com/mashroom/backend/common/log"
	"gitlab.okta-solutions.com/mashroom/backend/property"
	"google.golang.org/grpc"
)

const (
	AuthServiceEnv         = "AUTH_SERVICE_URL"
	AuthServiceDefault     = "auth:10000"
	AccountServiceEnv      = "ACCOUNT_SERVICE_URL"
	AccountServiceDefault  = "account:10000"
	PropertyServiceEnv     = "PROPERTY_SERVICE_URL"
	PropertyServiceDefault = "property:10000"
)

var (
	AuthService     = AuthServiceDefault
	AccountService  = AccountServiceDefault
	PropertyService = PropertyServiceDefault
)

func init() {
	if v, ok := os.LookupEnv(AuthServiceEnv); ok {
		AuthService = v
	}
	if v, ok := os.LookupEnv(AccountServiceEnv); ok {
		AccountService = v
	}
	if v, ok := os.LookupEnv(PropertyServiceEnv); ok {
		PropertyService = v
	}

	if v, ok := os.LookupEnv(IdealPostcodesApiKeyEnv); ok {
		IdealPostcodesApiKey = v
	}
	if v, ok := os.LookupEnv(IdealPostcodesQueryUrlEnv); ok {
		IdealPostcodesQueryUrl = v
	}
	if v, ok := os.LookupEnv(IdealPostcodesUrlEnv); ok {
		IdealPostcodesUrl = v
	}
	if v, ok := os.LookupEnv(IdealPostcodesTimeoutEnv); ok {
		i, err := strconv.Atoi(v)
		if err == nil {
			IdealPostcodesTimeout = i
		}
	}
}

type Server interface {
	verification.VerificationServiceServer
	Serve()
}

type serverImpl struct {
	AuthService     auth.AuthServiceClient
	AccountService  account.AccountServiceClient
	PropertyService property.PropertyServiceClient
}

func (server *serverImpl) VerifyField(ctx context.Context, field *verification.FieldDescription) (*verification.VerifyFieldResult, error) {
	if field == nil {
		return nil, errs.NilRequest
	}
	return verifiers.VerifyField(server, field)
}

func (server *serverImpl) VerifyStruct(ctx context.Context, strct *verification.StructDescription) (*verification.VerifyStructResult, error) {
	if strct == nil {
		return nil, errs.NilRequest
	}
	return verifiers.VerifyStruct(server, strct)
}

func (server *serverImpl) AuthClient() auth.AuthServiceClient {
	return server.AuthService
}

func (server *serverImpl) AccountClient() account.AccountServiceClient {
	return server.AccountService
}

func (server *serverImpl) PropertyClient() property.PropertyServiceClient {
	return server.PropertyService
}

func (server *serverImpl) Serve() {

}

func NewServer() Server {
	var authService auth.AuthServiceClient
	for {
		if conn, err := grpc.Dial(AuthService, grpc.WithInsecure()); err != nil {
			log.WithError(err).Error("Failed to connect to auth service")
		} else {
			authService = auth.NewAuthServiceClient(conn)
			break
		}
		time.Sleep(1 * time.Second)
	}

	var accountService account.AccountServiceClient
	for {
		if conn, err := grpc.Dial(AccountService, grpc.WithInsecure()); err != nil {
			log.WithError(err).Error("Failed to connect to account service")
		} else {
			accountService = account.NewAccountServiceClient(conn)
			break
		}
		time.Sleep(1 * time.Second)
	}

	var propertyService property.PropertyServiceClient
	for {
		if conn, err := grpc.Dial(PropertyService, grpc.WithInsecure()); err != nil {
			log.WithError(err).Error("Failed to connect to property service")
		} else {
			propertyService = property.NewPropertyServiceClient(conn)
			break
		}
		time.Sleep(1 * time.Second)
	}

	return &serverImpl{
		AuthService:     authService,
		AccountService:  accountService,
		PropertyService: propertyService,
	}
}
