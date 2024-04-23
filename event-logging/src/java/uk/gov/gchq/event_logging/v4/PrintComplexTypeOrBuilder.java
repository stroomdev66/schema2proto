// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/print_complex_type.proto

// Protobuf Java Version: 4.26.1
package uk.gov.gchq.event_logging.v4;

public interface PrintComplexTypeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:event_logging.v4.PrintComplexType)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * The print action that was performed, e.g. CreateJob.
   * </pre>
   *
   * <code>.event_logging.v4.PrintActionSimpleType action = 1 [json_name = "action", (.buf.validate.field) = { ... }</code>
   * @return The enum numeric value on the wire for action.
   */
  int getActionValue();
  /**
   * <pre>
   * The print action that was performed, e.g. CreateJob.
   * </pre>
   *
   * <code>.event_logging.v4.PrintActionSimpleType action = 1 [json_name = "action", (.buf.validate.field) = { ... }</code>
   * @return The action.
   */
  uk.gov.gchq.event_logging.v4.PrintActionSimpleType getAction();

  /**
   * <pre>
   * Describes the print job that the event relates to.
   * </pre>
   *
   * <code>.event_logging.v4.PrintComplexType.PrintJobType print_job = 2 [json_name = "printJob", (.buf.validate.field) = { ... }</code>
   * @return Whether the printJob field is set.
   */
  boolean hasPrintJob();
  /**
   * <pre>
   * Describes the print job that the event relates to.
   * </pre>
   *
   * <code>.event_logging.v4.PrintComplexType.PrintJobType print_job = 2 [json_name = "printJob", (.buf.validate.field) = { ... }</code>
   * @return The printJob.
   */
  uk.gov.gchq.event_logging.v4.PrintComplexType.PrintJobType getPrintJob();
  /**
   * <pre>
   * Describes the print job that the event relates to.
   * </pre>
   *
   * <code>.event_logging.v4.PrintComplexType.PrintJobType print_job = 2 [json_name = "printJob", (.buf.validate.field) = { ... }</code>
   */
  uk.gov.gchq.event_logging.v4.PrintComplexType.PrintJobTypeOrBuilder getPrintJobOrBuilder();

  /**
   * <pre>
   * The printer settings that are to be used for the print job.
   * </pre>
   *
   * <code>.event_logging.v4.PrintComplexType.PrintSettingsType print_settings = 3 [json_name = "printSettings"];</code>
   * @return Whether the printSettings field is set.
   */
  boolean hasPrintSettings();
  /**
   * <pre>
   * The printer settings that are to be used for the print job.
   * </pre>
   *
   * <code>.event_logging.v4.PrintComplexType.PrintSettingsType print_settings = 3 [json_name = "printSettings"];</code>
   * @return The printSettings.
   */
  uk.gov.gchq.event_logging.v4.PrintComplexType.PrintSettingsType getPrintSettings();
  /**
   * <pre>
   * The printer settings that are to be used for the print job.
   * </pre>
   *
   * <code>.event_logging.v4.PrintComplexType.PrintSettingsType print_settings = 3 [json_name = "printSettings"];</code>
   */
  uk.gov.gchq.event_logging.v4.PrintComplexType.PrintSettingsTypeOrBuilder getPrintSettingsOrBuilder();

  /**
   * <pre>
   * Describes the printer to use for the print job.
   * </pre>
   *
   * <code>.event_logging.v4.DeviceComplexType printer = 4 [json_name = "printer"];</code>
   * @return Whether the printer field is set.
   */
  boolean hasPrinter();
  /**
   * <pre>
   * Describes the printer to use for the print job.
   * </pre>
   *
   * <code>.event_logging.v4.DeviceComplexType printer = 4 [json_name = "printer"];</code>
   * @return The printer.
   */
  uk.gov.gchq.event_logging.v4.DeviceComplexType getPrinter();
  /**
   * <pre>
   * Describes the printer to use for the print job.
   * </pre>
   *
   * <code>.event_logging.v4.DeviceComplexType printer = 4 [json_name = "printer"];</code>
   */
  uk.gov.gchq.event_logging.v4.DeviceComplexTypeOrBuilder getPrinterOrBuilder();

  /**
   * <pre>
   * Used to determine if the action was successful. If omitted it is assumed that the event was successful and was permitted.
   * </pre>
   *
   * <code>.event_logging.v4.OutcomeComplexType outcome = 5 [json_name = "outcome"];</code>
   * @return Whether the outcome field is set.
   */
  boolean hasOutcome();
  /**
   * <pre>
   * Used to determine if the action was successful. If omitted it is assumed that the event was successful and was permitted.
   * </pre>
   *
   * <code>.event_logging.v4.OutcomeComplexType outcome = 5 [json_name = "outcome"];</code>
   * @return The outcome.
   */
  uk.gov.gchq.event_logging.v4.OutcomeComplexType getOutcome();
  /**
   * <pre>
   * Used to determine if the action was successful. If omitted it is assumed that the event was successful and was permitted.
   * </pre>
   *
   * <code>.event_logging.v4.OutcomeComplexType outcome = 5 [json_name = "outcome"];</code>
   */
  uk.gov.gchq.event_logging.v4.OutcomeComplexTypeOrBuilder getOutcomeOrBuilder();

  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 6 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<uk.gov.gchq.event_logging.v4.DataComplexType> 
      getDataList();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 6 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  uk.gov.gchq.event_logging.v4.DataComplexType getData(int index);
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 6 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  int getDataCount();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 6 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<? extends uk.gov.gchq.event_logging.v4.DataComplexTypeOrBuilder> 
      getDataOrBuilderList();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 6 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  uk.gov.gchq.event_logging.v4.DataComplexTypeOrBuilder getDataOrBuilder(
      int index);
}
