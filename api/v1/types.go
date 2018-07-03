// Copyright 2018 Augustin Husson
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

type PostAnnotations struct {
	DashboardId int64       `json:"dashboardId"`
	PanelId     int64       `json:"panelId"`
	Time        int64       `json:"time"`
	Text        string      `json:"text"`
	Tags        []string    `json:"tags"`
	Data        interface{} `json:"data"`
	IsRegion    bool        `json:"isRegion"`
	TimeEnd     int64       `json:"timeEnd"`
}

type UpdateAnnotations struct {
	Id       int64    `json:"id"`
	Time     int64    `json:"time"`
	Text     string   `json:"text"`
	Tags     []string `json:"tags"`
	IsRegion bool     `json:"isRegion"`
	TimeEnd  int64    `json:"timeEnd"`
}

type DeleteAnnotations struct {
	AlertId      int64 `json:"alertId"`
	DashboardId  int64 `json:"dashboardId"`
	PanelId      int64 `json:"panelId"`
	AnnotationId int64 `json:"annotationId"`
	RegionId     int64 `json:"regionId"`
}

type PostGraphiteAnnotations struct {
	When int64    `json:"when"`
	What string   `json:"what"`
	Data string   `json:"data"`
	Tags []string `json:"tags"`
}

type ResponseCreateAnnotation struct {
	Id      int64  `json:"id"`
	EndId   int64  `json:"endId"`
	Message string `json:"message"`
}

type ResponseCreateGraphiteAnnotation struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

type ResponseGetAnnotation struct {
	Id          int64       `json:"id"`
	AlertId     int64       `json:"alertId"`
	AlertName   string      `json:"alertName"`
	DashboardId int64       `json:"dashboardId"`
	PanelId     int64       `json:"panelId"`
	UserId      int64       `json:"userId"`
	NewState    string      `json:"newState"`
	PrevState   string      `json:"prevState"`
	Created     int64       `json:"created"`
	Updated     int64       `json:"updated"`
	Time        int64       `json:"time"`
	Text        string      `json:"text"`
	RegionId    int64       `json:"regionId"`
	Tags        []string    `json:"tags"`
	Login       string      `json:"login"`
	Email       string      `json:"email"`
	AvatarUrl   string      `json:"avatarUrl"`
	Data        interface{} `json:"data"`
}

type QueryParamAnnotation struct {
	// epoch datetime in milliseconds. Optional.
	from int64
	// epoch datetime in milliseconds. Optional.
	to int64
	// number. Optional. Find annotations created by a specific user
	userId int64
	// number. Optional. Find annotations for a specified alert.
	alertId int64
	// number. Optional. Find annotations that are scoped to a specific dashboard
	dashboardId int64
	// number. Optional. Find annotations that are scoped to a specific panel
	panelId int64
	// string. Optional. Use this to filter global annotations. Global annotations are annotations from an annotation data source that are not connected specifically to a dashboard or panel
	tags []string
	// string. Optional. alert|annotation Return alerts or user created annotations
	_type string
	// number. Optional - default is 100. Max limit for results returned.
	limit int64
}

func (query *QueryParamAnnotation) From(from int64) *QueryParamAnnotation {
	query.from = from
	return query
}

func (query *QueryParamAnnotation) To(to int64) *QueryParamAnnotation {
	query.to = to
	return query
}

func (query *QueryParamAnnotation) AlertID(alertID int64) *QueryParamAnnotation {
	query.alertId = alertID
	return query
}

func (query *QueryParamAnnotation) UserID(userID int64) *QueryParamAnnotation {
	query.userId = userID
	return query
}

func (query *QueryParamAnnotation) DashboardID(dashboardID int64) *QueryParamAnnotation {
	query.dashboardId = dashboardID
	return query
}

func (query *QueryParamAnnotation) PanelID(panelID int64) *QueryParamAnnotation {
	query.panelId = panelID
	return query
}

