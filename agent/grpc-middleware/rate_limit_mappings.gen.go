// generated by protoc-gen-consul-rate-limit; DO NOT EDIT.
package middleware

import "github.com/shulutkov/yellow-pages/agent/consul/rate"

var rpcRateLimitSpecs = map[string]rate.OperationSpec{
	"/hashicorp.consul.acl.ACLService/Login":                                     {Type: rate.OperationTypeWrite, Category: rate.OperationCategoryACL},
	"/hashicorp.consul.acl.ACLService/Logout":                                    {Type: rate.OperationTypeWrite, Category: rate.OperationCategoryACL},
	"/hashicorp.consul.connectca.ConnectCAService/Sign":                          {Type: rate.OperationTypeWrite, Category: rate.OperationCategoryConnectCA},
	"/hashicorp.consul.connectca.ConnectCAService/WatchRoots":                    {Type: rate.OperationTypeRead, Category: rate.OperationCategoryConnectCA},
	"/hashicorp.consul.dataplane.DataplaneService/GetEnvoyBootstrapParams":       {Type: rate.OperationTypeRead, Category: rate.OperationCategoryDataPlane},
	"/hashicorp.consul.dataplane.DataplaneService/GetSupportedDataplaneFeatures": {Type: rate.OperationTypeRead, Category: rate.OperationCategoryDataPlane},
	"/hashicorp.consul.dns.DNSService/Query":                                     {Type: rate.OperationTypeRead, Category: rate.OperationCategoryDNS},
	"/hashicorp.consul.internal.operator.OperatorService/TransferLeader":         {Type: rate.OperationTypeExempt, Category: rate.OperationCategoryOperator},
	"/hashicorp.consul.internal.peering.PeeringService/Establish":                {Type: rate.OperationTypeWrite, Category: rate.OperationCategoryPeering},
	"/hashicorp.consul.internal.peering.PeeringService/GenerateToken":            {Type: rate.OperationTypeWrite, Category: rate.OperationCategoryPeering},
	"/hashicorp.consul.internal.peering.PeeringService/PeeringDelete":            {Type: rate.OperationTypeWrite, Category: rate.OperationCategoryPeering},
	"/hashicorp.consul.internal.peering.PeeringService/PeeringList":              {Type: rate.OperationTypeRead, Category: rate.OperationCategoryPeering},
	"/hashicorp.consul.internal.peering.PeeringService/PeeringRead":              {Type: rate.OperationTypeRead, Category: rate.OperationCategoryPeering},
	"/hashicorp.consul.internal.peering.PeeringService/PeeringWrite":             {Type: rate.OperationTypeWrite, Category: rate.OperationCategoryPeering},
	"/hashicorp.consul.internal.peering.PeeringService/TrustBundleListByService": {Type: rate.OperationTypeRead, Category: rate.OperationCategoryPeering},
	"/hashicorp.consul.internal.peering.PeeringService/TrustBundleRead":          {Type: rate.OperationTypeRead, Category: rate.OperationCategoryPeering},
	"/hashicorp.consul.internal.peerstream.PeerStreamService/ExchangeSecret":     {Type: rate.OperationTypeWrite, Category: rate.OperationCategoryPeerStream},
	"/hashicorp.consul.internal.peerstream.PeerStreamService/StreamResources":    {Type: rate.OperationTypeRead, Category: rate.OperationCategoryPeerStream},
	"/hashicorp.consul.internal.storage.raft.ForwardingService/Delete":           {Type: rate.OperationTypeExempt, Category: rate.OperationCategoryResource},
	"/hashicorp.consul.internal.storage.raft.ForwardingService/List":             {Type: rate.OperationTypeExempt, Category: rate.OperationCategoryResource},
	"/hashicorp.consul.internal.storage.raft.ForwardingService/Read":             {Type: rate.OperationTypeExempt, Category: rate.OperationCategoryResource},
	"/hashicorp.consul.internal.storage.raft.ForwardingService/Write":            {Type: rate.OperationTypeExempt, Category: rate.OperationCategoryResource},
	"/hashicorp.consul.resource.ResourceService/Delete":                          {Type: rate.OperationTypeWrite, Category: rate.OperationCategoryResource},
	"/hashicorp.consul.resource.ResourceService/List":                            {Type: rate.OperationTypeRead, Category: rate.OperationCategoryResource},
	"/hashicorp.consul.resource.ResourceService/ListByOwner":                     {Type: rate.OperationTypeRead, Category: rate.OperationCategoryResource},
	"/hashicorp.consul.resource.ResourceService/Read":                            {Type: rate.OperationTypeRead, Category: rate.OperationCategoryResource},
	"/hashicorp.consul.resource.ResourceService/WatchList":                       {Type: rate.OperationTypeRead, Category: rate.OperationCategoryResource},
	"/hashicorp.consul.resource.ResourceService/Write":                           {Type: rate.OperationTypeWrite, Category: rate.OperationCategoryResource},
	"/hashicorp.consul.resource.ResourceService/WriteStatus":                     {Type: rate.OperationTypeWrite, Category: rate.OperationCategoryResource},
	"/hashicorp.consul.serverdiscovery.ServerDiscoveryService/WatchServers":      {Type: rate.OperationTypeRead, Category: rate.OperationCategoryServerDiscovery},
	"/subscribe.StateChangeSubscription/Subscribe":                               {Type: rate.OperationTypeRead, Category: rate.OperationCategorySubscribe},
}
