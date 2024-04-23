// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/system_complex_type.proto
// Protobuf Java Version: 4.26.1

package event-logging.v4.event_logging.v4;

public final class SystemComplexTypeProto {
  private SystemComplexTypeProto() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      SystemComplexTypeProto.class.getName());
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
    internal_static_event_logging_v4_SystemComplexType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_event_logging_v4_SystemComplexType_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_event_logging_v4_SystemComplexType_TagsType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_event_logging_v4_SystemComplexType_TagsType_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n*event_logging/v4/system_complex_type.p" +
      "roto\022\020event_logging.v4\0322event_logging/v4" +
      "/classification_complex_type.proto\032\033buf/" +
      "validate/validate.proto\"\241\003\n\021SystemComple" +
      "xType\022\032\n\004name\030\001 \001(\tB\006\272H\003\310\001\001R\004name\022 \n\013des" +
      "cription\030\002 \001(\tR\013description\022S\n\016classific" +
      "ation\030\003 \001(\0132+.event_logging.v4.Classific" +
      "ationComplexTypeR\016classification\022(\n\013envi" +
      "ronment\030\004 \001(\tB\006\272H\003\310\001\001R\013environment\022\"\n\014or" +
      "ganisation\030\005 \001(\tR\014organisation\022\'\n\017securi" +
      "ty_domain\030\006 \001(\tR\016securityDomain\022\030\n\007versi" +
      "on\030\007 \001(\tR\007version\022@\n\004tags\030\010 \001(\0132,.event_" +
      "logging.v4.SystemComplexType.TagsTypeR\004t" +
      "ags\032&\n\010TagsType\022\032\n\003tag\030\001 \003(\tB\010\272H\005\222\001\002\010\001R\003" +
      "tagB\337\001\n!event-logging.v4.event_logging.v" +
      "4B\026SystemComplexTypeProtoP\001ZEgithub.com/" +
      "gchq/event-logging-schema/event_logging/" +
      "v4;event_loggingv4\242\002\003EXX\252\002\017EventLogging." +
      "V4\312\002\017EventLogging\\V4\342\002\033EventLogging\\V4\\G" +
      "PBMetadata\352\002\020EventLogging::V4b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          event-logging.v4.event_logging.v4.ClassificationComplexTypeProto.getDescriptor(),
          build.buf.validate.ValidateProto.getDescriptor(),
        });
    internal_static_event_logging_v4_SystemComplexType_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_event_logging_v4_SystemComplexType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_event_logging_v4_SystemComplexType_descriptor,
        new java.lang.String[] { "Name", "Description", "Classification", "Environment", "Organisation", "SecurityDomain", "Version", "Tags", });
    internal_static_event_logging_v4_SystemComplexType_TagsType_descriptor =
      internal_static_event_logging_v4_SystemComplexType_descriptor.getNestedTypes().get(0);
    internal_static_event_logging_v4_SystemComplexType_TagsType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_event_logging_v4_SystemComplexType_TagsType_descriptor,
        new java.lang.String[] { "Tag", });
    descriptor.resolveAllFeaturesImmutable();
    event-logging.v4.event_logging.v4.ClassificationComplexTypeProto.getDescriptor();
    build.buf.validate.ValidateProto.getDescriptor();
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(build.buf.validate.ValidateProto.field);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
  }

  // @@protoc_insertion_point(outer_class_scope)
}
