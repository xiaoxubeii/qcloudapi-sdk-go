package bm

import (
	"github.com/dbdd4us/qcloudapi-sdk-go/common"
	"testing"
)

func TestDescribeDevice(t *testing.T) {
	invokeOpts := common.Opts{
		Region: "ap-guangzhou",
	}

	credential := common.Credential{
		SecretId:  "AKID52SKw5uMEy3jhpMUBqSylEBJBby6E0KC",
		SecretKey: "CIuaIXhppO3ZWGLUmVH7GYbgMJ1UAV2E",
	}

	client, _ := NewClient(credential, invokeOpts)

	lanIps := []string{"10.0.0.4"}
	req := DescribeDeviceArgs{
		LanIps: &lanIps,
	}

	if devInfo, err := client.DescribeDevice(&req); err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("DescribeDevice Pass devInfo=%v", devInfo)
	}

}
