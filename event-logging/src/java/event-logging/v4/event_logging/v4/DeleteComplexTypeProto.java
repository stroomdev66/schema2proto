// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/delete_complex_type.proto
// Protobuf Java Version: 4.26.1

package event-logging.v4.event_logging.v4;

public final class DeleteComplexTypeProto {
  private DeleteComplexTypeProto() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 1,
      /* suffix= */ "",
      DeleteComplexTypeProto.class.getName());
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
    internal_static_event_logging_v4_DeleteComplexType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_event_logging_v4_DeleteComplexType_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_event_logging_v4_DeleteComplexType_ChoiceWrapperDeleteComplexType_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_event_logging_v4_DeleteComplexType_ChoiceWrapperDeleteComplexType_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n*event_logging/v4/delete_complex_type.p" +
      "roto\022\020event_logging.v4\032*event_logging/v4" +
      "/banner_complex_type.proto\032(event_loggin" +
      "g/v4/chat_complex_type.proto\0321event_logg" +
      "ing/v4/configuration_complex_type.proto\032" +
      "(event_logging/v4/data_complex_type.prot" +
      "o\032,event_logging/v4/document_complex_typ" +
      "e.proto\032)event_logging/v4/email_complex_" +
      "type.proto\032(event_logging/v4/file_comple" +
      "x_type.proto\032*event_logging/v4/folder_co" +
      "mplex_type.proto\032.event_logging/v4/group" +
      "_chat_complex_type.proto\0320event_logging/" +
      "v4/multi_object_complex_type.proto\032*even" +
      "t_logging/v4/object_complex_type.proto\032+" +
      "event_logging/v4/outcome_complex_type.pr" +
      "oto\032,event_logging/v4/resource_complex_t" +
      "ype.proto\032,event_logging/v4/shortcut_com" +
      "plex_type.proto\032(event_logging/v4/user_c" +
      "omplex_type.proto\0323event_logging/v4/virt" +
      "ual_session_complex_type.proto\032(event_lo" +
      "gging/v4/voip_complex_type.proto\032\033buf/va" +
      "lidate/validate.proto\"\335\014\n\021DeleteComplexT" +
      "ype\022s\n\016choice_wrapper\030\001 \003(\0132B.event_logg" +
      "ing.v4.DeleteComplexType.ChoiceWrapperDe" +
      "leteComplexTypeB\010\272H\005\222\001\002\010\001R\rchoiceWrapper" +
      "\022>\n\007outcome\030\013 \001(\0132$.event_logging.v4.Out" +
      "comeComplexTypeR\007outcome\022?\n\004data\030\014 \003(\0132!" +
      ".event_logging.v4.DataComplexTypeB\010\272H\005\222\001" +
      "\002\010\000R\004data\032\321\n\n\036ChoiceWrapperDeleteComplex" +
      "Type\022R\n\013association\030\001 \001(\0132(.event_loggin" +
      "g.v4.AssociationComplexTypeB\006\272H\003\310\001\001R\013ass" +
      "ociation\022C\n\006banner\030\002 \001(\0132#.event_logging" +
      ".v4.BannerComplexTypeB\006\272H\003\310\001\001R\006banner\022=\n" +
      "\004chat\030\003 \001(\0132!.event_logging.v4.ChatCompl" +
      "exTypeB\006\272H\003\310\001\001R\004chat\022X\n\rconfiguration\030\004 " +
      "\001(\0132*.event_logging.v4.ConfigurationComp" +
      "lexTypeB\006\272H\003\310\001\001R\rconfiguration\022I\n\010criter" +
      "ia\030\005 \001(\0132%.event_logging.v4.CriteriaComp" +
      "lexTypeB\006\272H\003\310\001\001R\010criteria\022I\n\010document\030\006 " +
      "\001(\0132%.event_logging.v4.DocumentComplexTy" +
      "peB\006\272H\003\310\001\001R\010document\022@\n\005email\030\007 \001(\0132\".ev" +
      "ent_logging.v4.EmailComplexTypeB\006\272H\003\310\001\001R" +
      "\005email\022=\n\004file\030\010 \001(\0132!.event_logging.v4." +
      "FileComplexTypeB\006\272H\003\310\001\001R\004file\022C\n\006folder\030" +
      "\t \001(\0132#.event_logging.v4.FolderComplexTy" +
      "peB\006\272H\003\310\001\001R\006folder\022@\n\005group\030\n \001(\0132\".even" +
      "t_logging.v4.GroupComplexTypeB\006\272H\003\310\001\001R\005g" +
      "roup\022M\n\ngroup_chat\030\013 \001(\0132&.event_logging" +
      ".v4.GroupChatComplexTypeB\006\272H\003\310\001\001R\tgroupC" +
      "hat\022C\n\006object\030\014 \001(\0132#.event_logging.v4.O" +
      "bjectComplexTypeB\006\272H\003\310\001\001R\006object\022Y\n\016sear" +
      "ch_results\030\r \001(\0132*.event_logging.v4.Sear" +
      "chResultsComplexTypeB\006\272H\003\310\001\001R\rsearchResu" +
      "lts\022I\n\010shortcut\030\016 \001(\0132%.event_logging.v4" +
      ".ShortcutComplexTypeB\006\272H\003\310\001\001R\010shortcut\022=" +
      "\n\004user\030\017 \001(\0132!.event_logging.v4.UserComp" +
      "lexTypeB\006\272H\003\310\001\001R\004user\022\\\n\017virtual_session" +
      "\030\020 \001(\0132+.event_logging.v4.VirtualSession" +
      "ComplexTypeB\006\272H\003\310\001\001R\016virtualSession\022=\n\004v" +
      "oip\030\021 \001(\0132!.event_logging.v4.VoipComplex" +
      "TypeB\006\272H\003\310\001\001R\004voip\022I\n\010resource\030\022 \001(\0132%.e" +
      "vent_logging.v4.ResourceComplexTypeB\006\272H\003" +
      "\310\001\001R\010resourceB\337\001\n!event-logging.v4.event" +
      "_logging.v4B\026DeleteComplexTypeProtoP\001ZEg" +
      "ithub.com/gchq/event-logging-schema/even" +
      "t_logging/v4;event_loggingv4\242\002\003EXX\252\002\017Eve" +
      "ntLogging.V4\312\002\017EventLogging\\V4\342\002\033EventLo" +
      "gging\\V4\\GPBMetadata\352\002\020EventLogging::V4b" +
      "\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          event-logging.v4.event_logging.v4.BannerComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.ChatComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.ConfigurationComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.DataComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.DocumentComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.EmailComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.FileComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.FolderComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.GroupChatComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.MultiObjectComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.ObjectComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.OutcomeComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.ResourceComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.ShortcutComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.UserComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.VirtualSessionComplexTypeProto.getDescriptor(),
          event-logging.v4.event_logging.v4.VoipComplexTypeProto.getDescriptor(),
          build.buf.validate.ValidateProto.getDescriptor(),
        });
    internal_static_event_logging_v4_DeleteComplexType_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_event_logging_v4_DeleteComplexType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_event_logging_v4_DeleteComplexType_descriptor,
        new java.lang.String[] { "ChoiceWrapper", "Outcome", "Data", });
    internal_static_event_logging_v4_DeleteComplexType_ChoiceWrapperDeleteComplexType_descriptor =
      internal_static_event_logging_v4_DeleteComplexType_descriptor.getNestedTypes().get(0);
    internal_static_event_logging_v4_DeleteComplexType_ChoiceWrapperDeleteComplexType_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_event_logging_v4_DeleteComplexType_ChoiceWrapperDeleteComplexType_descriptor,
        new java.lang.String[] { "Association", "Banner", "Chat", "Configuration", "Criteria", "Document", "Email", "File", "Folder", "Group", "GroupChat", "Object", "SearchResults", "Shortcut", "User", "VirtualSession", "Voip", "Resource", });
    descriptor.resolveAllFeaturesImmutable();
    event-logging.v4.event_logging.v4.BannerComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.ChatComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.ConfigurationComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.DataComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.DocumentComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.EmailComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.FileComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.FolderComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.GroupChatComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.MultiObjectComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.ObjectComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.OutcomeComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.ResourceComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.ShortcutComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.UserComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.VirtualSessionComplexTypeProto.getDescriptor();
    event-logging.v4.event_logging.v4.VoipComplexTypeProto.getDescriptor();
    build.buf.validate.ValidateProto.getDescriptor();
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(build.buf.validate.ValidateProto.field);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
  }

  // @@protoc_insertion_point(outer_class_scope)
}
