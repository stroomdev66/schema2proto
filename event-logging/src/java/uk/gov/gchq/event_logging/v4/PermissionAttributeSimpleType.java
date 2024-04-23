// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/permission_attribute_simple_type.proto

// Protobuf Java Version: 4.26.1
package uk.gov.gchq.event_logging.v4;

/**
 * <pre>
 * The types of permission that can be assigned to an entity such as a document.
 * </pre>
 *
 * Protobuf enum {@code event_logging.v4.PermissionAttributeSimpleType}
 */
public enum PermissionAttributeSimpleType
    implements com.google.protobuf.ProtocolMessageEnum {
  /**
   * <pre>
   * Default
   * </pre>
   *
   * <code>PERMISSION_ATTRIBUTE_SIMPLE_TYPE_UNSPECIFIED = 0;</code>
   */
  PERMISSION_ATTRIBUTE_SIMPLE_TYPE_UNSPECIFIED(0),
  /**
   * <code>PERMISSION_ATTRIBUTE_SIMPLE_TYPE_AUTHOR = 1;</code>
   */
  PERMISSION_ATTRIBUTE_SIMPLE_TYPE_AUTHOR(1),
  /**
   * <code>PERMISSION_ATTRIBUTE_SIMPLE_TYPE_OWNER = 2;</code>
   */
  PERMISSION_ATTRIBUTE_SIMPLE_TYPE_OWNER(2),
  /**
   * <code>PERMISSION_ATTRIBUTE_SIMPLE_TYPE_READ = 3;</code>
   */
  PERMISSION_ATTRIBUTE_SIMPLE_TYPE_READ(3),
  /**
   * <code>PERMISSION_ATTRIBUTE_SIMPLE_TYPE_WRITE = 4;</code>
   */
  PERMISSION_ATTRIBUTE_SIMPLE_TYPE_WRITE(4),
  /**
   * <code>PERMISSION_ATTRIBUTE_SIMPLE_TYPE_EXECUTE = 5;</code>
   */
  PERMISSION_ATTRIBUTE_SIMPLE_TYPE_EXECUTE(5),
  UNRECOGNIZED(-1),
  ;

  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      PermissionAttributeSimpleType.class.getName());
  }
  /**
   * <pre>
   * Default
   * </pre>
   *
   * <code>PERMISSION_ATTRIBUTE_SIMPLE_TYPE_UNSPECIFIED = 0;</code>
   */
  public static final int PERMISSION_ATTRIBUTE_SIMPLE_TYPE_UNSPECIFIED_VALUE = 0;
  /**
   * <code>PERMISSION_ATTRIBUTE_SIMPLE_TYPE_AUTHOR = 1;</code>
   */
  public static final int PERMISSION_ATTRIBUTE_SIMPLE_TYPE_AUTHOR_VALUE = 1;
  /**
   * <code>PERMISSION_ATTRIBUTE_SIMPLE_TYPE_OWNER = 2;</code>
   */
  public static final int PERMISSION_ATTRIBUTE_SIMPLE_TYPE_OWNER_VALUE = 2;
  /**
   * <code>PERMISSION_ATTRIBUTE_SIMPLE_TYPE_READ = 3;</code>
   */
  public static final int PERMISSION_ATTRIBUTE_SIMPLE_TYPE_READ_VALUE = 3;
  /**
   * <code>PERMISSION_ATTRIBUTE_SIMPLE_TYPE_WRITE = 4;</code>
   */
  public static final int PERMISSION_ATTRIBUTE_SIMPLE_TYPE_WRITE_VALUE = 4;
  /**
   * <code>PERMISSION_ATTRIBUTE_SIMPLE_TYPE_EXECUTE = 5;</code>
   */
  public static final int PERMISSION_ATTRIBUTE_SIMPLE_TYPE_EXECUTE_VALUE = 5;


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
  public static PermissionAttributeSimpleType valueOf(int value) {
    return forNumber(value);
  }

  /**
   * @param value The numeric wire value of the corresponding enum entry.
   * @return The enum associated with the given numeric wire value.
   */
  public static PermissionAttributeSimpleType forNumber(int value) {
    switch (value) {
      case 0: return PERMISSION_ATTRIBUTE_SIMPLE_TYPE_UNSPECIFIED;
      case 1: return PERMISSION_ATTRIBUTE_SIMPLE_TYPE_AUTHOR;
      case 2: return PERMISSION_ATTRIBUTE_SIMPLE_TYPE_OWNER;
      case 3: return PERMISSION_ATTRIBUTE_SIMPLE_TYPE_READ;
      case 4: return PERMISSION_ATTRIBUTE_SIMPLE_TYPE_WRITE;
      case 5: return PERMISSION_ATTRIBUTE_SIMPLE_TYPE_EXECUTE;
      default: return null;
    }
  }

  public static com.google.protobuf.Internal.EnumLiteMap<PermissionAttributeSimpleType>
      internalGetValueMap() {
    return internalValueMap;
  }
  private static final com.google.protobuf.Internal.EnumLiteMap<
      PermissionAttributeSimpleType> internalValueMap =
        new com.google.protobuf.Internal.EnumLiteMap<PermissionAttributeSimpleType>() {
          public PermissionAttributeSimpleType findValueByNumber(int number) {
            return PermissionAttributeSimpleType.forNumber(number);
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
    return uk.gov.gchq.event_logging.v4.PermissionAttributeSimpleTypeProto.getDescriptor().getEnumTypes().get(0);
  }

  private static final PermissionAttributeSimpleType[] VALUES = values();

  public static PermissionAttributeSimpleType valueOf(
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

  private PermissionAttributeSimpleType(int value) {
    this.value = value;
  }

  // @@protoc_insertion_point(enum_scope:event_logging.v4.PermissionAttributeSimpleType)
}