func (query *QueryParamAnnotation) AddTag(tag string) *QueryParamAnnotation {
	query.tags = append(query.tags, tag)
	return query
}

func (query *QueryParamAnnotation) Type(t string) *QueryParamAnnotation {
	query._type = t
	return query
}

func (query *QueryParamAnnotation) Limit(limit int64) *QueryParamAnnotation {
	query.limit = limit
	return query
}

type AdminSettings struct {
	Default                       *DefaultSettings                       `json:"DEFAULT"`
	Alerting                      *AlertingSettings                      `json:"alerting"`
	Analytics                     *AnalyticsSettings                     `json:"analytics"`
	Auth                          *AuthSettings                          `json:"auth"`
	AuthAnonymous                 *AuthAnonymousSettings                 `json:"auth.anonymous"`
	AuthBasic                     *AuthBasicSettings                     `json:"auth.basic"`
	AuthGenericOauth              *AuthGenericOauthSettings              `json:"auth.generic_oauth"`
	AuthGithub                    *AuthGithubSettings                    `json:"auth.github"`
	AuthGoogle                    *AuthGoogleSettings                    `json:"auth.google"`
	AuthGrafanaCom                *AuthGrafanaComSettings                `json:"auth.grafana_com"`
	AuthGrafananet                *AuthGrafananetSettings                `json:"auth.grafananet"`
	AuthLdap                      *AuthLdapSettings                      `json:"auth.ldap"`
	AuthProxy                     *AuthProxySettings                     `json:"auth.proxy"`
	DashboardsJSON                *DashboardsJSONSettings                `json:"dashboards.json"`
	Database                      *DatabaseSettings                      `json:"database"`
	DataProxy                     *DataProxySettings                     `json:"dataproxy"`
	Emails                        *EmailsSettings                        `json:"emails"`
	Explore                       *ExploreSettings                       `json:"explore"`
	ExternalImageStorage          *ExternalImageStorageSettings          `json:"external_image_storage"`
	ExternalImageStorageAzureBlob *ExternalImageStorageAzureBlobSettings `json:"external_image_storage.azure_blob"`
	ExternalImageStorageGcs       *ExternalImageStorageGcsSettings       `json:"external_image_storage.gcs"`
	ExternalImageStorageS3        *ExternalImageStorageS3Settings        `json:"external_image_storage.s3"`
	ExternalImageStorageWebdav    *ExternalImageStorageWebdavSettings    `json:"external_image_storage.webdav"`
	GrafanaCom                    *GrafanaComSettings                    `json:"grafana_com"`
	GrafanaNet                    *GrafanaNetSettings                    `json:"grafana_net"`
	Log                           *LogSettings                           `json:"log"`
	LogConsole                    *LogConsoleSettings                    `json:"log.console"`
	LogFile                       *LogFileSettings                       `json:"log.file"`
	LogSyslog                     *LogSyslogSettings                     `json:"log.syslog"`
	Metrics                       *MetricsSettings                       `json:"metrics"`
	MetricsGraphite               *MetricsGraphiteSettings               `json:"metrics.graphite"`
	Paths                         *PathsSettings                         `json:"paths"`
	Plugins                       *PluginsSettings                       `json:"plugins"`
	Quota                         *QuotaSettings                         `json:"quota"`
	Rendering                     *RenderingSettings                     `json:"rendering"`
	Security                      *SecuritySettings                      `json:"security"`
	Server                        *ServerSettings                        `json:"server"`
	Session                       *SessionSettings                       `json:"session"`
	SMTP                          *SMTPSettings                          `json:"smtp"`
	Snapshots                     *SnapshotsSettings                     `json:"snapshots"`
	TracingJaeger                 *TracingJaegerSettings                 `json:"tracing.jaeger"`
	Users                         *UsersSettings                         `json:"users"`
}

type DefaultSettings struct {
	AppMode      string `json:"app_mode"`
	InstanceName string `json:"instance_name"`
}

