// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/move_complex_type.proto
// Protobuf Java Version: 4.26.1

package uk.gov.gchq.event_logging.v4;

public final class MoveComplexTypeProto {
  private MoveComplexTypeProto() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      MoveComplexTypeProto.class.getName());
  }
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_event_logging_v4_MoveComplexType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_event_logging_v4_MoveComplexType_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n(event_logging/v4/move_complex_type.pro" +
      "to\022\020event_logging.v4\0325event_logging/v4/c" +
      "opy_move_outcome_complex_type.proto\032(eve" +
      "nt_logging/v4/data_complex_type.proto\0320e" +
      "vent_logging/v4/multi_object_complex_typ" +
      "e.proto\032\033buf/validate/validate.proto\"\270\002\n" +
      "\017MoveComplexType\022H\n\006source\030\001 \001(\0132(.event" +
      "_logging.v4.MultiObjectComplexTypeB\006\272H\003\310" +
      "\001\001R\006source\022R\n\013destination\030\002 \001(\0132(.event_" +
      "logging.v4.MultiObjectComplexTypeB\006\272H\003\310\001" +
      "\001R\013destination\022F\n\007outcome\030\003 \001(\0132,.event_" +
      "logging.v4.CopyMoveOutcomeComplexTypeR\007o" +
      "utcome\022?\n\004data\030\004 \003(\0132!.event_logging.v4." +
      "DataComplexTypeB\010\272H\005\222\001\002\010\000R\004dataB\330\001\n\034uk.g" +
      "ov.gchq.event_logging.v4B\024MoveComplexTyp" +
      "eProtoP\001ZEgithub.com/gchq/event-logging-" +
      "schema/event_logging/v4;event_loggingv4\242" +
      "\002\003EXX\252\002\017EventLogging.V4\312\002\017EventLogging\\V" +
      "4\342\002\033EventLogging\\V4\\GPBMetadata\352\002\020EventL" +
      "ogging::V4b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          uk.gov.gchq.event_logging.v4.CopyMoveOutcomeComplexTypeProto.getDescriptor(),
          uk.gov.gchq.event_logging.v4.DataComplexTypeProto.getDescriptor(),
          uk.gov.gchq.event_logging.v4.MultiObjectComplexTypeProto.getDescriptor(),
          build.buf.validate.ValidateProto.getDescriptor(),
        });
    internal_static_event_logging_v4_MoveComplexType_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_event_logging_v4_MoveComplexType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_event_logging_v4_MoveComplexType_descriptor,
        new java.lang.String[] { "Source", "Destination", "Outcome", "Data", });
    descriptor.resolveAllFeaturesImmutable();
    uk.gov.gchq.event_logging.v4.CopyMoveOutcomeComplexTypeProto.getDescriptor();
    uk.gov.gchq.event_logging.v4.DataComplexTypeProto.getDescriptor();
    uk.gov.gchq.event_logging.v4.MultiObjectComplexTypeProto.getDescriptor();
    build.buf.validate.ValidateProto.getDescriptor();
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(build.buf.validate.ValidateProto.field);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
  }

  // @@protoc_insertion_point(outer_class_scope)
}
