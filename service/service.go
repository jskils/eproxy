package service

import (
	"eproxy/model"
	"eproxy/util"
	"fmt"
	"os"
	"time"
)

func Run(amount uint) ([]model.ResultResponse, error) {
	instanceIds, err := util.RunInstances(amount)
	if err != nil {
		return []model.ResultResponse{}, err
	}
	time.Sleep(10 * time.Second)
	var response []model.ResultResponse
	for _, instanceId := range instanceIds {
		allocationId, eip, err := util.AllocateEip()
		if err != nil {
			return []model.ResultResponse{}, err
		}
		err = util.AssociateEip(allocationId, *instanceId)
		if err != nil {
			return []model.ResultResponse{}, err
		}
		response = append(response, model.ResultResponse{
			InstanceId:   *instanceId,
			AllocationId: allocationId,
			EIP:          eip,
			Proxy:        fmt.Sprintf("%s:%s", eip, os.Getenv("PROXY_PORT")),
		})
	}
	return response, nil
}

func Change(request model.IdRequest) (model.ResultResponse, error) {
	instanceId := request.InstanceId
	oldAllocationId := request.AllocationId
	err := util.UnassociateEip(oldAllocationId, instanceId)
	if err != nil {
		return model.ResultResponse{}, err
	}
	newAllocationId, eip, err := util.AllocateEip()
	if err != nil {
		return model.ResultResponse{}, err
	}
	err = util.ReleaseEip(oldAllocationId)
	if err != nil {
		return model.ResultResponse{}, err
	}
	err = util.AssociateEip(newAllocationId, instanceId)
	if err != nil {
		return model.ResultResponse{}, err
	}
	return model.ResultResponse{
		InstanceId:   instanceId,
		AllocationId: newAllocationId,
		EIP:          eip,
		Proxy:        fmt.Sprintf("%s:%s", eip, os.Getenv("PROXY_PORT")),
	}, nil
}

func Release(request model.IdRequest) (model.IdResponse, error) {
	instanceId := request.InstanceId
	allocationId := request.AllocationId
	_ = util.UnassociateEip(allocationId, instanceId)
	ecsErr := util.DeleteInstance(instanceId)
	eipErr := util.ReleaseEip(allocationId)
	response := model.IdResponse{}
	if ecsErr == nil {
		response.InstanceId = instanceId
	}
	if eipErr == nil {
		response.AllocationId = allocationId
	}
	return response, nil
}
