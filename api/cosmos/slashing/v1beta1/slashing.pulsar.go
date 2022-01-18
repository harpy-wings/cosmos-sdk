package slashingv1beta1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_ValidatorSigningInfo                       protoreflect.MessageDescriptor
	fd_ValidatorSigningInfo_address               protoreflect.FieldDescriptor
	fd_ValidatorSigningInfo_start_height          protoreflect.FieldDescriptor
	fd_ValidatorSigningInfo_index_offset          protoreflect.FieldDescriptor
	fd_ValidatorSigningInfo_jailed_until          protoreflect.FieldDescriptor
	fd_ValidatorSigningInfo_tombstoned            protoreflect.FieldDescriptor
	fd_ValidatorSigningInfo_missed_blocks_counter protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_slashing_v1beta1_slashing_proto_init()
	md_ValidatorSigningInfo = File_cosmos_slashing_v1beta1_slashing_proto.Messages().ByName("ValidatorSigningInfo")
	fd_ValidatorSigningInfo_address = md_ValidatorSigningInfo.Fields().ByName("address")
	fd_ValidatorSigningInfo_start_height = md_ValidatorSigningInfo.Fields().ByName("start_height")
	fd_ValidatorSigningInfo_index_offset = md_ValidatorSigningInfo.Fields().ByName("index_offset")
	fd_ValidatorSigningInfo_jailed_until = md_ValidatorSigningInfo.Fields().ByName("jailed_until")
	fd_ValidatorSigningInfo_tombstoned = md_ValidatorSigningInfo.Fields().ByName("tombstoned")
	fd_ValidatorSigningInfo_missed_blocks_counter = md_ValidatorSigningInfo.Fields().ByName("missed_blocks_counter")
}

var _ protoreflect.Message = (*fastReflection_ValidatorSigningInfo)(nil)

type fastReflection_ValidatorSigningInfo ValidatorSigningInfo

func (x *ValidatorSigningInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorSigningInfo)(x)
}

