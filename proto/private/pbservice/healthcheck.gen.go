// Code generated by mog. DO NOT EDIT.

package pbservice

import "github.com/hashicorp/consul/agent/structs"

func CheckTypeToStructs(s *CheckType, t *structs.CheckType) {
	if s == nil {
		return
	}
	t.CheckID = CheckIDType(s.CheckID)
	t.Name = s.Name
	t.Status = s.Status
	t.Notes = s.Notes
	t.ScriptArgs = s.ScriptArgs
	t.HTTP = s.HTTP
	t.H2PING = s.H2PING
	t.H2PingUseTLS = s.H2PingUseTLS
	t.Header = MapHeadersToStructs(s.Header)
	t.Method = s.Method
	t.Body = s.Body
	t.DisableRedirects = s.DisableRedirects
	t.TCP = s.TCP
	t.TCPUseTLS = s.TCPUseTLS
	t.UDP = s.UDP
	t.Interval = structs.DurationFromProto(s.Interval)
	t.AliasNode = s.AliasNode
	t.AliasService = s.AliasService
	t.DockerContainerID = s.DockerContainerID
	t.Shell = s.Shell
	t.GRPC = s.GRPC
	t.GRPCUseTLS = s.GRPCUseTLS
	t.OSService = s.OSService
	t.TLSServerName = s.TLSServerName
	t.TLSSkipVerify = s.TLSSkipVerify
	t.Timeout = structs.DurationFromProto(s.Timeout)
	t.TTL = structs.DurationFromProto(s.TTL)
	t.SuccessBeforePassing = int(s.SuccessBeforePassing)
	t.FailuresBeforeWarning = int(s.FailuresBeforeWarning)
	t.FailuresBeforeCritical = int(s.FailuresBeforeCritical)
	t.ProxyHTTP = s.ProxyHTTP
	t.ProxyGRPC = s.ProxyGRPC
	t.DeregisterCriticalServiceAfter = structs.DurationFromProto(s.DeregisterCriticalServiceAfter)
	t.OutputMaxSize = int(s.OutputMaxSize)
}
func CheckTypeFromStructs(t *structs.CheckType, s *CheckType) {
	if s == nil {
		return
	}
	s.CheckID = string(t.CheckID)
	s.Name = t.Name
	s.Status = t.Status
	s.Notes = t.Notes
	s.ScriptArgs = t.ScriptArgs
	s.HTTP = t.HTTP
	s.H2PING = t.H2PING
	s.H2PingUseTLS = t.H2PingUseTLS
	s.Header = NewMapHeadersFromStructs(t.Header)
	s.Method = t.Method
	s.Body = t.Body
	s.DisableRedirects = t.DisableRedirects
	s.TCP = t.TCP
	s.TCPUseTLS = t.TCPUseTLS
	s.UDP = t.UDP
	s.Interval = structs.DurationToProto(t.Interval)
	s.AliasNode = t.AliasNode
	s.AliasService = t.AliasService
	s.DockerContainerID = t.DockerContainerID
	s.Shell = t.Shell
	s.GRPC = t.GRPC
	s.GRPCUseTLS = t.GRPCUseTLS
	s.OSService = t.OSService
	s.TLSServerName = t.TLSServerName
	s.TLSSkipVerify = t.TLSSkipVerify
	s.Timeout = structs.DurationToProto(t.Timeout)
	s.TTL = structs.DurationToProto(t.TTL)
	s.SuccessBeforePassing = int32(t.SuccessBeforePassing)
	s.FailuresBeforeWarning = int32(t.FailuresBeforeWarning)
	s.FailuresBeforeCritical = int32(t.FailuresBeforeCritical)
	s.ProxyHTTP = t.ProxyHTTP
	s.ProxyGRPC = t.ProxyGRPC
	s.DeregisterCriticalServiceAfter = structs.DurationToProto(t.DeregisterCriticalServiceAfter)
	s.OutputMaxSize = int32(t.OutputMaxSize)
}
func HealthCheckToStructs(s *HealthCheck, t *structs.HealthCheck) {
	if s == nil {
		return
	}
	t.Node = s.Node
	t.CheckID = CheckIDType(s.CheckID)
	t.Name = s.Name
	t.Status = s.Status
	t.Notes = s.Notes
	t.Output = s.Output
	t.ServiceID = s.ServiceID
	t.ServiceName = s.ServiceName
	t.ServiceTags = s.ServiceTags
	t.Type = s.Type
	t.Interval = s.Interval
	t.Timeout = s.Timeout
	t.ExposedPort = int(s.ExposedPort)
	t.PeerName = s.PeerName
	if s.Definition != nil {
		HealthCheckDefinitionToStructs(s.Definition, &t.Definition)
	}
	t.EnterpriseMeta = EnterpriseMetaToStructs(s.EnterpriseMeta)
	t.RaftIndex = RaftIndexToStructs(s.RaftIndex)
}
func HealthCheckFromStructs(t *structs.HealthCheck, s *HealthCheck) {
	if s == nil {
		return
	}
	s.Node = t.Node
	s.CheckID = string(t.CheckID)
	s.Name = t.Name
	s.Status = t.Status
	s.Notes = t.Notes
	s.Output = t.Output
	s.ServiceID = t.ServiceID
	s.ServiceName = t.ServiceName
	s.ServiceTags = t.ServiceTags
	s.Type = t.Type
	s.Interval = t.Interval
	s.Timeout = t.Timeout
	s.ExposedPort = int32(t.ExposedPort)
	s.PeerName = t.PeerName
	{
		var x HealthCheckDefinition
		HealthCheckDefinitionFromStructs(&t.Definition, &x)
		s.Definition = &x
	}
	s.EnterpriseMeta = NewEnterpriseMetaFromStructs(t.EnterpriseMeta)
	s.RaftIndex = NewRaftIndexFromStructs(t.RaftIndex)
}
func HealthCheckDefinitionToStructs(s *HealthCheckDefinition, t *structs.HealthCheckDefinition) {
	if s == nil {
		return
	}
	t.HTTP = s.HTTP
	t.TLSServerName = s.TLSServerName
	t.TLSSkipVerify = s.TLSSkipVerify
	t.Header = MapHeadersToStructs(s.Header)
	t.Method = s.Method
	t.Body = s.Body
	t.DisableRedirects = s.DisableRedirects
	t.TCP = s.TCP
	t.TCPUseTLS = s.TCPUseTLS
	t.UDP = s.UDP
	t.H2PING = s.H2PING
	t.OSService = s.OSService
	t.H2PingUseTLS = s.H2PingUseTLS
	t.Interval = structs.DurationFromProto(s.Interval)
	t.OutputMaxSize = uint(s.OutputMaxSize)
	t.Timeout = structs.DurationFromProto(s.Timeout)
	t.DeregisterCriticalServiceAfter = structs.DurationFromProto(s.DeregisterCriticalServiceAfter)
	t.ScriptArgs = s.ScriptArgs
	t.DockerContainerID = s.DockerContainerID
	t.Shell = s.Shell
	t.GRPC = s.GRPC
	t.GRPCUseTLS = s.GRPCUseTLS
	t.AliasNode = s.AliasNode
	t.AliasService = s.AliasService
	t.TTL = structs.DurationFromProto(s.TTL)
}
func HealthCheckDefinitionFromStructs(t *structs.HealthCheckDefinition, s *HealthCheckDefinition) {
	if s == nil {
		return
	}
	s.HTTP = t.HTTP
	s.TLSServerName = t.TLSServerName
	s.TLSSkipVerify = t.TLSSkipVerify
	s.Header = NewMapHeadersFromStructs(t.Header)
	s.Method = t.Method
	s.Body = t.Body
	s.DisableRedirects = t.DisableRedirects
	s.TCP = t.TCP
	s.TCPUseTLS = t.TCPUseTLS
	s.UDP = t.UDP
	s.H2PING = t.H2PING
	s.OSService = t.OSService
	s.H2PingUseTLS = t.H2PingUseTLS
	s.Interval = structs.DurationToProto(t.Interval)
	s.OutputMaxSize = uint32(t.OutputMaxSize)
	s.Timeout = structs.DurationToProto(t.Timeout)
	s.DeregisterCriticalServiceAfter = structs.DurationToProto(t.DeregisterCriticalServiceAfter)
	s.ScriptArgs = t.ScriptArgs
	s.DockerContainerID = t.DockerContainerID
	s.Shell = t.Shell
	s.GRPC = t.GRPC
	s.GRPCUseTLS = t.GRPCUseTLS
	s.AliasNode = t.AliasNode
	s.AliasService = t.AliasService
	s.TTL = structs.DurationToProto(t.TTL)
}
