// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/alert_complex_type.proto
// Protobuf Java Version: 4.26.1

package uk.gov.gchq.event_logging.v4;

public final class AlertComplexTypeProto {
  private AlertComplexTypeProto() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      AlertComplexTypeProto.class.getName());
  }
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_event_logging_v4_AlertComplexType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_event_logging_v4_AlertComplexType_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_event_logging_v4_AlertComplexType_IDSType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_event_logging_v4_AlertComplexType_IDSType_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_event_logging_v4_AlertComplexType_ChangeType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_event_logging_v4_AlertComplexType_ChangeType_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n)event_logging/v4/alert_complex_type.pr" +
      "oto\022\020event_logging.v4\0321event_logging/v4/" +
      "alert_priority_simple_type.proto\0321event_" +
      "logging/v4/alert_severity_simple_type.pr" +
      "oto\032-event_logging/v4/alert_type_simple_" +
      "type.proto\0327event_logging/v4/anti_malwar" +
      "e_threat_complex_type.proto\0320event_loggi" +
      "ng/v4/change_action_simple_type.proto\032(e" +
      "vent_logging/v4/data_complex_type.proto\032" +
      "0event_logging/v4/multi_object_complex_t" +
      "ype.proto\032+event_logging/v4/network_comp" +
      "lex_type.proto\0324event_logging/v4/network" +
      "_location_complex_type.proto\032\033buf/valida" +
      "te/validate.proto\"\356\010\n\020AlertComplexType\022A" +
      "\n\004type\030\001 \001(\0162%.event_logging.v4.AlertTyp" +
      "eSimpleTypeB\006\272H\003\310\001\001R\004type\022O\n\010severity\030\002 " +
      "\001(\0162).event_logging.v4.AlertSeveritySimp" +
      "leTypeB\010\272H\005\202\001\002\020\001R\010severity\022O\n\010priority\030\003" +
      " \001(\0162).event_logging.v4.AlertPrioritySim" +
      "pleTypeB\010\272H\005\202\001\002\020\001R\010priority\022\030\n\007subject\030\004" +
      " \001(\tR\007subject\022 \n\013description\030\005 \001(\tR\013desc" +
      "ription\022>\n\005i_d_s\030\006 \001(\0132*.event_logging.v" +
      "4.AlertComplexType.IDSTypeR\003iDS\022H\n\007malwa" +
      "re\030\007 \001(\0132..event_logging.v4.AntiMalwareT" +
      "hreatComplexTypeR\007malware\022>\n\007network\030\010 \001" +
      "(\0132$.event_logging.v4.NetworkComplexType" +
      "R\007network\022E\n\006change\030\t \001(\0132-.event_loggin" +
      "g.v4.AlertComplexType.ChangeTypeR\006change" +
      "\022?\n\004data\030\n \003(\0132!.event_logging.v4.DataCo" +
      "mplexTypeB\010\272H\005\222\001\002\010\000R\004data\032\367\001\n\007IDSType\022\022\n" +
      "\004rule\030\001 \001(\tR\004rule\022D\n\006source\030\002 \001(\0132,.even" +
      "t_logging.v4.NetworkLocationComplexTypeR" +
      "\006source\022N\n\013destination\030\003 \001(\0132,.event_log" +
      "ging.v4.NetworkLocationComplexTypeR\013dest" +
      "ination\022B\n\007payload\030\004 \001(\0132(.event_logging" +
      ".v4.MultiObjectComplexTypeR\007payload\032\354\001\n\n" +
      "ChangeType\022H\n\006action\030\001 \001(\0162(.event_loggi" +
      "ng.v4.ChangeActionSimpleTypeB\006\272H\003\310\001\001R\006ac" +
      "tion\022@\n\006before\030\002 \001(\0132(.event_logging.v4." +
      "MultiObjectComplexTypeR\006before\022>\n\005after\030" +
      "\003 \001(\0132(.event_logging.v4.MultiObjectComp" +
      "lexTypeR\005after\022\022\n\004rule\030\004 \001(\tR\004ruleB\331\001\n\034u" +
      "k.gov.gchq.event_logging.v4B\025AlertComple" +
      "xTypeProtoP\001ZEgithub.com/gchq/event-logg" +
      "ing-schema/event_logging/v4;event_loggin" +
      "gv4\242\002\003EXX\252\002\017EventLogging.V4\312\002\017EventLoggi" +
      "ng\\V4\342\002\033EventLogging\\V4\\GPBMetadata\352\002\020Ev" +
      "entLogging::V4b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          uk.gov.gchq.event_logging.v4.AlertPrioritySimpleTypeProto.getDescriptor(),
          uk.gov.gchq.event_logging.v4.AlertSeveritySimpleTypeProto.getDescriptor(),
          uk.gov.gchq.event_logging.v4.AlertTypeSimpleTypeProto.getDescriptor(),
          uk.gov.gchq.event_logging.v4.AntiMalwareThreatComplexTypeProto.getDescriptor(),
          uk.gov.gchq.event_logging.v4.ChangeActionSimpleTypeProto.getDescriptor(),
          uk.gov.gchq.event_logging.v4.DataComplexTypeProto.getDescriptor(),
          uk.gov.gchq.event_logging.v4.MultiObjectComplexTypeProto.getDescriptor(),
          uk.gov.gchq.event_logging.v4.NetworkComplexTypeProto.getDescriptor(),
          uk.gov.gchq.event_logging.v4.NetworkLocationComplexTypeProto.getDescriptor(),
          build.buf.validate.ValidateProto.getDescriptor(),
        });
    internal_static_event_logging_v4_AlertComplexType_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_event_logging_v4_AlertComplexType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_event_logging_v4_AlertComplexType_descriptor,
        new java.lang.String[] { "Type", "Severity", "Priority", "Subject", "Description", "IDS", "Malware", "Network", "Change", "Data", });
    internal_static_event_logging_v4_AlertComplexType_IDSType_descriptor =
      internal_static_event_logging_v4_AlertComplexType_descriptor.getNestedTypes().get(0);
    internal_static_event_logging_v4_AlertComplexType_IDSType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_event_logging_v4_AlertComplexType_IDSType_descriptor,
        new java.lang.String[] { "Rule", "Source", "Destination", "Payload", });
    internal_static_event_logging_v4_AlertComplexType_ChangeType_descriptor =
      internal_static_event_logging_v4_AlertComplexType_descriptor.getNestedTypes().get(1);
    internal_static_event_logging_v4_AlertComplexType_ChangeType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_event_logging_v4_AlertComplexType_ChangeType_descriptor,
        new java.lang.String[] { "Action", "Before", "After", "Rule", });
    descriptor.resolveAllFeaturesImmutable();
    uk.gov.gchq.event_logging.v4.AlertPrioritySimpleTypeProto.getDescriptor();
    uk.gov.gchq.event_logging.v4.AlertSeveritySimpleTypeProto.getDescriptor();
    uk.gov.gchq.event_logging.v4.AlertTypeSimpleTypeProto.getDescriptor();
    uk.gov.gchq.event_logging.v4.AntiMalwareThreatComplexTypeProto.getDescriptor();
    uk.gov.gchq.event_logging.v4.ChangeActionSimpleTypeProto.getDescriptor();
    uk.gov.gchq.event_logging.v4.DataComplexTypeProto.getDescriptor();
    uk.gov.gchq.event_logging.v4.MultiObjectComplexTypeProto.getDescriptor();
    uk.gov.gchq.event_logging.v4.NetworkComplexTypeProto.getDescriptor();
    uk.gov.gchq.event_logging.v4.NetworkLocationComplexTypeProto.getDescriptor();
    build.buf.validate.ValidateProto.getDescriptor();
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(build.buf.validate.ValidateProto.field);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
  }

  // @@protoc_insertion_point(outer_class_scope)
}
