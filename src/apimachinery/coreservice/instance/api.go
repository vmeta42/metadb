/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package instance

import (
	"configcenter/src/common/mapstr"
	"context"
	"net/http"

	"configcenter/src/common/blog"
	"configcenter/src/common/errors"
	"configcenter/src/common/metadata"
	"configcenter/src/common/util"
)

func (inst *instance) CreateInstance(ctx context.Context, h http.Header, objID string, input *metadata.CreateModelInstance) (resp *metadata.CreatedOneOptionResult, err error) {
	resp = new(metadata.CreatedOneOptionResult)
	subPath := "/create/model/%s/instance"

	err = inst.client.Post().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

func (inst *instance) CreateManyInstance(ctx context.Context, h http.Header, objID string, input *metadata.CreateManyModelInstance) (resp *metadata.CreatedManyOptionResult, err error) {
	resp = new(metadata.CreatedManyOptionResult)
	subPath := "/createmany/model/%s/instance"

	err = inst.client.Post().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

func (inst *instance) InsertManyInstance(ctx context.Context, h http.Header, objID string, input *metadata.CreateManyModelInstance) (resp *metadata.DelAndCreatedManyOptionResult, err error) {
	resp = new(metadata.DelAndCreatedManyOptionResult)
	subPath := "insertmany/model/%s/instances/cache"

	err = inst.client.Post().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

func (inst *instance) SetManyInstance(ctx context.Context, h http.Header, objID string, input *metadata.SetManyModelInstance) (resp *metadata.SetOptionResult, err error) {
	resp = new(metadata.SetOptionResult)
	subPath := "/setmany/model/%s/instances"

	err = inst.client.Post().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

func (inst *instance) UpdateInstance(ctx context.Context, h http.Header, objID string, input *metadata.UpdateOption) (resp *metadata.UpdatedOptionResult, err error) {
	resp = new(metadata.UpdatedOptionResult)
	subPath := "/update/model/%s/instance"

	err = inst.client.Put().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

func (inst *instance) ReadInstance(ctx context.Context, h http.Header, objID string, input *metadata.QueryCondition) (resp *metadata.QueryConditionResult, err error) {
	resp = new(metadata.QueryConditionResult)
	subPath := "/read/model/%s/instances"

	err = inst.client.Post().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

func (inst *instance) ReadInstanceAsst(ctx context.Context, h http.Header, objID string, input *metadata.QueryAsstCondition) (resp *metadata.MapArrayResponse, err error) {
	resp = new(metadata.MapArrayResponse)
	subPath := "/read/model/%s/instances/asst"

	err = inst.client.Post().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

func (inst *instance) ReadInstanceCache(ctx context.Context, h http.Header, objID string,
	input mapstr.MapStr) (resp *metadata.QueryConditionResult, header http.Header, err error) {
	resp = new(metadata.QueryConditionResult)
	subPath := "/read/model/%s/instances/cache"

	err, header = inst.client.Post().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		IntoBodyHeader(resp)
	return
}

func (inst *instance) UpdateInstanceCache(ctx context.Context, h http.Header, objID string,
	input mapstr.MapStr) (resp *metadata.ResponseDataMapStr, header http.Header, err error) {
	resp = new(metadata.ResponseDataMapStr)
	subPath := "/update/model/%s/instance/cache"

	err, header = inst.client.Put().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		IntoBodyHeader(resp)
	return
}

func (inst *instance) UpdateInstanceUnique(ctx context.Context, h http.Header, objID string,
	input mapstr.MapStr) (resp *metadata.ResponseDataMapStr, header http.Header, err error) {
	resp = new(metadata.ResponseDataMapStr)
	subPath := "/update/model/%s/instance/unique"

	err, header = inst.client.Put().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		IntoBodyHeader(resp)
	return
}

func (inst *instance) UpdateManyInstance(ctx context.Context, h http.Header, objID string,
	input mapstr.MapStr) (resp *metadata.ResponseDataMapStr, header http.Header, err error) {
	resp = new(metadata.ResponseDataMapStr)
	subPath := "/updatemany/model/%s/instance"

	err, header = inst.client.Post().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		IntoBodyHeader(resp)
	return
}

func (inst *instance) DeleteInstance(ctx context.Context, h http.Header, objID string, input *metadata.DeleteOption) (resp *metadata.DeletedOptionResult, err error) {
	resp = new(metadata.DeletedOptionResult)
	subPath := "/delete/model/%s/instance"

	err = inst.client.Delete().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

func (inst *instance) DeleteSkipArchiveInstance(ctx context.Context, h http.Header, objID string, input *metadata.DeleteOption) (resp *metadata.DeletedOptionResult, err error) {
	resp = new(metadata.DeletedOptionResult)
	subPath := "/deleteskiparchive/model/%s/instance/cache"

	err = inst.client.Delete().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

func (inst *instance) DeleteInstanceCascade(ctx context.Context, h http.Header, objID string, input *metadata.DeleteOption) (resp *metadata.DeletedOptionResult, err error) {
	resp = new(metadata.DeletedOptionResult)
	subPath := "/delete/model/%s/instance/cascade"

	err = inst.client.Delete().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

//  ReadInstanceStruct 按照结构体返回实例数据
func (inst *instance) ReadInstanceStruct(ctx context.Context, h http.Header, objID string,
	input *metadata.QueryCondition, result interface{}) errors.CCErrorCoder {

	rid := util.GetHTTPCCRequestID(h)
	subPath := "/read/model/%s/instances"

	err := inst.client.Post().
		WithContext(ctx).
		Body(input).
		SubResourcef(subPath, objID).
		WithHeaders(h).
		Do().
		Into(result)

	if err != nil {
		blog.ErrorJSON("ReadInstanceStruct failed, http request failed, err: %s, filter: %s, rid: %s", err, input, rid)
		return errors.CCHttpError
	}

	return nil
}