func (x *ValidatorSigningInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_slashing_v1beta1_slashing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorSigningInfo_messageType fastReflection_ValidatorSigningInfo_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorSigningInfo_messageType{}

type fastReflection_ValidatorSigningInfo_messageType struct{}

func (x fastReflection_ValidatorSigningInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorSigningInfo)(nil)
}
func (x fastReflection_ValidatorSigningInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorSigningInfo)
}
func (x fastReflection_ValidatorSigningInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorSigningInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorSigningInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorSigningInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorSigningInfo) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorSigningInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorSigningInfo) New() protoreflect.Message {
	return new(fastReflection_ValidatorSigningInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorSigningInfo) Interface() protoreflect.ProtoMessage {
	return (*ValidatorSigningInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorSigningInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_ValidatorSigningInfo_address, value) {
			return
		}
	}
	if x.StartHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.StartHeight)
		if !f(fd_ValidatorSigningInfo_start_height, value) {
			return
		}
	}
	if x.IndexOffset != int64(0) {
		value := protoreflect.ValueOfInt64(x.IndexOffset)
		if !f(fd_ValidatorSigningInfo_index_offset, value) {
			return
		}
	}
	if x.JailedUntil != nil {
		value := protoreflect.ValueOfMessage(x.JailedUntil.ProtoReflect())
		if !f(fd_ValidatorSigningInfo_jailed_until, value) {
			return
		}
	}
	if x.Tombstoned != false {
		value := protoreflect.ValueOfBool(x.Tombstoned)
		if !f(fd_ValidatorSigningInfo_tombstoned, value) {
			return
		}
	}
	if x.MissedBlocksCounter != int64(0) {
		value := protoreflect.ValueOfInt64(x.MissedBlocksCounter)
		if !f(fd_ValidatorSigningInfo_missed_blocks_counter, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ValidatorSigningInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.address":
		return x.Address != ""
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.start_height":
		return x.StartHeight != int64(0)
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.index_offset":
		return x.IndexOffset != int64(0)
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.jailed_until":
		return x.JailedUntil != nil
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.tombstoned":
		return x.Tombstoned != false
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.missed_blocks_counter":
		return x.MissedBlocksCounter != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.ValidatorSigningInfo"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.ValidatorSigningInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSigningInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.address":
		x.Address = ""
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.start_height":
		x.StartHeight = int64(0)
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.index_offset":
		x.IndexOffset = int64(0)
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.jailed_until":
		x.JailedUntil = nil
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.tombstoned":
		x.Tombstoned = false
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.missed_blocks_counter":
		x.MissedBlocksCounter = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.ValidatorSigningInfo"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.ValidatorSigningInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorSigningInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.start_height":
		value := x.StartHeight
		return protoreflect.ValueOfInt64(value)
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.index_offset":
		value := x.IndexOffset
		return protoreflect.ValueOfInt64(value)
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.jailed_until":
		value := x.JailedUntil
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.tombstoned":
		value := x.Tombstoned
		return protoreflect.ValueOfBool(value)
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.missed_blocks_counter":
		value := x.MissedBlocksCounter
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.ValidatorSigningInfo"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.ValidatorSigningInfo does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSigningInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.address":
		x.Address = value.Interface().(string)
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.start_height":
		x.StartHeight = value.Int()
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.index_offset":
		x.IndexOffset = value.Int()
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.jailed_until":
		x.JailedUntil = value.Message().Interface().(*timestamppb.Timestamp)
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.tombstoned":
		x.Tombstoned = value.Bool()
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.missed_blocks_counter":
		x.MissedBlocksCounter = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.ValidatorSigningInfo"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.ValidatorSigningInfo does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSigningInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.jailed_until":
		if x.JailedUntil == nil {
			x.JailedUntil = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.JailedUntil.ProtoReflect())
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.address":
		panic(fmt.Errorf("field address of message cosmos.slashing.v1beta1.ValidatorSigningInfo is not mutable"))
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.start_height":
		panic(fmt.Errorf("field start_height of message cosmos.slashing.v1beta1.ValidatorSigningInfo is not mutable"))
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.index_offset":
		panic(fmt.Errorf("field index_offset of message cosmos.slashing.v1beta1.ValidatorSigningInfo is not mutable"))
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.tombstoned":
		panic(fmt.Errorf("field tombstoned of message cosmos.slashing.v1beta1.ValidatorSigningInfo is not mutable"))
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.missed_blocks_counter":
		panic(fmt.Errorf("field missed_blocks_counter of message cosmos.slashing.v1beta1.ValidatorSigningInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.ValidatorSigningInfo"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.ValidatorSigningInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorSigningInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.address":
		return protoreflect.ValueOfString("")
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.start_height":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.index_offset":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.jailed_until":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.tombstoned":
		return protoreflect.ValueOfBool(false)
	case "cosmos.slashing.v1beta1.ValidatorSigningInfo.missed_blocks_counter":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.ValidatorSigningInfo"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.ValidatorSigningInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorSigningInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.slashing.v1beta1.ValidatorSigningInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorSigningInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSigningInfo) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ValidatorSigningInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorSigningInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorSigningInfo)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.StartHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.StartHeight))
		}
		if x.IndexOffset != 0 {
			n += 1 + runtime.Sov(uint64(x.IndexOffset))
		}
		if x.JailedUntil != nil {
			l = options.Size(x.JailedUntil)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Tombstoned {
			n += 2
		}
		if x.MissedBlocksCounter != 0 {
			n += 1 + runtime.Sov(uint64(x.MissedBlocksCounter))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorSigningInfo)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.MissedBlocksCounter != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MissedBlocksCounter))
			i--
			dAtA[i] = 0x30
		}
		if x.Tombstoned {
			i--
			if x.Tombstoned {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x28
		}
		if x.JailedUntil != nil {
			encoded, err := options.Marshal(x.JailedUntil)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if x.IndexOffset != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.IndexOffset))
			i--
			dAtA[i] = 0x18
		}
		if x.StartHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.StartHeight))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorSigningInfo)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorSigningInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorSigningInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
				}
				x.StartHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.StartHeight |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field IndexOffset", wireType)
				}
				x.IndexOffset = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.IndexOffset |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field JailedUntil", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.JailedUntil == nil {
					x.JailedUntil = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.JailedUntil); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tombstoned", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Tombstoned = bool(v != 0)
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MissedBlocksCounter", wireType)
				}
				x.MissedBlocksCounter = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MissedBlocksCounter |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Params                            protoreflect.MessageDescriptor
	fd_Params_signed_blocks_window       protoreflect.FieldDescriptor
	fd_Params_min_signed_per_window      protoreflect.FieldDescriptor
	fd_Params_downtime_jail_duration     protoreflect.FieldDescriptor
	fd_Params_slash_fraction_double_sign protoreflect.FieldDescriptor
	fd_Params_slash_fraction_downtime    protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_slashing_v1beta1_slashing_proto_init()
	md_Params = File_cosmos_slashing_v1beta1_slashing_proto.Messages().ByName("Params")
	fd_Params_signed_blocks_window = md_Params.Fields().ByName("signed_blocks_window")
	fd_Params_min_signed_per_window = md_Params.Fields().ByName("min_signed_per_window")
	fd_Params_downtime_jail_duration = md_Params.Fields().ByName("downtime_jail_duration")
	fd_Params_slash_fraction_double_sign = md_Params.Fields().ByName("slash_fraction_double_sign")
	fd_Params_slash_fraction_downtime = md_Params.Fields().ByName("slash_fraction_downtime")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_slashing_v1beta1_slashing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Params_messageType fastReflection_Params_messageType
