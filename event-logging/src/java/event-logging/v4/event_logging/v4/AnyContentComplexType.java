// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/any_content_complex_type.proto

// Protobuf Java Version: 4.26.1
package event-logging.v4.event_logging.v4;

/**
 * <pre>
 * This type is used to contain any content conforming to an agreed format/specification that is defined outside this XML Schema.
 * </pre>
 *
 * Protobuf type {@code event_logging.v4.AnyContentComplexType}
 */
public final class AnyContentComplexType extends
    com.google.protobuf.GeneratedMessage implements
    // @@protoc_insertion_point(message_implements:event_logging.v4.AnyContentComplexType)
    AnyContentComplexTypeOrBuilder {
private static final long serialVersionUID = 0L;
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      AnyContentComplexType.class.getName());
  }
  // Use AnyContentComplexType.newBuilder() to construct.
  private AnyContentComplexType(com.google.protobuf.GeneratedMessage.Builder<?> builder) {
    super(builder);
  }
  private AnyContentComplexType() {
    contentType_ = "";
    version_ = "";
    any_ = "";
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return event-logging.v4.event_logging.v4.AnyContentComplexTypeProto.internal_static_event_logging_v4_AnyContentComplexType_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return event-logging.v4.event_logging.v4.AnyContentComplexTypeProto.internal_static_event_logging_v4_AnyContentComplexType_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            event-logging.v4.event_logging.v4.AnyContentComplexType.class, event-logging.v4.event_logging.v4.AnyContentComplexType.Builder.class);
  }

  public static final int CONTENT_TYPE_FIELD_NUMBER = 1;
  @SuppressWarnings("serial")
  private volatile java.lang.Object contentType_ = "";
  /**
   * <pre>
   * String to describe the format type and specification of the content, e.g. JSON or XML. The valid values are defined outside this XML Schema.
   * </pre>
   *
   * <code>string content_type = 1 [json_name = "contentType"];</code>
   * @return The contentType.
   */
  @java.lang.Override
  public java.lang.String getContentType() {
    java.lang.Object ref = contentType_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      contentType_ = s;
      return s;
    }
  }
  /**
   * <pre>
   * String to describe the format type and specification of the content, e.g. JSON or XML. The valid values are defined outside this XML Schema.
   * </pre>
   *
   * <code>string content_type = 1 [json_name = "contentType"];</code>
   * @return The bytes for contentType.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getContentTypeBytes() {
    java.lang.Object ref = contentType_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      contentType_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int VERSION_FIELD_NUMBER = 2;
  @SuppressWarnings("serial")
  private volatile java.lang.Object version_ = "";
  /**
   * <pre>
   * Defines the version of data structure specification.
   * </pre>
   *
   * <code>string version = 2 [json_name = "version"];</code>
   * @return The version.
   */
  @java.lang.Override
  public java.lang.String getVersion() {
    java.lang.Object ref = version_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      version_ = s;
      return s;
    }
  }
  /**
   * <pre>
   * Defines the version of data structure specification.
   * </pre>
   *
   * <code>string version = 2 [json_name = "version"];</code>
   * @return The bytes for version.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getVersionBytes() {
    java.lang.Object ref = version_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      version_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int ANY_FIELD_NUMBER = 3;
  @SuppressWarnings("serial")
  private volatile java.lang.Object any_ = "";
  /**
   * <code>string any = 3 [json_name = "any"];</code>
   * @return The any.
   */
  @java.lang.Override
  public java.lang.String getAny() {
    java.lang.Object ref = any_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      any_ = s;
      return s;
    }
  }
  /**
   * <code>string any = 3 [json_name = "any"];</code>
   * @return The bytes for any.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getAnyBytes() {
    java.lang.Object ref = any_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      any_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (!com.google.protobuf.GeneratedMessage.isStringEmpty(contentType_)) {
      com.google.protobuf.GeneratedMessage.writeString(output, 1, contentType_);
    }
    if (!com.google.protobuf.GeneratedMessage.isStringEmpty(version_)) {
      com.google.protobuf.GeneratedMessage.writeString(output, 2, version_);
    }
    if (!com.google.protobuf.GeneratedMessage.isStringEmpty(any_)) {
      com.google.protobuf.GeneratedMessage.writeString(output, 3, any_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!com.google.protobuf.GeneratedMessage.isStringEmpty(contentType_)) {
      size += com.google.protobuf.GeneratedMessage.computeStringSize(1, contentType_);
    }
    if (!com.google.protobuf.GeneratedMessage.isStringEmpty(version_)) {
      size += com.google.protobuf.GeneratedMessage.computeStringSize(2, version_);
    }
    if (!com.google.protobuf.GeneratedMessage.isStringEmpty(any_)) {
      size += com.google.protobuf.GeneratedMessage.computeStringSize(3, any_);
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof event-logging.v4.event_logging.v4.AnyContentComplexType)) {
      return super.equals(obj);
    }
    event-logging.v4.event_logging.v4.AnyContentComplexType other = (event-logging.v4.event_logging.v4.AnyContentComplexType) obj;

    if (!getContentType()
        .equals(other.getContentType())) return false;
    if (!getVersion()
        .equals(other.getVersion())) return false;
    if (!getAny()
        .equals(other.getAny())) return false;
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + CONTENT_TYPE_FIELD_NUMBER;
    hash = (53 * hash) + getContentType().hashCode();
    hash = (37 * hash) + VERSION_FIELD_NUMBER;
    hash = (53 * hash) + getVersion().hashCode();
    hash = (37 * hash) + ANY_FIELD_NUMBER;
    hash = (53 * hash) + getAny().hashCode();
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static event-logging.v4.event_logging.v4.AnyContentComplexType parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static event-logging.v4.event_logging.v4.AnyContentComplexType parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static event-logging.v4.event_logging.v4.AnyContentComplexType parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static event-logging.v4.event_logging.v4.AnyContentComplexType parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static event-logging.v4.event_logging.v4.AnyContentComplexType parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static event-logging.v4.event_logging.v4.AnyContentComplexType parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static event-logging.v4.event_logging.v4.AnyContentComplexType parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseWithIOException(PARSER, input);
  }
  public static event-logging.v4.event_logging.v4.AnyContentComplexType parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static event-logging.v4.event_logging.v4.AnyContentComplexType parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static event-logging.v4.event_logging.v4.AnyContentComplexType parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static event-logging.v4.event_logging.v4.AnyContentComplexType parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseWithIOException(PARSER, input);
  }
  public static event-logging.v4.event_logging.v4.AnyContentComplexType parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(event-logging.v4.event_logging.v4.AnyContentComplexType prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessage.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * <pre>
   * This type is used to contain any content conforming to an agreed format/specification that is defined outside this XML Schema.
   * </pre>
   *
   * Protobuf type {@code event_logging.v4.AnyContentComplexType}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessage.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:event_logging.v4.AnyContentComplexType)
      event-logging.v4.event_logging.v4.AnyContentComplexTypeOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return event-logging.v4.event_logging.v4.AnyContentComplexTypeProto.internal_static_event_logging_v4_AnyContentComplexType_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return event-logging.v4.event_logging.v4.AnyContentComplexTypeProto.internal_static_event_logging_v4_AnyContentComplexType_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              event-logging.v4.event_logging.v4.AnyContentComplexType.class, event-logging.v4.event_logging.v4.AnyContentComplexType.Builder.class);
    }

    // Construct using event-logging.v4.event_logging.v4.AnyContentComplexType.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessage.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      bitField0_ = 0;
      contentType_ = "";
      version_ = "";
      any_ = "";
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return event-logging.v4.event_logging.v4.AnyContentComplexTypeProto.internal_static_event_logging_v4_AnyContentComplexType_descriptor;
    }

    @java.lang.Override
    public event-logging.v4.event_logging.v4.AnyContentComplexType getDefaultInstanceForType() {
      return event-logging.v4.event_logging.v4.AnyContentComplexType.getDefaultInstance();
    }

    @java.lang.Override
    public event-logging.v4.event_logging.v4.AnyContentComplexType build() {
      event-logging.v4.event_logging.v4.AnyContentComplexType result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public event-logging.v4.event_logging.v4.AnyContentComplexType buildPartial() {
      event-logging.v4.event_logging.v4.AnyContentComplexType result = new event-logging.v4.event_logging.v4.AnyContentComplexType(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(event-logging.v4.event_logging.v4.AnyContentComplexType result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.contentType_ = contentType_;
      }
      if (((from_bitField0_ & 0x00000002) != 0)) {
        result.version_ = version_;
      }
      if (((from_bitField0_ & 0x00000004) != 0)) {
        result.any_ = any_;
      }
    }

    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof event-logging.v4.event_logging.v4.AnyContentComplexType) {
        return mergeFrom((event-logging.v4.event_logging.v4.AnyContentComplexType)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(event-logging.v4.event_logging.v4.AnyContentComplexType other) {
      if (other == event-logging.v4.event_logging.v4.AnyContentComplexType.getDefaultInstance()) return this;
      if (!other.getContentType().isEmpty()) {
        contentType_ = other.contentType_;
        bitField0_ |= 0x00000001;
        onChanged();
      }
      if (!other.getVersion().isEmpty()) {
        version_ = other.version_;
        bitField0_ |= 0x00000002;
        onChanged();
      }
      if (!other.getAny().isEmpty()) {
        any_ = other.any_;
        bitField0_ |= 0x00000004;
        onChanged();
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              contentType_ = input.readStringRequireUtf8();
              bitField0_ |= 0x00000001;
              break;
            } // case 10
            case 18: {
              version_ = input.readStringRequireUtf8();
              bitField0_ |= 0x00000002;
              break;
            } // case 18
            case 26: {
              any_ = input.readStringRequireUtf8();
              bitField0_ |= 0x00000004;
              break;
            } // case 26
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }
    private int bitField0_;

    private java.lang.Object contentType_ = "";
    /**
     * <pre>
     * String to describe the format type and specification of the content, e.g. JSON or XML. The valid values are defined outside this XML Schema.
     * </pre>
     *
     * <code>string content_type = 1 [json_name = "contentType"];</code>
     * @return The contentType.
     */
    public java.lang.String getContentType() {
      java.lang.Object ref = contentType_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        contentType_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     * String to describe the format type and specification of the content, e.g. JSON or XML. The valid values are defined outside this XML Schema.
     * </pre>
     *
     * <code>string content_type = 1 [json_name = "contentType"];</code>
     * @return The bytes for contentType.
     */
    public com.google.protobuf.ByteString
        getContentTypeBytes() {
      java.lang.Object ref = contentType_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        contentType_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     * String to describe the format type and specification of the content, e.g. JSON or XML. The valid values are defined outside this XML Schema.
     * </pre>
     *
     * <code>string content_type = 1 [json_name = "contentType"];</code>
     * @param value The contentType to set.
     * @return This builder for chaining.
     */
    public Builder setContentType(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      contentType_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * String to describe the format type and specification of the content, e.g. JSON or XML. The valid values are defined outside this XML Schema.
     * </pre>
     *
     * <code>string content_type = 1 [json_name = "contentType"];</code>
     * @return This builder for chaining.
     */
    public Builder clearContentType() {
      contentType_ = getDefaultInstance().getContentType();
      bitField0_ = (bitField0_ & ~0x00000001);
      onChanged();
      return this;
    }
    /**
     * <pre>
     * String to describe the format type and specification of the content, e.g. JSON or XML. The valid values are defined outside this XML Schema.
     * </pre>
     *
     * <code>string content_type = 1 [json_name = "contentType"];</code>
     * @param value The bytes for contentType to set.
     * @return This builder for chaining.
     */
    public Builder setContentTypeBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      contentType_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }

    private java.lang.Object version_ = "";
    /**
     * <pre>
     * Defines the version of data structure specification.
     * </pre>
     *
     * <code>string version = 2 [json_name = "version"];</code>
     * @return The version.
     */
    public java.lang.String getVersion() {
      java.lang.Object ref = version_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        version_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <pre>
     * Defines the version of data structure specification.
     * </pre>
     *
     * <code>string version = 2 [json_name = "version"];</code>
     * @return The bytes for version.
     */
    public com.google.protobuf.ByteString
        getVersionBytes() {
      java.lang.Object ref = version_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        version_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <pre>
     * Defines the version of data structure specification.
     * </pre>
     *
     * <code>string version = 2 [json_name = "version"];</code>
     * @param value The version to set.
     * @return This builder for chaining.
     */
    public Builder setVersion(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      version_ = value;
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * Defines the version of data structure specification.
     * </pre>
     *
     * <code>string version = 2 [json_name = "version"];</code>
     * @return This builder for chaining.
     */
    public Builder clearVersion() {
      version_ = getDefaultInstance().getVersion();
      bitField0_ = (bitField0_ & ~0x00000002);
      onChanged();
      return this;
    }
    /**
     * <pre>
     * Defines the version of data structure specification.
     * </pre>
     *
     * <code>string version = 2 [json_name = "version"];</code>
     * @param value The bytes for version to set.
     * @return This builder for chaining.
     */
    public Builder setVersionBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      version_ = value;
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }

    private java.lang.Object any_ = "";
    /**
     * <code>string any = 3 [json_name = "any"];</code>
     * @return The any.
     */
    public java.lang.String getAny() {
      java.lang.Object ref = any_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        any_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string any = 3 [json_name = "any"];</code>
     * @return The bytes for any.
     */
    public com.google.protobuf.ByteString
        getAnyBytes() {
      java.lang.Object ref = any_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        any_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string any = 3 [json_name = "any"];</code>
     * @param value The any to set.
     * @return This builder for chaining.
     */
    public Builder setAny(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      any_ = value;
      bitField0_ |= 0x00000004;
      onChanged();
      return this;
    }
    /**
     * <code>string any = 3 [json_name = "any"];</code>
     * @return This builder for chaining.
     */
    public Builder clearAny() {
      any_ = getDefaultInstance().getAny();
      bitField0_ = (bitField0_ & ~0x00000004);
      onChanged();
      return this;
    }
    /**
     * <code>string any = 3 [json_name = "any"];</code>
     * @param value The bytes for any to set.
     * @return This builder for chaining.
     */
    public Builder setAnyBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      any_ = value;
      bitField0_ |= 0x00000004;
      onChanged();
      return this;
    }

    // @@protoc_insertion_point(builder_scope:event_logging.v4.AnyContentComplexType)
  }

  // @@protoc_insertion_point(class_scope:event_logging.v4.AnyContentComplexType)
  private static final event-logging.v4.event_logging.v4.AnyContentComplexType DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new event-logging.v4.event_logging.v4.AnyContentComplexType();
  }

  public static event-logging.v4.event_logging.v4.AnyContentComplexType getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<AnyContentComplexType>
      PARSER = new com.google.protobuf.AbstractParser<AnyContentComplexType>() {
    @java.lang.Override
    public AnyContentComplexType parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<AnyContentComplexType> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<AnyContentComplexType> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public event-logging.v4.event_logging.v4.AnyContentComplexType getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

