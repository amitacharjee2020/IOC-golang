/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package watch

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/alibaba/ioc-golang/autowire/util"
	"github.com/alibaba/ioc-golang/debug/api/ioc_golang/debug"
	"github.com/alibaba/ioc-golang/debug/interceptor/common"

	"github.com/stretchr/testify/assert"
)

func TestWatchInterceptor(t *testing.T) {
	watchInterceptor := GetWatchInterceptor()
	sdid := util.GetSDIDByStructPtr(&common.ServiceFoo{})
	methodName := "Invoke"
	sendCh := make(chan *debug.WatchResponse, 10)
	controlCh := make(chan *debug.WatchResponse, 10)
	go func() {
		info := <-sendCh
		controlCh <- info
	}()
	watchInterceptor.Watch(&Context{
		SDID:       sdid,
		MethodName: methodName,
		Ch:         sendCh,
	})

	service := &common.ServiceFoo{}
	ctx := context.Background()
	param := &common.RequestParam{
		User: &common.User{
			Name: "laurence",
		},
	}

	watchInterceptor.BeforeInvoke(sdid, methodName,
		[]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(param)})
	rsp, err := service.Invoke(ctx, param)
	watchInterceptor.AfterInvoke(sdid, methodName,
		[]reflect.Value{reflect.ValueOf(rsp), reflect.ValueOf(err)})
	info := <-controlCh
	assert.Equal(t, sdid, info.Sdid)
	assert.Equal(t, "Invoke", info.MethodName)
}

func TestWatchInterceptorWithCondition(t *testing.T) {
	watchInterceptor := GetWatchInterceptor()
	sdid := util.GetSDIDByStructPtr(&common.ServiceFoo{})
	methodName := "Invoke"
	sendCh := make(chan *debug.WatchResponse, 10)
	controlCh := make(chan *debug.WatchResponse, 10)
	go func() {
		for {
			info := <-sendCh
			controlCh <- info
		}
	}()
	watchCtx := &Context{
		SDID:       sdid,
		MethodName: methodName,
		Ch:         sendCh,
		FieldMatcher: &common.FieldMatcher{
			FieldIndex: 1,
			MatchRule:  "User.Name=lizhixin",
		},
	}
	watchInterceptor.Watch(watchCtx)

	service := &common.ServiceFoo{}
	ctx := context.Background()

	// not match
	param := &common.RequestParam{
		User: &common.User{
			Name: "laurence",
		},
	}
	watchInterceptor.BeforeInvoke(sdid, methodName,
		[]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(param)})
	rsp, err := service.Invoke(ctx, param)
	info := &debug.WatchResponse{}
	time.Sleep(time.Millisecond * 500)
	watchInterceptor.AfterInvoke(sdid, methodName,
		[]reflect.Value{reflect.ValueOf(rsp), reflect.ValueOf(err)})
	time.Sleep(time.Millisecond * 500)
	select {
	case info = <-controlCh:
	default:
	}
	assert.Equal(t, "", info.Sdid)

	// match
	param.User.Name = "lizhixin"
	watchInterceptor.BeforeInvoke(sdid, methodName,
		[]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(param)})
	rsp, err = service.Invoke(ctx, param)
	time.Sleep(time.Millisecond * 500)
	watchInterceptor.AfterInvoke(sdid, methodName,
		[]reflect.Value{reflect.ValueOf(rsp), reflect.ValueOf(err)})
	time.Sleep(time.Millisecond * 500)
	info = &debug.WatchResponse{}
	select {
	case info = <-controlCh:
	default:
	}
	assert.Equal(t, util.GetSDIDByStructPtr(&common.ServiceFoo{}), info.Sdid)

	// not watch
	param.User.Name = "lizhixin"
	watchInterceptor.UnWatch(watchCtx)
	watchInterceptor.BeforeInvoke(sdid, methodName,
		[]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(param)})
	_, _ = service.Invoke(ctx, param)
	watchInterceptor.AfterInvoke(sdid, methodName,
		[]reflect.Value{reflect.ValueOf(rsp), reflect.ValueOf(err)})
	time.Sleep(time.Millisecond * 500)
	info = &debug.WatchResponse{}
	select {
	case info = <-controlCh:
	default:
	}
	assert.Equal(t, "", info.Sdid)
}
