// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/authorise_complex_type.proto

// Protobuf Java Version: 4.26.1
package event-logging.v4.event_logging.v4;

public interface AuthoriseComplexTypeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:event_logging.v4.AuthoriseComplexType)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>repeated .event_logging.v4.AuthoriseComplexType.ChoiceWrapperAuthoriseComplexType choice_wrapper = 1 [json_name = "choiceWrapper", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<event-logging.v4.event_logging.v4.AuthoriseComplexType.ChoiceWrapperAuthoriseComplexType> 
      getChoiceWrapperList();
  /**
   * <code>repeated .event_logging.v4.AuthoriseComplexType.ChoiceWrapperAuthoriseComplexType choice_wrapper = 1 [json_name = "choiceWrapper", (.buf.validate.field) = { ... }</code>
   */
  event-logging.v4.event_logging.v4.AuthoriseComplexType.ChoiceWrapperAuthoriseComplexType getChoiceWrapper(int index);
  /**
   * <code>repeated .event_logging.v4.AuthoriseComplexType.ChoiceWrapperAuthoriseComplexType choice_wrapper = 1 [json_name = "choiceWrapper", (.buf.validate.field) = { ... }</code>
   */
  int getChoiceWrapperCount();
  /**
   * <code>repeated .event_logging.v4.AuthoriseComplexType.ChoiceWrapperAuthoriseComplexType choice_wrapper = 1 [json_name = "choiceWrapper", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<? extends event-logging.v4.event_logging.v4.AuthoriseComplexType.ChoiceWrapperAuthoriseComplexTypeOrBuilder> 
      getChoiceWrapperOrBuilderList();
  /**
   * <code>repeated .event_logging.v4.AuthoriseComplexType.ChoiceWrapperAuthoriseComplexType choice_wrapper = 1 [json_name = "choiceWrapper", (.buf.validate.field) = { ... }</code>
   */
  event-logging.v4.event_logging.v4.AuthoriseComplexType.ChoiceWrapperAuthoriseComplexTypeOrBuilder getChoiceWrapperOrBuilder(
      int index);

  /**
   * <pre>
   * The type of the event action, e.g. a modification to authorisation rules/groups or the request to be authorised.
   * </pre>
   *
   * <code>.event_logging.v4.AuthorisationSimpleType action = 11 [json_name = "action", (.buf.validate.field) = { ... }</code>
   * @return The enum numeric value on the wire for action.
   */
  int getActionValue();
  /**
   * <pre>
   * The type of the event action, e.g. a modification to authorisation rules/groups or the request to be authorised.
   * </pre>
   *
   * <code>.event_logging.v4.AuthorisationSimpleType action = 11 [json_name = "action", (.buf.validate.field) = { ... }</code>
   * @return The action.
   */
  event-logging.v4.event_logging.v4.AuthorisationSimpleType getAction();

  /**
   * <pre>
   * A list of roles or permissions that have been added to an object.
   * </pre>
   *
   * <code>.event_logging.v4.AuthoriseComplexType.AddGroupsType add_groups = 12 [json_name = "addGroups"];</code>
   * @return Whether the addGroups field is set.
   */
  boolean hasAddGroups();
  /**
   * <pre>
   * A list of roles or permissions that have been added to an object.
   * </pre>
   *
   * <code>.event_logging.v4.AuthoriseComplexType.AddGroupsType add_groups = 12 [json_name = "addGroups"];</code>
   * @return The addGroups.
   */
  event-logging.v4.event_logging.v4.AuthoriseComplexType.AddGroupsType getAddGroups();
  /**
   * <pre>
   * A list of roles or permissions that have been added to an object.
   * </pre>
   *
   * <code>.event_logging.v4.AuthoriseComplexType.AddGroupsType add_groups = 12 [json_name = "addGroups"];</code>
   */
  event-logging.v4.event_logging.v4.AuthoriseComplexType.AddGroupsTypeOrBuilder getAddGroupsOrBuilder();

  /**
   * <pre>
   * A list of roles or permissions that have been removed from an object.
   * </pre>
   *
   * <code>.event_logging.v4.AuthoriseComplexType.RemoveGroupsType remove_groups = 13 [json_name = "removeGroups"];</code>
   * @return Whether the removeGroups field is set.
   */
  boolean hasRemoveGroups();
  /**
   * <pre>
   * A list of roles or permissions that have been removed from an object.
   * </pre>
   *
   * <code>.event_logging.v4.AuthoriseComplexType.RemoveGroupsType remove_groups = 13 [json_name = "removeGroups"];</code>
   * @return The removeGroups.
   */
  event-logging.v4.event_logging.v4.AuthoriseComplexType.RemoveGroupsType getRemoveGroups();
  /**
   * <pre>
   * A list of roles or permissions that have been removed from an object.
   * </pre>
   *
   * <code>.event_logging.v4.AuthoriseComplexType.RemoveGroupsType remove_groups = 13 [json_name = "removeGroups"];</code>
   */
  event-logging.v4.event_logging.v4.AuthoriseComplexType.RemoveGroupsTypeOrBuilder getRemoveGroupsOrBuilder();

  /**
   * <pre>
   * Used to determine if the action was successful. If omitted it is assumed that the event was successful and was permitted.
   * </pre>
   *
   * <code>.event_logging.v4.OutcomeComplexType outcome = 14 [json_name = "outcome"];</code>
   * @return Whether the outcome field is set.
   */
  boolean hasOutcome();
  /**
   * <pre>
   * Used to determine if the action was successful. If omitted it is assumed that the event was successful and was permitted.
   * </pre>
   *
   * <code>.event_logging.v4.OutcomeComplexType outcome = 14 [json_name = "outcome"];</code>
   * @return The outcome.
   */
  event-logging.v4.event_logging.v4.OutcomeComplexType getOutcome();
  /**
   * <pre>
   * Used to determine if the action was successful. If omitted it is assumed that the event was successful and was permitted.
   * </pre>
   *
   * <code>.event_logging.v4.OutcomeComplexType outcome = 14 [json_name = "outcome"];</code>
   */
  event-logging.v4.event_logging.v4.OutcomeComplexTypeOrBuilder getOutcomeOrBuilder();

  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 15 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<event-logging.v4.event_logging.v4.DataComplexType> 
      getDataList();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 15 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  event-logging.v4.event_logging.v4.DataComplexType getData(int index);
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 15 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  int getDataCount();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 15 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  java.util.List<? extends event-logging.v4.event_logging.v4.DataComplexTypeOrBuilder> 
      getDataOrBuilderList();
  /**
   * <pre>
   * Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
   * </pre>
   *
   * <code>repeated .event_logging.v4.DataComplexType data = 15 [json_name = "data", (.buf.validate.field) = { ... }</code>
   */
  event-logging.v4.event_logging.v4.DataComplexTypeOrBuilder getDataOrBuilder(
      int index);
}
