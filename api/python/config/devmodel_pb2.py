# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: devmodel.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='devmodel.proto',
  package='',
  syntax='proto3',
  serialized_options=_b('\n\037com.zededa.cloud.uservice.protoZ$github.com/lf-edge/eve/api/go/config'),
  serialized_pb=_b('\n\x0e\x64\x65vmodel.proto\"n\n\x0fsWAdapterParams\x12\x1d\n\x05\x61Type\x18\x01 \x01(\x0e\x32\x0e.sWAdapterType\x12\x19\n\x11underlayInterface\x18\x08 \x01(\t\x12\x0e\n\x06vlanId\x18\t \x01(\r\x12\x11\n\tbondgroup\x18\n \x03(\t\"\xa1\x01\n\rSystemAdapter\x12\x0c\n\x04name\x18\x01 \x01(\t\x12&\n\x0c\x61llocDetails\x18\x14 \x01(\x0b\x32\x10.sWAdapterParams\x12\x12\n\nfreeUplink\x18\x02 \x01(\x08\x12\x0e\n\x06uplink\x18\x03 \x01(\x08\x12\x13\n\x0bnetworkUUID\x18\x04 \x01(\t\x12\x0c\n\x04\x61\x64\x64r\x18\x05 \x01(\t\x12\x13\n\x0blogicalName\x18\x06 \x01(\t*\\\n\x08ZCioType\x12\x0b\n\x07ZCioNop\x10\x00\x12\x0b\n\x07ZCioEth\x10\x01\x12\x0b\n\x07ZCioUSB\x10\x02\x12\x0b\n\x07ZCioCOM\x10\x03\x12\x0c\n\x08ZCioHDMI\x10\x04\x12\x0e\n\tZCioOther\x10\xff\x01*/\n\rsWAdapterType\x12\n\n\x06IGNORE\x10\x00\x12\x08\n\x04VLAN\x10\x01\x12\x08\n\x04\x42OND\x10\x02\x42G\n\x1f\x63om.zededa.cloud.uservice.protoZ$github.com/lf-edge/eve/api/go/configb\x06proto3')
)

_ZCIOTYPE = _descriptor.EnumDescriptor(
  name='ZCioType',
  full_name='ZCioType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ZCioNop', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ZCioEth', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ZCioUSB', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ZCioCOM', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ZCioHDMI', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ZCioOther', index=5, number=255,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=294,
  serialized_end=386,
)
_sym_db.RegisterEnumDescriptor(_ZCIOTYPE)

ZCioType = enum_type_wrapper.EnumTypeWrapper(_ZCIOTYPE)
_SWADAPTERTYPE = _descriptor.EnumDescriptor(
  name='sWAdapterType',
  full_name='sWAdapterType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='IGNORE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='VLAN', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BOND', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=388,
  serialized_end=435,
)
_sym_db.RegisterEnumDescriptor(_SWADAPTERTYPE)

sWAdapterType = enum_type_wrapper.EnumTypeWrapper(_SWADAPTERTYPE)
ZCioNop = 0
ZCioEth = 1
ZCioUSB = 2
ZCioCOM = 3
ZCioHDMI = 4
ZCioOther = 255
IGNORE = 0
VLAN = 1
BOND = 2



_SWADAPTERPARAMS = _descriptor.Descriptor(
  name='sWAdapterParams',
  full_name='sWAdapterParams',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='aType', full_name='sWAdapterParams.aType', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='underlayInterface', full_name='sWAdapterParams.underlayInterface', index=1,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='vlanId', full_name='sWAdapterParams.vlanId', index=2,
      number=9, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bondgroup', full_name='sWAdapterParams.bondgroup', index=3,
      number=10, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=18,
  serialized_end=128,
)


_SYSTEMADAPTER = _descriptor.Descriptor(
  name='SystemAdapter',
  full_name='SystemAdapter',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='SystemAdapter.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='allocDetails', full_name='SystemAdapter.allocDetails', index=1,
      number=20, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='freeUplink', full_name='SystemAdapter.freeUplink', index=2,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uplink', full_name='SystemAdapter.uplink', index=3,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='networkUUID', full_name='SystemAdapter.networkUUID', index=4,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='addr', full_name='SystemAdapter.addr', index=5,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='logicalName', full_name='SystemAdapter.logicalName', index=6,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=131,
  serialized_end=292,
)

_SWADAPTERPARAMS.fields_by_name['aType'].enum_type = _SWADAPTERTYPE
_SYSTEMADAPTER.fields_by_name['allocDetails'].message_type = _SWADAPTERPARAMS
DESCRIPTOR.message_types_by_name['sWAdapterParams'] = _SWADAPTERPARAMS
DESCRIPTOR.message_types_by_name['SystemAdapter'] = _SYSTEMADAPTER
DESCRIPTOR.enum_types_by_name['ZCioType'] = _ZCIOTYPE
DESCRIPTOR.enum_types_by_name['sWAdapterType'] = _SWADAPTERTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

sWAdapterParams = _reflection.GeneratedProtocolMessageType('sWAdapterParams', (_message.Message,), dict(
  DESCRIPTOR = _SWADAPTERPARAMS,
  __module__ = 'devmodel_pb2'
  # @@protoc_insertion_point(class_scope:sWAdapterParams)
  ))
_sym_db.RegisterMessage(sWAdapterParams)

SystemAdapter = _reflection.GeneratedProtocolMessageType('SystemAdapter', (_message.Message,), dict(
  DESCRIPTOR = _SYSTEMADAPTER,
  __module__ = 'devmodel_pb2'
  # @@protoc_insertion_point(class_scope:SystemAdapter)
  ))
_sym_db.RegisterMessage(SystemAdapter)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)