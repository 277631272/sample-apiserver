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
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"

	"git.woa.com/richardgu/sample-apisvc/pkg/apis/wardle"
)

// NewStrategy creates and returns a flunderStrategy instance
func NewStrategy(typer runtime.ObjectTyper) barStrategy {
	return barStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, and error in case the given runtime.Object is not a Flunder
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*wardle.Bar)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Bar")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

// MatchFlunder is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchBar(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *wardle.Bar) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type barStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (barStrategy) NamespaceScoped() bool {
	return true
}

func (barStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (barStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (barStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	//bar := obj.(*wardle.Bar)
	//return validation.ValidateBar(bar)
	return field.ErrorList{}
}

// WarningsOnCreate returns warnings for the creation of the given object.
func (barStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string { return nil }

func (barStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (barStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (barStrategy) Canonicalize(obj runtime.Object) {
}

func (barStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// WarningsOnUpdate returns warnings for the given update.
func (barStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}