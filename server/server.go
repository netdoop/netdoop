package server

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/netdoop/cwmp/acs"
	_ "github.com/netdoop/netdoop/docs"
	"github.com/netdoop/netdoop/models/omc"
	"github.com/netdoop/netdoop/server/api/develop"
	omcapi "github.com/netdoop/netdoop/server/api/omc"
	sysapi "github.com/netdoop/netdoop/server/api/system"
	"github.com/netdoop/netdoop/server/global"
	"github.com/netdoop/netdoop/store"
	"github.com/netdoop/netdoop/utils"

	"github.com/heypkg/iam"
	"github.com/heypkg/s3"
	"github.com/heypkg/store/echohandler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

// dev.netdoop.com

// @title NetDoop API
// @version 1.0
// @host dev.netdoop.com
// @BasePath /api/v1
// @schemes http
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
type Server struct {
	logger *zap.Logger
	env    *viper.Viper
	e      *echo.Echo

	iamsrv  *iam.IAMServer
	tasksrv *omc.TaskServer
	acssrv  *acs.AcsServer

	shutdown     bool
	shutdownCh   chan struct{}
	shutdownLock sync.Mutex
}

func DisableVaryHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			return err
		}
		c.Response().Header().Del(echo.HeaderVary)
		return nil
	}
}

