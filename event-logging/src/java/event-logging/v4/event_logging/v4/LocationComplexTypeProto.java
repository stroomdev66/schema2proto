// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/location_complex_type.proto
// Protobuf Java Version: 4.26.1

package event-logging.v4.event_logging.v4;

public final class LocationComplexTypeProto {
  private LocationComplexTypeProto() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      LocationComplexTypeProto.class.getName());
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
    internal_static_event_logging_v4_LocationComplexType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_event_logging_v4_LocationComplexType_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n,event_logging/v4/location_complex_type" +
      ".proto\022\020event_logging.v4\032/event_logging/" +
      "v4/coordinates_complex_type.proto\032\033buf/v" +
      "alidate/validate.proto\"\236\003\n\023LocationCompl" +
      "exType\022\030\n\007country\030\001 \001(\tR\007country\022\024\n\005stat" +
      "e\030\002 \001(\tR\005state\022\022\n\004city\030\003 \001(\tR\004city\022\022\n\004to" +
      "wn\030\004 \001(\tR\004town\022\022\n\004site\030\005 \001(\tR\004site\022\032\n\010bu" +
      "ilding\030\006 \001(\tR\010building\022\024\n\005floor\030\007 \001(\005R\005f" +
      "loor\022\022\n\004room\030\010 \001(\tR\004room\022\022\n\004desk\030\t \001(\tR\004" +
      "desk\022\022\n\004rack\030\n \001(\tR\004rack\022\032\n\010position\030\013 \001" +
      "(\tR\010position\022$\n\016time_zone_name\030) \001(\tR\014ti" +
      "meZoneName\022J\n\013coordinates\030* \001(\0132(.event_" +
      "logging.v4.CoordinatesComplexTypeR\013coord" +
      "inates\022\037\n\013access_zone\030+ \001(\tR\naccessZoneB" +
      "\341\001\n!event-logging.v4.event_logging.v4B\030L" +
      "ocationComplexTypeProtoP\001ZEgithub.com/gc" +
      "hq/event-logging-schema/event_logging/v4" +
      ";event_loggingv4\242\002\003EXX\252\002\017EventLogging.V4" +
      "\312\002\017EventLogging\\V4\342\002\033EventLogging\\V4\\GPB" +
      "Metadata\352\002\020EventLogging::V4b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          event-logging.v4.event_logging.v4.CoordinatesComplexTypeProto.getDescriptor(),
          build.buf.validate.ValidateProto.getDescriptor(),
        });
    internal_static_event_logging_v4_LocationComplexType_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_event_logging_v4_LocationComplexType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_event_logging_v4_LocationComplexType_descriptor,
        new java.lang.String[] { "Country", "State", "City", "Town", "Site", "Building", "Floor", "Room", "Desk", "Rack", "Position", "TimeZoneName", "Coordinates", "AccessZone", });
    descriptor.resolveAllFeaturesImmutable();
    event-logging.v4.event_logging.v4.CoordinatesComplexTypeProto.getDescriptor();
    build.buf.validate.ValidateProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
