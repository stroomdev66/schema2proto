// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/system_complex_type.proto

// Protobuf Java Version: 4.26.1
package uk.gov.gchq.event_logging.v4;

public interface SystemComplexTypeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:event_logging.v4.SystemComplexType)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * The name of the system.
   * </pre>
   *
   * <code>string name = 1 [json_name = "name", (.buf.validate.field) = { ... }</code>
   * @return The name.
   */
  java.lang.String getName();
  /**
   * <pre>
   * The name of the system.
   * </pre>
   *
   * <code>string name = 1 [json_name = "name", (.buf.validate.field) = { ... }</code>
   * @return The bytes for name.
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <pre>
   * An optional description of the system.
   * </pre>
   *
   * <code>string description = 2 [json_name = "description"];</code>
   * @return The description.
   */
  java.lang.String getDescription();
  /**
   * <pre>
   * An optional description of the system.
   * </pre>
   *
   * <code>string description = 2 [json_name = "description"];</code>
   * @return The bytes for description.
   */
  com.google.protobuf.ByteString
      getDescriptionBytes();

  /**
   * <pre>
   * An optional classification or protective marking of the overall system.
   * </pre>
   *
   * <code>.event_logging.v4.ClassificationComplexType classification = 3 [json_name = "classification"];</code>
   * @return Whether the classification field is set.
   */
  boolean hasClassification();
  /**
   * <pre>
   * An optional classification or protective marking of the overall system.
   * </pre>
   *
   * <code>.event_logging.v4.ClassificationComplexType classification = 3 [json_name = "classification"];</code>
   * @return The classification.
   */
  uk.gov.gchq.event_logging.v4.ClassificationComplexType getClassification();
  /**
   * <pre>
   * An optional classification or protective marking of the overall system.
   * </pre>
   *
   * <code>.event_logging.v4.ClassificationComplexType classification = 3 [json_name = "classification"];</code>
   */
  uk.gov.gchq.event_logging.v4.ClassificationComplexTypeOrBuilder getClassificationOrBuilder();

  /**
   * <pre>
   * The environment describes a specific instance of a system. A system may have multiple deployment for various purposes, e.g. a development, reference or operational deployment. An instance may also be site specific e.g. a deployment at a particular data center. The way an environment is described will differ depending on the system and the way it is deployed however a good example would be REF_DC1 to indicate that the environment is a reference deployment in data center 1.
   * </pre>
   *
   * <code>string environment = 4 [json_name = "environment", (.buf.validate.field) = { ... }</code>
   * @return The environment.
   */
  java.lang.String getEnvironment();
  /**
   * <pre>
   * The environment describes a specific instance of a system. A system may have multiple deployment for various purposes, e.g. a development, reference or operational deployment. An instance may also be site specific e.g. a deployment at a particular data center. The way an environment is described will differ depending on the system and the way it is deployed however a good example would be REF_DC1 to indicate that the environment is a reference deployment in data center 1.
   * </pre>
   *
   * <code>string environment = 4 [json_name = "environment", (.buf.validate.field) = { ... }</code>
   * @return The bytes for environment.
   */
  com.google.protobuf.ByteString
      getEnvironmentBytes();

  /**
   * <pre>
   * Describes the organisation that owns or has responsibility for the system.
   * </pre>
   *
   * <code>string organisation = 5 [json_name = "organisation"];</code>
   * @return The organisation.
   */
  java.lang.String getOrganisation();
  /**
   * <pre>
   * Describes the organisation that owns or has responsibility for the system.
   * </pre>
   *
   * <code>string organisation = 5 [json_name = "organisation"];</code>
   * @return The bytes for organisation.
   */
  com.google.protobuf.ByteString
      getOrganisationBytes();

  /**
   * <pre>
   * An optional element to define the domain that the system exists in.
   * </pre>
   *
   * <code>string security_domain = 6 [json_name = "securityDomain"];</code>
   * @return The securityDomain.
   */
  java.lang.String getSecurityDomain();
  /**
   * <pre>
   * An optional element to define the domain that the system exists in.
   * </pre>
   *
   * <code>string security_domain = 6 [json_name = "securityDomain"];</code>
   * @return The bytes for securityDomain.
   */
  com.google.protobuf.ByteString
      getSecurityDomainBytes();

  /**
   * <pre>
   * The version of the system's software, e.g. 1.3.2
   * </pre>
   *
   * <code>string version = 7 [json_name = "version"];</code>
   * @return The version.
   */
  java.lang.String getVersion();
  /**
   * <pre>
   * The version of the system's software, e.g. 1.3.2
   * </pre>
   *
   * <code>string version = 7 [json_name = "version"];</code>
   * @return The bytes for version.
   */
  com.google.protobuf.ByteString
      getVersionBytes();

  /**
   * <pre>
   * Optional tags that can be used for additional tagging or categorisation of the system. These tags allow for the grouping or filtering of similar systems.
   * </pre>
   *
   * <code>.event_logging.v4.SystemComplexType.TagsType tags = 8 [json_name = "tags"];</code>
   * @return Whether the tags field is set.
   */
  boolean hasTags();
  /**
   * <pre>
   * Optional tags that can be used for additional tagging or categorisation of the system. These tags allow for the grouping or filtering of similar systems.
   * </pre>
   *
   * <code>.event_logging.v4.SystemComplexType.TagsType tags = 8 [json_name = "tags"];</code>
   * @return The tags.
   */
  uk.gov.gchq.event_logging.v4.SystemComplexType.TagsType getTags();
  /**
   * <pre>
   * Optional tags that can be used for additional tagging or categorisation of the system. These tags allow for the grouping or filtering of similar systems.
   * </pre>
   *
   * <code>.event_logging.v4.SystemComplexType.TagsType tags = 8 [json_name = "tags"];</code>
   */
  uk.gov.gchq.event_logging.v4.SystemComplexType.TagsTypeOrBuilder getTagsOrBuilder();
}
