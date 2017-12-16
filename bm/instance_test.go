package bm

import (
	"github.com/dbdd4us/qcloudapi-sdk-go/common"
	"testing"
)

func TestDescribeDevice(t *testing.T) {

	client, _ := NewClientFromEnv()

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
