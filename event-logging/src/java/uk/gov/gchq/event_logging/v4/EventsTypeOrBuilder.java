// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/event_logging.proto

// Protobuf Java Version: 4.26.1
package uk.gov.gchq.event_logging.v4;

public interface EventsTypeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:event_logging.v4.EventsType)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * The version of the schema that this document conforms to.
   * </pre>
   *
   * <code>.event_logging.v4.VersionSimpleType version = 1 [json_name = "version", (.buf.validate.field) = { ... }</code>
   * @return The enum numeric value on the wire for version.
   */
  int getVersionValue();
  /**
   * <pre>
   * The version of the schema that this document conforms to.
   * </pre>
   *
   * <code>.event_logging.v4.VersionSimpleType version = 1 [json_name = "version", (.buf.validate.field) = { ... }</code>
   * @return The version.
   */
  uk.gov.gchq.event_logging.v4.VersionSimpleType getVersion();

  /**
   * <pre>
   * This element contains data relating to the sharing of a set of events between different systems or organisations. The data contained within this element will confirm to a specification defined outside of this schema.
   * </pre>
   *
   * <code>.event_logging.v4.AnyContentComplexType sharing_data = 2 [json_name = "sharingData"];</code>
   * @return Whether the sharingData field is set.
   */
  boolean hasSharingData();
  /**
   * <pre>
   * This element contains data relating to the sharing of a set of events between different systems or organisations. The data contained within this element will confirm to a specification defined outside of this schema.
   * </pre>
   *
   * <code>.event_logging.v4.AnyContentComplexType sharing_data = 2 [json_name = "sharingData"];</code>
   * @return The sharingData.
   */
  uk.gov.gchq.event_logging.v4.AnyContentComplexType getSharingData();
  /**
   * <pre>
   * This element contains data relating to the sharing of a set of events between different systems or organisations. The data contained within this element will confirm to a specification defined outside of this schema.
   * </pre>
   *
   * <code>.event_logging.v4.AnyContentComplexType sharing_data = 2 [json_name = "sharingData"];</code>
   */
  uk.gov.gchq.event_logging.v4.AnyContentComplexTypeOrBuilder getSharingDataOrBuilder();

  /**
   * <pre>
   * A single event that has occurred and been recorded.
   * </pre>
   *
   * <code>repeated .event_logging.v4.EventsType.EventType event = 3 [json_name = "event", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<uk.gov.gchq.event_logging.v4.EventsType.EventType> 
      getEventList();
  /**
   * <pre>
   * A single event that has occurred and been recorded.
   * </pre>
   *
   * <code>repeated .event_logging.v4.EventsType.EventType event = 3 [json_name = "event", (.buf.validate.field) = { ... }</code>
   */
  uk.gov.gchq.event_logging.v4.EventsType.EventType getEvent(int index);
  /**
   * <pre>
   * A single event that has occurred and been recorded.
   * </pre>
   *
   * <code>repeated .event_logging.v4.EventsType.EventType event = 3 [json_name = "event", (.buf.validate.field) = { ... }</code>
   */
  int getEventCount();
  /**
   * <pre>
   * A single event that has occurred and been recorded.
   * </pre>
   *
   * <code>repeated .event_logging.v4.EventsType.EventType event = 3 [json_name = "event", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<? extends uk.gov.gchq.event_logging.v4.EventsType.EventTypeOrBuilder> 
      getEventOrBuilderList();
  /**
   * <pre>
   * A single event that has occurred and been recorded.
   * </pre>
   *
   * <code>repeated .event_logging.v4.EventsType.EventType event = 3 [json_name = "event", (.buf.validate.field) = { ... }</code>
   */
  uk.gov.gchq.event_logging.v4.EventsType.EventTypeOrBuilder getEventOrBuilder(
      int index);
}
