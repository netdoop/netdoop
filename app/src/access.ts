export default (initialState: {
  name: string;
  avatar: string;
  loginUser: API.User | undefined;
}) => {
  const { loginUser } = initialState;
  const isLogin = loginUser !== undefined;
  const isAdmin = isLogin && loginUser?.Name === "admin";
  const rules = loginUser?.Rules || [];
  const isHidden = false;

  console.log("access", isLogin, loginUser)
  const apiAccess = {
    canListIAMRoles: rules.includes("api.iam.roles.list"),
    canGetIAMRole: rules.includes("api.iam.roles.get"),
    canCreateIAMRole: rules.includes("api.iam.roles.create"),
    canDeleteIAMRole: rules.includes("api.iam.roles.delete"),
    canUpdateIAMRoles: rules.includes("api.iam.roles.update"),
    canDisableIAMRole: rules.includes("api.iam.roles.disable"),
    canEnableIAMRole: rules.includes("api.iam.roles.enable"),
    canSetIAMRoleRules: rules.includes("api.iam.roles.set-rules"),

    canListIAMUsers: rules.includes("api.iam.users.list"),
    canGetIAMUser: rules.includes("api.iam.users.get"),
    canCreateIAMUser: rules.includes("api.iam.users.create"),
    canDeleteIAMUser: rules.includes("api.iam.users.delete"),
    canUpdateIAMUser: rules.includes("api.iam.users.update"),
    canDisableIAMUser: rules.includes("api.iam.users.disable"),
    canEnableIAMUser: rules.includes("api.iam.users.enable"),
    canResetIAMUserPassword: rules.includes("api.iam.users.reset-password"),
    canChangeIAMUserPassword: rules.includes("api.iam.users.change-password"),
    canListIAMUserRoles: rules.includes("api.iam.users.roles.list"),
    canCreateIAMUserRole: rules.includes("api.iam.users.roles.create"),
    canUpdateIAMUserRole: rules.includes("api.iam.users.roles.update"),

    canListAuditLogs: rules.includes("api.iam.audit-logs.list"),
    canGetAuditLog: rules.includes("api.iam.audit-logs.get"),

    canListS3Objects: rules.includes("api.s3.objects.list"),
    canGetS3Object: rules.includes("api.s3.objects.get"),
    canCreateS3Object: rules.includes("api.s3.objects.create"),
    canDeleteS3Objects: rules.includes("api.s3.objects.delete"),
    canUpdateS3Object: rules.includes("api.s3.objects.update"),
    canGetS3ObjectInfo: rules.includes("api.s3.objects.get-info"),

    canListOMCKPIMeasures: rules.includes("api.omc.kpi.measures.list"),
    canGetOMCKPIMeasure: rules.includes("api.omc.kpi.measures.get"),
    canCreateOMCKPIMeasure: rules.includes("api.omc.kpi.measures.create"),
    canDeleteOMCKPIMeasure: rules.includes("api.omc.kpi.measures.delete"),
    canUpdateOMCKPIMeasure: rules.includes("api.omc.kpi.measures.update"),
    canDisableOMCKPIMeasure: rules.includes("api.omc.kpi.measures.disable"),
    canEnableOMCKPIMeasure: rules.includes("api.omc.kpi.measures.enable"),

    canListOMCKPITemplates: rules.includes("api.omc.kpi.templates.list"),
    canGetOMCKPITemplate: rules.includes("api.omc.kpi.templates.get"),
    canCreateOMCKPITemplate: rules.includes("api.omc.kpi.templates.create"),
    canDeleteOMCKPITemplate: rules.includes("api.omc.kpi.templates.delete"),
    canUpdateOMCKPITemplate: rules.includes("api.omc.kpi.templates.update"),
    canListOMCKPITemplateRecords: rules.includes("api.omc.kpi.templates.list-records"),

    canListOMCDataModels: rules.includes("api.omc.datamodels.list"),
    canGetOMCDataModel: rules.includes("api.omc.datamodels.get"),
    canCreateOMCDataModel: rules.includes("api.omc.datamodels.create"),
    canDeleteOMCDataModel: rules.includes("api.omc.datamodels.delete"),

    canListOMCDataModelParameters: rules.includes("api.omc.datamodels.parameters.list"),
    canGetOMCDataModelParameter: rules.includes("api.omc.datamodels.parameters.get"),
    canCreateOMCDataModelParameter: rules.includes("api.omc.datamodels.parameters.create"),
    canDeleteOMCDataModelParameter: rules.includes("api.omc.datamodels.parameters.delete"),

    canListOMCDataModelTemplates: rules.includes("api.omc.datamodels.templates.list"),
    canGetOMCDataModelTemplate: rules.includes("api.omc.datamodels.templates.get"),
    canCreateOMCDataModelTemplate: rules.includes("api.omc.datamodels.templates.create"),
    canDeleteOMCDataModelTemplate: rules.includes("api.omc.datamodels.templates.delete"),

    canListOMCDeletedProducts: rules.includes("api.omc.deleted-products.list"),
    canDeleteOMCDeletedProducts: rules.includes("api.omc.deleted-products.delete"),

    canListOMCProducts: rules.includes("api.omc.products.list"),
    canGetOMCProduct: rules.includes("api.omc.products.get"),
    canCreateOMCProduct: rules.includes("api.omc.products.create"),
    canDeleteOMCProduct: rules.includes("api.omc.products.delete"),
    canUpdateOMCProduct: rules.includes("api.omc.products.update"),
    canDisableOMCProduct: rules.includes("api.omc.products.disable"),
    canEnableOMCProduct: rules.includes("api.omc.products.enable"),

    canListOMCGroups: rules.includes("api.omc.groups.list"),
    canGetOMCGroup: rules.includes("api.omc.groups.get"),
    canCreateOMCGroup: rules.includes("api.omc.groups.create"),
    canDeleteOMCGroup: rules.includes("api.omc.groups.delete"),
    canUpdateOMCGroup: rules.includes("api.omc.groups.update"),
    canListOMCGroupChildren: rules.includes("api.omc.groups.children.list"),
    canUpdateOMCGroupParent: rules.includes("api.omc.groups.parent.update"),

    canListOMCDeletedDevice: rules.includes("api.omc.deleted-devices.list"),
    canDeleteOMCDeletedDevice: rules.includes("api.omc.deleted-devices.delete"),
    canRecoverOMCDeletedDevice: rules.includes("api.omc.deleted-devices.recover"),

    canListOMCDevices: rules.includes("api.omc.devices.list"),
    canGetOMCDevice: rules.includes("api.omc.devices.get"),
    canCreateOMCDevice: rules.includes("api.omc.devices.create"),
    canDeleteOMCDevice: rules.includes("api.omc.devices.delete"),
    canUpdateOMCDevice: rules.includes("api.omc.devices.update"),
    canSetOMCDeviceParameterValues: rules.includes("api.omc.devices.set-parameter-values"),
    canGetOMCDeviceParameterValues: rules.includes("api.omc.devices.get-parameter-values"),
    canGetOMCDeviceParameters: rules.includes("api.omc.devices.get-parameters"),
    canGetOMCDeviceParameterNames: rules.includes("api.omc.devices.get-parameter-names"),
    canSetOMCDeviceGroup: rules.includes("api.omc.devices.set-group"),
    canRebootOMCDevice: rules.includes("api.omc.devices.reboot"),
    canDisableOMCDevice: rules.includes("api.omc.devices.disable"),
    canEnableOMCDevice: rules.includes("api.omc.devices.enable"),
    canGetOMCDeviceMethods: rules.includes("api.omc.devices.get-methods"),
    canAddObjectToOMCDevice: rules.includes("api.omc.devices.add-object"),
    canDeleteOMCDeviceObject: rules.includes("api.omc.devices.delete-object"),
    canUploadFileToOMCDevice: rules.includes("api.omc.devices.upload-file"),
    canUpgradeOMCDevice: rules.includes("api.omc.devices.upgrade"),

    canSetOMCDevicePerfEnable: rules.includes("api.omc.devices.perf-enable"),
    canSetOMCDevicePerfDisable: rules.includes("api.omc.devices.perf-disable"),

    canQueryOMCData: rules.includes("api.omc.data.query"),

    canListOMCDeviceMethodCalls: rules.includes("api.omc.device-method-calls.list"),
    canGetOMCDeviceMethodCall: rules.includes("api.omc.device-method-calls.get"),

    canListOMCDeviceEvents: rules.includes("api.omc.device-events.list"),
    canGetOMCDeviceEvent: rules.includes("api.omc.device-events.get"),

    canListOMCDeviceAlarms: rules.includes("api.omc.device-alarms.list"),
    canGetOMCDeviceAlarm: rules.includes("api.omc.device-alarms.get"),

    canListOMCTransferLogs: rules.includes("api.omc.transfer-logs.list"),
    canGetOMCTransferLog: rules.includes("api.omc.transfer-logs.get"),
    canDeleteOMCTransferLog: rules.includes("api.omc.transfer-logs.delete"),

    canListOMCFirmwares: rules.includes("api.omc.firmwares.list"),
    canGetOMCFirmware: rules.includes("api.omc.firmwares.get"),
    canCreateOMCFirmware: rules.includes("api.omc.firmwares.create"),
    canDeleteOMCFirmware: rules.includes("api.omc.firmwares.delete"),

    canListOMCTasks: rules.includes("api.omc.tasks.list"),
    canGetOMCTask: rules.includes("api.omc.tasks.get"),
    canCreateOMCTask: rules.includes("api.omc.tasks.create"),
    canListOMCTaskDeviceLogs: rules.includes("api.omc.task-device-logs.list"),
    canGetOMCTaskDeviceLog: rules.includes("api.omc.task-device-logs.get"),
  };

  const pageAccess = {
    canViewIAMPage: isAdmin || (apiAccess.canListIAMRoles || apiAccess.canListIAMUsers),
    canViewIAMRolesPage: isAdmin || (apiAccess.canListIAMRoles),
    canViewIAMUsersPage: isAdmin || (apiAccess.canListIAMUsers),
    canViewAuditPage: isAdmin || (apiAccess.canListAuditLogs),

    canViewOMCProductsPage: isAdmin || (apiAccess.canListOMCProducts && (apiAccess.canCreateOMCProduct || apiAccess.canUpdateOMCProduct || apiAccess.canDeleteOMCProduct)),
    canViewOMCGroupsPage: isAdmin || (apiAccess.canListOMCGroups && (apiAccess.canCreateOMCGroup || apiAccess.canUpdateOMCGroup || apiAccess.canDeleteOMCGroup || apiAccess.canUpdateOMCGroupParent)),

    canViewOMCCpeMonitorPage: isAdmin || (apiAccess.canListOMCDevices),
    canViewOMCCpeMonitorInformationPage: isAdmin || (apiAccess.canGetOMCDevice),
    canViewOMCCpeMonitorStatisticsPage: isAdmin || (apiAccess.canGetOMCDevice && apiAccess.canQueryOMCData),
    canViewOMCCpeMonitorSettingPage: isAdmin || (apiAccess.canGetOMCDevice && apiAccess.canSetOMCDeviceParameterValues),

    canViewOMCCpeMethodCallsPage: isAdmin || (apiAccess.canListOMCDeviceMethodCalls),
    canViewOMCCpeMethodCallDetailPage: isAdmin || (apiAccess.canGetOMCDeviceMethodCall),

    canViewOMCCpeEventsPage: isAdmin || (apiAccess.canListOMCDeviceEvents),
    canViewOMCCpeEventDetailPage: isAdmin || (apiAccess.canGetOMCDeviceEvent),
    canViewOMCCpeUploadFilesPage: isAdmin || (apiAccess.canListOMCTransferLogs),

    canViewOMCCpeFirmwaresPage: isAdmin || (apiAccess.canListOMCFirmwares),

    canViewOMCCpeInventoryPage: isAdmin || (apiAccess.canListOMCDevices),
    canViewOMCCpeInventoryInformationPage: isAdmin || (apiAccess.canGetOMCDevice),

    canViewOMCPerfPage: isAdmin || (apiAccess.canListOMCKPIMeasures && (apiAccess.canCreateOMCKPIMeasure || apiAccess.canUpdateOMCKPIMeasure|| apiAccess.canDeleteOMCKPIMeasure)),

    canViewOMCAlarmsPage: isAdmin || (apiAccess.canListOMCDeviceAlarms),
    canViewOMCAlarmDetailPage: isAdmin || (apiAccess.canGetOMCDeviceAlarm),
  };

  const subPageAccess = {
    canViewOMCDashboardPage: true,
    canViewOMCCpeLogsPage: isAdmin || (pageAccess.canViewOMCCpeEventsPage || pageAccess.canViewOMCCpeMethodCallsPage),
    canViewOMCCpeUpgradePage: isAdmin || (pageAccess.canViewOMCCpeFirmwaresPage),
  };

  const topPageAccess = {
    canViewOMCAdminPages: isAdmin || (pageAccess.canViewOMCProductsPage),
    canViewOMCCpePages: isAdmin || (pageAccess.canViewOMCCpeMonitorPage || subPageAccess.canViewOMCCpeLogsPage || subPageAccess.canViewOMCCpeUpgradePage),
    canViewSystemPages: isAdmin || (pageAccess.canViewIAMRolesPage || pageAccess.canViewIAMUsersPage || pageAccess.canViewAuditPage),

  };
  //   console.log(loginUser)
  // console.log(topPageAccess)
  // console.log(subPageAccess)
  // console.log(pageAccess)
  return {
    isLogin,
    isAdmin,
    isHidden,

    ...topPageAccess,
    ...subPageAccess,

    ...pageAccess,
    ...apiAccess,
  };
};