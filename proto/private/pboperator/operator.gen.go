// Code generated by mog. DO NOT EDIT.

package pboperator

import "github.com/arenadata/consul/api"

func TransferLeaderResponseToAPI(s *TransferLeaderResponse, t *api.TransferLeaderResponse) {
	if s == nil {
		return
	}
	t.Success = s.Success
}
func TransferLeaderResponseFromAPI(t *api.TransferLeaderResponse, s *TransferLeaderResponse) {
	if s == nil {
		return
	}
	s.Success = t.Success
}
