package catalogtest

import (
	"testing"

	svctest "github.com/arenadata/consul/agent/grpc-external/services/resource/testing"
	"github.com/arenadata/consul/internal/catalog"
	"github.com/arenadata/consul/internal/catalog/internal/controllers"
	"github.com/arenadata/consul/internal/controller"
	"github.com/arenadata/consul/internal/resource/reaper"
	"github.com/arenadata/consul/proto-public/pbresource"
	"github.com/arenadata/consul/sdk/testutil"
)

func runInMemResourceServiceAndControllers(t *testing.T, deps controllers.Dependencies) pbresource.ResourceServiceClient {
	t.Helper()

	ctx := testutil.TestContext(t)

	// Create the in-mem resource service
	client := svctest.RunResourceService(t, catalog.RegisterTypes)

	// Setup/Run the controller manager
	mgr := controller.NewManager(client, testutil.Logger(t))
	catalog.RegisterControllers(mgr, deps)

	// We also depend on the reaper to take care of cleaning up owned health statuses and
	// service endpoints so we must enable that controller as well
	reaper.RegisterControllers(mgr)
	mgr.SetRaftLeader(true)
	go mgr.Run(ctx)

	return client
}

func TestControllers_Integration(t *testing.T) {
	client := runInMemResourceServiceAndControllers(t, catalog.DefaultControllerDependencies())
	RunCatalogV1Alpha1IntegrationTest(t, client)
}
