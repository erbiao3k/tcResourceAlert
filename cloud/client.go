package cloud

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"tcResourceAlert/config"
)

var Credential = common.NewCredential(
	config.SecretId,
	config.SecretKey,
)

func NewCfg(cpfEndpoint string, region string) (*common.Credential, string, *profile.ClientProfile) {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = cpfEndpoint

	return Credential, region, cpf
}
