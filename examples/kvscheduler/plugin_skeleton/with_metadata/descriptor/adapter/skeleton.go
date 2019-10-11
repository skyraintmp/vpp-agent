// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/golang/protobuf/proto"
	. "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/examples/kvscheduler/plugin_skeleton/with_metadata/model"
	"github.com/ligato/vpp-agent/examples/kvscheduler/plugin_skeleton/with_metadata/metaidx"
)

////////// type-safe key-value pair with metadata //////////

type SkeletonKVWithMetadata struct {
	Key      string
	Value    *model.ValueSkeleton
	Metadata *metaidx.SkeletonMetadata
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type SkeletonDescriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *model.ValueSkeleton) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *model.ValueSkeleton) error
	Create               func(key string, value *model.ValueSkeleton) (metadata *metaidx.SkeletonMetadata, err error)
	Delete               func(key string, value *model.ValueSkeleton, metadata *metaidx.SkeletonMetadata) error
	Update               func(key string, oldValue, newValue *model.ValueSkeleton, oldMetadata *metaidx.SkeletonMetadata) (newMetadata *metaidx.SkeletonMetadata, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *model.ValueSkeleton, metadata *metaidx.SkeletonMetadata) bool
	Retrieve             func(correlate []SkeletonKVWithMetadata) ([]SkeletonKVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *model.ValueSkeleton) []KeyValuePair
	Dependencies         func(key string, value *model.ValueSkeleton) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type SkeletonDescriptorAdapter struct {
	descriptor *SkeletonDescriptor
}

func NewSkeletonDescriptor(typedDescriptor *SkeletonDescriptor) *KVDescriptor {
	adapter := &SkeletonDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:                 typedDescriptor.Name,
		KeySelector:          typedDescriptor.KeySelector,
		ValueTypeName:        typedDescriptor.ValueTypeName,
		KeyLabel:             typedDescriptor.KeyLabel,
		NBKeyPrefix:          typedDescriptor.NBKeyPrefix,
		WithMetadata:         typedDescriptor.WithMetadata,
		MetadataMapFactory:   typedDescriptor.MetadataMapFactory,
		IsRetriableFailure:   typedDescriptor.IsRetriableFailure,
		RetrieveDependencies: typedDescriptor.RetrieveDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Validate != nil {
		descriptor.Validate = adapter.Validate
	}
	if typedDescriptor.Create != nil {
		descriptor.Create = adapter.Create
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.UpdateWithRecreate != nil {
		descriptor.UpdateWithRecreate = adapter.UpdateWithRecreate
	}
	if typedDescriptor.Retrieve != nil {
		descriptor.Retrieve = adapter.Retrieve
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	return descriptor
}

func (da *SkeletonDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castSkeletonValue(key, oldValue)
	typedNewValue, err2 := castSkeletonValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *SkeletonDescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castSkeletonValue(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *SkeletonDescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castSkeletonValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *SkeletonDescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castSkeletonValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castSkeletonValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castSkeletonMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *SkeletonDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castSkeletonValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castSkeletonMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *SkeletonDescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castSkeletonValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castSkeletonValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castSkeletonMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *SkeletonDescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []SkeletonKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castSkeletonValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castSkeletonMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			SkeletonKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedValues, err := da.descriptor.Retrieve(correlateWithType)
	if err != nil {
		return nil, err
	}
	var values []KVWithMetadata
	for _, typedKVWithMetadata := range typedValues {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		values = append(values, kvWithMetadata)
	}
	return values, err
}

func (da *SkeletonDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castSkeletonValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *SkeletonDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castSkeletonValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castSkeletonValue(key string, value proto.Message) (*model.ValueSkeleton, error) {
	typedValue, ok := value.(*model.ValueSkeleton)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castSkeletonMetadata(key string, metadata Metadata) (*metaidx.SkeletonMetadata, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(*metaidx.SkeletonMetadata)
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
