// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/any_content_complex_type.proto
// Protobuf Java Version: 4.26.1

package event-logging.v4.event_logging.v4;

public final class AnyContentComplexTypeProto {
  private AnyContentComplexTypeProto() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      AnyContentComplexTypeProto.class.getName());
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
    internal_static_event_logging_v4_AnyContentComplexType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_event_logging_v4_AnyContentComplexType_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n/event_logging/v4/any_content_complex_t" +
      "ype.proto\022\020event_logging.v4\032\033buf/validat" +
      "e/validate.proto\"f\n\025AnyContentComplexTyp" +
      "e\022!\n\014content_type\030\001 \001(\tR\013contentType\022\030\n\007" +
      "version\030\002 \001(\tR\007version\022\020\n\003any\030\003 \001(\tR\003any" +
      "B\343\001\n!event-logging.v4.event_logging.v4B\032" +
      "AnyContentComplexTypeProtoP\001ZEgithub.com" +
      "/gchq/event-logging-schema/event_logging" +
      "/v4;event_loggingv4\242\002\003EXX\252\002\017EventLogging" +
      ".V4\312\002\017EventLogging\\V4\342\002\033EventLogging\\V4\\" +
      "GPBMetadata\352\002\020EventLogging::V4b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          build.buf.validate.ValidateProto.getDescriptor(),
        });
    internal_static_event_logging_v4_AnyContentComplexType_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_event_logging_v4_AnyContentComplexType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_event_logging_v4_AnyContentComplexType_descriptor,
        new java.lang.String[] { "ContentType", "Version", "Any", });
    descriptor.resolveAllFeaturesImmutable();
    build.buf.validate.ValidateProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