type AlertingSettings struct {
	Enabled       string `json:"enabled"`
	ExecuteAlerts string `json:"execute_alerts"`
}
type AnalyticsSettings struct {
	CheckForUpdates     string `json:"check_for_updates"`
	GoogleAnalyticsUaID string `json:"google_analytics_ua_id"`
	GoogleTagManagerID  string `json:"google_tag_manager_id"`
	ReportingEnabled    string `json:"reporting_enabled"`
}

type AuthSettings struct {
	DisableLoginForm   string `json:"disable_login_form"`
	DisableSignoutMenu string `json:"disable_signout_menu"`
	SignoutRedirectURL string `json:"signout_redirect_url"`
}

type AuthAnonymousSettings struct {
	Enabled string `json:"enabled"`
	OrgName string `json:"org_name"`
	OrgRole string `json:"org_role"`
}

type AuthBasicSettings struct {
	Enabled string `json:"enabled"`
}

type AuthGenericOauthSettings struct {
	AllowSignUp           string `json:"allow_sign_up"`
	AllowedDomains        string `json:"allowed_domains"`
	AllowedOrganizations  string `json:"allowed_organizations"`
	APIURL                string `json:"api_url"`
	AuthURL               string `json:"auth_url"`
	ClientID              string `json:"client_id"`
	ClientSecret          string `json:"client_secret"`
	Enabled               string `json:"enabled"`
	HostedDomain          string `json:"hosted_domain"`
	Name                  string `json:"name"`
	Scopes                string `json:"scopes"`
	TeamIds               string `json:"team_ids"`
	TLSClientCa           string `json:"tls_client_ca"`
	TLSClientCert         string `json:"tls_client_cert"`
	TLSClientKey          string `json:"tls_client_key"`
	TLSSkipVerifyInsecure string `json:"tls_skip_verify_insecure"`
	TokenURL              string `json:"token_url"`
}

type AuthGithubSettings struct {
	AllowSignUp           string `json:"allow_sign_up"`
	AllowedDomains        string `json:"allowed_domains"`
	AllowedOrganizations  string `json:"allowed_organizations"`
	APIURL                string `json:"api_url"`
	AuthURL               string `json:"auth_url"`
	ClientID              string `json:"client_id"`
	ClientSecret          string `json:"client_secret"`
	Enabled               string `json:"enabled"`
	HostedDomain          string `json:"hosted_domain"`
	Name                  string `json:"name"`
	Scopes                string `json:"scopes"`
	TeamIds               string `json:"team_ids"`
	TLSClientCa           string `json:"tls_client_ca"`
	TLSClientCert         string `json:"tls_client_cert"`
	TLSClientKey          string `json:"tls_client_key"`
	TLSSkipVerifyInsecure string `json:"tls_skip_verify_insecure"`
	TokenURL              string `json:"token_url"`
}

type AuthGoogleSettings struct {
	AllowSignUp           string `json:"allow_sign_up"`
	AllowedDomains        string `json:"allowed_domains"`
	APIURL                string `json:"api_url"`
	AuthURL               string `json:"auth_url"`
	ClientID              string `json:"client_id"`
	ClientSecret          string `json:"client_secret"`
	Enabled               string `json:"enabled"`
	HostedDomain          string `json:"hosted_domain"`
	Name                  string `json:"name"`
	Scopes                string `json:"scopes"`
	TLSClientCa           string `json:"tls_client_ca"`
	TLSClientCert         string `json:"tls_client_cert"`
	TLSClientKey          string `json:"tls_client_key"`
	TLSSkipVerifyInsecure string `json:"tls_skip_verify_insecure"`
	TokenURL              string `json:"token_url"`
}

