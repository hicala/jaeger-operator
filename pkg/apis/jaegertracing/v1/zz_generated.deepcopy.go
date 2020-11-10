// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScaleSpec) DeepCopyInto(out *AutoScaleSpec) {
	*out = *in
	if in.Autoscale != nil {
		in, out := &in.Autoscale, &out.Autoscale
		*out = new(bool)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScaleSpec.
func (in *AutoScaleSpec) DeepCopy() *AutoScaleSpec {
	if in == nil {
		return nil
	}
	out := new(AutoScaleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchSpec) DeepCopyInto(out *ElasticsearchSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Storage.DeepCopyInto(&out.Storage)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchSpec.
func (in *ElasticsearchSpec) DeepCopy() *ElasticsearchSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FreeForm) DeepCopyInto(out *FreeForm) {
	*out = *in
	if in.json != nil {
		in, out := &in.json, &out.json
		*out = new([]byte)
		if **in != nil {
			in, out := *in, *out
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FreeForm.
func (in *FreeForm) DeepCopy() *FreeForm {
	if in == nil {
		return nil
	}
	out := new(FreeForm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jaeger) DeepCopyInto(out *Jaeger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jaeger.
func (in *Jaeger) DeepCopy() *Jaeger {
	if in == nil {
		return nil
	}
	out := new(Jaeger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Jaeger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerAgentSpec) DeepCopyInto(out *JaegerAgentSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	in.Options.DeepCopyInto(&out.Options)
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	in.Config.DeepCopyInto(&out.Config)
	if in.SidecarSecurityContext != nil {
		in, out := &in.SidecarSecurityContext, &out.SidecarSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.HostNetwork != nil {
		in, out := &in.HostNetwork, &out.HostNetwork
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerAgentSpec.
func (in *JaegerAgentSpec) DeepCopy() *JaegerAgentSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerAllInOneSpec) DeepCopyInto(out *JaegerAllInOneSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	in.Config.DeepCopyInto(&out.Config)
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	if in.TracingEnabled != nil {
		in, out := &in.TracingEnabled, &out.TracingEnabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerAllInOneSpec.
func (in *JaegerAllInOneSpec) DeepCopy() *JaegerAllInOneSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerAllInOneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerCassandraCreateSchemaSpec) DeepCopyInto(out *JaegerCassandraCreateSchemaSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerCassandraCreateSchemaSpec.
func (in *JaegerCassandraCreateSchemaSpec) DeepCopy() *JaegerCassandraCreateSchemaSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerCassandraCreateSchemaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerCollectorSpec) DeepCopyInto(out *JaegerCollectorSpec) {
	*out = *in
	in.AutoScaleSpec.DeepCopyInto(&out.AutoScaleSpec)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Options.DeepCopyInto(&out.Options)
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerCollectorSpec.
func (in *JaegerCollectorSpec) DeepCopy() *JaegerCollectorSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerCollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerCommonSpec) DeepCopyInto(out *JaegerCommonSpec) {
	*out = *in
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerCommonSpec.
func (in *JaegerCommonSpec) DeepCopy() *JaegerCommonSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerCommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerDependenciesSpec) DeepCopyInto(out *JaegerDependenciesSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.SuccessfulJobsHistoryLimit != nil {
		in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.ElasticsearchClientNodeOnly != nil {
		in, out := &in.ElasticsearchClientNodeOnly, &out.ElasticsearchClientNodeOnly
		*out = new(bool)
		**out = **in
	}
	if in.ElasticsearchNodesWanOnly != nil {
		in, out := &in.ElasticsearchNodesWanOnly, &out.ElasticsearchNodesWanOnly
		*out = new(bool)
		**out = **in
	}
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerDependenciesSpec.
func (in *JaegerDependenciesSpec) DeepCopy() *JaegerDependenciesSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerDependenciesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerEsIndexCleanerSpec) DeepCopyInto(out *JaegerEsIndexCleanerSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.NumberOfDays != nil {
		in, out := &in.NumberOfDays, &out.NumberOfDays
		*out = new(int)
		**out = **in
	}
	if in.SuccessfulJobsHistoryLimit != nil {
		in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerEsIndexCleanerSpec.
func (in *JaegerEsIndexCleanerSpec) DeepCopy() *JaegerEsIndexCleanerSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerEsIndexCleanerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerEsRolloverSpec) DeepCopyInto(out *JaegerEsRolloverSpec) {
	*out = *in
	if in.SuccessfulJobsHistoryLimit != nil {
		in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerEsRolloverSpec.
func (in *JaegerEsRolloverSpec) DeepCopy() *JaegerEsRolloverSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerEsRolloverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerIngesterSpec) DeepCopyInto(out *JaegerIngesterSpec) {
	*out = *in
	in.AutoScaleSpec.DeepCopyInto(&out.AutoScaleSpec)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Options.DeepCopyInto(&out.Options)
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	in.Config.DeepCopyInto(&out.Config)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerIngesterSpec.
func (in *JaegerIngesterSpec) DeepCopy() *JaegerIngesterSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerIngesterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerIngressOpenShiftSpec) DeepCopyInto(out *JaegerIngressOpenShiftSpec) {
	*out = *in
	if in.SkipLogout != nil {
		in, out := &in.SkipLogout, &out.SkipLogout
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerIngressOpenShiftSpec.
func (in *JaegerIngressOpenShiftSpec) DeepCopy() *JaegerIngressOpenShiftSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerIngressOpenShiftSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerIngressSpec) DeepCopyInto(out *JaegerIngressSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.Openshift.DeepCopyInto(&out.Openshift)
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]JaegerIngressTLSSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerIngressSpec.
func (in *JaegerIngressSpec) DeepCopy() *JaegerIngressSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerIngressTLSSpec) DeepCopyInto(out *JaegerIngressTLSSpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerIngressTLSSpec.
func (in *JaegerIngressTLSSpec) DeepCopy() *JaegerIngressTLSSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerIngressTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerList) DeepCopyInto(out *JaegerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Jaeger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerList.
func (in *JaegerList) DeepCopy() *JaegerList {
	if in == nil {
		return nil
	}
	out := new(JaegerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JaegerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerQuerySpec) DeepCopyInto(out *JaegerQuerySpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Options.DeepCopyInto(&out.Options)
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	if in.TracingEnabled != nil {
		in, out := &in.TracingEnabled, &out.TracingEnabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerQuerySpec.
func (in *JaegerQuerySpec) DeepCopy() *JaegerQuerySpec {
	if in == nil {
		return nil
	}
	out := new(JaegerQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerSamplingSpec) DeepCopyInto(out *JaegerSamplingSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerSamplingSpec.
func (in *JaegerSamplingSpec) DeepCopy() *JaegerSamplingSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerSamplingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerSpec) DeepCopyInto(out *JaegerSpec) {
	*out = *in
	in.AllInOne.DeepCopyInto(&out.AllInOne)
	in.Query.DeepCopyInto(&out.Query)
	in.Collector.DeepCopyInto(&out.Collector)
	in.Ingester.DeepCopyInto(&out.Ingester)
	in.Agent.DeepCopyInto(&out.Agent)
	in.UI.DeepCopyInto(&out.UI)
	in.Sampling.DeepCopyInto(&out.Sampling)
	in.Storage.DeepCopyInto(&out.Storage)
	in.Ingress.DeepCopyInto(&out.Ingress)
	in.JaegerCommonSpec.DeepCopyInto(&out.JaegerCommonSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerSpec.
func (in *JaegerSpec) DeepCopy() *JaegerSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerStatus) DeepCopyInto(out *JaegerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerStatus.
func (in *JaegerStatus) DeepCopy() *JaegerStatus {
	if in == nil {
		return nil
	}
	out := new(JaegerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerStorageSpec) DeepCopyInto(out *JaegerStorageSpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	in.CassandraCreateSchema.DeepCopyInto(&out.CassandraCreateSchema)
	in.Dependencies.DeepCopyInto(&out.Dependencies)
	in.EsIndexCleaner.DeepCopyInto(&out.EsIndexCleaner)
	in.EsRollover.DeepCopyInto(&out.EsRollover)
	in.Elasticsearch.DeepCopyInto(&out.Elasticsearch)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerStorageSpec.
func (in *JaegerStorageSpec) DeepCopy() *JaegerStorageSpec {
	if in == nil {
		return nil
	}
	out := new(JaegerStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerUISpec) DeepCopyInto(out *JaegerUISpec) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerUISpec.
func (in *JaegerUISpec) DeepCopy() *JaegerUISpec {
	if in == nil {
		return nil
	}
	out := new(JaegerUISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Options) DeepCopyInto(out *Options) {
	*out = *in
	if in.opts != nil {
		in, out := &in.opts, &out.opts
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.json != nil {
		in, out := &in.json, &out.json
		*out = new([]byte)
		if **in != nil {
			in, out := *in, *out
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Options.
func (in *Options) DeepCopy() *Options {
	if in == nil {
		return nil
	}
	out := new(Options)
	in.DeepCopyInto(out)
	return out
}