var _ protoreflect.MessageType = fastReflection_Params_messageType{}

type fastReflection_Params_messageType struct{}

func (x fastReflection_Params_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Params)(nil)
}
func (x fastReflection_Params_messageType) New() protoreflect.Message {
	return new(fastReflection_Params)
}
func (x fastReflection_Params_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Params) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Params) Type() protoreflect.MessageType {
	return _fastReflection_Params_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Params) New() protoreflect.Message {
	return new(fastReflection_Params)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Params) Interface() protoreflect.ProtoMessage {
	return (*Params)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Params) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.SignedBlocksWindow != int64(0) {
		value := protoreflect.ValueOfInt64(x.SignedBlocksWindow)
		if !f(fd_Params_signed_blocks_window, value) {
			return
		}
	}
	if len(x.MinSignedPerWindow) != 0 {
		value := protoreflect.ValueOfBytes(x.MinSignedPerWindow)
		if !f(fd_Params_min_signed_per_window, value) {
			return
		}
	}
	if x.DowntimeJailDuration != nil {
		value := protoreflect.ValueOfMessage(x.DowntimeJailDuration.ProtoReflect())
		if !f(fd_Params_downtime_jail_duration, value) {
			return
		}
	}
	if len(x.SlashFractionDoubleSign) != 0 {
		value := protoreflect.ValueOfBytes(x.SlashFractionDoubleSign)
		if !f(fd_Params_slash_fraction_double_sign, value) {
			return
		}
	}
	if len(x.SlashFractionDowntime) != 0 {
		value := protoreflect.ValueOfBytes(x.SlashFractionDowntime)
		if !f(fd_Params_slash_fraction_downtime, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Params) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.Params.signed_blocks_window":
		return x.SignedBlocksWindow != int64(0)
	case "cosmos.slashing.v1beta1.Params.min_signed_per_window":
		return len(x.MinSignedPerWindow) != 0
	case "cosmos.slashing.v1beta1.Params.downtime_jail_duration":
		return x.DowntimeJailDuration != nil
	case "cosmos.slashing.v1beta1.Params.slash_fraction_double_sign":
		return len(x.SlashFractionDoubleSign) != 0
	case "cosmos.slashing.v1beta1.Params.slash_fraction_downtime":
		return len(x.SlashFractionDowntime) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.Params.signed_blocks_window":
		x.SignedBlocksWindow = int64(0)
	case "cosmos.slashing.v1beta1.Params.min_signed_per_window":
		x.MinSignedPerWindow = nil
	case "cosmos.slashing.v1beta1.Params.downtime_jail_duration":
		x.DowntimeJailDuration = nil
	case "cosmos.slashing.v1beta1.Params.slash_fraction_double_sign":
		x.SlashFractionDoubleSign = nil
	case "cosmos.slashing.v1beta1.Params.slash_fraction_downtime":
		x.SlashFractionDowntime = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Params) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.slashing.v1beta1.Params.signed_blocks_window":
		value := x.SignedBlocksWindow
		return protoreflect.ValueOfInt64(value)
	case "cosmos.slashing.v1beta1.Params.min_signed_per_window":
		value := x.MinSignedPerWindow
		return protoreflect.ValueOfBytes(value)
	case "cosmos.slashing.v1beta1.Params.downtime_jail_duration":
		value := x.DowntimeJailDuration
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.slashing.v1beta1.Params.slash_fraction_double_sign":
		value := x.SlashFractionDoubleSign
		return protoreflect.ValueOfBytes(value)
	case "cosmos.slashing.v1beta1.Params.slash_fraction_downtime":
		value := x.SlashFractionDowntime
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.Params does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.Params.signed_blocks_window":
		x.SignedBlocksWindow = value.Int()
	case "cosmos.slashing.v1beta1.Params.min_signed_per_window":
		x.MinSignedPerWindow = value.Bytes()
	case "cosmos.slashing.v1beta1.Params.downtime_jail_duration":
		x.DowntimeJailDuration = value.Message().Interface().(*durationpb.Duration)
	case "cosmos.slashing.v1beta1.Params.slash_fraction_double_sign":
		x.SlashFractionDoubleSign = value.Bytes()
	case "cosmos.slashing.v1beta1.Params.slash_fraction_downtime":
		x.SlashFractionDowntime = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.Params.downtime_jail_duration":
		if x.DowntimeJailDuration == nil {
			x.DowntimeJailDuration = new(durationpb.Duration)
		}
		return protoreflect.ValueOfMessage(x.DowntimeJailDuration.ProtoReflect())
	case "cosmos.slashing.v1beta1.Params.signed_blocks_window":
		panic(fmt.Errorf("field signed_blocks_window of message cosmos.slashing.v1beta1.Params is not mutable"))
	case "cosmos.slashing.v1beta1.Params.min_signed_per_window":
		panic(fmt.Errorf("field min_signed_per_window of message cosmos.slashing.v1beta1.Params is not mutable"))
	case "cosmos.slashing.v1beta1.Params.slash_fraction_double_sign":
		panic(fmt.Errorf("field slash_fraction_double_sign of message cosmos.slashing.v1beta1.Params is not mutable"))
	case "cosmos.slashing.v1beta1.Params.slash_fraction_downtime":
		panic(fmt.Errorf("field slash_fraction_downtime of message cosmos.slashing.v1beta1.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.slashing.v1beta1.Params.signed_blocks_window":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.slashing.v1beta1.Params.min_signed_per_window":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.slashing.v1beta1.Params.downtime_jail_duration":
		m := new(durationpb.Duration)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.slashing.v1beta1.Params.slash_fraction_double_sign":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.slashing.v1beta1.Params.slash_fraction_downtime":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.slashing.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.slashing.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.slashing.v1beta1.Params", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Params) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Params) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Params) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.SignedBlocksWindow != 0 {
			n += 1 + runtime.Sov(uint64(x.SignedBlocksWindow))
		}
		l = len(x.MinSignedPerWindow)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DowntimeJailDuration != nil {
			l = options.Size(x.DowntimeJailDuration)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.SlashFractionDoubleSign)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.SlashFractionDowntime)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.SlashFractionDowntime) > 0 {
			i -= len(x.SlashFractionDowntime)
			copy(dAtA[i:], x.SlashFractionDowntime)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.SlashFractionDowntime)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.SlashFractionDoubleSign) > 0 {
			i -= len(x.SlashFractionDoubleSign)
			copy(dAtA[i:], x.SlashFractionDoubleSign)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.SlashFractionDoubleSign)))
			i--
			dAtA[i] = 0x22
		}
		if x.DowntimeJailDuration != nil {
			encoded, err := options.Marshal(x.DowntimeJailDuration)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.MinSignedPerWindow) > 0 {
			i -= len(x.MinSignedPerWindow)
			copy(dAtA[i:], x.MinSignedPerWindow)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MinSignedPerWindow)))
			i--
			dAtA[i] = 0x12
		}
		if x.SignedBlocksWindow != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.SignedBlocksWindow))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SignedBlocksWindow", wireType)
				}
				x.SignedBlocksWindow = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.SignedBlocksWindow |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MinSignedPerWindow", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.MinSignedPerWindow = append(x.MinSignedPerWindow[:0], dAtA[iNdEx:postIndex]...)
				if x.MinSignedPerWindow == nil {
					x.MinSignedPerWindow = []byte{}
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DowntimeJailDuration", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.DowntimeJailDuration == nil {
					x.DowntimeJailDuration = &durationpb.Duration{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.DowntimeJailDuration); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SlashFractionDoubleSign", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SlashFractionDoubleSign = append(x.SlashFractionDoubleSign[:0], dAtA[iNdEx:postIndex]...)
				if x.SlashFractionDoubleSign == nil {
					x.SlashFractionDoubleSign = []byte{}
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SlashFractionDowntime", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SlashFractionDowntime = append(x.SlashFractionDowntime[:0], dAtA[iNdEx:postIndex]...)
				if x.SlashFractionDowntime == nil {
					x.SlashFractionDowntime = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/slashing/v1beta1/slashing.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ValidatorSigningInfo defines a validator's signing info for monitoring their
// liveness activity.
type ValidatorSigningInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Height at which validator was first a candidate OR was unjailed
	StartHeight int64 `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	// Index which is incremented each time the validator was a bonded
	// in a block and may have signed a precommit or not. This in conjunction with the
	// `SignedBlocksWindow` param determines the index in the `MissedBlocksBitArray`.
	IndexOffset int64 `protobuf:"varint,3,opt,name=index_offset,json=indexOffset,proto3" json:"index_offset,omitempty"`
	// Timestamp until which the validator is jailed due to liveness downtime.
	JailedUntil *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=jailed_until,json=jailedUntil,proto3" json:"jailed_until,omitempty"`
	// Whether or not a validator has been tombstoned (killed out of validator set). It is set
	// once the validator commits an equivocation or for any other configured misbehiavor.
	Tombstoned bool `protobuf:"varint,5,opt,name=tombstoned,proto3" json:"tombstoned,omitempty"`
	// A counter kept to avoid unnecessary array reads.
	// Note that `Sum(MissedBlocksBitArray)` always equals `MissedBlocksCounter`.
	MissedBlocksCounter int64 `protobuf:"varint,6,opt,name=missed_blocks_counter,json=missedBlocksCounter,proto3" json:"missed_blocks_counter,omitempty"`
}

func (x *ValidatorSigningInfo) Reset() {
	*x = ValidatorSigningInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_slashing_v1beta1_slashing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorSigningInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorSigningInfo) ProtoMessage() {}

// Deprecated: Use ValidatorSigningInfo.ProtoReflect.Descriptor instead.
func (*ValidatorSigningInfo) Descriptor() ([]byte, []int) {
	return file_cosmos_slashing_v1beta1_slashing_proto_rawDescGZIP(), []int{0}
}

func (x *ValidatorSigningInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ValidatorSigningInfo) GetStartHeight() int64 {
	if x != nil {
		return x.StartHeight
	}
	return 0
}

func (x *ValidatorSigningInfo) GetIndexOffset() int64 {
	if x != nil {
		return x.IndexOffset
	}
	return 0
}

func (x *ValidatorSigningInfo) GetJailedUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.JailedUntil
	}
	return nil
}

