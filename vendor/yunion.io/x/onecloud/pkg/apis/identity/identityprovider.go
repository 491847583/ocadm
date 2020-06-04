// Copyright 2019 Yunion
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

package identity

import "yunion.io/x/onecloud/pkg/apis"

type IdentityProviderDetails struct {
	apis.StandaloneResourceDetails

	SyncIntervalSeconds int    `json:"sync_interval_seconds"`
	TargetDomain        string `json:"target_domain"`

	RoleCount    int `json:"role_count,allowempty"`
	UserCount    int `json:"user_count,allowempty"`
	PolicyCount  int `json:"policy_count,allowempty"`
	DomainCount  int `json:"domain_count,allowempty"`
	ProjectCount int `json:"project_count,allowempty"`
	GroupCount   int `json:"group_count,allowempty"`
}

type IdentityProviderCreateInput struct {
	apis.EnabledStatusStandaloneResourceCreateInput

	// 后端驱动名称
	Driver string `json:"driver"`

	// 模板名称
	Template string `json:"template"`

	// 默认导入用户和组的域
	TargetDomain string `json:"target_domain"`
	// swagger:ignore
	// Deprecated
	TargetDomainId string `json:"target_domain_id" "yunion:deprecated-by":"target_domain"`

	// 新建域的时候是否自动新建第一个项目
	AutoCreateProject *bool `json:"auto_create_project"`

	// 自动同步间隔，单位：秒
	SyncIntervalSeconds *int `json:"sync_interval_seconds"`

	// 配置信息
	Config TConfigs `json:"config"`
}
