// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"strings"
	"testing"
	"time"

	"github.com/arenadata/consul/sdk/testutil/retry"

	"github.com/hashicorp/go-uuid"

	"github.com/stretchr/testify/require"
)

func TestAPI_ClientTxn(t *testing.T) {
	t.Parallel()
	c, s := makeClient(t)
	defer s.Stop()

	s.WaitForSerfCheck(t)

	session := c.Session()
	txn := c.Txn()

	// Set up a test service and health check.
	nodeID, err := uuid.GenerateUUID()
	require.NoError(t, err)

	catalog := c.Catalog()
	reg := &CatalogRegistration{
		ID:      nodeID,
		Node:    "foo",
		Address: "2.2.2.2",
		Service: &AgentService{
			ID:      "foo1",
			Service: "foo",
		},
		Checks: HealthChecks{
			{
				CheckID: "bar",
				Status:  "critical",
				Definition: HealthCheckDefinition{
					TCP:                                    "1.1.1.1",
					IntervalDuration:                       5 * time.Second,
					TimeoutDuration:                        10 * time.Second,
					DeregisterCriticalServiceAfterDuration: 20 * time.Second,
				},
			},
			{
				CheckID: "baz",
				Status:  "passing",
				Definition: HealthCheckDefinition{
					TCP:                            "2.2.2.2",
					Interval:                       ReadableDuration(40 * time.Second),
					Timeout:                        ReadableDuration(80 * time.Second),
					DeregisterCriticalServiceAfter: ReadableDuration(160 * time.Second),
				},
			},
			{
				CheckID: "bor",
				Status:  "critical",
				Definition: HealthCheckDefinition{
					UDP:                            "1.1.1.1",
					Interval:                       ReadableDuration(5 * time.Second),
					Timeout:                        ReadableDuration(10 * time.Second),
					DeregisterCriticalServiceAfter: ReadableDuration(20 * time.Second),
				},
				Type: "udp",
			},
			{
				CheckID: "bur",
				Status:  "passing",
				Definition: HealthCheckDefinition{
					UDP:                            "2.2.2.2",
					Interval:                       ReadableDuration(5 * time.Second),
					Timeout:                        ReadableDuration(10 * time.Second),
					DeregisterCriticalServiceAfter: ReadableDuration(20 * time.Second),
				},
				Type: "udp",
			},
		},
	}
	_, err = catalog.Register(reg, nil)
	require.NoError(t, err)

	node, _, err := catalog.Node("foo", nil)
	require.NoError(t, err)
	require.Equal(t, nodeID, node.Node.ID)

	// Make a session.
	id, _, err := session.CreateNoChecks(nil, nil)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	defer session.Destroy(id, nil)

	// Acquire and get the key via a transaction, but don't supply a valid
	// session.
	key := testKey()
	value := []byte("test")
	ops := TxnOps{
		&TxnOp{
			KV: &KVTxnOp{
				Verb:  KVLock,
				Key:   key,
				Value: value,
			},
		},
		&TxnOp{
			KV: &KVTxnOp{
				Verb: KVGet,
				Key:  key,
			},
		},
		&TxnOp{
			Node: &NodeTxnOp{
				Verb: NodeGet,
				Node: Node{Node: "foo"},
			},
		},
		&TxnOp{
			Service: &ServiceTxnOp{
				Verb:    ServiceGet,
				Node:    "foo",
				Service: AgentService{ID: "foo1"},
			},
		},
		&TxnOp{
			Check: &CheckTxnOp{
				Verb:  CheckGet,
				Check: HealthCheck{Node: "foo", CheckID: "bar"},
			},
		},
		&TxnOp{
			Check: &CheckTxnOp{
				Verb:  CheckGet,
				Check: HealthCheck{Node: "foo", CheckID: "baz"},
			},
		},
		&TxnOp{
			Check: &CheckTxnOp{
				Verb:  CheckGet,
				Check: HealthCheck{Node: "foo", CheckID: "bor"},
			},
		},
		&TxnOp{
			Check: &CheckTxnOp{
				Verb:  CheckGet,
				Check: HealthCheck{Node: "foo", CheckID: "bur"},
			},
		},
	}
	ok, ret, _, err := txn.Txn(ops, nil)
	if err != nil {
		t.Fatalf("err: %v", err)
	} else if ok {
		t.Fatalf("transaction should have failed")
	}

	if ret == nil || len(ret.Errors) != 2 || len(ret.Results) != 0 {
		t.Fatalf("bad: %v", ret.Errors[2])
	}
	if ret.Errors[0].OpIndex != 0 ||
		!strings.Contains(ret.Errors[0].What, "missing session") ||
		!strings.Contains(ret.Errors[1].What, "doesn't exist") {
		t.Fatalf("bad: %v", ret.Errors[0])
	}

	// Now poke in a real session and try again.
	ops[0].KV.Session = id
	ok, ret, _, err = txn.Txn(ops, nil)
	if err != nil {
		t.Fatalf("err: %v", err)
	} else if !ok {
		t.Fatalf("transaction failure")
	}

	if ret == nil || len(ret.Errors) != 0 || len(ret.Results) != 8 {
		t.Fatalf("bad: %v", ret)
	}
	expected := TxnResults{
		&TxnResult{
			KV: &KVPair{
				Key:         key,
				Session:     id,
				LockIndex:   1,
				CreateIndex: ret.Results[0].KV.CreateIndex,
				ModifyIndex: ret.Results[0].KV.ModifyIndex,
				Namespace:   ret.Results[0].KV.Namespace,
				Partition:   defaultPartition,
			},
		},
		&TxnResult{
			KV: &KVPair{
				Key:         key,
				Session:     id,
				Value:       []byte("test"),
				LockIndex:   1,
				CreateIndex: ret.Results[1].KV.CreateIndex,
				ModifyIndex: ret.Results[1].KV.ModifyIndex,
				Namespace:   ret.Results[0].KV.Namespace,
				Partition:   defaultPartition,
			},
		},
		&TxnResult{
			Node: &Node{
				ID:          nodeID,
				Node:        "foo",
				Partition:   defaultPartition,
				Address:     "2.2.2.2",
				Datacenter:  "dc1",
				CreateIndex: ret.Results[2].Node.CreateIndex,
				ModifyIndex: ret.Results[2].Node.CreateIndex,
			},
		},
		&TxnResult{
			Service: &CatalogService{
				ID:          "foo1",
				CreateIndex: ret.Results[3].Service.CreateIndex,
				ModifyIndex: ret.Results[3].Service.CreateIndex,
				Partition:   defaultPartition,
				Namespace:   defaultNamespace,
			},
		},
		&TxnResult{
			Check: &HealthCheck{
				Node:    "foo",
				CheckID: "bar",
				Status:  "critical",
				Definition: HealthCheckDefinition{
					TCP:                                    "1.1.1.1",
					Interval:                               ReadableDuration(5 * time.Second),
					IntervalDuration:                       5 * time.Second,
					Timeout:                                ReadableDuration(10 * time.Second),
					TimeoutDuration:                        10 * time.Second,
					DeregisterCriticalServiceAfter:         ReadableDuration(20 * time.Second),
					DeregisterCriticalServiceAfterDuration: 20 * time.Second,
				},
				Type:        "tcp",
				Partition:   defaultPartition,
				Namespace:   defaultNamespace,
				CreateIndex: ret.Results[4].Check.CreateIndex,
				ModifyIndex: ret.Results[4].Check.CreateIndex,
			},
		},
		&TxnResult{
			Check: &HealthCheck{
				Node:    "foo",
				CheckID: "baz",
				Status:  "passing",
				Definition: HealthCheckDefinition{
					TCP:                                    "2.2.2.2",
					Interval:                               ReadableDuration(40 * time.Second),
					IntervalDuration:                       40 * time.Second,
					Timeout:                                ReadableDuration(80 * time.Second),
					TimeoutDuration:                        80 * time.Second,
					DeregisterCriticalServiceAfter:         ReadableDuration(160 * time.Second),
					DeregisterCriticalServiceAfterDuration: 160 * time.Second,
				},
				Type:        "tcp",
				Partition:   defaultPartition,
				Namespace:   defaultNamespace,
				CreateIndex: ret.Results[4].Check.CreateIndex,
				ModifyIndex: ret.Results[4].Check.CreateIndex,
			},
		},
		&TxnResult{
			Check: &HealthCheck{
				Node:    "foo",
				CheckID: "bor",
				Status:  "critical",
				Definition: HealthCheckDefinition{
					UDP:                                    "1.1.1.1",
					Interval:                               ReadableDuration(5 * time.Second),
					IntervalDuration:                       5 * time.Second,
					Timeout:                                ReadableDuration(10 * time.Second),
					TimeoutDuration:                        10 * time.Second,
					DeregisterCriticalServiceAfter:         ReadableDuration(20 * time.Second),
					DeregisterCriticalServiceAfterDuration: 20 * time.Second,
				},
				Type:        "udp",
				Partition:   defaultPartition,
				Namespace:   defaultNamespace,
				CreateIndex: ret.Results[4].Check.CreateIndex,
				ModifyIndex: ret.Results[4].Check.CreateIndex,
			},
		},
		&TxnResult{
			Check: &HealthCheck{
				Node:    "foo",
				CheckID: "bur",
				Status:  "passing",
				Definition: HealthCheckDefinition{
					UDP:                                    "2.2.2.2",
					Interval:                               ReadableDuration(5 * time.Second),
					IntervalDuration:                       5 * time.Second,
					Timeout:                                ReadableDuration(10 * time.Second),
					TimeoutDuration:                        10 * time.Second,
					DeregisterCriticalServiceAfter:         ReadableDuration(20 * time.Second),
					DeregisterCriticalServiceAfterDuration: 20 * time.Second,
				},
				Type:        "udp",
				Partition:   defaultPartition,
				Namespace:   defaultNamespace,
				CreateIndex: ret.Results[4].Check.CreateIndex,
				ModifyIndex: ret.Results[4].Check.CreateIndex,
			},
		},
	}
	require.Equal(t, expected, ret.Results)

	retry.Run(t, func(r *retry.R) {
		// Run a read-only transaction.
		ops = TxnOps{
			&TxnOp{
				KV: &KVTxnOp{
					Verb: KVGet,
					Key:  key,
				},
			},
			&TxnOp{
				Node: &NodeTxnOp{
					Verb: NodeGet,
					Node: Node{ID: s.Config.NodeID, Node: s.Config.NodeName},
				},
			},
		}
		ok, ret, _, err = txn.Txn(ops, nil)
		if err != nil {
			r.Fatalf("err: %v", err)
		} else if !ok {
			r.Fatalf("transaction failure")
		}

		expected = TxnResults{
			&TxnResult{
				KV: &KVPair{
					Key:         key,
					Session:     id,
					Value:       []byte("test"),
					LockIndex:   1,
					CreateIndex: ret.Results[0].KV.CreateIndex,
					ModifyIndex: ret.Results[0].KV.ModifyIndex,
					Namespace:   ret.Results[0].KV.Namespace,
					Partition:   defaultPartition,
				},
			},
			&TxnResult{
				Node: &Node{
					ID:         s.Config.NodeID,
					Node:       s.Config.NodeName,
					Partition:  defaultPartition,
					Address:    "127.0.0.1",
					Datacenter: "dc1",
					TaggedAddresses: map[string]string{
						"lan":      s.Config.Bind,
						"lan_ipv4": s.Config.Bind,
						"wan":      s.Config.Bind,
						"wan_ipv4": s.Config.Bind,
					},
					Meta: map[string]string{
						"consul-network-segment": "",
						"consul-version":         s.Config.Version,
					},
					CreateIndex: ret.Results[1].Node.CreateIndex,
					ModifyIndex: ret.Results[1].Node.ModifyIndex,
				},
			},
		}
		require.Equal(r, expected, ret.Results)
	})

	// Sanity check using the regular GET API.
	kv := c.KV()
	pair, meta, err := kv.Get(key, nil)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if pair == nil {
		t.Fatalf("expected value: %#v", pair)
	}
	if pair.LockIndex != 1 {
		t.Fatalf("Expected lock: %v", pair)
	}
	if pair.Session != id {
		t.Fatalf("Expected lock: %v", pair)
	}
	if meta.LastIndex == 0 {
		t.Fatalf("unexpected value: %#v", meta)
	}
}