func NewServer() *Server {
	s := &Server{
		shutdownCh: make(chan struct{}),
	}
	s.logger = utils.GetLogger().Named("server")

	s.env = utils.GetEnv()
	db := store.GetDB()
	omc.Setup()

	s.iamsrv = NewIAMServer()
	if s.iamsrv == nil {
		return nil
	}
	s.tasksrv = omc.GetTaskServer()
	if s.tasksrv == nil {
		return nil
	}

	s.e = echo.New()
	s.e.HideBanner = true
	s.e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogError:  true,
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			if v.Status >= 200 && v.Status < 300 {
				if utils.VerboseMode {
					s.logger.Debug("request",
						zap.String("method", v.Method),
						zap.String("URI", v.URI),
						zap.Int("status", v.Status),
					)
				}
			} else {
				s.logger.Error("request",
					zap.String("method", v.Method),
					zap.String("URI", v.URI),
					zap.Int("status", v.Status),
					zap.Error(v.Error),
				)
			}
			return nil
		},
	}))
	if s.env.GetInt("server_cors_enable") == 1 {
		s.e.Use(middleware.CORS())
	}

	s.e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	s.e.Static("/", s.env.GetString("server_public_path"))
	s.e.GET("/swagger/*", echoSwagger.WrapHandler)

	developGroup := s.e.Group("/dev")
	developGroup.POST("/post", develop.HandlePost)

	sessionStore := store.GetRedisSessionStore()
	acsGroup := s.e.Group("/acs")
	s.acssrv.SetupPostEchoGroupWithOptions(acsGroup, sessionStore, acs.Options{DumpBody: true})
	uploadGroup := s.e.Group("/upload")
	s.acssrv.SetupUploadEchoGroup(uploadGroup)

	handlers := []echo.MiddlewareFunc{
		s.iamsrv.MakeJwtHandler(),
		s.iamsrv.MakeLoginHandler(),
	}
	if s.env.GetInt("iam_audit_log_enable") != 0 {
		handlers = append(handlers, s.iamsrv.MakeAuditLogHandler())
	}

	apiGroup := s.e.Group("/api/v1")
	apiGroup.POST("/auth", s.iamsrv.HandleAuthenticate, s.iamsrv.MakeAuditLogHandler())
	apiGroup.GET("/routes", s.handleListApiRoutes)

	sysGroup := apiGroup.Group("/system")
	sysGroup.GET("/info", sysapi.HandleGetSystemInfo)
	sysGroup.GET("/time", sysapi.HandleGetSystemTime)

	iamGroup := apiGroup.Group("/iam")
	iamGroup.Use(handlers...)
	s.iamsrv.SetupEchoGroup(iamGroup)

	s3Group := apiGroup.Group("/s3")
	s3Group.Use(handlers...)
	s3Group.GET("/objects", s3.HandleListObjects)
	s3Group.POST("/objects/:bucket/:key", s3.HandlePutObject)
	s3Group.PUT("/objects/:bucket/:key", s3.HandlePutObject)
	s3Group.GET("/objects/:bucket/:key/info", s3.HandleGetObjectInfo, s3.S3ObjectHandler)
	s3Group.GET("/objects/:bucket/:key", s3.HandleGetObject, s3.S3ObjectHandler)
	s3Group.DELETE("/objects/:bucket/:key", s3.HandleDeleteObject, s3.S3ObjectHandler)

	omcGroup := apiGroup.Group("/omc")
	omcGroup.Use(handlers...)
	omcGroup.POST("/data", omcapi.HandleQueryData)

	omcGroup.GET("/datamodels", omcapi.HandleListDataModel)
	omcGroup.POST("/datamodels", omcapi.HandleCreateDataModel)
	omcGroup.DELETE("/datamodels/:id", omcapi.HandleDeleteDataModel, echohandler.ObjectHandler[omc.DataModel](db))
	omcGroup.GET("/datamodels/:id", omcapi.HandleGetDataModel, echohandler.ObjectHandler[omc.DataModel](db))

	omcGroup.GET("/datamodels/:id/parameters", omcapi.HandleListDataModelParameters)
	omcGroup.POST("/datamodels/:id/parameters", omcapi.HandleCreateDataModelParameter)
	omcGroup.DELETE("/datamodels//:id/parameters/:parameter_id", omcapi.HandleDeleteDataModelParameter, echohandler.ObjectHandler[omc.DataModel](db))
	omcGroup.GET("/datamodels/:id/parameters/:parameter_id", omcapi.HandleGetDataModelParameter, echohandler.ObjectHandler[omc.DataModel](db))

	omcGroup.GET("/datamodels/:id/templates", omcapi.HandleListDataModelTemplates)
	omcGroup.POST("/datamodels/:id/templates", omcapi.HandleCreateDataModelTemplate)
	omcGroup.DELETE("/datamodels//:id/templates/:template_id", omcapi.HandleDeleteDataModelTemplate, echohandler.ObjectHandler[omc.DataModel](db))
	omcGroup.GET("/datamodels/:id/templates/:template_id", omcapi.HandleGetDataModelTemplate, echohandler.ObjectHandler[omc.DataModel](db))

	omcGroup.GET("/products", omcapi.HandleListProducts)
	omcGroup.POST("/products", omcapi.HandleCreateProduct)
	omcGroup.DELETE("/products/:id", omcapi.HandleDeleteProduct, echohandler.ObjectHandler[omc.Product](db))
	omcGroup.GET("/products/:id", omcapi.HandleGetProduct, echohandler.ObjectHandler[omc.Product](db))

	omcGroup.GET("/products", omcapi.HandleListProducts)
	omcGroup.POST("/products", omcapi.HandleCreateProduct)
	omcGroup.DELETE("/products/:id", omcapi.HandleDeleteProduct, echohandler.ObjectHandler[omc.Product](db))
	omcGroup.GET("/products/:id", omcapi.HandleGetProduct, echohandler.ObjectHandler[omc.Product](db))
	omcGroup.PUT("/products/:id", omcapi.HandleUpdateProductInfo, echohandler.ObjectHandler[omc.Product](db))
	omcGroup.PUT("/products/:id/enable", omcapi.HandleSetProductEnable, echohandler.ObjectHandler[omc.Product](db))
	omcGroup.PUT("/products/:id/disable", omcapi.HandleSetProductDisable, echohandler.ObjectHandler[omc.Product](db))
	omcGroup.GET("/products/:id/firmwares", omcapi.HandleListProductFirmwares, echohandler.ObjectHandler[omc.Product](db))

	omcGroup.GET("/deleted-products", omcapi.HandleListDeletedProducts)
	omcGroup.DELETE("/deleted-products/:id", omcapi.HandleDeleteDeletedProduct, echohandler.DeletedObjectHandler[omc.Product](db))

	omcGroup.GET("/kpi/measures", omcapi.HandleListKPIMeasures)
	omcGroup.POST("/kpi/measures", omcapi.HandleCreateKPIMeasure)
	omcGroup.DELETE("/kpi/measures/:id", omcapi.HandleDeleteKPIMeasure, echohandler.ObjectHandler[omc.KPIMeas](db))
	omcGroup.GET("/kpi/measures/:id", omcapi.HandleGetKPIMeasure, echohandler.ObjectHandler[omc.KPIMeas](db))
	omcGroup.PUT("/kpi/measures/:id", omcapi.HandleUpdateKPIMeasureInfo, echohandler.ObjectHandler[omc.KPIMeas](db))
	omcGroup.PUT("/kpi/measures/:id/enable", omcapi.HandleSetKPIMeasureEnable, echohandler.ObjectHandler[omc.KPIMeas](db))
	omcGroup.PUT("/kpi/measures/:id/disable", omcapi.HandleSetKPIMeasureDisable, echohandler.ObjectHandler[omc.KPIMeas](db))

	omcGroup.GET("/kpi/templates", omcapi.HandleListKPITemplates)
	omcGroup.POST("/kpi/templates", omcapi.HandleCreateKPITemplate)
	omcGroup.DELETE("/kpi/templates/:id", omcapi.HandleDeleteKPITemplate, echohandler.ObjectHandler[omc.KPITemplate](db))
	omcGroup.GET("/kpi/templates/:id", omcapi.HandleGetKPITemplate, echohandler.ObjectHandler[omc.KPITemplate](db))
	omcGroup.PUT("/kpi/templates/:id", omcapi.HandleUpdateKPITemplateInfo, echohandler.ObjectHandler[omc.KPITemplate](db))
	omcGroup.GET("/kpi/templates/:id/records", omcapi.HandleListKPITemplateRecords, echohandler.ObjectHandler[omc.KPITemplate](db))

	omcGroup.GET("/devices", omcapi.HandleListDevices)
	omcGroup.POST("/devices", omcapi.HandleCreateDevice)
	omcGroup.DELETE("/devices/:id", omcapi.HandleDeleteDevice, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.GET("/devices/:id", omcapi.HandleGetDevice, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.PUT("/devices/:id", omcapi.HandleUpdateDeviceInfo, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.PUT("/devices/:id/group", omcapi.HandleSetGroupForDevice, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.PUT("/devices/:id/enable", omcapi.HandlesSetDeviceEnable, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.PUT("/devices/:id/disable", omcapi.HandlesSetDeviceDisable, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.GET("/devices/:id/methods", omcapi.HandleGetDeviceMethods, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.GET("/devices/:id/parameters", omcapi.HandleGetDeviceParameters, echohandler.ObjectHandler[omc.Device](db))

	omcGroup.POST("/devices/:id/reboot", omcapi.HandleRebootDevice, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.POST("/devices/:id/get-parameter-names", omcapi.HandleGetDeviceParameterNames, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.POST("/devices/:id/set-parameter-values", omcapi.HandleSetDeviceParameterValues, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.POST("/devices/:id/get-parameter-values", omcapi.HandleGetDeviceParameterValues, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.POST("/devices/:id/add-object", omcapi.HandleAddDeviceObject, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.POST("/devices/:id/delete-object", omcapi.HandleDeleteDeviceObject, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.POST("/devices/:id/upload-file", omcapi.HandleUploadDeviceFile, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.POST("/devices/:id/upgrade", omcapi.HandleUpgradeDevice, echohandler.ObjectHandler[omc.Device](db))

	omcGroup.POST("/devices/:id/perf-disable", omcapi.HandleSetDevicePerfDisable, echohandler.ObjectHandler[omc.Device](db))
	omcGroup.POST("/devices/:id/perf-enable", omcapi.HandleSetDevicePerfEnable, echohandler.ObjectHandler[omc.Device](db))

	omcGroup.GET("/deleted-devices", omcapi.HandleListDeletedDevices)
	omcGroup.DELETE("/deleted-devices/:id", omcapi.HandleDeleteDeletedDevice, echohandler.DeletedObjectHandler[omc.Device](db))
	omcGroup.POST("/deleted-devices/:id/recover", omcapi.HandleRecoverDeletedDevice, echohandler.DeletedObjectHandler[omc.Device](db))

	omcGroup.GET("/groups", omcapi.HandleListGroups)
	omcGroup.POST("/groups", omcapi.HandleCreateGroup)
	omcGroup.DELETE("/groups/:id", omcapi.HandleDeleteGroup, echohandler.ObjectHandler[omc.Group](db))
	omcGroup.GET("/groups/:id", omcapi.HandleGetGroup, echohandler.ObjectHandler[omc.Group](db))
	omcGroup.GET("/groups/:id/children", omcapi.HandleGetGroupChild, echohandler.ObjectHandler[omc.Group](db))
	omcGroup.PUT("/groups/:id", omcapi.HandleUpdateGroupInfo, echohandler.ObjectHandler[omc.Group](db))
	omcGroup.PUT("/groups/:id/parent", omcapi.HandleSetGroupParent, echohandler.ObjectHandler[omc.Group](db))

	omcGroup.GET("/device-method-calls", omcapi.HandleListDeviceMethodCalls)
	omcGroup.GET("/device-method-calls/:ts", omcapi.HandleGetDeviceMethodCall, echohandler.TSObjectHandler[omc.DeviceMethodCall](db))

	omcGroup.GET("/device-events", omcapi.HandleListDeviceEvents)
	omcGroup.GET("/device-events/:ts", omcapi.HandleGetDeviceEvent, echohandler.TSObjectHandler[omc.DeviceEvent](db))

	omcGroup.GET("/device-alarms", omcapi.HandleListDeviceAlarms)
	omcGroup.GET("/device-alarms/:ts", omcapi.HandleGetDeviceAlarm, echohandler.TSObjectHandler[omc.DeviceAlarm](db))

	omcGroup.GET("/firmwares", omcapi.HandleListFirmwares)
	omcGroup.POST("/firmwares", omcapi.HandleCreateFirmware)
	omcGroup.GET("/firmwares/:id", omcapi.HandleGetFirmware, echohandler.ObjectHandler[omc.Firmware](db))
	omcGroup.DELETE("/firmwares/:id", omcapi.HandleDeleteFirmware, echohandler.ObjectHandler[omc.Firmware](db))
	omcGroup.PUT("/firmwares/:id/products", omcapi.HandleSetFirmwareProducts, echohandler.ObjectHandler[omc.Firmware](db))

	omcGroup.GET("/transfer-logs", omcapi.HandleListDeviceTransferLogs)
	omcGroup.GET("/transfer-logs/:ts", omcapi.HandleGetDeviceTransferLog, echohandler.ObjectHandler[omc.DeviceTransferLog](db))
	omcGroup.DELETE("/transfer-logs/:ts", omcapi.HandleDeleteDeviceTransferLog, echohandler.ObjectHandler[omc.DeviceTransferLog](db))

	omcGroup.GET("/tasks", omcapi.HandleListTasks)
	omcGroup.POST("/tasks", omcapi.HandleCreateTask)
	omcGroup.GET("/tasks/:id", omcapi.HandleGetTask, echohandler.ObjectHandler[omc.Task](db))

	omcGroup.GET("/task-device-logs", omcapi.HandleListTaskDeviceLogs)
	omcGroup.GET("/task-device-logs/:ts", omcapi.HandleGetTaskDeviceLog, echohandler.TSObjectHandler[omc.TaskDeviceLog](db))

	return s
}

func (s *Server) Close() {
	s.shutdownLock.Lock()
	defer s.shutdownLock.Unlock()

	if s.shutdown {
		return
	}
	s.logger.Debug("server close")
	s.shutdown = true
	close(s.shutdownCh)
}

func (s *Server) Run() {
	s.logger.Debug("server running")
	defer s.logger.Debug("server stopped")

	var wg sync.WaitGroup
	defer wg.Wait()
	colorer := utils.GetColorer()
	colorer.Printf(global.Banner, colorer.Red("v"+global.Version+" ("+global.Build+")"), colorer.Blue(global.Website))

	time.Sleep(2 * time.Second)
	s.Setup()
	omc.ReloadAllDataModels()
	omc.ReloadAllProducts()
	omc.ReloadAllKPIMeansure()

	omc.StartAllActiveDeviceKeepaliveTimers()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.tasksrv.Run()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// s.stund.Run()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		addr := s.env.GetString("server_address")
		if err := s.e.Start(addr); err != nil {
			if strings.Contains(err.Error(), "Server closed") {
				return
			}
			s.logger.Fatal("start server", zap.Error(err))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		interval := s.env.GetDuration("data_retention_period")*time.Second + time.Hour*48
		for {
			if s.shutdown {
				break
			}
			time.Sleep(time.Second)
			now := time.Now()
			if now.Unix()%1800 == 0 {
				if err := s3.DropObjectsBeforeInterval(interval); err != nil {
					s.logger.Error("drop objects before interval", zap.Error(err))
				}
			}
			if now.Unix()%3600 == 0 {
				if err := s3.DropInvalidObjectFiles("", omc.DeviceUploadBucket); err != nil {
					s.logger.Error("drop objects before interval", zap.Error(err))
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if s.shutdown {
				break
			}
			time.Sleep(time.Second)
			now := time.Now()
			if now.Unix()%30 == 0 {
				if err := omc.FetchAllDeviceStatus(now); err != nil {
					s.logger.Error("fetch device status", zap.Error(err))
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cleared := true
		alarmIdentifier := ""

		db := store.GetDB()
		for {
			if s.shutdown {
				break
			}
			time.Sleep(time.Minute * 2)
			cleared = !cleared
			if !cleared || alarmIdentifier == "" {
				alarmIdentifier = fmt.Sprintf("%v", time.Now().Unix())
			}
			omc.DebugPostAlarm(db, cleared, alarmIdentifier)
		}
	}()

	<-s.shutdownCh

	s.tasksrv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.e.Shutdown(ctx); err != nil {
		s.logger.Error("shutdown server", zap.Error(err))
	}
}

func (s *Server) Setup() {
	db := store.GetDB()
	schema := ""

	omc.SetupDefaultDataModel(db, schema)
	omc.SetupProducts(db, schema)
	omc.SetupTasks(db, schema)

	group := omc.Group{ID: 0, Name: "default"}
	db.Exec(`INSERT INTO "groups" ("id","schema","name") VALUES (0,'','root')`)
	db.Create(&omc.Group{Name: "test1", ParentID: &group.ID})
	db.Create(&omc.Group{Name: "test2", ParentID: &group.ID})
}

type listRoutesBody struct {
	Data  []iam.ApiRule `json:"Data"`
	Total int           `json:"Total"`
}

// handleListApiRoutes godoc
// @Summary List routes
// @ID list-routes
// @Tags System
// @Security Bearer
// @Success 200 {object} listRoutesBody
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /routes [get]
func (s *Server) handleListApiRoutes(c echo.Context) error {
	routes := []iam.ApiRule{}
	for _, route := range s.e.Routes() {
		if strings.HasPrefix(route.Name, "github.com/netdoop/netdoop/server/api/") {
			if strings.HasPrefix(route.Path, "/api/v1") {
				routes = append(routes, iam.ApiRule{Method: strings.ToUpper(route.Method), Path: route.Path})
			}
		}
	}
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Path < routes[j].Path
	})
	return c.JSON(http.StatusOK, listRoutesBody{Data: routes, Total: len(routes)})
}
