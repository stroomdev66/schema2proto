// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/authenticate_outcome_reason_simple_type.proto
// Protobuf Java Version: 4.26.1

package event_logging.v4;

public final class AuthenticateOutcomeReasonSimpleTypeOuterClass {
  private AuthenticateOutcomeReasonSimpleTypeOuterClass() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      AuthenticateOutcomeReasonSimpleTypeOuterClass.class.getName());
  }
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  /**
   * <pre>
   * The types of outcome from an authentication event.
   * </pre>
   *
   * Protobuf enum {@code event_logging.v4.AuthenticateOutcomeReasonSimpleType}
   */
  public enum AuthenticateOutcomeReasonSimpleType
      implements com.google.protobuf.ProtocolMessageEnum {
    /**
     * <pre>
     * Default
     * </pre>
     *
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_UNSPECIFIED = 0;</code>
     */
    AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_UNSPECIFIED(0),
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_USERNAME_OR_PASSWORD = 1;</code>
     */
    AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_USERNAME_OR_PASSWORD(1),
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_USERNAME = 2;</code>
     */
    AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_USERNAME(2),
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_PASSWORD = 3;</code>
     */
    AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_PASSWORD(3),
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_EXPIRED_CERTIFICATE = 4;</code>
     */
    AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_EXPIRED_CERTIFICATE(4),
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_REVOKED_CERTIFICATE = 5;</code>
     */
    AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_REVOKED_CERTIFICATE(5),
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_CA = 6;</code>
     */
    AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_CA(6),
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_EXPIRED_CA = 7;</code>
     */
    AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_EXPIRED_CA(7),
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_ACCOUNT_LOCKED = 8;</code>
     */
    AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_ACCOUNT_LOCKED(8),
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_ACCOUNT_NOT_VALID_FOR_LOGIN_TYPE = 9;</code>
     */
    AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_ACCOUNT_NOT_VALID_FOR_LOGIN_TYPE(9),
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_OTHER = 10;</code>
     */
    AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_OTHER(10),
    UNRECOGNIZED(-1),
    ;

    static {
      com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
        com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
        /* major= */ 4,
        /* minor= */ 26,
        /* patch= */ 1,
        /* suffix= */ "",
        AuthenticateOutcomeReasonSimpleType.class.getName());
    }
    /**
     * <pre>
     * Default
     * </pre>
     *
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_UNSPECIFIED = 0;</code>
     */
    public static final int AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_UNSPECIFIED_VALUE = 0;
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_USERNAME_OR_PASSWORD = 1;</code>
     */
    public static final int AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_USERNAME_OR_PASSWORD_VALUE = 1;
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_USERNAME = 2;</code>
     */
    public static final int AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_USERNAME_VALUE = 2;
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_PASSWORD = 3;</code>
     */
    public static final int AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_PASSWORD_VALUE = 3;
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_EXPIRED_CERTIFICATE = 4;</code>
     */
    public static final int AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_EXPIRED_CERTIFICATE_VALUE = 4;
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_REVOKED_CERTIFICATE = 5;</code>
     */
    public static final int AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_REVOKED_CERTIFICATE_VALUE = 5;
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_CA = 6;</code>
     */
    public static final int AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_CA_VALUE = 6;
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_EXPIRED_CA = 7;</code>
     */
    public static final int AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_EXPIRED_CA_VALUE = 7;
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_ACCOUNT_LOCKED = 8;</code>
     */
    public static final int AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_ACCOUNT_LOCKED_VALUE = 8;
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_ACCOUNT_NOT_VALID_FOR_LOGIN_TYPE = 9;</code>
     */
    public static final int AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_ACCOUNT_NOT_VALID_FOR_LOGIN_TYPE_VALUE = 9;
    /**
     * <code>AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_OTHER = 10;</code>
     */
    public static final int AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_OTHER_VALUE = 10;


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
    public static AuthenticateOutcomeReasonSimpleType valueOf(int value) {
      return forNumber(value);
    }

    /**
     * @param value The numeric wire value of the corresponding enum entry.
     * @return The enum associated with the given numeric wire value.
     */
    public static AuthenticateOutcomeReasonSimpleType forNumber(int value) {
      switch (value) {
        case 0: return AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_UNSPECIFIED;
        case 1: return AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_USERNAME_OR_PASSWORD;
        case 2: return AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_USERNAME;
        case 3: return AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_PASSWORD;
        case 4: return AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_EXPIRED_CERTIFICATE;
        case 5: return AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_REVOKED_CERTIFICATE;
        case 6: return AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_INCORRECT_CA;
        case 7: return AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_EXPIRED_CA;
        case 8: return AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_ACCOUNT_LOCKED;
        case 9: return AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_ACCOUNT_NOT_VALID_FOR_LOGIN_TYPE;
        case 10: return AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_OTHER;
        default: return null;
      }
    }

    public static com.google.protobuf.Internal.EnumLiteMap<AuthenticateOutcomeReasonSimpleType>
        internalGetValueMap() {
      return internalValueMap;
    }
    private static final com.google.protobuf.Internal.EnumLiteMap<
        AuthenticateOutcomeReasonSimpleType> internalValueMap =
          new com.google.protobuf.Internal.EnumLiteMap<AuthenticateOutcomeReasonSimpleType>() {
            public AuthenticateOutcomeReasonSimpleType findValueByNumber(int number) {
              return AuthenticateOutcomeReasonSimpleType.forNumber(number);
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
      return event_logging.v4.AuthenticateOutcomeReasonSimpleTypeOuterClass.getDescriptor().getEnumTypes().get(0);
    }

    private static final AuthenticateOutcomeReasonSimpleType[] VALUES = values();

    public static AuthenticateOutcomeReasonSimpleType valueOf(
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

    private AuthenticateOutcomeReasonSimpleType(int value) {
      this.value = value;
    }

    // @@protoc_insertion_point(enum_scope:event_logging.v4.AuthenticateOutcomeReasonSimpleType)
  }


  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n>event_logging/v4/authenticate_outcome_" +
      "reason_simple_type.proto\022\020event_logging." +
      "v4\032\033buf/validate/validate.proto*\333\005\n#Auth" +
      "enticateOutcomeReasonSimpleType\0227\n3AUTHE" +
      "NTICATE_OUTCOME_REASON_SIMPLE_TYPE_UNSPE" +
      "CIFIED\020\000\022J\nFAUTHENTICATE_OUTCOME_REASON_" +
      "SIMPLE_TYPE_INCORRECT_USERNAME_OR_PASSWO" +
      "RD\020\001\022>\n:AUTHENTICATE_OUTCOME_REASON_SIMP" +
      "LE_TYPE_INCORRECT_USERNAME\020\002\022>\n:AUTHENTI" +
      "CATE_OUTCOME_REASON_SIMPLE_TYPE_INCORREC" +
      "T_PASSWORD\020\003\022?\n;AUTHENTICATE_OUTCOME_REA" +
      "SON_SIMPLE_TYPE_EXPIRED_CERTIFICATE\020\004\022?\n" +
      ";AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE" +
      "_REVOKED_CERTIFICATE\020\005\0228\n4AUTHENTICATE_O" +
      "UTCOME_REASON_SIMPLE_TYPE_INCORRECT_CA\020\006" +
      "\0226\n2AUTHENTICATE_OUTCOME_REASON_SIMPLE_T" +
      "YPE_EXPIRED_CA\020\007\022:\n6AUTHENTICATE_OUTCOME" +
      "_REASON_SIMPLE_TYPE_ACCOUNT_LOCKED\020\010\022L\nH" +
      "AUTHENTICATE_OUTCOME_REASON_SIMPLE_TYPE_" +
      "ACCOUNT_NOT_VALID_FOR_LOGIN_TYPE\020\t\0221\n-AU" +
      "THENTICATE_OUTCOME_REASON_SIMPLE_TYPE_OT" +
      "HER\020\nb\006proto3"
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
