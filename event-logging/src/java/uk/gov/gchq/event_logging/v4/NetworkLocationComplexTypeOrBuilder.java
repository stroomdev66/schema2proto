// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: event_logging/v4/network_location_complex_type.proto

// Protobuf Java Version: 4.26.1
package uk.gov.gchq.event_logging.v4;

public interface NetworkLocationComplexTypeOrBuilder extends
    // @@protoc_insertion_point(interface_extends:event_logging.v4.NetworkLocationComplexType)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * A device at the source or destination involved in the network activity.
   * </pre>
   *
   * <code>.event_logging.v4.DeviceComplexType device = 1 [json_name = "device"];</code>
   * @return Whether the device field is set.
   */
  boolean hasDevice();
  /**
   * <pre>
   * A device at the source or destination involved in the network activity.
   * </pre>
   *
   * <code>.event_logging.v4.DeviceComplexType device = 1 [json_name = "device"];</code>
   * @return The device.
   */
  uk.gov.gchq.event_logging.v4.DeviceComplexType getDevice();
  /**
   * <pre>
   * A device at the source or destination involved in the network activity.
   * </pre>
   *
   * <code>.event_logging.v4.DeviceComplexType device = 1 [json_name = "device"];</code>
   */
  uk.gov.gchq.event_logging.v4.DeviceComplexTypeOrBuilder getDeviceOrBuilder();

  /**
   * <pre>
   * The application being used by the source or destination.
   * </pre>
   *
   * <code>string application = 2 [json_name = "application"];</code>
   * @return The application.
   */
  java.lang.String getApplication();
  /**
   * <pre>
   * The application being used by the source or destination.
   * </pre>
   *
   * <code>string application = 2 [json_name = "application"];</code>
   * @return The bytes for application.
   */
  com.google.protobuf.ByteString
      getApplicationBytes();

  /**
   * <pre>
   * The transport protocol being used by the source or destination.
   * </pre>
   *
   * <code>.event_logging.v4.NetworkProtocolSimpleType transport_protocol = 3 [json_name = "transportProtocol", (.buf.validate.field) = { ... }</code>
   * @return The enum numeric value on the wire for transportProtocol.
   */
  int getTransportProtocolValue();
  /**
   * <pre>
   * The transport protocol being used by the source or destination.
   * </pre>
   *
   * <code>.event_logging.v4.NetworkProtocolSimpleType transport_protocol = 3 [json_name = "transportProtocol", (.buf.validate.field) = { ... }</code>
   * @return The transportProtocol.
   */
  uk.gov.gchq.event_logging.v4.NetworkProtocolSimpleType getTransportProtocol();

  /**
   * <pre>
   * The Internet Control Message Protocol type number. See https://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml
   * </pre>
   *
   * <code>int32 icmp_type = 4 [json_name = "icmpType"];</code>
   * @return The icmpType.
   */
  int getIcmpType();

  /**
   * <pre>
   * The HTTP method, e.g. GET, POST, DELETE, PUT etc
   * </pre>
   *
   * <code>string http_method = 5 [json_name = "httpMethod"];</code>
   * @return The httpMethod.
   */
  java.lang.String getHttpMethod();
  /**
   * <pre>
   * The HTTP method, e.g. GET, POST, DELETE, PUT etc
   * </pre>
   *
   * <code>string http_method = 5 [json_name = "httpMethod"];</code>
   * @return The bytes for httpMethod.
   */
  com.google.protobuf.ByteString
      getHttpMethodBytes();

  /**
   * <pre>
   * The application protocol being used by the source or destination.
   * </pre>
   *
   * <code>string application_protocol = 6 [json_name = "applicationProtocol"];</code>
   * @return The applicationProtocol.
   */
  java.lang.String getApplicationProtocol();
  /**
   * <pre>
   * The application protocol being used by the source or destination.
   * </pre>
   *
   * <code>string application_protocol = 6 [json_name = "applicationProtocol"];</code>
   * @return The bytes for applicationProtocol.
   */
  com.google.protobuf.ByteString
      getApplicationProtocolBytes();

  /**
   * <pre>
   * The port being used by the source or destination.
   * </pre>
   *
   * <code>int32 port = 7 [json_name = "port"];</code>
   * @return The port.
   */
  int getPort();

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
