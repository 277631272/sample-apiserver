//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	wardle "git.woa.com/richardgu/sample-apisvc/pkg/apis/wardle"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Flunder)(nil), (*wardle.Flunder)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Flunder_To_wardle_Flunder(a.(*Flunder), b.(*wardle.Flunder), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*wardle.Flunder)(nil), (*Flunder)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_wardle_Flunder_To_v1beta1_Flunder(a.(*wardle.Flunder), b.(*Flunder), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FlunderList)(nil), (*wardle.FlunderList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_FlunderList_To_wardle_FlunderList(a.(*FlunderList), b.(*wardle.FlunderList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*wardle.FlunderList)(nil), (*FlunderList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_wardle_FlunderList_To_v1beta1_FlunderList(a.(*wardle.FlunderList), b.(*FlunderList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FlunderSpec)(nil), (*wardle.FlunderSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_FlunderSpec_To_wardle_FlunderSpec(a.(*FlunderSpec), b.(*wardle.FlunderSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*wardle.FlunderSpec)(nil), (*FlunderSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_wardle_FlunderSpec_To_v1beta1_FlunderSpec(a.(*wardle.FlunderSpec), b.(*FlunderSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*FlunderStatus)(nil), (*wardle.FlunderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_FlunderStatus_To_wardle_FlunderStatus(a.(*FlunderStatus), b.(*wardle.FlunderStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*wardle.FlunderStatus)(nil), (*FlunderStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_wardle_FlunderStatus_To_v1beta1_FlunderStatus(a.(*wardle.FlunderStatus), b.(*FlunderStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_Flunder_To_wardle_Flunder(in *Flunder, out *wardle.Flunder, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_FlunderSpec_To_wardle_FlunderSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_FlunderStatus_To_wardle_FlunderStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Flunder_To_wardle_Flunder is an autogenerated conversion function.
func Convert_v1beta1_Flunder_To_wardle_Flunder(in *Flunder, out *wardle.Flunder, s conversion.Scope) error {
	return autoConvert_v1beta1_Flunder_To_wardle_Flunder(in, out, s)
}

func autoConvert_wardle_Flunder_To_v1beta1_Flunder(in *wardle.Flunder, out *Flunder, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_wardle_FlunderSpec_To_v1beta1_FlunderSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_wardle_FlunderStatus_To_v1beta1_FlunderStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_wardle_Flunder_To_v1beta1_Flunder is an autogenerated conversion function.
func Convert_wardle_Flunder_To_v1beta1_Flunder(in *wardle.Flunder, out *Flunder, s conversion.Scope) error {
	return autoConvert_wardle_Flunder_To_v1beta1_Flunder(in, out, s)
}

func autoConvert_v1beta1_FlunderList_To_wardle_FlunderList(in *FlunderList, out *wardle.FlunderList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]wardle.Flunder)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_FlunderList_To_wardle_FlunderList is an autogenerated conversion function.
func Convert_v1beta1_FlunderList_To_wardle_FlunderList(in *FlunderList, out *wardle.FlunderList, s conversion.Scope) error {
	return autoConvert_v1beta1_FlunderList_To_wardle_FlunderList(in, out, s)
}

func autoConvert_wardle_FlunderList_To_v1beta1_FlunderList(in *wardle.FlunderList, out *FlunderList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Flunder)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_wardle_FlunderList_To_v1beta1_FlunderList is an autogenerated conversion function.
func Convert_wardle_FlunderList_To_v1beta1_FlunderList(in *wardle.FlunderList, out *FlunderList, s conversion.Scope) error {
	return autoConvert_wardle_FlunderList_To_v1beta1_FlunderList(in, out, s)
}

func autoConvert_v1beta1_FlunderSpec_To_wardle_FlunderSpec(in *FlunderSpec, out *wardle.FlunderSpec, s conversion.Scope) error {
	out.FlunderReference = in.FlunderReference
	out.FischerReference = in.FischerReference
	out.ReferenceType = wardle.ReferenceType(in.ReferenceType)
	return nil
}

// Convert_v1beta1_FlunderSpec_To_wardle_FlunderSpec is an autogenerated conversion function.
func Convert_v1beta1_FlunderSpec_To_wardle_FlunderSpec(in *FlunderSpec, out *wardle.FlunderSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_FlunderSpec_To_wardle_FlunderSpec(in, out, s)
}

func autoConvert_wardle_FlunderSpec_To_v1beta1_FlunderSpec(in *wardle.FlunderSpec, out *FlunderSpec, s conversion.Scope) error {
	out.FlunderReference = in.FlunderReference
	out.FischerReference = in.FischerReference
	out.ReferenceType = ReferenceType(in.ReferenceType)
	return nil
}

// Convert_wardle_FlunderSpec_To_v1beta1_FlunderSpec is an autogenerated conversion function.
func Convert_wardle_FlunderSpec_To_v1beta1_FlunderSpec(in *wardle.FlunderSpec, out *FlunderSpec, s conversion.Scope) error {
	return autoConvert_wardle_FlunderSpec_To_v1beta1_FlunderSpec(in, out, s)
}

func autoConvert_v1beta1_FlunderStatus_To_wardle_FlunderStatus(in *FlunderStatus, out *wardle.FlunderStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1beta1_FlunderStatus_To_wardle_FlunderStatus is an autogenerated conversion function.
func Convert_v1beta1_FlunderStatus_To_wardle_FlunderStatus(in *FlunderStatus, out *wardle.FlunderStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_FlunderStatus_To_wardle_FlunderStatus(in, out, s)
}

func autoConvert_wardle_FlunderStatus_To_v1beta1_FlunderStatus(in *wardle.FlunderStatus, out *FlunderStatus, s conversion.Scope) error {
	return nil
}

// Convert_wardle_FlunderStatus_To_v1beta1_FlunderStatus is an autogenerated conversion function.
func Convert_wardle_FlunderStatus_To_v1beta1_FlunderStatus(in *wardle.FlunderStatus, out *FlunderStatus, s conversion.Scope) error {
	return autoConvert_wardle_FlunderStatus_To_v1beta1_FlunderStatus(in, out, s)
}
