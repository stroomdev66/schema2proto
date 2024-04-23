// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/media_type_simple_type.proto

// Protobuf Java Version: 4.26.1
package uk.gov.gchq.event_logging.v4;

/**
 * <pre>
 * The types of removable media.
 * </pre>
 *
 * Protobuf enum {@code event_logging.v4.MediaTypeSimpleType}
 */
public enum MediaTypeSimpleType
    implements com.google.protobuf.ProtocolMessageEnum {
  /**
   * <pre>
   * Default
   * </pre>
   *
   * <code>MEDIA_TYPE_SIMPLE_TYPE_UNSPECIFIED = 0;</code>
   */
  MEDIA_TYPE_SIMPLE_TYPE_UNSPECIFIED(0),
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_HARD_DISK = 1;</code>
   */
  MEDIA_TYPE_SIMPLE_TYPE_HARD_DISK(1),
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_USB_MASS_STORAGE = 2;</code>
   */
  MEDIA_TYPE_SIMPLE_TYPE_USB_MASS_STORAGE(2),
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_BLU_RAY = 3;</code>
   */
  MEDIA_TYPE_SIMPLE_TYPE_BLU_RAY(3),
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_DVD = 4;</code>
   */
  MEDIA_TYPE_SIMPLE_TYPE_DVD(4),
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_CD = 5;</code>
   */
  MEDIA_TYPE_SIMPLE_TYPE_CD(5),
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_FLOPPY_DISK = 6;</code>
   */
  MEDIA_TYPE_SIMPLE_TYPE_FLOPPY_DISK(6),
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_TAPE = 7;</code>
   */
  MEDIA_TYPE_SIMPLE_TYPE_TAPE(7),
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_MEMORY_CARD = 8;</code>
   */
  MEDIA_TYPE_SIMPLE_TYPE_MEMORY_CARD(8),
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_OTHER = 9;</code>
   */
  MEDIA_TYPE_SIMPLE_TYPE_OTHER(9),
  UNRECOGNIZED(-1),
  ;

  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      MediaTypeSimpleType.class.getName());
  }
  /**
   * <pre>
   * Default
   * </pre>
   *
   * <code>MEDIA_TYPE_SIMPLE_TYPE_UNSPECIFIED = 0;</code>
   */
  public static final int MEDIA_TYPE_SIMPLE_TYPE_UNSPECIFIED_VALUE = 0;
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_HARD_DISK = 1;</code>
   */
  public static final int MEDIA_TYPE_SIMPLE_TYPE_HARD_DISK_VALUE = 1;
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_USB_MASS_STORAGE = 2;</code>
   */
  public static final int MEDIA_TYPE_SIMPLE_TYPE_USB_MASS_STORAGE_VALUE = 2;
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_BLU_RAY = 3;</code>
   */
  public static final int MEDIA_TYPE_SIMPLE_TYPE_BLU_RAY_VALUE = 3;
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_DVD = 4;</code>
   */
  public static final int MEDIA_TYPE_SIMPLE_TYPE_DVD_VALUE = 4;
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_CD = 5;</code>
   */
  public static final int MEDIA_TYPE_SIMPLE_TYPE_CD_VALUE = 5;
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_FLOPPY_DISK = 6;</code>
   */
  public static final int MEDIA_TYPE_SIMPLE_TYPE_FLOPPY_DISK_VALUE = 6;
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_TAPE = 7;</code>
   */
  public static final int MEDIA_TYPE_SIMPLE_TYPE_TAPE_VALUE = 7;
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_MEMORY_CARD = 8;</code>
   */
  public static final int MEDIA_TYPE_SIMPLE_TYPE_MEMORY_CARD_VALUE = 8;
  /**
   * <code>MEDIA_TYPE_SIMPLE_TYPE_OTHER = 9;</code>
   */
  public static final int MEDIA_TYPE_SIMPLE_TYPE_OTHER_VALUE = 9;


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
  public static MediaTypeSimpleType valueOf(int value) {
    return forNumber(value);
  }

  /**
   * @param value The numeric wire value of the corresponding enum entry.
   * @return The enum associated with the given numeric wire value.
   */
  public static MediaTypeSimpleType forNumber(int value) {
    switch (value) {
      case 0: return MEDIA_TYPE_SIMPLE_TYPE_UNSPECIFIED;
      case 1: return MEDIA_TYPE_SIMPLE_TYPE_HARD_DISK;
      case 2: return MEDIA_TYPE_SIMPLE_TYPE_USB_MASS_STORAGE;
      case 3: return MEDIA_TYPE_SIMPLE_TYPE_BLU_RAY;
      case 4: return MEDIA_TYPE_SIMPLE_TYPE_DVD;
      case 5: return MEDIA_TYPE_SIMPLE_TYPE_CD;
      case 6: return MEDIA_TYPE_SIMPLE_TYPE_FLOPPY_DISK;
      case 7: return MEDIA_TYPE_SIMPLE_TYPE_TAPE;
      case 8: return MEDIA_TYPE_SIMPLE_TYPE_MEMORY_CARD;
      case 9: return MEDIA_TYPE_SIMPLE_TYPE_OTHER;
      default: return null;
    }
  }

  public static com.google.protobuf.Internal.EnumLiteMap<MediaTypeSimpleType>
      internalGetValueMap() {
    return internalValueMap;
  }
  private static final com.google.protobuf.Internal.EnumLiteMap<
      MediaTypeSimpleType> internalValueMap =
        new com.google.protobuf.Internal.EnumLiteMap<MediaTypeSimpleType>() {
          public MediaTypeSimpleType findValueByNumber(int number) {
            return MediaTypeSimpleType.forNumber(number);
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
    return uk.gov.gchq.event_logging.v4.MediaTypeSimpleTypeProto.getDescriptor().getEnumTypes().get(0);
  }

  private static final MediaTypeSimpleType[] VALUES = values();

  public static MediaTypeSimpleType valueOf(
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

  private MediaTypeSimpleType(int value) {
    this.value = value;
  }

  // @@protoc_insertion_point(enum_scope:event_logging.v4.MediaTypeSimpleType)
}

