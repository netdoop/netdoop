declare namespace API {
  type addDeviceObejectBody = {
    ObjectName?: string;
    ParameterKey?: string;
  };

  type addDeviceObjectParams = {
    /** Device ID */
    id: number;
  };

  type addUserRolesBody = {
    roleNames?: string[];
  };

  type addUserRolesParams = {
    /** User ID */
    id: number;
  };

  type ApiRule = {
    Method?: string;
    Path?: string;
  };

  type AuditLog = {
    ApiName?: string;
    MetaData?: Tags;
    Method?: string;
    Path?: string;
    Schema?: string;
    Status?: number;
    Time?: string;
    Updated?: string;
    User?: User;
    UserAlias?: string;
    UserId?: number;
    UserName?: string;
  };

  type authBody = {
    password?: string;
    username?: string;
  };

  type authResponseBody = {
    token?: string;
  };

  type changePasswordBody = {
    NewPassword?: string;
    Password?: string;
  };

  type changeUserPasswordParams = {
    /** User ID */
    id: number;
  };

  type createDataModelBody = {
    metaData?: Tags;
    name?: string;
    parameterPath?: string;
    productType?: string;
  };

  type createDataModelParameterBody = {
    defaultValue?: string;
    description?: string;
    metaData?: Tags;
    name?: string;
    type?: string;
    writable?: boolean;
  };

  type createDataModelTemplateBody = {
    metaData?: Tags;
    name?: string;
    parameterNames?: Tags;
  };

  type createDeviceBody = {
    Name?: string;
    Oui?: string;
    ProductClass?: string;
    SerailNumber?: string;
  };

  type createGroupBody = {
    name?: string;
    parentID?: number;
  };

  type createKPIMeasureBody = {
    formula?: string;
    measTypeID?: string;
    measTypeSet?: string;
    name?: string;
    productType?: string;
    statsType?: string;
    unit?: string;
  };

  type createKPITemplateBody = {
    measTypeIds?: string[];
    name?: string;
    periodicInterval?: number;
    selectIds?: number[];
    selectType?: string;
  };

  type createProductBody = {
    manufacturer?: string;
    modelName?: string;
    oui?: string;
    productClass?: string;
  };

  type createRoleBody = {
    alias?: string;
    name?: string;
  };

  type createTaskBody = {
    Creater?: string;
    DeviceIds?: number[];
    ExecEndTime?: string;
    ExecInterval?: number;
    ExecMode?: TaskExecMode;
    ExecProcess?: number;
    ExecStartTime?: string;
    ExecTimes?: number;
    MetaData?: Tags;
    ParameterNames?: string[];
    TaskName?: string;
    TaskStatus?: TaskStatus;
    TaskType?: string;
  };

  type createUserBody = {
    alias?: string;
    name?: string;
    password?: string;
  };

  type DataModel = {
    Created?: string;
    Default?: boolean;
    Deleted?: DeletedAt;
    Id?: number;
    MetaData?: Tags;
    Name?: string;
    ParameterPath?: string;
    Parameters?: DataModelParameter[];
    ProductType?: string;
    Schema?: string;
    Templates?: DataModelTemplate[];
    Updated?: string;
  };

  type DataModelParameter = {
    Created?: string;
    DataModelId?: number;
    Default?: boolean;
    DefaultValue?: string;
    Deleted?: DeletedAt;
    Description?: string;
    Id?: number;
    MetaData?: Tags;
    Name?: string;
    Type?: string;
    Updated?: string;
    Writable?: boolean;
  };

  type DataModelTemplate = {
    Created?: string;
    DataModelId?: number;
    Default?: boolean;
    Deleted?: DeletedAt;
    Id?: number;
    MetaData?: Tags;
    Name?: string;
    ParameterNames?: Tags;
    Updated?: string;
  };

  type DeletedAt = {
    time?: string;
    /** Valid is true if Time is not NULL */
    valid?: boolean;
  };

  type deleteDatamodelParameterParams = {
    /** ID */
    id: number;
    /** Parameter ID */
    parameter_id: string;
  };

  type deleteDatamodelParams = {
    /** ID */
    id: number;
  };

  type deleteDatamodelTemplateParams = {
    /** ID */
    id: number;
    /** TemplateID */
    template_id: string;
  };

  type deleteDeletedDeviceParams = {
    /** Device ID */
    id: number;
  };

  type deleteDeletedProductParams = {
    /** Product ID */
    id: number;
  };

  type deleteDeviceObejectBody = {
    ObjectName?: string;
    ParameterKey?: string;
  };

  type deleteDeviceObjectParams = {
    /** Device ID */
    id: number;
  };

  type deleteDeviceParams = {
    /** Device ID */
    id: number;
  };

  type deleteDeviceTransferLogParams = {
    /** Timestamp */
    ts: number;
  };

  type deleteFirmwareParams = {
    /** Firmware ID */
    id: number;
  };

  type deleteGroupParams = {
    /** Group ID */
    id: number;
  };

  type deleteKpiMeasureParams = {
    /** ID */
    id: number;
  };

  type deleteKpiTemplateParams = {
    /** ID */
    id: number;
  };

  type deleteObjectParams = {
    /** S3 objet bucket */
    bucket: string;
    /** S3 object key */
    key: string;
  };

  type deleteProductParams = {
    /** Product ID */
    id: number;
  };

  type deleteRoleParams = {
    /** Role ID */
    id: number;
  };

  type deleteTaskParams = {
    /** Task ID */
    id: number;
  };

  type deleteUserParams = {
    /** User ID */
    id: number;
  };

  type Device = {
    ActiveStatus?: string;
    Created?: string;
    Deleted?: DeletedAt;
    Enable?: boolean;
    Group?: Group;
    GroupId?: number;
    Id?: number;
    LastInformTime?: string;
    MetaData?: Tags;
    Methods?: string[];
    Name?: string;
    Online?: boolean;
    Oui?: string;
    ParameterNotifications?: ParameterNotifications;
    ParameterValues?: ParameterValues;
    ParameterWritables?: ParameterWritables;
    Product?: Product;
    ProductClass?: string;
    ProductId?: number;
    ProductType?: string;
    Properties?: Tags;
    Schema?: string;
    SerialNumber?: string;
    Updated?: string;
  };

  type DeviceAlarm = {
    AdditionalInformation?: string;
    AdditionalText?: string;
    AlarmChangedTime?: string;
    AlarmCleared?: boolean;
    AlarmClearedTime?: string;
    AlarmConfirmed?: boolean;
    AlarmConfirmedTime?: string;
    AlarmIdentifier?: string;
    AlarmRaisedTime?: string;
    Device?: Device;
    DeviceId?: number;
    EventType?: string;
    ManagedObjectInstance?: string;
    Oui?: string;
    PerceivedSeverity?: string;
    ProbableCause?: string;
    ProductClass?: string;
    ProductType?: string;
    Schema?: string;
    SerialNumber?: string;
    SpecificProblem?: string;
    Time?: string;
    Updated?: string;
  };

  type DeviceEvent = {
    CurrentTime?: string;
    Device?: Device;
    DeviceId?: number;
    EventType?: string;
    MetaData?: Tags;
    Oui?: string;
    ProductClass?: string;
    ProductType?: string;
    Schema?: string;
    SerialNumber?: string;
    Time?: string;
    Updated?: string;
  };

  type DeviceMethodCall = {
    CommandKey?: string;
    Device?: Device;
    DeviceId?: number;
    FaultCode?: number;
    FaultString?: string;
    MethodName?: string;
    Oui?: string;
    ProductClass?: string;
    ProductType?: string;
    RequestValues?: Tags;
    ResponseValues?: Tags;
    Schema?: string;
    SerialNumber?: string;
    State?: DeviceMethodCallState;
    Time?: string;
    Updated?: string;
  };

  type DeviceMethodCallState = 0 | 1 | 2 | 3 | 100;

  type DeviceTransferLog = {
    CompleteTime?: string;
    Device?: Device;
    DeviceId?: number;
    FaultCode?: number;
    FaultString?: string;
    FileName?: string;
    FileType?: string;
    Firmware?: Firmware;
    FirmwareId?: number;
    ObjectBucket?: string;
    ObjectKey?: string;
    Oui?: string;
    ProductClass?: string;
    ProductType?: string;
    S3Object?: S3Object;
    S3ObjectId?: number;
    Schema?: string;
    SerialNumber?: string;
    StartTime?: string;
    Time?: string;
    TransferType?: string;
    Updated?: string;
  };

  type downloadObjectParams = {
    /** S3 objet bucket */
    bucket: string;
    /** S3 object key */
    key: string;
    /** Token for the object */
    token: string;
  };

  type Firmware = {
    Created?: string;
    Deleted?: DeletedAt;
    Id?: number;
    Name?: string;
    ProductType?: string;
    S3Object?: S3Object;
    S3ObjectId?: number;
    Schema?: string;
    Updated?: string;
    UploadTime?: string;
    Uploader?: string;
    Version?: string;
    products?: Product[];
  };

  type getAuditLogParams = {
    /** Timestamp */
    ts: number;
  };

  type getDatamodelParameterParams = {
    /** ID */
    id: number;
    /** Parameter ID */
    parameter_id: string;
  };

  type getDatamodelParams = {
    /** ID */
    id: number;
  };

  type getDatamodelTemplateParams = {
    /** ID */
    id: number;
    /** TemplateID */
    template_id: string;
  };

  type getDeviceAlarmParams = {
    /** Timestamp */
    ts: number;
  };

  type getDeviceEventParams = {
    /** Timestamp */
    ts: number;
  };

  type getDeviceLogParams = {
    /** Timestamp */
    ts: number;
  };

  type getDeviceMethodCallParams = {
    /** Timestamp */
    ts: number;
  };

  type getDeviceParameterNamesBody = {
    NextLevel?: boolean;
    ParameterPath?: string;
  };

  type getDeviceParameterNamesParams = {
    /** Device ID */
    id: number;
  };

  type getDeviceParameterValuesBody = {
    Names?: string[];
  };

  type getDeviceParameterValuesParams = {
    /** Device ID */
    id: number;
  };

  type getDeviceParams = {
    /** Device ID */
    id: number;
  };

  type getDeviceTransferLogParams = {
    /** Timestamp */
    ts: number;
  };

  type getFirmwareParams = {
    /** Firmware ID */
    id: number;
  };

  type getGroupChildrenParams = {
    /** Group ID */
    id: number;
  };

  type getGroupParams = {
    /** Group ID */
    id: number;
  };

  type getKpiMeasureParams = {
    /** ID */
    id: number;
  };

  type getKpiTemplateParams = {
    /** ID */
    id: number;
  };

  type getObjectInfoParams = {
    /** S3 objet bucket */
    bucket: string;
    /** S3 object key */
    key: string;
  };

  type getProductParams = {
    /** Product ID */
    id: number;
  };

  type getRoleParams = {
    /** Role ID */
    id: number;
  };

  type getTaskParams = {
    /** Task ID */
    id: number;
  };

  type getUserParams = {
    /** User ID */
    id: number;
  };

  type getUserRolesParams = {
    /** User ID */
    id: number;
  };

  type Group = {
    Children?: Group[];
    Created?: string;
    Id?: number;
    MetaData?: Tags;
    Name?: string;
    Parent?: Group;
    ParentId?: number;
    Schema?: string;
    Updated?: string;
  };

  type HTTPError = {
    message?: any;
  };

  type KPIMeas = {
    Created?: string;
    Default?: boolean;
    Deleted?: DeletedAt;
    Enable?: boolean;
    Formula?: string;
    Id?: number;
    MeasTypeID?: string;
    MeasTypeSet?: string;
    MetaData?: Tags;
    Name?: string;
    ProductType?: string;
    Schema?: string;
    StatsType?: string;
    Unit?: string;
    Updated?: string;
  };

  type KPITemplate = {
    Created?: string;
    Default?: boolean;
    Deleted?: DeletedAt;
    Id?: number;
    MeasTypeIds?: string[];
    MetaData?: Tags;
    Name?: string;
    PeriodicInterval?: number;
    ProductType?: string;
    Schema?: string;
    SelectIds?: number[];
    SelectType?: string;
    Updated?: string;
  };

  type listAuditLogsData = {
    Data?: AuditLog[];
    Total?: number;
  };

  type listAuditLogsParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listDataModelData = {
    Data?: DataModel[];
    Total?: number;
  };

  type listDatamodelParametersParams = {
    /** Device ID */
    id: string;
  };

  type listDatamodelsParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listDatamodelTemplatesParams = {
    /** Device ID */
    id: string;
  };

  type listDeletedDevicesData = {
    Data?: Device[];
    Total?: number;
  };

  type listDeletedDevicesParams = {
    /** Query */
    q?: string;
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
  };

  type listDeletedProductsData = {
    Data?: Product[];
    Total?: number;
  };

  type listDeletedProductsParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listDeviceAlarmsData = {
    Data?: DeviceAlarm[];
    Total?: number;
  };

  type listDeviceAlarmsParams = {
    /** Offset */
    offset?: number;
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Query */
    q?: string;
    /** Sort order */
    order_by?: string;
  };

  type listDeviceEventsData = {
    Data?: DeviceEvent[];
    Total?: number;
  };

  type listDeviceEventsParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listDeviceLogsParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listDeviceMethodCallsData = {
    Data?: DeviceMethodCall[];
    Total?: number;
  };

  type listDeviceMethodCallsParams = {
    /** Offset */
    offset?: number;
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Query */
    q?: string;
    /** Sort order */
    order_by?: string;
  };

  type listDeviceMethodsParams = {
    /** Device ID */
    id: string;
  };

  type listDeviceParametersParams = {
    /** Device ID */
    id: string;
  };

  type listDevicesData = {
    Data?: Device[];
    Total?: number;
  };

  type listDevicesParams = {
    /** Query */
    q?: string;
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
  };

  type listDeviceTransferLogsData = {
    Data?: DeviceTransferLog[];
    Total?: number;
  };

  type listDeviceTransferLogsParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listFirmwaresData = {
    Data?: Firmware[];
    Total?: number;
  };

  type listFirmwaresParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listKPIMeasuresData = {
    Data?: KPIMeas[];
    Total?: number;
  };

  type listKpiMeasuresParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listKPITemplateRecordsData = {
    Data?: Record<string, any>[];
    Total?: number;
  };

  type listKpiTemplateRecordsParams = {
    /** ID */
    id: number;
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listKPITemplatesData = {
    Data?: KPITemplate[];
    Total?: number;
  };

  type listKpiTemplatesParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listObjectsParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listProductFirmwaresParams = {
    /** Product ID */
    id: number;
  };

  type listProductsData = {
    Data?: Product[];
    Total?: number;
  };

  type listProductsParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listRolesData = {
    Data?: Role[];
    Total?: number;
  };

  type listRolesParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listRoutesBody = {
    Data?: ApiRule[];
    Total?: number;
  };

  type listRulesBody = {
    Data?: string[];
    Total?: number;
  };

  type listS3ObjectsData = {
    Data?: S3Object[];
    Total?: number;
  };

  type listTaskDeviceLogsData = {
    Data?: TaskDeviceLog[];
    Total?: number;
  };

  type listTasksData = {
    Data?: Task[];
    Total?: number;
  };

  type listTasksParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type listUsersData = {
    Data?: User[];
    Total?: number;
  };

  type listUsersParams = {
    /** Page */
    page?: number;
    /** Page size */
    page_size?: number;
    /** Sort order */
    order_by?: string;
    /** Query */
    q?: string;
  };

  type Map = true;

  type ParameterNotifications = true;

  type ParameterValues = true;

  type ParameterWritables = true;

  type Product = {
    Created?: string;
    Default?: boolean;
    Deleted?: DeletedAt;
    Enable?: boolean;
    Firmwares?: Firmware[];
    Id?: number;
    Manufacturer?: string;
    MetaData?: Tags;
    ModelName?: string;
    Oui?: string;
    ParameterPath?: string;
    PerformanceValueDefines?: Tags;
    ProductClass?: string;
    ProductType?: string;
    Schema?: string;
    SupportedAlarms?: ProductSupportedAlarm[];
    Updated?: string;
  };

  type ProductSupportedAlarm = {
    AlarmIdentifier?: string;
    Created?: string;
    EventType?: string;
    Id?: number;
    Oui?: string;
    PerceivedSeverity?: string;
    ProbableCause?: string;
    Product?: Product;
    ProductClass?: string;
    ProductId?: number;
    ProductType?: string;
    Schema?: string;
    SerialNumber?: string;
    SpecificProblem?: string;
    Updated?: string;
  };

  type putObjectParams = {
    /** S3 objet bucket */
    bucket: string;
    /** S3 object key */
    key: string;
  };

  type rebootDeviceParams = {
    /** Device ID */
    id: number;
  };

  type recoverDeletedDeviceParams = {
    /** Device ID */
    id: number;
  };

  type resetUserPasswordParams = {
    /** User ID */
    id: number;
  };

  type Role = {
    Alias?: string;
    Created?: number;
    Default?: boolean;
    Deleted?: DeletedAt;
    Enable?: boolean;
    Id?: number;
    MetaData?: Tags;
    Name?: string;
    Rules?: string[];
    Schema?: string;
    Updated?: number;
  };

  type S3Object = {
    Bucket?: string;
    Created?: string;
    Deleted?: DeletedAt;
    DownloadUrl?: string;
    FileName?: string;
    FileSize?: number;
    Id?: number;
    Key?: string;
    MetaData?: Tags;
    Schema?: string;
    SignMethod?: string;
    Signature?: string;
    Updated?: string;
  };

  type setDeviceDisableParams = {
    /** Device ID */
    id: number;
  };

  type setDeviceEnableParams = {
    /** Device ID */
    id: number;
  };

  type setDeviceParameterValuesBody = {
    Values?: Record<string, any>;
  };

  type setDeviceParameterValuesParams = {
    /** Device ID */
    id: number;
  };

  type setDevicePerfDisableParams = {
    /** Device ID */
    id: number;
  };

  type setDevicePerfEnableParams = {
    /** Device ID */
    id: number;
  };

  type setFirmwareProductsBody = {
    ModelNames?: string[];
  };

  type setFirmwareProductsParams = {
    /** Device ID */
    id: number;
  };

  type setGroupForDeviceBody = {
    GroupId?: number;
  };

  type setGroupForDeviceParams = {
    /** Device ID */
    id: number;
  };

  type setGroupParentBody = {
    parentID?: number;
  };

  type setGroupParentParams = {
    /** Group ID */
    id: number;
  };

  type setKpiMeasureDisableParams = {
    /** ID */
    id: number;
  };

  type setKpiMeasureEnableParams = {
    /** ID */
    id: number;
  };

  type setProductDisableParams = {
    /** Product ID */
    id: number;
  };

  type setProductEnableParams = {
    /** Product ID */
    id: number;
  };

  type setRoleApiRulesBody = {
    RoleIds?: string[];
  };

  type setRoleDisableParams = {
    /** Role ID */
    id: number;
  };

  type setRoleEnableParams = {
    /** Role ID */
    id: number;
  };

  type setRoleRulesParams = {
    /** Role ID */
    id: number;
  };

  type setUserDisableParams = {
    /** User ID */
    id: number;
  };

  type setUserEnableParams = {
    /** User ID */
    id: number;
  };

  type setUserRoles = {
    roleNames?: string[];
  };

  type setUserRolesParams = {
    /** User ID */
    id: number;
  };

  type systemInfoData = {
    Build?: string;
    Name?: string;
    Version?: string;
  };

  type systemTimeData = {
    Current?: number;
  };

  type Tags = true;

  type Tags = true;

  type Task = {
    Created?: string;
    Creater?: string;
    Default?: boolean;
    Deleted?: DeletedAt;
    DeviceIds?: Tags;
    Enable?: boolean;
    ExecEndTime?: string;
    ExecInterval?: number;
    ExecLastTime?: string;
    ExecMode?: TaskExecMode;
    ExecNextTime?: string;
    ExecProcess?: number;
    ExecRate?: number;
    ExecStartTime?: string;
    ExecTimes?: number;
    Id?: number;
    MetaData?: Tags;
    RetryInterval?: number;
    RetryTimes?: number;
    Schema?: string;
    TaskName?: string;
    TaskStatus?: TaskStatus;
    TaskType?: string;
    Updated?: string;
  };

  type TaskDeviceLog = {
    Code?: number;
    Device?: Device;
    DeviceId?: number;
    DeviceIds?: Tags;
    EndTime?: string;
    ExecMode?: number;
    Info?: string;
    Schema?: string;
    StartTime?: string;
    Task?: Task;
    TaskId?: number;
    TaskName?: string;
    TaskType?: string;
    Time?: string;
    Updated?: string;
  };

  type TaskExecMode = 0 | 1;

  type TaskStatus = 0 | 1 | 2 | 3 | 4;

  type TSPoint = true;

  type TSQuery = {
    Format?: string;
    GroupBy?: string[];
    Id?: string;
    Interval?: number;
    Limit?: number;
    Offset?: number;
    OrderBy?: string[];
    Reverse?: boolean;
    Search?: string;
    Select?: TSQuerySelect[];
    Source?: string;
  };

  type TSQueryCommand = {
    From?: number;
    Query?: TSQuery;
    TimeZone?: string;
    To?: number;
  };

  type TSQuerySelect = {
    Func?: string;
    Name?: string;
    Source?: string;
  };

  type TSResult = {
    Data?: TSSeries[];
    From?: number;
    To?: number;
  };

  type TSSeries = {
    GroupBy?: Record<string, any>;
    Points?: TSPoint[];
  };

  type ugradeDeviceParams = {
    /** Device ID */
    id: number;
  };

  type updateDeviceInfoBody = {
    Name?: string;
  };

  type updateDeviceInfoParams = {
    /** Device ID */
    id: number;
  };

  type updateGroupInfoBody = {
    name?: string;
  };

  type updateGroupInfoParams = {
    /** Group ID */
    id: number;
  };

  type updateKPIMeasureInfoBody = {
    name?: string;
  };

  type updateKpiMeasureInfoParams = {
    /** ID */
    id: number;
  };

  type updateKPITemplateBody = {
    measTypeIds?: string[];
    periodicInterval?: number;
    selectIds?: number[];
    selectType?: string;
  };

  type updateKpiTemplateParams = {
    /** ID */
    id: number;
  };

  type updateProductInfoBody = {
    manufacturer?: string;
  };

  type updateProductInfoParams = {
    /** Product ID */
    id: number;
  };

  type updateRoleBody = {
    alias?: string;
  };

  type updateRoleParams = {
    /** Role ID */
    id: number;
  };

  type updateUserBody = {
    alias?: string;
  };

  type updateUserParams = {
    /** User ID */
    id: number;
  };

  type upgradeDeviceBody = {
    DelaySeconds?: number;
    FirmwareID?: number;
  };

  type uploadDeviceFileBody = {
    DelaySeconds?: number;
    /** Url          string `json:"Url"`
Username     string `json:"Username"`
Password     string `json:"Password"` */
    FileType?: string;
  };

  type uploadDeviceFileParams = {
    /** Device ID */
    id: number;
  };

  type User = {
    Alias?: string;
    Created?: number;
    Default?: boolean;
    Deleted?: DeletedAt;
    Enable?: boolean;
    Id?: number;
    MetaData?: Tags;
    Name?: string;
    PasswordExpireAt?: number;
    Roles?: string[];
    Rules?: string[];
    Schema?: string;
    Updated?: number;
  };
}
