// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/outcome_complex_type.proto
// Protobuf Java Version: 4.26.1

package event-logging.v4.event_logging.v4;

public final class OutcomeComplexTypeProto {
  private OutcomeComplexTypeProto() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      OutcomeComplexTypeProto.class.getName());
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
    internal_static_event_logging_v4_OutcomeComplexType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_event_logging_v4_OutcomeComplexType_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_event_logging_v4_OutcomeComplexType_AuthServiceType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_event_logging_v4_OutcomeComplexType_AuthServiceType_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n+event_logging/v4/outcome_complex_type." +
      "proto\022\020event_logging.v4\032(event_logging/v" +
      "4/data_complex_type.proto\032\033buf/validate/" +
      "validate.proto\"\331\002\n\022OutcomeComplexType\022\030\n" +
      "\007success\030\001 \001(\010R\007success\022\034\n\tpermitted\030\002 \001" +
      "(\010R\tpermitted\022W\n\014auth_service\030\003 \001(\01324.ev" +
      "ent_logging.v4.OutcomeComplexType.AuthSe" +
      "rviceTypeR\013authService\022 \n\013description\030\004 " +
      "\001(\tR\013description\022?\n\004data\030\005 \003(\0132!.event_l" +
      "ogging.v4.DataComplexTypeB\010\272H\005\222\001\002\010\000R\004dat" +
      "a\032O\n\017AuthServiceType\022\016\n\002id\030\001 \001(\tR\002id\022,\n\r" +
      "cache_timeout\030\002 \001(\rB\007\272H\004*\002 \000R\014cacheTimeo" +
      "utB\340\001\n!event-logging.v4.event_logging.v4" +
      "B\027OutcomeComplexTypeProtoP\001ZEgithub.com/" +
      "gchq/event-logging-schema/event_logging/" +
      "v4;event_loggingv4\242\002\003EXX\252\002\017EventLogging." +
      "V4\312\002\017EventLogging\\V4\342\002\033EventLogging\\V4\\G" +
      "PBMetadata\352\002\020EventLogging::V4b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          event-logging.v4.event_logging.v4.DataComplexTypeProto.getDescriptor(),
          build.buf.validate.ValidateProto.getDescriptor(),
        });
    internal_static_event_logging_v4_OutcomeComplexType_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_event_logging_v4_OutcomeComplexType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_event_logging_v4_OutcomeComplexType_descriptor,
        new java.lang.String[] { "Success", "Permitted", "AuthService", "Description", "Data", });
    internal_static_event_logging_v4_OutcomeComplexType_AuthServiceType_descriptor =
      internal_static_event_logging_v4_OutcomeComplexType_descriptor.getNestedTypes().get(0);
    internal_static_event_logging_v4_OutcomeComplexType_AuthServiceType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_event_logging_v4_OutcomeComplexType_AuthServiceType_descriptor,
        new java.lang.String[] { "Id", "CacheTimeout", });
    descriptor.resolveAllFeaturesImmutable();
    event-logging.v4.event_logging.v4.DataComplexTypeProto.getDescriptor();
    build.buf.validate.ValidateProto.getDescriptor();
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(build.buf.validate.ValidateProto.field);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
  }

  // @@protoc_insertion_point(outer_class_scope)
}
