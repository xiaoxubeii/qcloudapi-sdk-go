package bmvpc

import (
	"github.com/dbdd4us/qcloudapi-sdk-go/common"
	"time"
)

const (
	TASK_STATE_SUCCESS = 0
	TASK_STATE_FAILED  = 1
	TASK_STATE_DOING   = 2
	TASK_STATE_UNKNOWN = 3
	TASK_STATE_TIMEOUT = 4

	TASK_QUERY_INTERVAL = 1
)

type DescribeBmSubnetRequest struct {
	UnVpcId    *string `qcloud_arg:"unVpcId,omitempty"`
	UnSubnetId *string `qcloud_arg:"unSubnetId,omitempty"`
	SubnetName *string `qcloud_arg:"subnetName,omitempty"`
	VlanId     *int    `qcloud_arg:"vlanId,omitempty"`
	Limit      *int    `qcloud_arg:"limit,omitempty"`
	Offset     *int    `qcloud_arg:"offset,omitempty"`
}

type BmSubnetDetail struct {
	VpcId            int    `json:"vpcId"`
	UnVpcId          string `json:"unVpcId"`
	SubnetId         int    `json:"subnetId"`
	UnSubnetId       string `json:"unSubnetId"`
	SubnetName       string `json:"subnetName"`
	CidrBlock        string `json:"cidrBlock"`
	ZoneId           string `json:"zoneId"`
	VlanId           string `json:"vlanId"`
	DhcpEnable       int    `json:"dhcpEnable"`
	IpReserved       int    `json:"ipReserve"`
	DistributeedFlag int    `json:"distributedFlag"`
}

type BmDescribeSubnetResponse struct {
	TotalCount    int              `json:"totalCount"`
	SubnetDetails []BmSubnetDetail `json:"data"`
}

//查询子网列表：https://cloud.tencent.com/document/product/386/6648
func (client *Client) DescribeBmSubnetEx(req *DescribeBmSubnetRequest) (*BmDescribeSubnetResponse, error) {
	rsp := &BmDescribeSubnetResponse{}
	err := client.Invoke("DescribeBmSubnetEx", req, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

type DescribeBmCpmRequest struct {
	UnVpcId    string `qcloud_arg:"vpcId"`
	UnSubnetId string `qcloud_arg:"subnetId"`
}

type CpmInfo struct {
	InstanceId string `json:"instanceId"`
}

type DescribeBmCpmResponse struct {
	CpmSet []CpmInfo `json:"cpmSet"`
}

//根据子网subnetID和vpcId，查询子网中的主机instanceID
//https://cloud.tencent.com/document/product/386/9319
func (client *Client) DescribeBmCpmBySubnetId(req *DescribeBmCpmRequest) (*[]CpmInfo, error) {
	rsp := &DescribeBmCpmResponse{}
	err := client.Invoke("DescribeBmCpmBySubnetId", req, rsp)
	if err != nil {
		return nil, err
	}
	return &rsp.CpmSet, nil
}

type BmSubnetCreateParam struct {
	SubnetName      string `qcloud_arg:"subnetName"`
	CidrBlock       string `qcloud_arg:"cidrBlock"`
	DistributedFlag *int   `qcloud_arg:"distributedFlag"`
}

type CreateBmSubnetRequest struct {
	UnVpcId   string                `qcloud_arg"unVpcId"`
	VLanId    *int                  `qcloud_arg:"vlanId"`
	SubnetSet []BmSubnetCreateParam `Aqcloud_arg:"subnetSet"`
}

type BmSubnetInfo struct {
	SubnetId   string `json:"subnetId"`
	UnSubnetId string `json:"unSubnetId"`
	SubnetName string `json:"subnetName"`
	CidrBlock  string `json:"cidrBlock"`
}

type CreateBmSubnetResponse struct {
	SubnetSet []BmSubnetInfo `json:"subnetSet"`
}

//创建子网：https://cloud.tencent.com/document/product/386/9263
func (client *Client) CreateBmSubnet(req *CreateBmSubnetRequest) (*[]BmSubnetInfo, error) {
	rsp := &CreateBmSubnetResponse{}
	err := client.Invoke("CreateBmSubnetd", req, rsp)
	if err != nil {
		return nil, err
	}
	return &rsp.SubnetSet, nil
}

//物理机加入和移除的时候，都使用下面这两个数据结构
type CreateBmInterfaceRequest struct {
	UnVpcId     string   `qcloud_arg:"unVpcId"`
	UnSubnetId  string   `qcloud_arg:"unSubnetId"`
	InstanceIds []string `qcloud_arg:"instanceIds"`
}

type BmVpcTask struct {
	TaskId      int      `json:"taskId"`
	ResourceIds []string `json:"instanceIds"`
}

//将物理机添加到子网：https://cloud.tencent.com/document/product/386/9265
func (client *Client) CreateBmInterface(req *CreateBmInterfaceRequest) (int, error) {
	bmVpcTask := &BmVpcTask{}
	rsp := &common.DataResponse{
		Response: bmVpcTask,
	}

	err := client.Invoke("CreateBmInterface", req, rsp)
	if err != nil {
		return 0, err
	}

	return bmVpcTask.TaskId, nil
}

type DelBmInterfaceRequest CreateBmInterfaceRequest

//物理机中移除子网：https://cloud.tencent.com/document/product/386/9266
func (client *Client) DelBmInterface(req *DelBmInterfaceRequest) (int, error) {
	bmVpcTask := &BmVpcTask{}
	rsp := &common.DataResponse{
		Response: bmVpcTask,
	}

	err := client.Invoke("DelBmInterface", req, rsp)
	if err != nil {
		return 0, err
	}

	return bmVpcTask.TaskId, nil
}

type DeleteBmSubnetRequest struct {
	UnVpcId    string `qcloud_arg:"unVpcId"`
	UnSubnetId string `qcloud_arg:"unSubnetId"`
}

//https://cloud.tencent.com/document/product/386/9264
func (client *Client) DeleteBmSubnet(req *DeleteBmSubnetRequest) error {
	rsp := &common.Response{}
	err := client.Invoke("DeleteBmSubnet", req, rsp)
	if err != nil {
		return 0, err
	}
	if rsp.Code == 0 {
		return nil
	} else {
		return error.New(rsp.Message)
	}
}

type BmVpcQueryTaskRequest struct {
	TaskId int `qcloud_arg:"taskId"`
}

type BmVpcTaskStatus struct {
	Status int `json:"status"`
}

//https://cloud.tencent.com/document/product/386/9267
func (client *Client) QueryBmTaskResult(taskId int) (int, error) {
	req := BmVpcQueryTaskRequest{
		TaskId: taskId,
	}

	taskStatus := &BmVpcTaskStatus{}
	rsp := &common.DataResponse{
		Response: taskStatus,
	}

	err := client.Invoke("QueryBmTaskResult", req, rsp)
	if err != nil {
		return common.TASK_STATE_UNKNOWN, err
	}
	return taskStatus.Status, nil
}

func (client *Client) WaitUntiTaskDone(taskId int, timeout int) error {

	count := 0
	for {
		time.Sleep(TASK_QUERY_INTERVAL * time.Second)
		count++

		state := client.QueryBmTaskResult(taskId)
		if state == TASK_STATE_SUCCESS {
			return nil
		} else if state == TASK_STATE_FAILED {
			return error.New("bmVpc waitUntilTaskDone task failed")
		}

		if count*TASK_QUERY_INTERVAL < timeout {
			continue
		} else {
			return error.New("bmVpc waitUntilTaskDone task timeout")
		}
	}
}
