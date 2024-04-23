// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/approval_complex_type.proto

// Protobuf Java Version: 4.26.1
package uk.gov.gchq.event_logging.v4;

public interface ApprovalComplexTypeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:event_logging.v4.ApprovalComplexType)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * The action that the approval event is capturing, e.g. Approve, Reject, etc.
   * </pre>
   *
   * <code>.event_logging.v4.ApprovalActionSimpleType action = 1 [json_name = "action", (.buf.validate.field) = { ... }</code>
   * @return The enum numeric value on the wire for action.
   */
  int getActionValue();
  /**
   * <pre>
   * The action that the approval event is capturing, e.g. Approve, Reject, etc.
   * </pre>
   *
   * <code>.event_logging.v4.ApprovalActionSimpleType action = 1 [json_name = "action", (.buf.validate.field) = { ... }</code>
   * @return The action.
   */
  uk.gov.gchq.event_logging.v4.ApprovalActionSimpleType getAction();

  /**
   * <pre>
   * An identifier associated with the approval step/process.
   * </pre>
   *
   * <code>string id = 2 [json_name = "id"];</code>
   * @return The id.
   */
  java.lang.String getId();
  /**
   * <pre>
   * An identifier associated with the approval step/process.
   * </pre>
   *
   * <code>string id = 2 [json_name = "id"];</code>
   * @return The bytes for id.
   */
  com.google.protobuf.ByteString
      getIdBytes();

  /**
   * <pre>
   * The object that this approval step (or request for approval) relates to, e.g. the document being approved.
   * </pre>
   *
   * <code>.event_logging.v4.MultiObjectComplexType subject = 3 [json_name = "subject"];</code>
   * @return Whether the subject field is set.
   */
  boolean hasSubject();
  /**
   * <pre>
   * The object that this approval step (or request for approval) relates to, e.g. the document being approved.
   * </pre>
   *
   * <code>.event_logging.v4.MultiObjectComplexType subject = 3 [json_name = "subject"];</code>
   * @return The subject.
   */
  uk.gov.gchq.event_logging.v4.MultiObjectComplexType getSubject();
  /**
   * <pre>
   * The object that this approval step (or request for approval) relates to, e.g. the document being approved.
   * </pre>
   *
   * <code>.event_logging.v4.MultiObjectComplexType subject = 3 [json_name = "subject"];</code>
   */
  uk.gov.gchq.event_logging.v4.MultiObjectComplexTypeOrBuilder getSubjectOrBuilder();

  /**
   * <pre>
   * The user(s) that requested the approval, e.g. jbloggs requesting approval from a manager for his finance report.
   * </pre>
   *
   * <code>.event_logging.v4.ApprovalComplexType.RequestorsType requestors = 4 [json_name = "requestors"];</code>
   * @return Whether the requestors field is set.
   */
  boolean hasRequestors();
  /**
   * <pre>
   * The user(s) that requested the approval, e.g. jbloggs requesting approval from a manager for his finance report.
   * </pre>
   *
   * <code>.event_logging.v4.ApprovalComplexType.RequestorsType requestors = 4 [json_name = "requestors"];</code>
   * @return The requestors.
   */
  uk.gov.gchq.event_logging.v4.ApprovalComplexType.RequestorsType getRequestors();
  /**
   * <pre>
   * The user(s) that requested the approval, e.g. jbloggs requesting approval from a manager for his finance report.
   * </pre>
   *
   * <code>.event_logging.v4.ApprovalComplexType.RequestorsType requestors = 4 [json_name = "requestors"];</code>
   */
  uk.gov.gchq.event_logging.v4.ApprovalComplexType.RequestorsTypeOrBuilder getRequestorsOrBuilder();

  /**
   * <pre>
   * The user(s) that are providing the approval, e.g. a user requesting approval from manager fsmith for his finance report.
   * </pre>
   *
   * <code>.event_logging.v4.ApprovalComplexType.ApproversType approvers = 5 [json_name = "approvers"];</code>
   * @return Whether the approvers field is set.
   */
  boolean hasApprovers();
  /**
   * <pre>
   * The user(s) that are providing the approval, e.g. a user requesting approval from manager fsmith for his finance report.
   * </pre>
   *
   * <code>.event_logging.v4.ApprovalComplexType.ApproversType approvers = 5 [json_name = "approvers"];</code>
   * @return The approvers.
   */
  uk.gov.gchq.event_logging.v4.ApprovalComplexType.ApproversType getApprovers();
  /**
   * <pre>
   * The user(s) that are providing the approval, e.g. a user requesting approval from manager fsmith for his finance report.
   * </pre>
   *
   * <code>.event_logging.v4.ApprovalComplexType.ApproversType approvers = 5 [json_name = "approvers"];</code>
   */
  uk.gov.gchq.event_logging.v4.ApprovalComplexType.ApproversTypeOrBuilder getApproversOrBuilder();

  /**
   * <pre>
   * The reason for the approval/rejection/request, e.g. why the approval was rejected.
   * </pre>
   *
   * <code>string reason = 6 [json_name = "reason"];</code>
   * @return The reason.
   */
  java.lang.String getReason();
  /**
   * <pre>
   * The reason for the approval/rejection/request, e.g. why the approval was rejected.
   * </pre>
   *
   * <code>string reason = 6 [json_name = "reason"];</code>
   * @return The bytes for reason.
   */
  com.google.protobuf.ByteString
      getReasonBytes();

  /**
   * <pre>
   * Used to determine if the action was successful. If omitted it is assumed that the event was successful and was permitted.
   * </pre>
   *
   * <code>.event_logging.v4.OutcomeComplexType outcome = 7 [json_name = "outcome"];</code>
   * @return Whether the outcome field is set.
   */
  boolean hasOutcome();
  /**
   * <pre>
   * Used to determine if the action was successful. If omitted it is assumed that the event was successful and was permitted.
   * </pre>
   *
   * <code>.event_logging.v4.OutcomeComplexType outcome = 7 [json_name = "outcome"];</code>
   * @return The outcome.
   */
  uk.gov.gchq.event_logging.v4.OutcomeComplexType getOutcome();
  /**
   * <pre>
   * Used to determine if the action was successful. If omitted it is assumed that the event was successful and was permitted.
   * </pre>
   *
   * <code>.event_logging.v4.OutcomeComplexType outcome = 7 [json_name = "outcome"];</code>
   */
  uk.gov.gchq.event_logging.v4.OutcomeComplexTypeOrBuilder getOutcomeOrBuilder();

  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 8 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<uk.gov.gchq.event_logging.v4.DataComplexType> 
      getDataList();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 8 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  uk.gov.gchq.event_logging.v4.DataComplexType getData(int index);
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 8 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  int getDataCount();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 8 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<? extends uk.gov.gchq.event_logging.v4.DataComplexTypeOrBuilder> 
      getDataOrBuilderList();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 8 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  uk.gov.gchq.event_logging.v4.DataComplexTypeOrBuilder getDataOrBuilder(
      int index);
}
