package server

import (
	"fmt"
	"time"

	"github.com/netdoop/netdoop/store"
	"github.com/netdoop/netdoop/utils"

	"github.com/heypkg/iam"
	"github.com/pkg/errors"
)

var apiRulesMap = map[string]iam.ApiRule{
	// "api.system.time.get": {Method: "GET", Path: "/api/v1/system/time"},

	"api.omc.data.query": {Method: "POST", Path: "/api/v1/omc/data"},

	"api.omc.groups.list":          {Method: "GET", Path: "/api/v1/omc/groups"},
	"api.omc.groups.create":        {Method: "POST", Path: "/api/v1/omc/groups"},
	"api.omc.groups.get":           {Method: "GET", Path: "/api/v1/omc/groups/:id"},
	"api.omc.groups.update":        {Method: "PUT", Path: "/api/v1/omc/groups/:id"},
	"api.omc.groups.delete":        {Method: "DELETE", Path: "/api/v1/omc/groups/:id"},
	"api.omc.groups.children.list": {Method: "GET", Path: "/api/v1/omc/groups/:id/children"},
	"api.omc.groups.parent.update": {Method: "PUT", Path: "/api/v1/omc/groups/:id/parent"},

	"api.omc.kpi.measures.list":    {Method: "GET", Path: "/api/v1/omc/kpi/measures"},
	"api.omc.kpi.measures.create":  {Method: "POST", Path: "/api/v1/omc/kpi/measures"},
	"api.omc.kpi.measures.delete":  {Method: "DELETE", Path: "/api/v1/omc/kpi/measures/:id"},
	"api.omc.kpi.measures.get":     {Method: "GET", Path: "/api/v1/omc/kpi/measures/:id"},
	"api.omc.kpi.measures.update":  {Method: "PUT", Path: "/api/v1/omc/kpi/measures/:id"},
	"api.omc.kpi.measures.enable":  {Method: "PUT", Path: "/api/v1/omc/kpi/measures/:id/enable"},
	"api.omc.kpi.measures.disable": {Method: "PUT", Path: "/api/v1/omc/kpi/measures/:id/disable"},

	"api.omc.kpi.templates.list":         {Method: "GET", Path: "/api/v1/omc/kpi/templates"},
	"api.omc.kpi.templates.create":       {Method: "POST", Path: "/api/v1/omc/kpi/templates"},
	"api.omc.kpi.templates.delete":       {Method: "DELETE", Path: "/api/v1/omc/kpi/templates/:id"},
	"api.omc.kpi.templates.get":          {Method: "GET", Path: "/api/v1/omc/kpi/templates/:id"},
	"api.omc.kpi.templates.update":       {Method: "PUT", Path: "/api/v1/omc/kpi/templates/:id"},
	"api.omc.kpi.templates.list-records": {Method: "GET", Path: "/api/v1/omc/kpi/templates/:id/records"},

	"api.omc.datamodels.list":   {Method: "GET", Path: "/api/v1/omc/datamodels"},
	"api.omc.datamodels.create": {Method: "POST", Path: "/api/v1/omc/datamodels"},
	"api.omc.datamodels.delete": {Method: "DELETE", Path: "/api/v1/omc/datamodels/:id"},
	"api.omc.datamodels.get":    {Method: "GET", Path: "/api/v1/omc/datamodels/:id"},

	"api.omc.datamodels.parameters.list":   {Method: "GET", Path: "/api/v1/omc/datamodels/:id/parameters"},
	"api.omc.datamodels.parameters.create": {Method: "POST", Path: "/api/v1/omc/datamodels/:id/parameters"},
	"api.omc.datamodels.parameters.delete": {Method: "DELETE", Path: "/api/v1/omc/datamodels/:id/parameters/:parameter_id"},
	"api.omc.datamodels.parameters.get":    {Method: "GET", Path: "/api/v1/omc/datamodels/:id/parameters/:parameter_id"},

	"api.omc.datamodels.templates.list":   {Method: "GET", Path: "/api/v1/omc/datamodels/:id/templates"},
	"api.omc.datamodels.templates.create": {Method: "POST", Path: "/api/v1/omc/datamodels/:id/templates"},
	"api.omc.datamodels.templates.delete": {Method: "DELETE", Path: "/api/v1/omc/datamodels/:id/templates/:template_id"},
	"api.omc.datamodels.templates.get":    {Method: "GET", Path: "/api/v1/omc/datamodels/:id/templates/:template_id"},

	"api.omc.products.list":           {Method: "GET", Path: "/api/v1/omc/products"},
	"api.omc.products.create":         {Method: "POST", Path: "/api/v1/omc/products"},
	"api.omc.products.delete":         {Method: "DELETE", Path: "/api/v1/omc/products/:id"},
	"api.omc.products.get":            {Method: "GET", Path: "/api/v1/omc/products/:id"},
	"api.omc.products.update":         {Method: "PUT", Path: "/api/v1/omc/products/:id"},
	"api.omc.products.enable":         {Method: "PUT", Path: "/api/v1/omc/products/:id/enable"},
	"api.omc.products.disable":        {Method: "PUT", Path: "/api/v1/omc/products/:id/disable"},
	"api.omc.products.firmwares.list": {Method: "GET", Path: "/api/v1/omc/products/:id/firmwares"},

	"api.omc.deleted-products.list":   {Method: "GET", Path: "/api/v1/omc/deleted-products"},
	"api.omc.deleted-products.delete": {Method: "DELETE", Path: "/api/v1/omc/deleted-products/:id"},

	"api.omc.devices.list":                 {Method: "GET", Path: "/api/v1/omc/devices"},
	"api.omc.devices.create":               {Method: "POST", Path: "/api/v1/omc/devices"},
	"api.omc.devices.update":               {Method: "PUT", Path: "/api/v1/omc/devices/:id"},
	"api.omc.devices.delete":               {Method: "DELETE", Path: "/api/v1/omc/devices/:id"},
	"api.omc.devices.get":                  {Method: "GET", Path: "/api/v1/omc/devices/:id"},
	"api.omc.devices.enable":               {Method: "PUT", Path: "/api/v1/omc/devices/:id/enable"},
	"api.omc.devices.disable":              {Method: "PUT", Path: "/api/v1/omc/devices/:id/disable"},
	"api.omc.devices.set-group":            {Method: "PUT", Path: "/api/v1/omc/devices/:id/group"},
	"api.omc.devices.add-object":           {Method: "POST", Path: "/api/v1/omc/devices/:id/add-object"},
	"api.omc.devices.delete-object":        {Method: "POST", Path: "/api/v1/omc/devices/:id/delete-object"},
	"api.omc.devices.get-parameter-names":  {Method: "POST", Path: "/api/v1/omc/devices/:id/get-parameter-names"},
	"api.omc.devices.get-parameter-values": {Method: "POST", Path: "/api/v1/omc/devices/:id/get-parameter-values"},
	"api.omc.devices.get-methods":          {Method: "GET", Path: "/api/v1/omc/devices/:id/methods"},
	"api.omc.devices.get-parameters":       {Method: "GET", Path: "/api/v1/omc/devices/:id/parameters"},
	"api.omc.devices.reboot":               {Method: "POST", Path: "/api/v1/omc/devices/:id/reboot"},
	"api.omc.devices.set-parameter-values": {Method: "POST", Path: "/api/v1/omc/devices/:id/set-parameter-values"},
	"api.omc.devices.upload-file":          {Method: "POST", Path: "/api/v1/omc/devices/:id/upload-file"},
	"api.omc.devices.upgrade":              {Method: "POST", Path: "/api/v1/omc/devices/:id/upgrade"},

	"api.omc.devices.perf-disable": {Method: "POST", Path: "/api/v1/omc/devices/:id/perf-disable"},
	"api.omc.devices.perf-enable":  {Method: "POST", Path: "/api/v1/omc/devices/:id/perf-enable"},

	"api.omc.deleted-devices.list":    {Method: "GET", Path: "/api/v1/omc/deleted-devices"},
	"api.omc.deleted-devices.delete":  {Method: "DELETE", Path: "/api/v1/omc/deleted-devices/:id"},
	"api.omc.deleted-devices.recover": {Method: "POST", Path: "/api/v1/omc/deleted-devices/:id/recover"},

	"api.omc.device-alarms.list":       {Method: "GET", Path: "/api/v1/omc/device-alarms"},
	"api.omc.device-alarms.get":        {Method: "GET", Path: "/api/v1/omc/device-alarms/:ts"},
	"api.omc.device-events.list":       {Method: "GET", Path: "/api/v1/omc/device-events"},
	"api.omc.device-events.get":        {Method: "GET", Path: "/api/v1/omc/device-events/:ts"},
	"api.omc.device-method-calls.list": {Method: "GET", Path: "/api/v1/omc/device-method-calls"},
	"api.omc.device-method-calls.get":  {Method: "GET", Path: "/api/v1/omc/device-method-calls/:ts"},

	"api.omc.transfer-logs.list":   {Method: "GET", Path: "/api/v1/omc/transfer-logs"},
	"api.omc.transfer-logs.delete": {Method: "DELETE", Path: "/api/v1/omc/transfer-logs/:ts"},
	"api.omc.transfer-logs.get":    {Method: "GET", Path: "/api/v1/omc/transfer-logs/:ts"},

	"api.omc.firmwares.create":       {Method: "POST", Path: "/api/v1/omc/firmwares"},
	"api.omc.firmwares.list":         {Method: "GET", Path: "/api/v1/omc/firmwares"},
	"api.omc.firmwares.get":          {Method: "GET", Path: "/api/v1/omc/firmwares/:id"},
	"api.omc.firmwares.delete":       {Method: "DELETE", Path: "/api/v1/omc/firmwares/:id"},
	"api.omc.firmwares.products.set": {Method: "PUT", Path: "/api/v1/omc/firmwares/:id/products"},

	"api.omc.task-device-logs.list": {Method: "GET", Path: "/api/v1/omc/task-device-logs"},
	"api.omc.task-device-logs.get":  {Method: "GET", Path: "/api/v1/omc/task-device-logs/:ts"},
	"api.omc.tasks.create":          {Method: "POST", Path: "/api/v1/omc/tasks"},
	"api.omc.tasks.list":            {Method: "GET", Path: "/api/v1/omc/tasks"},
	"api.omc.tasks.get":             {Method: "GET", Path: "/api/v1/omc/tasks/:id"},

	"api.s3.objects.list":     {Method: "GET", Path: "/api/v1/s3/objects"},
	"api.s3.objects.get":      {Method: "GET", Path: "/api/v1/s3/objects/:bucket/:key"},
	"api.s3.objects.create":   {Method: "POST", Path: "/api/v1/s3/objects/:bucket/:key"},
	"api.s3.objects.update":   {Method: "PUT", Path: "/api/v1/s3/objects/:bucket/:key"},
	"api.s3.objects.delete":   {Method: "DELETE", Path: "/api/v1/s3/objects/:bucket/:key"},
	"api.s3.objects.get-info": {Method: "GET", Path: "/api/v1/s3/objects/:bucket/:key/info"},
}