type AuthGrafanaComSettings struct {
	AllowSignUp           string `json:"allow_sign_up"`
	AllowedDomains        string `json:"allowed_domains"`
	AllowedOrganizations  string `json:"allowed_organizations"`
	APIURL                string `json:"api_url"`
	AuthURL               string `json:"auth_url"`
	ClientID              string `json:"client_id"`
	ClientSecret          string `json:"client_secret"`
	Enabled               string `json:"enabled"`
	HostedDomain          string `json:"hosted_domain"`
	Name                  string `json:"name"`
	Scopes                string `json:"scopes"`
	TLSClientCa           string `json:"tls_client_ca"`
	TLSClientCert         string `json:"tls_client_cert"`
	TLSClientKey          string `json:"tls_client_key"`
	TLSSkipVerifyInsecure string `json:"tls_skip_verify_insecure"`
	TokenURL              string `json:"token_url"`
}

type AuthGrafananetSettings struct {
	AllowSignUp           string `json:"allow_sign_up"`
	AllowedDomains        string `json:"allowed_domains"`
	AllowedOrganizations  string `json:"allowed_organizations"`
	APIURL                string `json:"api_url"`
	AuthURL               string `json:"auth_url"`
	ClientID              string `json:"client_id"`
	ClientSecret          string `json:"client_secret"`
	Enabled               string `json:"enabled"`
	HostedDomain          string `json:"hosted_domain"`
	Name                  string `json:"name"`
	Scopes                string `json:"scopes"`
	TLSClientCa           string `json:"tls_client_ca"`
	TLSClientCert         string `json:"tls_client_cert"`
	TLSClientKey          string `json:"tls_client_key"`
	TLSSkipVerifyInsecure string `json:"tls_skip_verify_insecure"`
	TokenURL              string `json:"token_url"`
}

type AuthLdapSettings struct {
	AllowSignUp string `json:"allow_sign_up"`
	ConfigFile  string `json:"config_file"`
	Enabled     string `json:"enabled"`
}

type AuthProxySettings struct {
	AutoSignUp     string `json:"auto_sign_up"`
	Enabled        string `json:"enabled"`
	HeaderName     string `json:"header_name"`
	HeaderProperty string `json:"header_property"`
	Headers        string `json:"headers"`
	LdapSyncTTL    string `json:"ldap_sync_ttl"`
	Whitelist      string `json:"whitelist"`
}

type DashboardsJSONSettings struct {
	Enabled string `json:"enabled"`
	Path    string `json:"path"`
}

type DatabaseSettings struct {
	CaCertPath       string `json:"ca_cert_path"`
	ClientCertPath   string `json:"client_cert_path"`
	ClientKeyPath    string `json:"client_key_path"`
	ConnMaxLifetime  string `json:"conn_max_lifetime"`
	ConnectionString string `json:"connection_string"`
	Host             string `json:"host"`
	LogQueries       string `json:"log_queries"`
	MaxIdleConn      string `json:"max_idle_conn"`
	MaxOpenConn      string `json:"max_open_conn"`
	Name             string `json:"name"`
	Password         string `json:"password"`
	Path             string `json:"path"`
	ServerCertName   string `json:"server_cert_name"`
	SslMode          string `json:"ssl_mode"`
	Type             string `json:"type"`
	URL              string `json:"url"`
	User             string `json:"user"`
}

type DataProxySettings struct {
	Logging string `json:"logging"`
}

type EmailsSettings struct {
	TemplatesPattern     string `json:"templates_pattern"`
	WelcomeEmailOnSignUp string `json:"welcome_email_on_sign_up"`
}

type ExploreSettings struct {
	Enabled string `json:"enabled"`
}

type ExternalImageStorageSettings struct {
	Provider string `json:"provider"`
}

type ExternalImageStorageAzureBlobSettings struct {
	AccountKey    string `json:"account_key"`
	AccountName   string `json:"account_name"`
	ContainerName string `json:"container_name"`
}

type ExternalImageStorageGcsSettings struct {
	Bucket  string `json:"bucket"`
	KeyFile string `json:"key_file"`
	Path    string `json:"path"`
}

type ExternalImageStorageS3Settings struct {
	AccessKey string `json:"access_key"`
	Bucket    string `json:"bucket"`
	BucketURL string `json:"bucket_url"`
	Path      string `json:"path"`
	Region    string `json:"region"`
	SecretKey string `json:"secret_key"`
}

