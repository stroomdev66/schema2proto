// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/process_type_simple_type.proto
// Protobuf Java Version: 4.26.1

package event-logging.v4.event_logging.v4;

public final class ProcessTypeSimpleTypeProto {
  private ProcessTypeSimpleTypeProto() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      ProcessTypeSimpleTypeProto.class.getName());
  }
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n/event_logging/v4/process_type_simple_t" +
      "ype.proto\022\020event_logging.v4\032\033buf/validat" +
      "e/validate.proto*\262\001\n\025ProcessTypeSimpleTy" +
      "pe\022(\n$PROCESS_TYPE_SIMPLE_TYPE_UNSPECIFI" +
      "ED\020\000\022\037\n\033PROCESS_TYPE_SIMPLE_TYPE_OS\020\001\022$\n" +
      " PROCESS_TYPE_SIMPLE_TYPE_SERVICE\020\002\022(\n$P" +
      "ROCESS_TYPE_SIMPLE_TYPE_APPLICATION\020\003B\343\001" +
      "\n!event-logging.v4.event_logging.v4B\032Pro" +
      "cessTypeSimpleTypeProtoP\001ZEgithub.com/gc" +
      "hq/event-logging-schema/event_logging/v4" +
      ";event_loggingv4\242\002\003EXX\252\002\017EventLogging.V4" +
      "\312\002\017EventLogging\\V4\342\002\033EventLogging\\V4\\GPB" +
      "Metadata\352\002\020EventLogging::V4b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          build.buf.validate.ValidateProto.getDescriptor(),
        });
    descriptor.resolveAllFeaturesImmutable();
    build.buf.validate.ValidateProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
