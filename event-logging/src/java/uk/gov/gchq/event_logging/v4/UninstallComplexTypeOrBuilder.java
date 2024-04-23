// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/uninstall_complex_type.proto

// Protobuf Java Version: 4.26.1
package uk.gov.gchq.event_logging.v4;

public interface UninstallComplexTypeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:event_logging.v4.UninstallComplexType)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * Description of the software that was installed/uninstalled.
   * </pre>
   *
   * <code>.event_logging.v4.SoftwareComplexType software = 1 [json_name = "software", (.buf.validate.field) = { ... }</code>
   * @return Whether the software field is set.
   */
  boolean hasSoftware();
  /**
   * <pre>
   * Description of the software that was installed/uninstalled.
   * </pre>
   *
   * <code>.event_logging.v4.SoftwareComplexType software = 1 [json_name = "software", (.buf.validate.field) = { ... }</code>
   * @return The software.
   */
  uk.gov.gchq.event_logging.v4.SoftwareComplexType getSoftware();
  /**
   * <pre>
   * Description of the software that was installed/uninstalled.
   * </pre>
   *
   * <code>.event_logging.v4.SoftwareComplexType software = 1 [json_name = "software", (.buf.validate.field) = { ... }</code>
   */
  uk.gov.gchq.event_logging.v4.SoftwareComplexTypeOrBuilder getSoftwareOrBuilder();

  /**
   * <pre>
   * Description of the hardware that was installed/uninstalled.
   * </pre>
   *
   * <code>.event_logging.v4.HardwareComplexType hardware = 2 [json_name = "hardware", (.buf.validate.field) = { ... }</code>
   * @return Whether the hardware field is set.
   */
  boolean hasHardware();
  /**
   * <pre>
   * Description of the hardware that was installed/uninstalled.
   * </pre>
   *
   * <code>.event_logging.v4.HardwareComplexType hardware = 2 [json_name = "hardware", (.buf.validate.field) = { ... }</code>
   * @return The hardware.
   */
  uk.gov.gchq.event_logging.v4.HardwareComplexType getHardware();
  /**
   * <pre>
   * Description of the hardware that was installed/uninstalled.
   * </pre>
   *
   * <code>.event_logging.v4.HardwareComplexType hardware = 2 [json_name = "hardware", (.buf.validate.field) = { ... }</code>
   */
  uk.gov.gchq.event_logging.v4.HardwareComplexTypeOrBuilder getHardwareOrBuilder();

  /**
   * <pre>
   * Description of the media that was installed/uninstalled.
   * </pre>
   *
   * <code>.event_logging.v4.MediaComplexType media = 3 [json_name = "media", (.buf.validate.field) = { ... }</code>
   * @return Whether the media field is set.
   */
  boolean hasMedia();
  /**
   * <pre>
   * Description of the media that was installed/uninstalled.
   * </pre>
   *
   * <code>.event_logging.v4.MediaComplexType media = 3 [json_name = "media", (.buf.validate.field) = { ... }</code>
   * @return The media.
   */
  uk.gov.gchq.event_logging.v4.MediaComplexType getMedia();
  /**
   * <pre>
   * Description of the media that was installed/uninstalled.
   * </pre>
   *
   * <code>.event_logging.v4.MediaComplexType media = 3 [json_name = "media", (.buf.validate.field) = { ... }</code>
   */
  uk.gov.gchq.event_logging.v4.MediaComplexTypeOrBuilder getMediaOrBuilder();

  /**
   * <pre>
   * The outcome of the (un)installation.
   * </pre>
   *
   * <code>.event_logging.v4.OutcomeComplexType outcome = 11 [json_name = "outcome"];</code>
   * @return Whether the outcome field is set.
   */
  boolean hasOutcome();
  /**
   * <pre>
   * The outcome of the (un)installation.
   * </pre>
   *
   * <code>.event_logging.v4.OutcomeComplexType outcome = 11 [json_name = "outcome"];</code>
   * @return The outcome.
   */
  uk.gov.gchq.event_logging.v4.OutcomeComplexType getOutcome();
  /**
   * <pre>
   * The outcome of the (un)installation.
   * </pre>
   *
   * <code>.event_logging.v4.OutcomeComplexType outcome = 11 [json_name = "outcome"];</code>
   */
  uk.gov.gchq.event_logging.v4.OutcomeComplexTypeOrBuilder getOutcomeOrBuilder();

  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 12 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<uk.gov.gchq.event_logging.v4.DataComplexType> 
      getDataList();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 12 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  uk.gov.gchq.event_logging.v4.DataComplexType getData(int index);
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 12 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  int getDataCount();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 12 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<? extends uk.gov.gchq.event_logging.v4.DataComplexTypeOrBuilder> 
      getDataOrBuilderList();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 12 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  uk.gov.gchq.event_logging.v4.DataComplexTypeOrBuilder getDataOrBuilder(
      int index);
}
