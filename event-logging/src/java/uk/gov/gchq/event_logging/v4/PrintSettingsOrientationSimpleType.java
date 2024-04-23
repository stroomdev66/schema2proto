// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/print_settings_orientation_simple_type.proto

// Protobuf Java Version: 4.26.1
package uk.gov.gchq.event_logging.v4;

/**
 * <pre>
 * The orientation types of a document when it is printed.
 * </pre>
 *
 * Protobuf enum {@code event_logging.v4.PrintSettingsOrientationSimpleType}
 */
public enum PrintSettingsOrientationSimpleType
    implements com.google.protobuf.ProtocolMessageEnum {
  /**
   * <pre>
   * Default
   * </pre>
   *
   * <code>PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_UNSPECIFIED = 0;</code>
   */
  PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_UNSPECIFIED(0),
  /**
   * <code>PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_PORTRAIT = 1;</code>
   */
  PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_PORTRAIT(1),
  /**
   * <code>PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_LANDSCAPE = 2;</code>
   */
  PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_LANDSCAPE(2),
  UNRECOGNIZED(-1),
  ;

  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      PrintSettingsOrientationSimpleType.class.getName());
  }
  /**
   * <pre>
   * Default
   * </pre>
   *
   * <code>PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_UNSPECIFIED = 0;</code>
   */
  public static final int PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_UNSPECIFIED_VALUE = 0;
  /**
   * <code>PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_PORTRAIT = 1;</code>
   */
  public static final int PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_PORTRAIT_VALUE = 1;
  /**
   * <code>PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_LANDSCAPE = 2;</code>
   */
  public static final int PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_LANDSCAPE_VALUE = 2;


  public final int getNumber() {
    if (this == UNRECOGNIZED) {
      throw new java.lang.IllegalArgumentException(
          "Can't get the number of an unknown enum value.");
    }
    return value;
  }

  /**
   * @param value The numeric wire value of the corresponding enum entry.
   * @return The enum associated with the given numeric wire value.
   * @deprecated Use {@link #forNumber(int)} instead.
   */
  @java.lang.Deprecated
  public static PrintSettingsOrientationSimpleType valueOf(int value) {
    return forNumber(value);
  }

  /**
   * @param value The numeric wire value of the corresponding enum entry.
   * @return The enum associated with the given numeric wire value.
   */
  public static PrintSettingsOrientationSimpleType forNumber(int value) {
    switch (value) {
      case 0: return PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_UNSPECIFIED;
      case 1: return PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_PORTRAIT;
      case 2: return PRINT_SETTINGS_ORIENTATION_SIMPLE_TYPE_LANDSCAPE;
      default: return null;
    }
  }

  public static com.google.protobuf.Internal.EnumLiteMap<PrintSettingsOrientationSimpleType>
      internalGetValueMap() {
    return internalValueMap;
  }
  private static final com.google.protobuf.Internal.EnumLiteMap<
      PrintSettingsOrientationSimpleType> internalValueMap =
        new com.google.protobuf.Internal.EnumLiteMap<PrintSettingsOrientationSimpleType>() {
          public PrintSettingsOrientationSimpleType findValueByNumber(int number) {
            return PrintSettingsOrientationSimpleType.forNumber(number);
          }
        };

  public final com.google.protobuf.Descriptors.EnumValueDescriptor
      getValueDescriptor() {
    if (this == UNRECOGNIZED) {
      throw new java.lang.IllegalStateException(
          "Can't get the descriptor of an unrecognized enum value.");
    }
    return getDescriptor().getValues().get(ordinal());
  }
  public final com.google.protobuf.Descriptors.EnumDescriptor
      getDescriptorForType() {
    return getDescriptor();
  }
  public static final com.google.protobuf.Descriptors.EnumDescriptor
      getDescriptor() {
    return uk.gov.gchq.event_logging.v4.PrintSettingsOrientationSimpleTypeProto.getDescriptor().getEnumTypes().get(0);
  }

  private static final PrintSettingsOrientationSimpleType[] VALUES = values();

  public static PrintSettingsOrientationSimpleType valueOf(
      com.google.protobuf.Descriptors.EnumValueDescriptor desc) {
    if (desc.getType() != getDescriptor()) {
      throw new java.lang.IllegalArgumentException(
        "EnumValueDescriptor is not for this type.");
    }
    if (desc.getIndex() == -1) {
      return UNRECOGNIZED;
    }
    return VALUES[desc.getIndex()];
  }

  private final int value;

  private PrintSettingsOrientationSimpleType(int value) {
    this.value = value;
  }

  // @@protoc_insertion_point(enum_scope:event_logging.v4.PrintSettingsOrientationSimpleType)
}

