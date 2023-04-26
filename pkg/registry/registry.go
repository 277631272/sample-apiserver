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

package registry

import (
	"fmt"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
)

// REST implements a RESTStorage for API services against etcd
type REST struct {
	//// 实现了rest.Storage
	//	type Storage interface {
	//	// 当请求的数据存放到该方法创建的对象之后，可以调用Create/Update进行持久化
	//	// 必须返回一个适用于 Codec.DecodeInto([]byte, runtime.Object) 的指针类型
	//	New() runtime.Object
	//}

	//// 还实现了rest.StandardStorage
	//	type StandardStorage interface {
	//	Getter
	//	Lister
	//	CreaterUpdater
	//	GracefulDeleter
	//	CollectionDeleter
	//	Watcher
	//}
	*genericregistry.Store
}

// RESTInPeace is just a simple function that panics on error.
// Otherwise returns the given storage object. It is meant to be
// a wrapper for wardle registries.
func RESTInPeace(storage *REST, err error) *REST {
	if err != nil {
		err = fmt.Errorf("unable to create REST storage for a resource due to %v, will die", err)
		panic(err)
	}
	return storage
}

//func (rest *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
//	fmt.Printf("create....: %#v\n", obj)
//	return rest.Store.Create(ctx, obj, createValidation, options)
//	//return nil, errors.NewMethodNotSupported(wardle.Resource("bars"), "CREATE")
//}
