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
	vm "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/models"
)

type DescribeImageConstraintsBatchRequest struct {
	core.JDCloudRequest

	/* 地域ID  */
	RegionId string `json:"regionId"`

	/* 镜像ID列表 (Optional) */
	Ids []string `json:"ids"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeImageConstraintsBatchRequest(
	regionId string,
) *DescribeImageConstraintsBatchRequest {

	return &DescribeImageConstraintsBatchRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/imageConstraints",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
		RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param ids: 镜像ID列表 (Optional)
 */
func NewDescribeImageConstraintsBatchRequestWithAllParams(
	regionId string,
	ids []string,
) *DescribeImageConstraintsBatchRequest {

	return &DescribeImageConstraintsBatchRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/imageConstraints",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
		RegionId: regionId,
		Ids:      ids,
	}
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeImageConstraintsBatchRequestWithoutParam() *DescribeImageConstraintsBatchRequest {

	return &DescribeImageConstraintsBatchRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/imageConstraints",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/* param regionId: 地域ID(Required) */
func (r *DescribeImageConstraintsBatchRequest) SetRegionId(regionId string) {
	r.RegionId = regionId
}

/* param ids: 镜像ID列表(Optional) */
func (r *DescribeImageConstraintsBatchRequest) SetIds(ids []string) {
	r.Ids = ids
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeImageConstraintsBatchRequest) GetRegionId() string {
	return r.RegionId
}

type DescribeImageConstraintsBatchResponse struct {
	RequestID string                              `json:"requestId"`
	Error     core.ErrorResponse                  `json:"error"`
	Result    DescribeImageConstraintsBatchResult `json:"result"`
}

type DescribeImageConstraintsBatchResult struct {
	ImageConstraints []vm.ImageConstraint `json:"imageConstraints"`
}
