package bmvpc

import (
	"github.com/dbdd4us/qcloudapi-sdk-go/common"
	"testing"
)


func ListSubnet(client *Client,t *testing.T) error {
	limit := 20000
	subnetName:="ccs_bm_brank"
	req := &DescribeBmSubnetRequest{
		SubnetName: &subnetName,
		Limit:      &limit,
	}
	describeBmSubnetRsp, err := client.DescribeBmSubnetEx(req)
	if err != nil {
		t.Error(err.Error())
		return  err
	}else {
		t.Logf("describeBmSubnetEx ok rsp=%v\n",describeBmSubnetRsp)
		return nil 
	}
}


func tesBmInterface(client *Client,t *testing.T) {
    //将物理机加入子网
	unVpcId := "vpc-f5hfhmpo"
	unSubnetId :="subnet-2ge7vz9p"
	instanceIds := []string{"cpm-an5a9wv4"}
	req := &CreateBmInterfaceRequest {
		UnVpcId:     unVpcId,
		UnSubnetId:  unSubnetId,
		InstanceIds: instanceIds,
	}

	taskId, err := client.CreateBmInterface(req)
	if err != nil{
		t.Error(err.Error())
		return 
	}

	err = client.WaitUntiTaskDone(taskId,60)
	if err != nil{
		t.Error(err.Error())
		return 
	}else{
		t.Logf("createBmInterface ok")
	}

	//查询物理机列表
	describeBmReq := &DescribeBmCpmRequest{
			UnVpcId:    unVpcId,
			UnSubnetId: unSubnetId,
	}
	cpmSet, err := client.DescribeBmCpmBySubnetId(describeBmReq)
	if err != nil {
		t.Error(err.Error())
		return 
	}else{
		t.Logf("cpmSet=%v",cpmSet)
	}

	//将物理机移出子网
	delBmInterfaceReq := &DelBmInterfaceRequest{
		UnVpcId:     unVpcId,
		UnSubnetId:  unSubnetId,
		InstanceIds: instanceIds,
	}

	taskId, err = client.DelBmInterface(delBmInterfaceReq)
	if err != nil{
		t.Error(err.Error())
		return 
	}else {
		t.Log("DelBmInterface ok")
	}

    err=client.WaitUntiTaskDone(taskId,60)
	if err != nil{
		t.Error(err.Error())
		return 
	}else{
		t.Logf("DelBmInterface ok")
	}
}


func TestSubnet(t *testing.T) {
	unVpcId:="vpc-f5hfhmpo"
	subnetName:="ccs_bm_brank"
    invokeOpts := common.Opts{
		Region: "ap-guangzhou",
	}

	credential := common.Credential{
		SecretId:  "AKID52SKw5uMEy3jhpMUBqSylEBJBby6E0KC",
		SecretKey: "CIuaIXhppO3ZWGLUmVH7GYbgMJ1UAV2E",
	}

	client, _ := NewClient(credential, invokeOpts)

	//1、创建子网
	distributedFlag := 0
	vlanId := 2903
	subnetCreateParam := BmSubnetCreateParam{
		SubnetName:      subnetName,
		CidrBlock:       "10.20.61.0/24",
		DistributedFlag: &distributedFlag,
	}
	subnetSet := []BmSubnetCreateParam{subnetCreateParam}
	createSubnetReq := &CreateBmSubnetRequest{
		UnVpcId:   unVpcId,
		VLanId:    &vlanId,
		SubnetSet: subnetSet,
	}
	outputSubnetSet, err := client.CreateBmSubnet(createSubnetReq)
	if err != nil {
         t.Error(err.Error())
		 return 
	}else {
		t.Logf("CreateBmSubnet ok ouptuSubnetSet=%v",outputSubnetSet)
	}
	
	//2、查询子网
	err = ListSubnet(client,t)
	if err != nil{
		return 
	}

	tesBmInterface(client,t)

	//3、删除子网
	req := &DeleteBmSubnetRequest{
		UnVpcId:    unVpcId,
		UnSubnetId: (*outputSubnetSet)[0].UnSubnetId,
	}

	err = client.DeleteBmSubnet(req)
	if err != nil{
		 t.Error(err.Error())
		return 
	}else{
		t.Log("DleteBmSubnet OK ")
	}

	
}


