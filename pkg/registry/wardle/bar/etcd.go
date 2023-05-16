/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package bar

import (
	"git.woa.com/richardgu/sample-apisvc/pkg/apis/wardle"
	"git.woa.com/richardgu/sample-apisvc/pkg/registry"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

// NewREST returns a RESTStorage object that will work against API services.
func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)

	store := &genericregistry.Store{
		// 实例化资源的函数
		// 每次增删改查，都牵涉到结构的创建，因此在此打断点可以拦截所有请求
		NewFunc: func() runtime.Object { return &wardle.Bar{} },
		// 实例化资源列表的函数
		NewListFunc: func() runtime.Object { return &wardle.BarList{} },
		// 判断对象是否可以被该存储处理
		PredicateFunc: MatchBar,
		// 资源的复数名称，当上下文中缺少必要的请求信息时使用
		DefaultQualifiedResource:  wardle.Resource("bars"),
		SingularQualifiedResource: wardle.Resource("bar"),

		// 增删改的策略
		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		// TODO: define table converter that exposes more than name/creation timestamp
		TableConvertor: rest.NewDefaultTableConvertor(wardle.Resource("bars")),
	}
	// 填充默认字段
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &registry.REST{store}, nil
}