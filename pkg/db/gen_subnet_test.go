// nolint
package db

import (
	"context"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"
)

//For skip import error.
var _ = errors.New("")

func TestSubnet(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "subnet")
	// mutexProject := UseTable(db.DB, "subnet")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeSubnet()
	model.UUID = "subnet_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "subnet_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var VirtualMachineInterfacecreateref []*models.SubnetVirtualMachineInterfaceRef
	var VirtualMachineInterfacerefModel *models.VirtualMachineInterface
	VirtualMachineInterfacerefModel = models.MakeVirtualMachineInterface()
	VirtualMachineInterfacerefModel.UUID = "subnet_virtual_machine_interface_ref_uuid"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "subnet_virtual_machine_interface_ref_uuid"}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfacerefModel,
	})
	VirtualMachineInterfacerefModel.UUID = "subnet_virtual_machine_interface_ref_uuid1"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "subnet_virtual_machine_interface_ref_uuid1"}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfacerefModel,
	})
	VirtualMachineInterfacerefModel.UUID = "subnet_virtual_machine_interface_ref_uuid2"
	VirtualMachineInterfacerefModel.FQName = []string{"test", "subnet_virtual_machine_interface_ref_uuid2"}
	_, err = db.CreateVirtualMachineInterface(ctx, &models.CreateVirtualMachineInterfaceRequest{
		VirtualMachineInterface: VirtualMachineInterfacerefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.SubnetVirtualMachineInterfaceRef{UUID: "subnet_virtual_machine_interface_ref_uuid", To: []string{"test", "subnet_virtual_machine_interface_ref_uuid"}})
	VirtualMachineInterfacecreateref = append(VirtualMachineInterfacecreateref, &models.SubnetVirtualMachineInterfaceRef{UUID: "subnet_virtual_machine_interface_ref_uuid2", To: []string{"test", "subnet_virtual_machine_interface_ref_uuid2"}})
	model.VirtualMachineInterfaceRefs = VirtualMachineInterfacecreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "subnet_admin_project_uuid"
	projectModel.FQName = []string{"default-domain-test", "admin-test"}
	projectModel.Perms2.Owner = "admin"
	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:admin-test", TenantAccess: 7})
	model.Perms2.Share = createShare

	_, err = db.CreateProject(ctx, &models.CreateProjectRequest{
		Project: projectModel,
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	//    //populate update map
	//    updateMap := map[string]interface{}{}
	//
	//
	//    common.SetValueByPath(updateMap, ".UUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".SubnetIPPrefix.IPPrefixLen", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".SubnetIPPrefix.IPPrefix", ".", "test")
	//
	//
	//
	//    if ".Perms2.Share" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.UserVisible", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OtherAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.GroupAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Group", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.LastModified", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Enable", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Description", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Creator", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Created", ".", "test")
	//
	//
	//
	//    if ".FQName" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".FQName", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".FQName", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DisplayName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ConfigurationVersion", ".", 1.0)
	//
	//
	//
	//    if ".Annotations.KeyValuePair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
	//    }
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "subnet_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var VirtualMachineInterfaceref []interface{}
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"delete", "uuid":"subnet_virtual_machine_interface_ref_uuid", "to": []string{"test", "subnet_virtual_machine_interface_ref_uuid"}})
	//    VirtualMachineInterfaceref = append(VirtualMachineInterfaceref, map[string]interface{}{"operation":"add", "uuid":"subnet_virtual_machine_interface_ref_uuid1", "to": []string{"test", "subnet_virtual_machine_interface_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "VirtualMachineInterfaceRefs", ".", VirtualMachineInterfaceref)
	//
	//
	_, err = db.CreateSubnet(ctx,
		&models.CreateSubnetRequest{
			Subnet: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateSubnet(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_subnet_virtual_machine_interface` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs delete statement failed")
		}
		_, err = stmt.Exec("subnet_dummy_uuid", "subnet_virtual_machine_interface_ref_uuid")
		_, err = stmt.Exec("subnet_dummy_uuid", "subnet_virtual_machine_interface_ref_uuid1")
		_, err = stmt.Exec("subnet_dummy_uuid", "subnet_virtual_machine_interface_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteVirtualMachineInterface(ctx,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: "subnet_virtual_machine_interface_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref subnet_virtual_machine_interface_ref_uuid  failed", err)
	}
	_, err = db.DeleteVirtualMachineInterface(ctx,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: "subnet_virtual_machine_interface_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref subnet_virtual_machine_interface_ref_uuid1  failed", err)
	}
	_, err = db.DeleteVirtualMachineInterface(
		ctx,
		&models.DeleteVirtualMachineInterfaceRequest{
			ID: "subnet_virtual_machine_interface_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref subnet_virtual_machine_interface_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	response, err := db.ListSubnet(ctx, &models.ListSubnetRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.Subnets) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteSubnet(ctxDemo,
		&models.DeleteSubnetRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateSubnet(ctx,
		&models.CreateSubnetRequest{
			Subnet: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteSubnet(ctx,
		&models.DeleteSubnetRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	response, err = db.ListSubnet(ctx, &models.ListSubnetRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.Subnets) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
