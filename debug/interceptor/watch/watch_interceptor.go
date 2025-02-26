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
	"reflect"
	"sync"

	"github.com/petermattis/goid"

	"github.com/alibaba/ioc-golang/debug/interceptor/common"

	"github.com/alibaba/ioc-golang/debug/api/ioc_golang/debug"
)

type Interceptor struct {
	watch sync.Map
}

func (w *Interceptor) BeforeInvoke(sdid, methodName string, values []reflect.Value) []reflect.Value {
	if watchCtxInterface, ok := w.watch.Load(common.GetMethodUniqueKey(sdid, methodName)); ok {
		watchCtxInterface.(*Context).BeforeInvoke(values)
	}
	return values
}

func (w *Interceptor) AfterInvoke(sdid, methodName string, values []reflect.Value) []reflect.Value {
	if watchCtxInterface, ok := w.watch.Load(common.GetMethodUniqueKey(sdid, methodName)); ok {
		watchCtxInterface.(*Context).AfterInvoke(values)
	}
	return values
}

func (w *Interceptor) Watch(watchCtx *Context) {
	w.watch.Store(common.GetMethodUniqueKey(watchCtx.SDID, watchCtx.MethodName), watchCtx)
}

func (w *Interceptor) UnWatch(watchCtx *Context) {
	w.watch.Delete(common.GetMethodUniqueKey(watchCtx.SDID, watchCtx.MethodName))
}

type Context struct {
	SDID              string
	MethodName        string
	Ch                chan *debug.WatchResponse
	FieldMatcher      *common.FieldMatcher
	watchGRRequestMap sync.Map
}

func (w *Context) BeforeInvoke(params []reflect.Value) {
	if w.FieldMatcher != nil && !w.FieldMatcher.Match(params) {
		// doesn't match
		return
	}
	grID := goid.Get()
	w.watchGRRequestMap.Store(grID, params)
}

func (w *Context) AfterInvoke(returnValues []reflect.Value) {
	grID := goid.Get()
	paramValues, ok := w.watchGRRequestMap.Load(grID)
	if !ok {
		return
	}
	invokeDetail := &debug.WatchResponse{
		Sdid:         w.SDID,
		MethodName:   w.MethodName,
		Params:       common.ReflectValues2Strings(paramValues.([]reflect.Value)),
		ReturnValues: common.ReflectValues2Strings(returnValues),
	}
	w.Ch <- invokeDetail
}

var watchInterceptorSingleton *Interceptor

func GetWatchInterceptor() *Interceptor {
	if watchInterceptorSingleton == nil {
		watchInterceptorSingleton = &Interceptor{}
	}
	return watchInterceptorSingleton
}
