// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/hardware_type_simple_type.proto
// Protobuf Java Version: 4.26.1

package event-logging.v4.event_logging.v4;

public final class HardwareTypeSimpleTypeProto {
  private HardwareTypeSimpleTypeProto() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      HardwareTypeSimpleTypeProto.class.getName());
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
      "\n0event_logging/v4/hardware_type_simple_" +
      "type.proto\022\020event_logging.v4\032\033buf/valida" +
      "te/validate.proto*\255\003\n\026HardwareTypeSimple" +
      "Type\022)\n%HARDWARE_TYPE_SIMPLE_TYPE_UNSPEC" +
      "IFIED\020\000\022+\n\'HARDWARE_TYPE_SIMPLE_TYPE_OPT" +
      "ICAL_DRIVE\020\001\022\'\n#HARDWARE_TYPE_SIMPLE_TYP" +
      "E_HARD_DISK\020\002\022.\n*HARDWARE_TYPE_SIMPLE_TY" +
      "PE_USB_MASS_STORAGE\020\003\022%\n!HARDWARE_TYPE_S" +
      "IMPLE_TYPE_PRINTER\020\004\022#\n\037HARDWARE_TYPE_SI" +
      "MPLE_TYPE_MODEM\020\005\022&\n\"HARDWARE_TYPE_SIMPL" +
      "E_TYPE_KEYBOARD\020\006\022#\n\037HARDWARE_TYPE_SIMPL" +
      "E_TYPE_MOUSE\020\007\022$\n HARDWARE_TYPE_SIMPLE_T" +
      "YPE_WEBCAM\020\010\022#\n\037HARDWARE_TYPE_SIMPLE_TYP" +
      "E_OTHER\020\tB\344\001\n!event-logging.v4.event_log" +
      "ging.v4B\033HardwareTypeSimpleTypeProtoP\001ZE" +
      "github.com/gchq/event-logging-schema/eve" +
      "nt_logging/v4;event_loggingv4\242\002\003EXX\252\002\017Ev" +
      "entLogging.V4\312\002\017EventLogging\\V4\342\002\033EventL" +
      "ogging\\V4\\GPBMetadata\352\002\020EventLogging::V4" +
      "b\006proto3"
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
