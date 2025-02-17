// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/config/enforcer/config.proto

package org.wso2.choreo.connect.discovery.config.enforcer;

public final class ConfigProto {
  private ConfigProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_wso2_discovery_config_enforcer_Config_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_wso2_discovery_config_enforcer_Config_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n+wso2/discovery/config/enforcer/config." +
      "proto\022\036wso2.discovery.config.enforcer\032)w" +
      "so2/discovery/config/enforcer/cert.proto" +
      "\032,wso2/discovery/config/enforcer/service" +
      ".proto\0322wso2/discovery/config/enforcer/j" +
      "wt_generator.proto\032/wso2/discovery/confi" +
      "g/enforcer/jwt_issuer.proto\032/wso2/discov" +
      "ery/config/enforcer/throttling.proto\032*ws" +
      "o2/discovery/config/enforcer/cache.proto" +
      "\032.wso2/discovery/config/enforcer/analyti" +
      "cs.proto\032-wso2/discovery/config/enforcer" +
      "/security.proto\032/wso2/discovery/config/e" +
      "nforcer/management.proto\0320wso2/discovery" +
      "/config/enforcer/rest_server.proto\"\264\005\n\006C" +
      "onfig\022:\n\010security\030\001 \001(\0132(.wso2.discovery" +
      ".config.enforcer.Security\022;\n\010keystore\030\002 " +
      "\001(\0132).wso2.discovery.config.enforcer.Cer" +
      "tStore\022=\n\ntruststore\030\003 \001(\0132).wso2.discov" +
      "ery.config.enforcer.CertStore\022<\n\013authSer" +
      "vice\030\004 \001(\0132\'.wso2.discovery.config.enfor" +
      "cer.Service\022B\n\014jwtGenerator\030\005 \001(\0132,.wso2" +
      ".discovery.config.enforcer.JWTGenerator\022" +
      ">\n\nthrottling\030\006 \001(\0132*.wso2.discovery.con" +
      "fig.enforcer.Throttling\0224\n\005cache\030\007 \001(\0132%" +
      ".wso2.discovery.config.enforcer.Cache\022<\n" +
      "\tjwtIssuer\030\010 \001(\0132).wso2.discovery.config" +
      ".enforcer.JWTIssuer\022<\n\tanalytics\030\t \001(\0132)" +
      ".wso2.discovery.config.enforcer.Analytic" +
      "s\022>\n\nmanagement\030\n \001(\0132*.wso2.discovery.c" +
      "onfig.enforcer.Management\022>\n\nrestServer\030" +
      "\013 \001(\0132*.wso2.discovery.config.enforcer.R" +
      "estServerB\222\001\n1org.wso2.choreo.connect.di" +
      "scovery.config.enforcerB\013ConfigProtoP\001ZN" +
      "github.com/envoyproxy/go-control-plane/w" +
      "so2/discovery/config/enforcer;enforcerb\006" +
      "proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.wso2.choreo.connect.discovery.config.enforcer.CertStoreProto.getDescriptor(),
          org.wso2.choreo.connect.discovery.config.enforcer.ServiceProto.getDescriptor(),
          org.wso2.choreo.connect.discovery.config.enforcer.JWTGeneratorProto.getDescriptor(),
          org.wso2.choreo.connect.discovery.config.enforcer.JWTIssuerProto.getDescriptor(),
          org.wso2.choreo.connect.discovery.config.enforcer.ThrottlingProto.getDescriptor(),
          org.wso2.choreo.connect.discovery.config.enforcer.CacheProto.getDescriptor(),
          org.wso2.choreo.connect.discovery.config.enforcer.AnalyticsProto.getDescriptor(),
          org.wso2.choreo.connect.discovery.config.enforcer.SecurityProto.getDescriptor(),
          org.wso2.choreo.connect.discovery.config.enforcer.ManagementProto.getDescriptor(),
          org.wso2.choreo.connect.discovery.config.enforcer.RestServerProto.getDescriptor(),
        });
    internal_static_wso2_discovery_config_enforcer_Config_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_wso2_discovery_config_enforcer_Config_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_wso2_discovery_config_enforcer_Config_descriptor,
        new java.lang.String[] { "Security", "Keystore", "Truststore", "AuthService", "JwtGenerator", "Throttling", "Cache", "JwtIssuer", "Analytics", "Management", "RestServer", });
    org.wso2.choreo.connect.discovery.config.enforcer.CertStoreProto.getDescriptor();
    org.wso2.choreo.connect.discovery.config.enforcer.ServiceProto.getDescriptor();
    org.wso2.choreo.connect.discovery.config.enforcer.JWTGeneratorProto.getDescriptor();
    org.wso2.choreo.connect.discovery.config.enforcer.JWTIssuerProto.getDescriptor();
    org.wso2.choreo.connect.discovery.config.enforcer.ThrottlingProto.getDescriptor();
    org.wso2.choreo.connect.discovery.config.enforcer.CacheProto.getDescriptor();
    org.wso2.choreo.connect.discovery.config.enforcer.AnalyticsProto.getDescriptor();
    org.wso2.choreo.connect.discovery.config.enforcer.SecurityProto.getDescriptor();
    org.wso2.choreo.connect.discovery.config.enforcer.ManagementProto.getDescriptor();
    org.wso2.choreo.connect.discovery.config.enforcer.RestServerProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
