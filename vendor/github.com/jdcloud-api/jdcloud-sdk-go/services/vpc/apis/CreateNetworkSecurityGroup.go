// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type CreateNetworkSecurityGroupRequest struct {
	core.JDCloudRequest

	/* Region ID  */
	RegionId string `json:"regionId"`

	/* 私有网络ID  */
	VpcId string `json:"vpcId"`

	/* 安全组名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。  */
	NetworkSecurityGroupName string `json:"networkSecurityGroupName"`

	/* 描述,​ 允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
	Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param vpcId: 私有网络ID (Required)
 * param networkSecurityGroupName: 安全组名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateNetworkSecurityGroupRequest(
	regionId string,
	vpcId string,
	networkSecurityGroupName string,
) *CreateNetworkSecurityGroupRequest {

	return &CreateNetworkSecurityGroupRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkSecurityGroups/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
		RegionId:                 regionId,
		VpcId:                    vpcId,
		NetworkSecurityGroupName: networkSecurityGroupName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param vpcId: 私有网络ID (Required)
 * param networkSecurityGroupName: 安全组名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Required)
 * param description: 描述,​ 允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 */
func NewCreateNetworkSecurityGroupRequestWithAllParams(
	regionId string,
	vpcId string,
	networkSecurityGroupName string,
	description *string,
) *CreateNetworkSecurityGroupRequest {

	return &CreateNetworkSecurityGroupRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkSecurityGroups/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
		RegionId:                 regionId,
		VpcId:                    vpcId,
		NetworkSecurityGroupName: networkSecurityGroupName,
		Description:              description,
	}
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateNetworkSecurityGroupRequestWithoutParam() *CreateNetworkSecurityGroupRequest {

	return &CreateNetworkSecurityGroupRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkSecurityGroups/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/* param regionId: Region ID(Required) */
func (r *CreateNetworkSecurityGroupRequest) SetRegionId(regionId string) {
	r.RegionId = regionId
}

/* param vpcId: 私有网络ID(Required) */
func (r *CreateNetworkSecurityGroupRequest) SetVpcId(vpcId string) {
	r.VpcId = vpcId
}

/* param networkSecurityGroupName: 安全组名称，只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。(Required) */
func (r *CreateNetworkSecurityGroupRequest) SetNetworkSecurityGroupName(networkSecurityGroupName string) {
	r.NetworkSecurityGroupName = networkSecurityGroupName
}

/* param description: 描述,​ 允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *CreateNetworkSecurityGroupRequest) SetDescription(description string) {
	r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateNetworkSecurityGroupRequest) GetRegionId() string {
	return r.RegionId
}

type CreateNetworkSecurityGroupResponse struct {
	RequestID string                           `json:"requestId"`
	Error     core.ErrorResponse               `json:"error"`
	Result    CreateNetworkSecurityGroupResult `json:"result"`
}

type CreateNetworkSecurityGroupResult struct {
	NetworkSecurityGroupId string `json:"networkSecurityGroupId"`
}
