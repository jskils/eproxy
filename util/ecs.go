package util

import (
	"errors"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"os"
	"time"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// CreateClient 创建阿里云Client
func create20140526Client() (_result *ecs20140526.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")),
		AccessKeySecret: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")),
	}
	config.Endpoint = tea.String(os.Getenv("ECS_ENDPOINT"))
	_result = &ecs20140526.Client{}
	_result, _err = ecs20140526.NewClient(config)
	return _result, _err
}

func RunInstances(amount uint) ([]*string, error) {
	ecsClient, err := create20140526Client()
	if err != nil {
		return nil, errors.New("AliyunEcs.RunInstances Failed, err:" + err.Error())
	}
	now := time.Now()
	formattedTime := now.Format("20060102150405")

	runInstancesRequest := &ecs20140526.RunInstancesRequest{
		RegionId:           tea.String(os.Getenv("REGION_ID")),
		LaunchTemplateName: tea.String(os.Getenv("ECS_TEMPLATE_NAME")),
		Amount:             tea.Int32(int32(amount)),
		MinAmount:          tea.Int32(int32(amount)),
		InstanceName:       tea.String("eproxy"),
		Description:        tea.String(fmt.Sprintf("eproxy%s", formattedTime)),
	}
	runtime := &util.RuntimeOptions{}
	response, err := ecsClient.RunInstancesWithOptions(runInstancesRequest, runtime)
	if err != nil {
		return nil, err
	}
	instanceIdSet := response.Body.InstanceIdSets.InstanceIdSet
	if len(instanceIdSet) == 0 {
		return nil, errors.New("AliyunEcs.RunInstances Failed")
	}
	return instanceIdSet, nil
}

func DeleteInstance(instanceId string) (err error) {
	ecsClient, err := create20140526Client()
	if err != nil {
		return errors.New("AliyunEcs.DeleteInstance Failed, err:" + err.Error())
	}
	deleteInstanceRequest := &ecs20140526.DeleteInstanceRequest{
		InstanceId: tea.String(instanceId),
		Force:      tea.Bool(true),
	}
	runtime := &util.RuntimeOptions{}
	_, err = ecsClient.DeleteInstanceWithOptions(deleteInstanceRequest, runtime)
	if err != nil {
		return err
	}
	return
}
