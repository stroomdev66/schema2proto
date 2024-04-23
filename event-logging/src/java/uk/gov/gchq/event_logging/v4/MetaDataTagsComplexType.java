// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/meta_data_tags_complex_type.proto

// Protobuf Java Version: 4.26.1
package uk.gov.gchq.event_logging.v4;

/**
 * <pre>
 * Metadata tags that can be used for additional object tagging or categorisation. Object tagging allows for the labelling (or filtering) of objects using words that label, categorise or group similar items, using a taxonomy defined outside this schema. For example, an email could be tagged with tags like 'internal', 'spam', 'external', 'rich-content', etc.
 * </pre>
 *
 * Protobuf type {@code event_logging.v4.MetaDataTagsComplexType}
 */
public final class MetaDataTagsComplexType extends
    com.google.protobuf.GeneratedMessage implements
    // @@protoc_insertion_point(message_implements:event_logging.v4.MetaDataTagsComplexType)
    MetaDataTagsComplexTypeOrBuilder {
private static final long serialVersionUID = 0L;
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      MetaDataTagsComplexType.class.getName());
  }
  // Use MetaDataTagsComplexType.newBuilder() to construct.
  private MetaDataTagsComplexType(com.google.protobuf.GeneratedMessage.Builder<?> builder) {
    super(builder);
  }
  private MetaDataTagsComplexType() {
    tag_ =
        com.google.protobuf.LazyStringArrayList.emptyList();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return uk.gov.gchq.event_logging.v4.MetaDataTagsComplexTypeProto.internal_static_event_logging_v4_MetaDataTagsComplexType_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return uk.gov.gchq.event_logging.v4.MetaDataTagsComplexTypeProto.internal_static_event_logging_v4_MetaDataTagsComplexType_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType.class, uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType.Builder.class);
  }

  public static final int TAG_FIELD_NUMBER = 1;
  @SuppressWarnings("serial")
  private com.google.protobuf.LazyStringArrayList tag_ =
      com.google.protobuf.LazyStringArrayList.emptyList();
  /**
   * <pre>
   * A categorisation tag or label
   * </pre>
   *
   * <code>repeated string tag = 1 [json_name = "tag", (.buf.validate.field) = { ... }</code>
   * @return A list containing the tag.
   */
  public com.google.protobuf.ProtocolStringList
      getTagList() {
    return tag_;
  }
  /**
   * <pre>
   * A categorisation tag or label
   * </pre>
   *
   * <code>repeated string tag = 1 [json_name = "tag", (.buf.validate.field) = { ... }</code>
   * @return The count of tag.
   */
  public int getTagCount() {
    return tag_.size();
  }
  /**
   * <pre>
   * A categorisation tag or label
   * </pre>
   *
   * <code>repeated string tag = 1 [json_name = "tag", (.buf.validate.field) = { ... }</code>
   * @param index The index of the element to return.
   * @return The tag at the given index.
   */
  public java.lang.String getTag(int index) {
    return tag_.get(index);
  }
  /**
   * <pre>
   * A categorisation tag or label
   * </pre>
   *
   * <code>repeated string tag = 1 [json_name = "tag", (.buf.validate.field) = { ... }</code>
   * @param index The index of the value to return.
   * @return The bytes of the tag at the given index.
   */
  public com.google.protobuf.ByteString
      getTagBytes(int index) {
    return tag_.getByteString(index);
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
    for (int i = 0; i < tag_.size(); i++) {
      com.google.protobuf.GeneratedMessage.writeString(output, 1, tag_.getRaw(i));
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    {
      int dataSize = 0;
      for (int i = 0; i < tag_.size(); i++) {
        dataSize += computeStringSizeNoTag(tag_.getRaw(i));
      }
      size += dataSize;
      size += 1 * getTagList().size();
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
    if (!(obj instanceof uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType)) {
      return super.equals(obj);
    }
    uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType other = (uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType) obj;

    if (!getTagList()
        .equals(other.getTagList())) return false;
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
    if (getTagCount() > 0) {
      hash = (37 * hash) + TAG_FIELD_NUMBER;
      hash = (53 * hash) + getTagList().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseWithIOException(PARSER, input);
  }
  public static uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessage
        .parseWithIOException(PARSER, input);
  }
  public static uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType parseFrom(
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
  public static Builder newBuilder(uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType prototype) {
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
   * Metadata tags that can be used for additional object tagging or categorisation. Object tagging allows for the labelling (or filtering) of objects using words that label, categorise or group similar items, using a taxonomy defined outside this schema. For example, an email could be tagged with tags like 'internal', 'spam', 'external', 'rich-content', etc.
   * </pre>
   *
   * Protobuf type {@code event_logging.v4.MetaDataTagsComplexType}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessage.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:event_logging.v4.MetaDataTagsComplexType)
      uk.gov.gchq.event_logging.v4.MetaDataTagsComplexTypeOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return uk.gov.gchq.event_logging.v4.MetaDataTagsComplexTypeProto.internal_static_event_logging_v4_MetaDataTagsComplexType_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return uk.gov.gchq.event_logging.v4.MetaDataTagsComplexTypeProto.internal_static_event_logging_v4_MetaDataTagsComplexType_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType.class, uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType.Builder.class);
    }

    // Construct using uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType.newBuilder()
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
      tag_ =
          com.google.protobuf.LazyStringArrayList.emptyList();
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return uk.gov.gchq.event_logging.v4.MetaDataTagsComplexTypeProto.internal_static_event_logging_v4_MetaDataTagsComplexType_descriptor;
    }

    @java.lang.Override
    public uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType getDefaultInstanceForType() {
      return uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType.getDefaultInstance();
    }

    @java.lang.Override
    public uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType build() {
      uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType buildPartial() {
      uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType result = new uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        tag_.makeImmutable();
        result.tag_ = tag_;
      }
    }

    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType) {
        return mergeFrom((uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType other) {
      if (other == uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType.getDefaultInstance()) return this;
      if (!other.tag_.isEmpty()) {
        if (tag_.isEmpty()) {
          tag_ = other.tag_;
          bitField0_ |= 0x00000001;
        } else {
          ensureTagIsMutable();
          tag_.addAll(other.tag_);
        }
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
              java.lang.String s = input.readStringRequireUtf8();
              ensureTagIsMutable();
              tag_.add(s);
              break;
            } // case 10
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

    private com.google.protobuf.LazyStringArrayList tag_ =
        com.google.protobuf.LazyStringArrayList.emptyList();
    private void ensureTagIsMutable() {
      if (!tag_.isModifiable()) {
        tag_ = new com.google.protobuf.LazyStringArrayList(tag_);
      }
      bitField0_ |= 0x00000001;
    }
    /**
     * <pre>
     * A categorisation tag or label
     * </pre>
     *
     * <code>repeated string tag = 1 [json_name = "tag", (.buf.validate.field) = { ... }</code>
     * @return A list containing the tag.
     */
    public com.google.protobuf.ProtocolStringList
        getTagList() {
      tag_.makeImmutable();
      return tag_;
    }
    /**
     * <pre>
     * A categorisation tag or label
     * </pre>
     *
     * <code>repeated string tag = 1 [json_name = "tag", (.buf.validate.field) = { ... }</code>
     * @return The count of tag.
     */
    public int getTagCount() {
      return tag_.size();
    }
    /**
     * <pre>
     * A categorisation tag or label
     * </pre>
     *
     * <code>repeated string tag = 1 [json_name = "tag", (.buf.validate.field) = { ... }</code>
     * @param index The index of the element to return.
     * @return The tag at the given index.
     */
    public java.lang.String getTag(int index) {
      return tag_.get(index);
    }
    /**
     * <pre>
     * A categorisation tag or label
     * </pre>
     *
     * <code>repeated string tag = 1 [json_name = "tag", (.buf.validate.field) = { ... }</code>
     * @param index The index of the value to return.
     * @return The bytes of the tag at the given index.
     */
    public com.google.protobuf.ByteString
        getTagBytes(int index) {
      return tag_.getByteString(index);
    }
    /**
     * <pre>
     * A categorisation tag or label
     * </pre>
     *
     * <code>repeated string tag = 1 [json_name = "tag", (.buf.validate.field) = { ... }</code>
     * @param index The index to set the value at.
     * @param value The tag to set.
     * @return This builder for chaining.
     */
    public Builder setTag(
        int index, java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      ensureTagIsMutable();
      tag_.set(index, value);
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * A categorisation tag or label
     * </pre>
     *
     * <code>repeated string tag = 1 [json_name = "tag", (.buf.validate.field) = { ... }</code>
     * @param value The tag to add.
     * @return This builder for chaining.
     */
    public Builder addTag(
        java.lang.String value) {
      if (value == null) { throw new NullPointerException(); }
      ensureTagIsMutable();
      tag_.add(value);
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * A categorisation tag or label
     * </pre>
     *
     * <code>repeated string tag = 1 [json_name = "tag", (.buf.validate.field) = { ... }</code>
     * @param values The tag to add.
     * @return This builder for chaining.
     */
    public Builder addAllTag(
        java.lang.Iterable<java.lang.String> values) {
      ensureTagIsMutable();
      com.google.protobuf.AbstractMessageLite.Builder.addAll(
          values, tag_);
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * A categorisation tag or label
     * </pre>
     *
     * <code>repeated string tag = 1 [json_name = "tag", (.buf.validate.field) = { ... }</code>
     * @return This builder for chaining.
     */
    public Builder clearTag() {
      tag_ =
        com.google.protobuf.LazyStringArrayList.emptyList();
      bitField0_ = (bitField0_ & ~0x00000001);;
      onChanged();
      return this;
    }
    /**
     * <pre>
     * A categorisation tag or label
     * </pre>
     *
     * <code>repeated string tag = 1 [json_name = "tag", (.buf.validate.field) = { ... }</code>
     * @param value The bytes of the tag to add.
     * @return This builder for chaining.
     */
    public Builder addTagBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      checkByteStringIsUtf8(value);
      ensureTagIsMutable();
      tag_.add(value);
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }

    // @@protoc_insertion_point(builder_scope:event_logging.v4.MetaDataTagsComplexType)
  }

  // @@protoc_insertion_point(class_scope:event_logging.v4.MetaDataTagsComplexType)
  private static final uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType();
  }

  public static uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<MetaDataTagsComplexType>
      PARSER = new com.google.protobuf.AbstractParser<MetaDataTagsComplexType>() {
    @java.lang.Override
    public MetaDataTagsComplexType parsePartialFrom(
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

  public static com.google.protobuf.Parser<MetaDataTagsComplexType> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<MetaDataTagsComplexType> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public uk.gov.gchq.event_logging.v4.MetaDataTagsComplexType getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

