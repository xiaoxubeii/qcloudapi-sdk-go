package bmvpc

import (
	"github.com/dbdd4us/qcloudapi-sdk-go/common"
)

const (
	BmVpcHost = "bmvpc.api.qcloud.com"
	BmVpcPath = "/v2/index.php"
)

type Client struct {
	*common.Client
}

func NewClient(credential common.CredentialInterface, opts common.Opts) (*Client, error) {
	if opts.Host == "" {
		opts.Host = BmVpcHost
	}
	if opts.Path == "" {
		opts.Path = BmVpcPath
	}

	client, err := common.NewClient(credential, opts)
	if err != nil {
		return &Client{}, err
	}
	return &Client{client}, nil
}
