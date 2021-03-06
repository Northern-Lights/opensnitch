# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: opensnitch/network/network.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='opensnitch/network/network.proto',
  package='opensnitch.network',
  syntax='proto3',
  serialized_options=_b('Z(github.com/evilsocket/opensnitch/network'),
  serialized_pb=_b('\n opensnitch/network/network.proto\x12\x12opensnitch.network\"\x99\x02\n\nConnection\x12\x39\n\x08protocol\x18\x01 \x01(\x0e\x32\'.opensnitch.network.Connection.Protocol\x12\x0e\n\x06src_ip\x18\x02 \x01(\t\x12\x10\n\x08src_port\x18\x03 \x01(\r\x12\x0e\n\x06\x64st_ip\x18\x04 \x01(\t\x12\x10\n\x08\x64st_host\x18\x05 \x01(\t\x12\x10\n\x08\x64st_port\x18\x06 \x01(\r\x12\x0f\n\x07user_id\x18\x07 \x01(\r\x12\x12\n\nprocess_id\x18\x08 \x01(\r\x12\x14\n\x0cprocess_path\x18\t \x01(\t\x12\x14\n\x0cprocess_args\x18\n \x03(\t\")\n\x08Protocol\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x07\n\x03TCP\x10\x01\x12\x07\n\x03UDP\x10\x02\x42*Z(github.com/evilsocket/opensnitch/networkb\x06proto3')
)



_CONNECTION_PROTOCOL = _descriptor.EnumDescriptor(
  name='Protocol',
  full_name='opensnitch.network.Connection.Protocol',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TCP', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UDP', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=297,
  serialized_end=338,
)
_sym_db.RegisterEnumDescriptor(_CONNECTION_PROTOCOL)


_CONNECTION = _descriptor.Descriptor(
  name='Connection',
  full_name='opensnitch.network.Connection',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='protocol', full_name='opensnitch.network.Connection.protocol', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='src_ip', full_name='opensnitch.network.Connection.src_ip', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='src_port', full_name='opensnitch.network.Connection.src_port', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dst_ip', full_name='opensnitch.network.Connection.dst_ip', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dst_host', full_name='opensnitch.network.Connection.dst_host', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dst_port', full_name='opensnitch.network.Connection.dst_port', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='opensnitch.network.Connection.user_id', index=6,
      number=7, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='process_id', full_name='opensnitch.network.Connection.process_id', index=7,
      number=8, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='process_path', full_name='opensnitch.network.Connection.process_path', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='process_args', full_name='opensnitch.network.Connection.process_args', index=9,
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
    _CONNECTION_PROTOCOL,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=57,
  serialized_end=338,
)

_CONNECTION.fields_by_name['protocol'].enum_type = _CONNECTION_PROTOCOL
_CONNECTION_PROTOCOL.containing_type = _CONNECTION
DESCRIPTOR.message_types_by_name['Connection'] = _CONNECTION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Connection = _reflection.GeneratedProtocolMessageType('Connection', (_message.Message,), dict(
  DESCRIPTOR = _CONNECTION,
  __module__ = 'opensnitch.network.network_pb2'
  # @@protoc_insertion_point(class_scope:opensnitch.network.Connection)
  ))
_sym_db.RegisterMessage(Connection)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
