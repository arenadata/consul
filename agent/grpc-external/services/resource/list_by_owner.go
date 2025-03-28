// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arenadata/consul/acl"
	"github.com/arenadata/consul/proto-public/pbresource"
)

func (s *Server) ListByOwner(ctx context.Context, req *pbresource.ListByOwnerRequest) (*pbresource.ListByOwnerResponse, error) {
	if err := validateListByOwnerRequest(req); err != nil {
		return nil, err
	}

	_, err := s.resolveType(req.Owner.Type)
	if err != nil {
		return nil, err
	}

	children, err := s.Backend.ListByOwner(ctx, req.Owner)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed list by owner: %v", err)
	}

	authz, err := s.getAuthorizer(tokenFromContext(ctx))
	if err != nil {
		return nil, err
	}

	result := make([]*pbresource.Resource, 0)
	for _, child := range children {
		reg, err := s.resolveType(child.Id.Type)
		if err != nil {
			return nil, err
		}

		// ACL filter
		err = reg.ACLs.Read(authz, child.Id)
		switch {
		case acl.IsErrPermissionDenied(err):
			continue
		case err != nil:
			return nil, status.Errorf(codes.Internal, "failed read acl: %v", err)
		}

		result = append(result, child)
	}
	return &pbresource.ListByOwnerResponse{Resources: result}, nil
}

func validateListByOwnerRequest(req *pbresource.ListByOwnerRequest) error {
	if req.Owner == nil {
		return status.Errorf(codes.InvalidArgument, "owner is required")
	}

	if err := validateId(req.Owner, "owner"); err != nil {
		return err
	}

	if req.Owner.Uid == "" {
		return status.Errorf(codes.InvalidArgument, "owner uid is required")
	}
	return nil
}