func NewIAMServer() *iam.IAMServer {
	env := utils.GetEnv()
	secret := utils.GetEnv().GetString("secret")
	driverName := env.GetString("db_driver")
	dataSourceName := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=%v&search_path=%v",
		env.GetString("db_driver"),
		env.GetString("db_username"),
		env.GetString("db_password"),
		env.GetString("db_host"),
		env.GetString("db_port"),
		env.GetString("db_db"),
		env.GetString("db_sslmode"),
		env.GetString("db_domain"),
	)

	duration := time.Duration(time.Second * env.GetDuration("data_retention_period"))

	db := store.GetDB()
	schema := ""

	s := iam.NewIAMServer(db, duration, driverName, dataSourceName, secret, apiRulesMap)
	SetupAdmin(s, schema)
	SetupViewer(s, schema)
	return s
}

func SetupAdmin(s *iam.IAMServer, schema string) error {
	env := utils.GetEnv()
	return s.SetupAdmin(schema, env.GetString("admin_password"))
}

func SetupViewer(s *iam.IAMServer, schema string) error {
	role, err := s.CreateDefaultRole(schema, "viewer", "viewer", []string{
		"api.omc.*.get",
		"api.omc.*.list",
		"api.omc.data.query",
	})
	if err != nil {
		return errors.Wrap(err, "create viewer role")
	}
	user, err := s.CreateDefaultUser(schema, "guest", "guest", "guest2023")
	if err != nil {
		return errors.Wrap(err, "create guest user")
	}
	if _, err := s.AddRoleForUser(user.Schema, user.Name, role.Name); err != nil {
		return err
	}
	return nil
}