type ExternalImageStorageWebdavSettings struct {
	Password  string `json:"password"`
	PublicURL string `json:"public_url"`
	URL       string `json:"url"`
	Username  string `json:"username"`
}

type GrafanaComSettings struct {
	URL string `json:"url"`
}

type GrafanaNetSettings struct {
	URL string `json:"url"`
}

type LogSettings struct {
	Filters string `json:"filters"`
	Level   string `json:"level"`
	Mode    string `json:"mode"`
}

type LogConsoleSettings struct {
	Format string `json:"format"`
	Level  string `json:"level"`
}

type LogFileSettings struct {
	DailyRotate  string `json:"daily_rotate"`
	Format       string `json:"format"`
	Level        string `json:"level"`
	LogRotate    string `json:"log_rotate"`
	MaxDays      string `json:"max_days"`
	MaxLines     string `json:"max_lines"`
	MaxSizeShift string `json:"max_size_shift"`
}

type LogSyslogSettings struct {
	Address  string `json:"address"`
	Facility string `json:"facility"`
	Format   string `json:"format"`
	Level    string `json:"level"`
	Network  string `json:"network"`
	Tag      string `json:"tag"`
}

type MetricsSettings struct {
	Enabled         string `json:"enabled"`
	IntervalSeconds string `json:"interval_seconds"`
}

type MetricsGraphiteSettings struct {
	Address string `json:"address"`
	Prefix  string `json:"prefix"`
}

type PathsSettings struct {
	Data         string `json:"data"`
	Logs         string `json:"logs"`
	Plugins      string `json:"plugins"`
	Provisioning string `json:"provisioning"`
}

type PluginsSettings struct {
	AppTLSSkipVerifyInsecure string `json:"app_tls_skip_verify_insecure"`
}

type QuotaSettings struct {
	Enabled          string `json:"enabled"`
	GlobalAPIKey     string `json:"global_api_key"`
	GlobalDashboard  string `json:"global_dashboard"`
	GlobalDataSource string `json:"global_data_source"`
	GlobalOrg        string `json:"global_org"`
	GlobalSession    string `json:"global_session"`
	GlobalUser       string `json:"global_user"`
	OrgAPIKey        string `json:"org_api_key"`
	OrgDashboard     string `json:"org_dashboard"`
	OrgDataSource    string `json:"org_data_source"`
	OrgUser          string `json:"org_user"`
	UserOrg          string `json:"user_org"`
}

type RenderingSettings struct {
	ServerURL string `json:"server_url"`
}

type SecuritySettings struct {
	AdminPassword                    string `json:"admin_password"`
	AdminUser                        string `json:"admin_user"`
	CookieRememberName               string `json:"cookie_remember_name"`
	CookieUsername                   string `json:"cookie_username"`
	DataSourceProxyWhitelist         string `json:"data_source_proxy_whitelist"`
	DisableBruteForceLoginProtection string `json:"disable_brute_force_login_protection"`
	DisableGravatar                  string `json:"disable_gravatar"`
	LoginRememberDays                string `json:"login_remember_days"`
	SecretKey                        string `json:"secret_key"`
}

type ServerSettings struct {
	CertFile       string `json:"cert_file"`
	CertKey        string `json:"cert_key"`
	Domain         string `json:"domain"`
	EnableGzip     string `json:"enable_gzip"`
	EnforceDomain  string `json:"enforce_domain"`
	HTTPAddr       string `json:"http_addr"`
	HTTPPort       string `json:"http_port"`
	Protocol       string `json:"protocol"`
	RootURL        string `json:"root_url"`
	RouterLogging  string `json:"router_logging"`
	Socket         string `json:"socket"`
	StaticRootPath string `json:"static_root_path"`
}

