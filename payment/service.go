package payment

import (
	"bwastartup/user"
	midtrans "github.com/veritrans/go-midtrans"
	"strconv"
)

type service struct {
}

type Service interface {
	GetPaymentUrl(Transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentUrl(Transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-c0k4TjHEgWcPFIhArfwCbSZY"
	midclient.ClientKey = "SB-Mid-client-7_xbLFRodEzvtkL8"
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{CustomerDetail: &midtrans.CustDetail{
		Email: user.Email,
		FName: user.Name,
	},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(Transaction.ID),
			GrossAmt: int64(Transaction.Amount),
		},
	}
	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil
}
