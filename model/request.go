package model

type RunRequest struct {
	Amount uint `json:"amount" form:"amount" binding:"required"`
}

type IdRequest struct {
	InstanceId   string `json:"instanceId" form:"instanceId" binding:"required"`
	AllocationId string `json:"allocationId" form:"allocationId" binding:"required"`
}
