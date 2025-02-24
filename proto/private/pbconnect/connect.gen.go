// Code generated by mog. DO NOT EDIT.

package pbconnect

import "github.com/shulutkov/yellow-pages/agent/structs"

func CARootToStructsCARoot(s *CARoot, t *structs.CARoot) {
	if s == nil {
		return
	}
	t.ID = s.ID
	t.Name = s.Name
	t.SerialNumber = s.SerialNumber
	t.SigningKeyID = s.SigningKeyID
	t.ExternalTrustDomain = s.ExternalTrustDomain
	t.NotBefore = structs.TimeFromProto(s.NotBefore)
	t.NotAfter = structs.TimeFromProto(s.NotAfter)
	t.RootCert = s.RootCert
	t.IntermediateCerts = s.IntermediateCerts
	t.SigningCert = s.SigningCert
	t.SigningKey = s.SigningKey
	t.Active = s.Active
	t.RotatedOutAt = structs.TimeFromProto(s.RotatedOutAt)
	t.PrivateKeyType = s.PrivateKeyType
	t.PrivateKeyBits = int(s.PrivateKeyBits)
	t.RaftIndex = RaftIndexTo(s.RaftIndex)
}
func CARootFromStructsCARoot(t *structs.CARoot, s *CARoot) {
	if s == nil {
		return
	}
	s.ID = t.ID
	s.Name = t.Name
	s.SerialNumber = t.SerialNumber
	s.SigningKeyID = t.SigningKeyID
	s.ExternalTrustDomain = t.ExternalTrustDomain
	s.NotBefore = structs.TimeToProto(t.NotBefore)
	s.NotAfter = structs.TimeToProto(t.NotAfter)
	s.RootCert = t.RootCert
	s.IntermediateCerts = t.IntermediateCerts
	s.SigningCert = t.SigningCert
	s.SigningKey = t.SigningKey
	s.Active = t.Active
	s.RotatedOutAt = structs.TimeToProto(t.RotatedOutAt)
	s.PrivateKeyType = t.PrivateKeyType
	s.PrivateKeyBits = int32(t.PrivateKeyBits)
	s.RaftIndex = RaftIndexFrom(t.RaftIndex)
}
func CARootsToStructsIndexedCARoots(s *CARoots, t *structs.IndexedCARoots) {
	if s == nil {
		return
	}
	t.ActiveRootID = s.ActiveRootID
	t.TrustDomain = s.TrustDomain
	{
		t.Roots = make([]*structs.CARoot, len(s.Roots))
		for i := range s.Roots {
			if s.Roots[i] != nil {
				var x structs.CARoot
				CARootToStructsCARoot(s.Roots[i], &x)
				t.Roots[i] = &x
			}
		}
	}
	t.QueryMeta = QueryMetaTo(s.QueryMeta)
}
func CARootsFromStructsIndexedCARoots(t *structs.IndexedCARoots, s *CARoots) {
	if s == nil {
		return
	}
	s.ActiveRootID = t.ActiveRootID
	s.TrustDomain = t.TrustDomain
	{
		s.Roots = make([]*CARoot, len(t.Roots))
		for i := range t.Roots {
			if t.Roots[i] != nil {
				var x CARoot
				CARootFromStructsCARoot(t.Roots[i], &x)
				s.Roots[i] = &x
			}
		}
	}
	s.QueryMeta = QueryMetaFrom(t.QueryMeta)
}
func IssuedCertToStructsIssuedCert(s *IssuedCert, t *structs.IssuedCert) {
	if s == nil {
		return
	}
	t.SerialNumber = s.SerialNumber
	t.CertPEM = s.CertPEM
	t.PrivateKeyPEM = s.PrivateKeyPEM
	t.Service = s.Service
	t.ServiceURI = s.ServiceURI
	t.Agent = s.Agent
	t.AgentURI = s.AgentURI
	t.ServerURI = s.ServerURI
	t.Kind = structs.ServiceKind(s.Kind)
	t.KindURI = s.KindURI
	t.ValidAfter = structs.TimeFromProto(s.ValidAfter)
	t.ValidBefore = structs.TimeFromProto(s.ValidBefore)
	t.EnterpriseMeta = EnterpriseMetaTo(s.EnterpriseMeta)
	t.RaftIndex = RaftIndexTo(s.RaftIndex)
}
func IssuedCertFromStructsIssuedCert(t *structs.IssuedCert, s *IssuedCert) {
	if s == nil {
		return
	}
	s.SerialNumber = t.SerialNumber
	s.CertPEM = t.CertPEM
	s.PrivateKeyPEM = t.PrivateKeyPEM
	s.Service = t.Service
	s.ServiceURI = t.ServiceURI
	s.Agent = t.Agent
	s.AgentURI = t.AgentURI
	s.ServerURI = t.ServerURI
	s.Kind = string(t.Kind)
	s.KindURI = t.KindURI
	s.ValidAfter = structs.TimeToProto(t.ValidAfter)
	s.ValidBefore = structs.TimeToProto(t.ValidBefore)
	s.EnterpriseMeta = EnterpriseMetaFrom(t.EnterpriseMeta)
	s.RaftIndex = RaftIndexFrom(t.RaftIndex)
}
