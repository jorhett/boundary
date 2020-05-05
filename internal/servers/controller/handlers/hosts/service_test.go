package hosts_test

import (
	"context"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/watchtower/internal/gen/controller/api"
	"github.com/hashicorp/watchtower/internal/servers/controller/handlers/hosts"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDelete(t *testing.T) {
	toMerge := &api.DeleteHostRequest{
		OrgId:         "1",
		ProjectId:     "2",
		HostCatalogId: "3",
		Id:            "4",
	}

	s := hosts.Service{}
	cases := []struct {
		name    string
		req     *api.DeleteHostRequest
		res     *api.DeleteHostResponse
		errCode codes.Code
	}{
		{
			name:    "Success even when doesn't exist",
			req:     &api.DeleteHostRequest{},
			res:     &api.DeleteHostResponse{},
			errCode: codes.OK,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := proto.Clone(toMerge).(*api.DeleteHostRequest)
			proto.Merge(req, tc.req)
			got, gErr := s.DeleteHost(context.Background(), req)

			if status.Code(gErr) != tc.errCode {
				t.Errorf("DeleteHost(%+v) got error %v, wanted %v", req, gErr, tc.errCode)
			}

			if !proto.Equal(got, tc.res) {
				t.Errorf("DeleteHost(%q) got response %q, wanted %q", req, got, tc.res)
			}
		})
	}
}

func TestList(t *testing.T) {
	toMerge := &api.ListHostsRequest{
		OrgId:         "1",
		ProjectId:     "2",
		HostCatalogId: "3",
	}

	s := hosts.Service{}
	cases := []struct {
		name    string
		req     *api.ListHostsRequest
		res     *api.ListHostsResponse
		errCode codes.Code
	}{
		{
			name: "List from a valid catalog id",
			req:  &api.ListHostsRequest{},
			// TODO: Update this when the List method is implemented
			res:     nil,
			errCode: codes.NotFound,
		},
		{
			name:    "Non Existant Host Catalog",
			req:     &api.ListHostsRequest{HostCatalogId: "this doesnt exist"},
			res:     nil,
			errCode: codes.NotFound,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := proto.Clone(toMerge).(*api.ListHostsRequest)
			proto.Merge(req, tc.req)
			got, gErr := s.ListHosts(context.Background(), req)
			if status.Code(gErr) != tc.errCode {
				t.Errorf("ListHosts(%+v) got error %v, wanted %v", req, gErr, tc.errCode)
			}
			if !proto.Equal(got, tc.res) {
				t.Errorf("ListHosts(%q) got response %q, wanted %q", req, got, tc.res)
			}
		})
	}
}

func TestGet(t *testing.T) {
	toMerge := &api.GetHostRequest{
		OrgId:         "1",
		ProjectId:     "2",
		HostCatalogId: "3",
	}

	s := hosts.Service{}
	cases := []struct {
		name    string
		req     *api.GetHostRequest
		res     *api.GetHostResponse
		errCode codes.Code
	}{
		// TODO: These cases need to be updated as the handlers get implemented.
		{
			name:    "Default request",
			req:     &api.GetHostRequest{},
			res:     nil,
			errCode: codes.Unimplemented,
		},
		{
			name: "Non Existant Host Catalog",
			req:  &api.GetHostRequest{HostCatalogId: "this doesnt exist"},
			// The response and error will need to change when this is implemented to be a 404 error
			res:     nil,
			errCode: codes.Unimplemented,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := proto.Clone(toMerge).(*api.GetHostRequest)
			proto.Merge(req, tc.req)
			got, gErr := s.GetHost(context.Background(), req)
			if status.Code(gErr) != tc.errCode {
				t.Errorf("ListHosts(%+v) got error %v, wanted %v", req, gErr, tc.errCode)
			}
			if !proto.Equal(got, tc.res) {
				t.Errorf("ListHosts(%q) got response %q, wanted %q", req, got, tc.res)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	toMerge := &api.UpdateHostRequest{
		OrgId:         "1",
		ProjectId:     "2",
		HostCatalogId: "3",
	}

	s := hosts.Service{}
	cases := []struct {
		name    string
		req     *api.UpdateHostRequest
		res     *api.UpdateHostResponse
		errCode codes.Code
	}{
		// TODO: These cases need to be updated as the handlers get implemented.
		{
			name:    "Default request",
			req:     &api.UpdateHostRequest{},
			res:     nil,
			errCode: codes.Unimplemented,
		},
		{
			name: "Non Existant Host Catalog",
			req:  &api.UpdateHostRequest{HostCatalogId: "this doesnt exist"},
			// The response and error will need to change when this is implemented to be a 404 error
			res:     nil,
			errCode: codes.Unimplemented,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := proto.Clone(toMerge).(*api.UpdateHostRequest)
			proto.Merge(req, tc.req)
			got, gErr := s.UpdateHost(context.Background(), req)
			if status.Code(gErr) != tc.errCode {
				t.Errorf("UpdateHost(%+v) got error %v, wanted %v", req, gErr, tc.errCode)
			}
			if !proto.Equal(got, tc.res) {
				t.Errorf("UpdateHost(%q) got response %q, wanted %q", req, got, tc.res)
			}
		})
	}
}