package bm

import (
	"errors"
	"time"
)

const (
	TASK_QUERY_INTERVAL = 2
)

type BmContainerSubnetIpReq struct {
	UnVPcId    string `qcloud_arg:"unVpcId"`
	UnSubnetId string `qcloud_arg:"unSubnetId"`
}

type BmTask struct {
	TaskId int `json:"taskId"`
}

//herb新加接口，注册容器网段RegisterContainerSubnetIp
func (bmClient *Client) RegisterContainerSubnetIp(unVpcId, unSubnetId string) (int, error) {
	task := &BmTask{}
	rsp := &BmResponse{
		Response: task,
	}

	req := &BmContainerSubnetIpReq{
		UnVPcId:    unVpcId,
		UnSubnetId: unSubnetId,
	}

	err := bmClient.Invoke("RegisterContainerSubnetIp", req, rsp)
	if err != nil {
		return 0, err
	}

	return task.TaskId, nil
}

//herb新加接口，退还容器网段ReleaseContainerSubnetIp
func (bmClient *Client) ReleaseContainerSubnetIp(unVpcId, unSubnetId string) (int, error) {
	task := &BmTask{}
	rsp := &BmResponse{
		Response: task,
	}
	req := &BmContainerSubnetIpReq{
		UnVPcId:    unVpcId,
		UnSubnetId: unSubnetId,
	}

	err := bmClient.Invoke("ReleaseContainerSubnetIp", req, rsp)
	if err != nil {
		return 0, err
	}

	return task.TaskId, nil
}

type BmQueryTaskResultRequest struct {
	TaskId int `qcloud_arg:"taskId"`
}

/*
1：成功
2：失败
3：部分成功，部分失败
4：未完成
5：部分成功，部分未完成
6：部分未完成，部分失败
7：部分未完成，部分失败，部分成功
*/
type BmQueryTaskStatus struct {
	Status int `json:"status"`
}

//查询异步接口DescriptionOperationResult
func (bmClient *Client) DescriptionOperationResult(taskId int) (int, error) {
	req := &BmQueryTaskResultRequest{
		TaskId: taskId,
	}
	status := &BmQueryTaskStatus{}
	rsp := &BmResponse{
		Response: status,
	}

	err := bmClient.Invoke("DescriptionOperationResult", req, rsp)
	if err != nil {
		return 0, err
	}
	return status.Status, nil

}

func (client *Client) WaitUntiTaskDone(taskId int, timeout int) error {
	count := 0
	for {
		time.Sleep(TASK_QUERY_INTERVAL * time.Second)
		count++

		state, err := client.DescriptionOperationResult(taskId)
		if err != nil {
			return err
		}
		if state == 2 || state == 3 || state == 5 || state == 6 || state == 7 {
			return errors.New("bm waitUntilTaskDone task failed")
		} else if state == 1 {
			return nil
		}

		if count*TASK_QUERY_INTERVAL < timeout {
			continue
		} else {
			return errors.New("bm waitUntilTaskDone task timeout")
		}
	}
}
