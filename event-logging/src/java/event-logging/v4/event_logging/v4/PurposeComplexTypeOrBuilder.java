// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/purpose_complex_type.proto

// Protobuf Java Version: 4.26.1
package event-logging.v4.event_logging.v4;

public interface PurposeComplexTypeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:event_logging.v4.PurposeComplexType)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * The classification of this task.
   * </pre>
   *
   * <code>.event_logging.v4.ClassificationComplexType classification = 1 [json_name = "classification"];</code>
   * @return Whether the classification field is set.
   */
  boolean hasClassification();
  /**
   * <pre>
   * The classification of this task.
   * </pre>
   *
   * <code>.event_logging.v4.ClassificationComplexType classification = 1 [json_name = "classification"];</code>
   * @return The classification.
   */
  event-logging.v4.event_logging.v4.ClassificationComplexType getClassification();
  /**
   * <pre>
   * The classification of this task.
   * </pre>
   *
   * <code>.event_logging.v4.ClassificationComplexType classification = 1 [json_name = "classification"];</code>
   */
  event-logging.v4.event_logging.v4.ClassificationComplexTypeOrBuilder getClassificationOrBuilder();

  /**
   * <pre>
   * A description of the task.
   * </pre>
   *
   * <code>string subject = 2 [json_name = "subject"];</code>
   * @return The subject.
   */
  java.lang.String getSubject();
  /**
   * <pre>
   * A description of the task.
   * </pre>
   *
   * <code>string subject = 2 [json_name = "subject"];</code>
   * @return The bytes for subject.
   */
  com.google.protobuf.ByteString
      getSubjectBytes();

  /**
   * <pre>
   * Names of stakeholders.
   * </pre>
   *
   * <code>string stakeholders = 3 [json_name = "stakeholders"];</code>
   * @return The stakeholders.
   */
  java.lang.String getStakeholders();
  /**
   * <pre>
   * Names of stakeholders.
   * </pre>
   *
   * <code>string stakeholders = 3 [json_name = "stakeholders"];</code>
   * @return The bytes for stakeholders.
   */
  com.google.protobuf.ByteString
      getStakeholdersBytes();

  /**
   * <pre>
   * Business case supporting task.
   * </pre>
   *
   * <code>string justification = 4 [json_name = "justification"];</code>
   * @return The justification.
   */
  java.lang.String getJustification();
  /**
   * <pre>
   * Business case supporting task.
   * </pre>
   *
   * <code>string justification = 4 [json_name = "justification"];</code>
   * @return The bytes for justification.
   */
  com.google.protobuf.ByteString
      getJustificationBytes();

  /**
   * <pre>
   * Expected outcome from task.
   * </pre>
   *
   * <code>string expected_outcome = 5 [json_name = "expectedOutcome"];</code>
   * @return The expectedOutcome.
   */
  java.lang.String getExpectedOutcome();
  /**
   * <pre>
   * Expected outcome from task.
   * </pre>
   *
   * <code>string expected_outcome = 5 [json_name = "expectedOutcome"];</code>
   * @return The bytes for expectedOutcome.
   */
  com.google.protobuf.ByteString
      getExpectedOutcomeBytes();

  /**
   * <pre>
   * The authorisations that were granted to allow this event action to take place.
   * </pre>
   *
   * <code>.event_logging.v4.PurposeComplexType.AuthorisationsType authorisations = 6 [json_name = "authorisations"];</code>
   * @return Whether the authorisations field is set.
   */
  boolean hasAuthorisations();
  /**
   * <pre>
   * The authorisations that were granted to allow this event action to take place.
   * </pre>
   *
   * <code>.event_logging.v4.PurposeComplexType.AuthorisationsType authorisations = 6 [json_name = "authorisations"];</code>
   * @return The authorisations.
   */
  event-logging.v4.event_logging.v4.PurposeComplexType.AuthorisationsType getAuthorisations();
  /**
   * <pre>
   * The authorisations that were granted to allow this event action to take place.
   * </pre>
   *
   * <code>.event_logging.v4.PurposeComplexType.AuthorisationsType authorisations = 6 [json_name = "authorisations"];</code>
   */
  event-logging.v4.event_logging.v4.PurposeComplexType.AuthorisationsTypeOrBuilder getAuthorisationsOrBuilder();

  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 7 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<event-logging.v4.event_logging.v4.DataComplexType> 
      getDataList();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 7 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  event-logging.v4.event_logging.v4.DataComplexType getData(int index);
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 7 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  int getDataCount();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 7 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<? extends event-logging.v4.event_logging.v4.DataComplexTypeOrBuilder> 
      getDataOrBuilderList();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 7 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  event-logging.v4.event_logging.v4.DataComplexTypeOrBuilder getDataOrBuilder(
      int index);
}
