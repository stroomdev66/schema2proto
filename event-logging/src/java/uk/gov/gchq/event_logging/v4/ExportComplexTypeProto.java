// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/export_complex_type.proto
// Protobuf Java Version: 4.26.1

package uk.gov.gchq.event_logging.v4;

public final class ExportComplexTypeProto {
  private ExportComplexTypeProto() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      ExportComplexTypeProto.class.getName());
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
    internal_static_event_logging_v4_ExportComplexType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_event_logging_v4_ExportComplexType_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n*event_logging/v4/export_complex_type.p" +
      "roto\022\020event_logging.v4\032(event_logging/v4" +
      "/data_complex_type.proto\0320event_logging/" +
      "v4/multi_object_complex_type.proto\032+even" +
      "t_logging/v4/outcome_complex_type.proto\032" +
      "\033buf/validate/validate.proto\"\242\002\n\021ExportC" +
      "omplexType\022@\n\006source\030\001 \001(\0132(.event_loggi" +
      "ng.v4.MultiObjectComplexTypeR\006source\022J\n\013" +
      "destination\030\002 \001(\0132(.event_logging.v4.Mul" +
      "tiObjectComplexTypeR\013destination\022>\n\007outc" +
      "ome\030\003 \001(\0132$.event_logging.v4.OutcomeComp" +
      "lexTypeR\007outcome\022?\n\004data\030\004 \003(\0132!.event_l" +
      "ogging.v4.DataComplexTypeB\010\272H\005\222\001\002\010\000R\004dat" +
      "aB\332\001\n\034uk.gov.gchq.event_logging.v4B\026Expo" +
      "rtComplexTypeProtoP\001ZEgithub.com/gchq/ev" +
      "ent-logging-schema/event_logging/v4;even" +
      "t_loggingv4\242\002\003EXX\252\002\017EventLogging.V4\312\002\017Ev" +
      "entLogging\\V4\342\002\033EventLogging\\V4\\GPBMetad" +
      "ata\352\002\020EventLogging::V4b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          uk.gov.gchq.event_logging.v4.DataComplexTypeProto.getDescriptor(),
          uk.gov.gchq.event_logging.v4.MultiObjectComplexTypeProto.getDescriptor(),
          uk.gov.gchq.event_logging.v4.OutcomeComplexTypeProto.getDescriptor(),
          build.buf.validate.ValidateProto.getDescriptor(),
        });
    internal_static_event_logging_v4_ExportComplexType_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_event_logging_v4_ExportComplexType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_event_logging_v4_ExportComplexType_descriptor,
        new java.lang.String[] { "Source", "Destination", "Outcome", "Data", });
    descriptor.resolveAllFeaturesImmutable();
    uk.gov.gchq.event_logging.v4.DataComplexTypeProto.getDescriptor();
    uk.gov.gchq.event_logging.v4.MultiObjectComplexTypeProto.getDescriptor();
    uk.gov.gchq.event_logging.v4.OutcomeComplexTypeProto.getDescriptor();
    build.buf.validate.ValidateProto.getDescriptor();
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(build.buf.validate.ValidateProto.field);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
  }

  // @@protoc_insertion_point(outer_class_scope)
}
