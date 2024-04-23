// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/print_action_simple_type.proto

// Protobuf Java Version: 4.26.1
package uk.gov.gchq.event_logging.v4;

/**
 * <pre>
 * The types of action relating to the use of a printer.
 * </pre>
 *
 * Protobuf enum {@code event_logging.v4.PrintActionSimpleType}
 */
public enum PrintActionSimpleType
    implements com.google.protobuf.ProtocolMessageEnum {
  /**
   * <pre>
   * Default
   * </pre>
   *
   * <code>PRINT_ACTION_SIMPLE_TYPE_UNSPECIFIED = 0;</code>
   */
  PRINT_ACTION_SIMPLE_TYPE_UNSPECIFIED(0),
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_CREATE_JOB = 1;</code>
   */
  PRINT_ACTION_SIMPLE_TYPE_CREATE_JOB(1),
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_CANCEL_JOB = 2;</code>
   */
  PRINT_ACTION_SIMPLE_TYPE_CANCEL_JOB(2),
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_PAUSE_JOB = 3;</code>
   */
  PRINT_ACTION_SIMPLE_TYPE_PAUSE_JOB(3),
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_RESUME_JOB = 4;</code>
   */
  PRINT_ACTION_SIMPLE_TYPE_RESUME_JOB(4),
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_START_PRINT = 5;</code>
   */
  PRINT_ACTION_SIMPLE_TYPE_START_PRINT(5),
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_FINISH_PRINT = 6;</code>
   */
  PRINT_ACTION_SIMPLE_TYPE_FINISH_PRINT(6),
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_CANCEL_PRINT = 7;</code>
   */
  PRINT_ACTION_SIMPLE_TYPE_CANCEL_PRINT(7),
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_FAILED_PRINT = 8;</code>
   */
  PRINT_ACTION_SIMPLE_TYPE_FAILED_PRINT(8),
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_OTHER = 9;</code>
   */
  PRINT_ACTION_SIMPLE_TYPE_OTHER(9),
  UNRECOGNIZED(-1),
  ;

  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      PrintActionSimpleType.class.getName());
  }
  /**
   * <pre>
   * Default
   * </pre>
   *
   * <code>PRINT_ACTION_SIMPLE_TYPE_UNSPECIFIED = 0;</code>
   */
  public static final int PRINT_ACTION_SIMPLE_TYPE_UNSPECIFIED_VALUE = 0;
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_CREATE_JOB = 1;</code>
   */
  public static final int PRINT_ACTION_SIMPLE_TYPE_CREATE_JOB_VALUE = 1;
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_CANCEL_JOB = 2;</code>
   */
  public static final int PRINT_ACTION_SIMPLE_TYPE_CANCEL_JOB_VALUE = 2;
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_PAUSE_JOB = 3;</code>
   */
  public static final int PRINT_ACTION_SIMPLE_TYPE_PAUSE_JOB_VALUE = 3;
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_RESUME_JOB = 4;</code>
   */
  public static final int PRINT_ACTION_SIMPLE_TYPE_RESUME_JOB_VALUE = 4;
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_START_PRINT = 5;</code>
   */
  public static final int PRINT_ACTION_SIMPLE_TYPE_START_PRINT_VALUE = 5;
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_FINISH_PRINT = 6;</code>
   */
  public static final int PRINT_ACTION_SIMPLE_TYPE_FINISH_PRINT_VALUE = 6;
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_CANCEL_PRINT = 7;</code>
   */
  public static final int PRINT_ACTION_SIMPLE_TYPE_CANCEL_PRINT_VALUE = 7;
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_FAILED_PRINT = 8;</code>
   */
  public static final int PRINT_ACTION_SIMPLE_TYPE_FAILED_PRINT_VALUE = 8;
  /**
   * <code>PRINT_ACTION_SIMPLE_TYPE_OTHER = 9;</code>
   */
  public static final int PRINT_ACTION_SIMPLE_TYPE_OTHER_VALUE = 9;


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
  public static PrintActionSimpleType valueOf(int value) {
    return forNumber(value);
  }

  /**
   * @param value The numeric wire value of the corresponding enum entry.
   * @return The enum associated with the given numeric wire value.
   */
  public static PrintActionSimpleType forNumber(int value) {
    switch (value) {
      case 0: return PRINT_ACTION_SIMPLE_TYPE_UNSPECIFIED;
      case 1: return PRINT_ACTION_SIMPLE_TYPE_CREATE_JOB;
      case 2: return PRINT_ACTION_SIMPLE_TYPE_CANCEL_JOB;
      case 3: return PRINT_ACTION_SIMPLE_TYPE_PAUSE_JOB;
      case 4: return PRINT_ACTION_SIMPLE_TYPE_RESUME_JOB;
      case 5: return PRINT_ACTION_SIMPLE_TYPE_START_PRINT;
      case 6: return PRINT_ACTION_SIMPLE_TYPE_FINISH_PRINT;
      case 7: return PRINT_ACTION_SIMPLE_TYPE_CANCEL_PRINT;
      case 8: return PRINT_ACTION_SIMPLE_TYPE_FAILED_PRINT;
      case 9: return PRINT_ACTION_SIMPLE_TYPE_OTHER;
      default: return null;
    }
  }

  public static com.google.protobuf.Internal.EnumLiteMap<PrintActionSimpleType>
      internalGetValueMap() {
    return internalValueMap;
  }
  private static final com.google.protobuf.Internal.EnumLiteMap<
      PrintActionSimpleType> internalValueMap =
        new com.google.protobuf.Internal.EnumLiteMap<PrintActionSimpleType>() {
          public PrintActionSimpleType findValueByNumber(int number) {
            return PrintActionSimpleType.forNumber(number);
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
    return uk.gov.gchq.event_logging.v4.PrintActionSimpleTypeProto.getDescriptor().getEnumTypes().get(0);
  }

  private static final PrintActionSimpleType[] VALUES = values();

  public static PrintActionSimpleType valueOf(
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

  private PrintActionSimpleType(int value) {
    this.value = value;
  }

  // @@protoc_insertion_point(enum_scope:event_logging.v4.PrintActionSimpleType)
}

