package util

import (
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	vpc20160428 "github.com/alibabacloud-go/vpc-20160428/v6/client"
	"os"
)

func createVpc20160428Client() (_result *vpc20160428.Client, _err error) {
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Vpc
	config.Endpoint = tea.String(os.Getenv("VPC_ENDPOINT"))
	_result = &vpc20160428.Client{}
	_result, _err = vpc20160428.NewClient(config)
	return _result, _err
}

func AllocateEip() (string, string, error) {
	client, err := createVpc20160428Client()
	if err != nil {
		return "", "", errors.New("allocateEip failed, err:" + err.Error())
	}
	allocateEipAddressRequest := &vpc20160428.AllocateEipAddressRequest{
		RegionId:           tea.String(os.Getenv("REGION_ID")),
		Bandwidth:          tea.String("100"),
		AutoPay:            tea.Bool(true),
		InstanceChargeType: tea.String("PostPaid"),
		InternetChargeType: tea.String("PayByTraffic"),
		Name:               tea.String("eproxy"),
		Description:        tea.String("eproxy"),
	}
	runtime := &util.RuntimeOptions{}
	result, err := client.AllocateEipAddressWithOptions(allocateEipAddressRequest, runtime)
	if err != nil {
		return "", "", err
	}
	body := result.Body
	if body == nil || body.AllocationId == nil || body.EipAddress == nil {
		return "", "", err
	}
	return *(body.AllocationId), *(body.EipAddress), nil
}

func ReleaseEip(allocationId string) (err error) {
	client, err := createVpc20160428Client()
	if err != nil {
		return errors.New("ReleaseEip Failed, err:" + err.Error())
	}

	releaseEipAddressRequest := &vpc20160428.ReleaseEipAddressRequest{
		RegionId:     tea.String(os.Getenv("REGION_ID")),
		AllocationId: tea.String(allocationId),
	}
	runtime := &util.RuntimeOptions{}
	_, err = client.ReleaseEipAddressWithOptions(releaseEipAddressRequest, runtime)
	if err != nil {
		return err
	}
	return
}

func AssociateEip(allocationId string, instanceId string) error {
	client, err := createVpc20160428Client()
	if err != nil {
		return errors.New("allocateEip failed, err:" + err.Error())
	}
	associateEipAddressRequest := &vpc20160428.AssociateEipAddressRequest{
		RegionId:     tea.String(os.Getenv("REGION_ID")),
		AllocationId: tea.String(allocationId),
		InstanceId:   tea.String(instanceId),
	}
	runtime := &util.RuntimeOptions{}
	_, err = client.AssociateEipAddressWithOptions(associateEipAddressRequest, runtime)
	if err != nil {
		return err
	}
	return nil
}

func UnassociateEip(allocationId string, instanceId string) error {
	client, err := createVpc20160428Client()
	if err != nil {
		return errors.New("unassociateEip failed, err:" + err.Error())
	}
	unassociateEipAddressRequest := &vpc20160428.UnassociateEipAddressRequest{
		RegionId:     tea.String(os.Getenv("REGION_ID")),
		AllocationId: tea.String(allocationId),
		InstanceId:   tea.String(instanceId),
	}
	runtime := &util.RuntimeOptions{}
	_, err = client.UnassociateEipAddressWithOptions(unassociateEipAddressRequest, runtime)
	if err != nil {
		return err
	}
	return nil
}
