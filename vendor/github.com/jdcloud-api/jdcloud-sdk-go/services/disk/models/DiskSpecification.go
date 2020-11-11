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

package models

type DiskSpecification struct {

	/* 云硬盘类型 (Optional) */
	DiskType string `json:"diskType"`

	/* 支持的最小尺寸，单位为 GiB (Optional) */
	MinSizeGB int `json:"minSizeGB"`

	/* 支持的最大尺寸，单位为 GiB (Optional) */
	MaxSizeGB int `json:"maxSizeGB"`

	/* 步长尺寸，单位为 GiB (Optional) */
	StepSizeGB int `json:"stepSizeGB"`

	/* 描述信息 (Optional) */
	Description string `json:"description"`

	/* 默认的iops数量(基础iops数量) (Optional) */
	DefaultIOPS int `json:"defaultIOPS"`

	/* iops步长增量 (Optional) */
	StepIOPS float32 `json:"stepIOPS"`

	/* 最大iops数量 (Optional) */
	MaxIOPS int `json:"maxIOPS"`

	/* 默认的吞吐量 (Optional) */
	DefaultThroughput int `json:"defaultThroughput"`

	/* 吞吐量步长增量 (Optional) */
	StepThroughput float32 `json:"stepThroughput"`

	/* 最大吞吐量 (Optional) */
	MaxThroughput int `json:"maxThroughput"`
}