type SessionSettings struct {
	ConnMaxLifetime string `json:"conn_max_lifetime"`
	CookieName      string `json:"cookie_name"`
	CookieSecure    string `json:"cookie_secure"`
	GcIntervalTime  string `json:"gc_interval_time"`
	Provider        string `json:"provider"`
	ProviderConfig  string `json:"provider_config"`
	SessionLifeTime string `json:"session_life_time"`
}

type SMTPSettings struct {
	CertFile     string `json:"cert_file"`
	EhloIdentity string `json:"ehlo_identity"`
	Enabled      string `json:"enabled"`
	FromAddress  string `json:"from_address"`
	FromName     string `json:"from_name"`
	Host         string `json:"host"`
	KeyFile      string `json:"key_file"`
	Password     string `json:"password"`
	SkipVerify   string `json:"skip_verify"`
	User         string `json:"user"`
}

type SnapshotsSettings struct {
	ExternalEnabled       string `json:"external_enabled"`
	ExternalSnapshotName  string `json:"external_snapshot_name"`
	ExternalSnapshotURL   string `json:"external_snapshot_url"`
	SnapshotRemoveExpired string `json:"snapshot_remove_expired"`
}

type TracingJaegerSettings struct {
	Address           string `json:"address"`
	AlwaysIncludedTag string `json:"always_included_tag"`
	SamplerParam      string `json:"sampler_param"`
	SamplerType       string `json:"sampler_type"`
}

type UsersSettings struct {
	AllowOrgCreate         string `json:"allow_org_create"`
	AllowSignUp            string `json:"allow_sign_up"`
	AutoAssignOrg          string `json:"auto_assign_org"`
	AutoAssignOrgRole      string `json:"auto_assign_org_role"`
	DefaultTheme           string `json:"default_theme"`
	ExternalManageInfo     string `json:"external_manage_info"`
	ExternalManageLinkName string `json:"external_manage_link_name"`
	ExternalManageLinkURL  string `json:"external_manage_link_url"`
	LoginHint              string `json:"login_hint"`
	VerifyEmailEnabled     string `json:"verify_email_enabled"`
	ViewersCanEdit         string `json:"viewers_can_edit"`
}

type AdminCreateUserForm struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Name     string `json:"name"`
	Password string `json:"password" binding:"Required"`
}

type AdminCreateUserResponse struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

type AdminUpdateUserPasswordForm struct {
	Password string `json:"password" binding:"Required"`
}

type AdminUpdateUserPermissionsForm struct {
	IsGrafanaAdmin bool `json:"isGrafanaAdmin"`
}

type UserQuotaResponse struct {
	UserId int64  `json:"user_id"`
	Target string `json:"target"`
	Limit  int64  `json:"limit"`
	Used   int64  `json:"used"`
}

type UpdateUserQuotaForm struct {
	Target string `json:"target"`
	Limit  int64  `json:"limit"`
	UserId int64  `json:"-"`
}

type AdminStats struct {
	Users       int `json:"users"`
	Orgs        int `json:"orgs"`
	Dashboards  int `json:"dashboards"`
	Snapshots   int `json:"snapshots"`
	Tags        int `json:"tags"`
	Datasources int `json:"datasources"`
	Playlists   int `json:"playlists"`
	Stars       int `json:"stars"`
	Alerts      int `json:"alerts"`
	ActiveUsers int `json:"activeUsers"`
}

type PauseAllAlertsForm struct {
	Paused bool `json:"paused"`
}

type PauseAllAlertsResponse struct {
	AlertAffected int64  `json:"alertAffected"`
	State         string `json:"state"`
	Message       string `json:"message"`
}

type RoleType string

const (
	RoleViewer RoleType = "Viewer"
	RoleEditor RoleType = "Editor"
	RoleAdmin  RoleType = "Admin"
)

type GetAPIKeyResponse struct {
	Id   int64    `json:"id"`
	Name string   `json:"name"`
	Role RoleType `json:"role"`
}

type APIKeyForm struct {
	Name string   `json:"name" binding:"Required"`
	Role RoleType `json:"role" binding:"Required"`
}

type CreateAPIKeyResponse struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}