func (x *ValidatorSigningInfo) GetTombstoned() bool {
	if x != nil {
		return x.Tombstoned
	}
	return false
}

func (x *ValidatorSigningInfo) GetMissedBlocksCounter() int64 {
	if x != nil {
		return x.MissedBlocksCounter
	}
	return 0
}

// Params represents the parameters used for by the slashing module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignedBlocksWindow      int64                `protobuf:"varint,1,opt,name=signed_blocks_window,json=signedBlocksWindow,proto3" json:"signed_blocks_window,omitempty"`
	MinSignedPerWindow      []byte               `protobuf:"bytes,2,opt,name=min_signed_per_window,json=minSignedPerWindow,proto3" json:"min_signed_per_window,omitempty"`
	DowntimeJailDuration    *durationpb.Duration `protobuf:"bytes,3,opt,name=downtime_jail_duration,json=downtimeJailDuration,proto3" json:"downtime_jail_duration,omitempty"`
	SlashFractionDoubleSign []byte               `protobuf:"bytes,4,opt,name=slash_fraction_double_sign,json=slashFractionDoubleSign,proto3" json:"slash_fraction_double_sign,omitempty"`
	SlashFractionDowntime   []byte               `protobuf:"bytes,5,opt,name=slash_fraction_downtime,json=slashFractionDowntime,proto3" json:"slash_fraction_downtime,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_slashing_v1beta1_slashing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_cosmos_slashing_v1beta1_slashing_proto_rawDescGZIP(), []int{1}
}

func (x *Params) GetSignedBlocksWindow() int64 {
	if x != nil {
		return x.SignedBlocksWindow
	}
	return 0
}

func (x *Params) GetMinSignedPerWindow() []byte {
	if x != nil {
		return x.MinSignedPerWindow
	}
	return nil
}

func (x *Params) GetDowntimeJailDuration() *durationpb.Duration {
	if x != nil {
		return x.DowntimeJailDuration
	}
	return nil
}

func (x *Params) GetSlashFractionDoubleSign() []byte {
	if x != nil {
		return x.SlashFractionDoubleSign
	}
	return nil
}

func (x *Params) GetSlashFractionDowntime() []byte {
	if x != nil {
		return x.SlashFractionDowntime
	}
	return nil
}

var File_cosmos_slashing_v1beta1_slashing_proto protoreflect.FileDescriptor

var file_cosmos_slashing_v1beta1_slashing_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x02, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2,
	0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x47, 0x0a, 0x0c, 0x6a, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf,
	0x1f, 0x01, 0x52, 0x0b, 0x6a, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12,
	0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x74, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x64, 0x12,
	0x32, 0x0a, 0x15, 0x6d, 0x69, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13,
	0x6d, 0x69, 0x73, 0x73, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x3a, 0x08, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x01, 0x22, 0xcd, 0x03,
	0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x61, 0x0a, 0x15, 0x6d, 0x69,
	0x6e, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x77, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x2e, 0xc8, 0xde, 0x1f, 0x00, 0xda,
	0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x12, 0x6d, 0x69, 0x6e, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x50, 0x65, 0x72, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x59, 0x0a,
	0x16, 0x64, 0x6f, 0x77, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6a, 0x61, 0x69, 0x6c, 0x5f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x98, 0xdf,
	0x1f, 0x01, 0x52, 0x14, 0x64, 0x6f, 0x77, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x4a, 0x61, 0x69, 0x6c,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6b, 0x0a, 0x1a, 0x73, 0x6c, 0x61, 0x73,
	0x68, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x2e, 0xc8, 0xde,
	0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x17, 0x73, 0x6c,
	0x61, 0x73, 0x68, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x66, 0x0a, 0x17, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x5f, 0x66,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x2e, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x15, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x46, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x77, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x42, 0xf8, 0x01,
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x6c, 0x61,
	0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0d, 0x53,
	0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x73, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02,
	0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x17, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xe2, 0x02, 0x23, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x6c, 0x61, 0x73,
	0x68, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x3a, 0x3a, 0x53, 0x6c, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xa8, 0xe2, 0x1e, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_slashing_v1beta1_slashing_proto_rawDescOnce sync.Once
	file_cosmos_slashing_v1beta1_slashing_proto_rawDescData = file_cosmos_slashing_v1beta1_slashing_proto_rawDesc
)

func file_cosmos_slashing_v1beta1_slashing_proto_rawDescGZIP() []byte {
	file_cosmos_slashing_v1beta1_slashing_proto_rawDescOnce.Do(func() {
		file_cosmos_slashing_v1beta1_slashing_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_slashing_v1beta1_slashing_proto_rawDescData)
	})
	return file_cosmos_slashing_v1beta1_slashing_proto_rawDescData
}

var file_cosmos_slashing_v1beta1_slashing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cosmos_slashing_v1beta1_slashing_proto_goTypes = []interface{}{
	(*ValidatorSigningInfo)(nil),  // 0: cosmos.slashing.v1beta1.ValidatorSigningInfo
	(*Params)(nil),                // 1: cosmos.slashing.v1beta1.Params
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 3: google.protobuf.Duration
}
var file_cosmos_slashing_v1beta1_slashing_proto_depIdxs = []int32{
	2, // 0: cosmos.slashing.v1beta1.ValidatorSigningInfo.jailed_until:type_name -> google.protobuf.Timestamp
	3, // 1: cosmos.slashing.v1beta1.Params.downtime_jail_duration:type_name -> google.protobuf.Duration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cosmos_slashing_v1beta1_slashing_proto_init() }
func file_cosmos_slashing_v1beta1_slashing_proto_init() {
	if File_cosmos_slashing_v1beta1_slashing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_slashing_v1beta1_slashing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorSigningInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_slashing_v1beta1_slashing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_slashing_v1beta1_slashing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_slashing_v1beta1_slashing_proto_goTypes,
		DependencyIndexes: file_cosmos_slashing_v1beta1_slashing_proto_depIdxs,
		MessageInfos:      file_cosmos_slashing_v1beta1_slashing_proto_msgTypes,
	}.Build()
	File_cosmos_slashing_v1beta1_slashing_proto = out.File
	file_cosmos_slashing_v1beta1_slashing_proto_rawDesc = nil
	file_cosmos_slashing_v1beta1_slashing_proto_goTypes = nil
	file_cosmos_slashing_v1beta1_slashing_proto_depIdxs = nil
}