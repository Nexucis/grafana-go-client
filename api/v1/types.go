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
	AppMode        string                  `json:"app_mode"`
	Analytics      *AnalyticsSettings      `json:"analytics"`
	AuthAnonymous  *AuthAnonymousSettings  `json:"auth.anonymous"`
	AuthBasic      *AuthBasicSettings      `json:"auth.basic"`
	AuthGithub     *AuthGithubSettings     `json:"auth.github"`
	AuthGoogle     *AuthGoogleSettings     `json:"auth.google"`
	AuthLdap       *AuthLdapSettings       `json:"auth.ldap"`
	AuthProxy      *AuthProxySettings      `json:"auth.proxy"`
	DashboardsJSON *DashboardsJSONSettings `json:"dashboards.json"`
	Database       *DatabaseSettings       `json:"database"`
	Emails         *EmailsSettings         `json:"emails"`
	Log            *LogSettings            `json:"log"`
	LogConsole     *LogConsoleSettings     `json:"log.console"`
	LogFile        *LogFileSettings        `json:"log.file"`
	Paths          *PathsSettings          `json:"paths"`
	Security       *SecuritySettings       `json:"security"`
	Server         *ServerSettings         `json:"server"`
	Session        *SessionSettings        `json:"session"`
	SMTP           *SMTPSettings           `json:"smtp"`
	Users          *UsersSettings          `json:"users"`
}

type AnalyticsSettings struct {
	GoogleAnalyticsUAID string `json:"google_analytics_ua_id"`
	ReportingEnabled    bool   `json:"reporting_enabled"`
}

type AuthAnonymousSettings struct {
	Enabled string `json:"enabled"`
	OrgName string `json:"org_name"`
	OrgRole string `json:"org_role"`
}

type AuthBasicSettings struct {
	Enabled string `json:"enabled"`
}

type AuthGithubSettings struct {
	AllowSignUp          string `json:"allow_sign_up"`
	AllowedDomains       string `json:"allowed_domains"`
	AllowedOrganizations string `json:"allowed_organizations"`
	APIURL               string `json:"api_url"`
	AuthURL              string `json:"auth_url"`
	ClientID             string `json:"client_id"`
	ClientSecret         string `json:"client_secret"`
	Enabled              string `json:"enabled"`
	Scopes               string `json:"scopes"`
	TeamIds              string `json:"team_ids"`
	TokenURL             string `json:"token_url"`
}
type AuthGoogleSettings struct {
	AllowSignUp    string `json:"allow_sign_up"`
	AllowedDomains string `json:"allowed_domains"`
	APIURL         string `json:"api_url"`
	AuthURL        string `json:"auth_url"`
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	Enabled        string `json:"enabled"`
	Scopes         string `json:"scopes"`
	TokenURL       string `json:"token_url"`
}
type AuthLdapSettings struct {
	ConfigFile string `json:"config_file"`
	Enabled    string `json:"enabled"`
}
type AuthProxySettings struct {
	AutoSignUp     string `json:"auto_sign_up"`
	Enabled        string `json:"enabled"`
	HeaderName     string `json:"header_name"`
	HeaderProperty string `json:"header_property"`
}
type DashboardsJSONSettings struct {
	Enabled string `json:"enabled"`
	Path    string `json:"path"`
}
type DatabaseSettings struct {
	Host     string `json:"host"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Path     string `json:"path"`
	SslMode  string `json:"ssl_mode"`
	Type     string `json:"type"`
	User     string `json:"user"`
}
type EmailsSettings struct {
	TemplatesPattern     string `json:"templates_pattern"`
	WelcomeEmailOnSignUp string `json:"welcome_email_on_sign_up"`
}
type LogSettings struct {
	BufferLen string `json:"buffer_len"`
	Level     string `json:"level"`
	Mode      string `json:"mode"`
}
type LogConsoleSettings struct {
	Level string `json:"level"`
}
type LogFileSettings struct {
	DailyRotate   string `json:"daily_rotate"`
	FileName      string `json:"file_name"`
	Level         string `json:"level"`
	LogRotate     string `json:"log_rotate"`
	MaxDays       string `json:"max_days"`
	MaxLines      string `json:"max_lines"`
	MaxLinesShift string `json:"max_lines_shift"`
	MaxSizeShift  string `json:"max_size_shift"`
}
type PathsSettings struct {
	Data string `json:"data"`
	Logs string `json:"logs"`
}
type SecuritySettings struct {
	AdminPassword      string `json:"admin_password"`
	AdminUser          string `json:"admin_user"`
	CookieRememberName string `json:"cookie_remember_name"`
	CookieUsername     string `json:"cookie_username"`
	DisableGravatar    string `json:"disable_gravatar"`
	LoginRememberDays  string `json:"login_remember_days"`
	SecretKey          string `json:"secret_key"`
}
type ServerSettings struct {
	CertFile         string `json:"cert_file"`
	CertKey          string `json:"cert_key"`
	Domain           string `json:"domain"`
	EnableGzip       string `json:"enable_gzip"`
	EnforceDomain    string `json:"enforce_domain"`
	HTTPAddr         string `json:"http_addr"`
	HTTPPort         string `json:"http_port"`
	Protocol         string `json:"protocol"`
	RootURL          string `json:"root_url"`
	RouterLogging    string `json:"router_logging"`
	DataProxyLogging string `json:"data_proxy_logging"`
	StaticRootPath   string `json:"static_root_path"`
}
type SessionSettings struct {
	CookieName      string `json:"cookie_name"`
	CookieSecure    string `json:"cookie_secure"`
	GcIntervalTime  string `json:"gc_interval_time"`
	Provider        string `json:"provider"`
	ProviderConfig  string `json:"provider_config"`
	SessionLifeTime string `json:"session_life_time"`
}
type SMTPSettings struct {
	CertFile     string `json:"cert_file"`
	Enabled      string `json:"enabled"`
	FromAddress  string `json:"from_address"`
	FromName     string `json:"from_name"`
	EhloIdentity string `json:"ehlo_identity"`
	Host         string `json:"host"`
	KeyFile      string `json:"key_file"`
	Password     string `json:"password"`
	SkipVerify   string `json:"skip_verify"`
	User         string `json:"user"`
}
type UsersSettings struct {
	AllowOrgCreate    string `json:"allow_org_create"`
	AllowSignUp       string `json:"allow_sign_up"`
	AutoAssignOrg     string `json:"auto_assign_org"`
	AutoAssignOrgRole string `json:"auto_assign_org_role"`
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

type GetAPIKeyResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type APIKeyForm struct {
	Name string `json:"name" binding:"Required"`
	Role string `json:"role" binding:"Required"`
}

type CreateAPIKeyResponse struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}
