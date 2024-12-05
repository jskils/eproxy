package model

type ResultResponse struct {
	InstanceId   string `json:"instanceId"`
	AllocationId string `json:"allocationId"`
	EIP          string `json:"eip"`
	Proxy        string `json:"proxy"`
}

type IdResponse struct {
	InstanceId   string `json:"instanceId" form:"instanceId" binding:"required"`
	AllocationId string `json:"allocationId" form:"allocationId" binding:"required"`
}
