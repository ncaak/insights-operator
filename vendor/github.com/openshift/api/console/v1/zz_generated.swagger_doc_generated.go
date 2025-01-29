package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_Link = map[string]string{
	"":     "Represents a standard link that could be generated in HTML",
	"text": "text is the display text for the link",
	"href": "href is the absolute secure URL for the link (must use https)",
}

func (Link) SwaggerDoc() map[string]string {
	return map_Link
}

var map_CLIDownloadLink = map[string]string{
	"text": "text is the display text for the link",
	"href": "href is the absolute secure URL for the link (must use https)",
}

func (CLIDownloadLink) SwaggerDoc() map[string]string {
	return map_CLIDownloadLink
}

var map_ConsoleCLIDownload = map[string]string{
	"":         "ConsoleCLIDownload is an extension for configuring openshift web console command line interface (CLI) downloads.\n\nCompatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsoleCLIDownload) SwaggerDoc() map[string]string {
	return map_ConsoleCLIDownload
}

var map_ConsoleCLIDownloadList = map[string]string{
	"":         "Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsoleCLIDownloadList) SwaggerDoc() map[string]string {
	return map_ConsoleCLIDownloadList
}

var map_ConsoleCLIDownloadSpec = map[string]string{
	"":            "ConsoleCLIDownloadSpec is the desired cli download configuration.",
	"displayName": "displayName is the display name of the CLI download.",
	"description": "description is the description of the CLI download (can include markdown).",
	"links":       "links is a list of objects that provide CLI download link details.",
}

func (ConsoleCLIDownloadSpec) SwaggerDoc() map[string]string {
	return map_ConsoleCLIDownloadSpec
}

var map_ConsoleExternalLogLink = map[string]string{
	"":         "ConsoleExternalLogLink is an extension for customizing OpenShift web console log links.\n\nCompatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsoleExternalLogLink) SwaggerDoc() map[string]string {
	return map_ConsoleExternalLogLink
}

var map_ConsoleExternalLogLinkList = map[string]string{
	"":         "Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsoleExternalLogLinkList) SwaggerDoc() map[string]string {
	return map_ConsoleExternalLogLinkList
}

var map_ConsoleExternalLogLinkSpec = map[string]string{
	"":                "ConsoleExternalLogLinkSpec is the desired log link configuration. The log link will appear on the logs tab of the pod details page.",
	"text":            "text is the display text for the link",
	"hrefTemplate":    "hrefTemplate is an absolute secure URL (must use https) for the log link including variables to be replaced. Variables are specified in the URL with the format ${variableName}, for instance, ${containerName} and will be replaced with the corresponding values from the resource. Resource is a pod. Supported variables are: - ${resourceName} - name of the resource which containes the logs - ${resourceUID} - UID of the resource which contains the logs\n              - e.g. `11111111-2222-3333-4444-555555555555`\n- ${containerName} - name of the resource's container that contains the logs - ${resourceNamespace} - namespace of the resource that contains the logs - ${resourceNamespaceUID} - namespace UID of the resource that contains the logs - ${podLabels} - JSON representation of labels matching the pod with the logs\n            - e.g. `{\"key1\":\"value1\",\"key2\":\"value2\"}`\n\ne.g., https://example.com/logs?resourceName=${resourceName}&containerName=${containerName}&resourceNamespace=${resourceNamespace}&podLabels=${podLabels}",
	"namespaceFilter": "namespaceFilter is a regular expression used to restrict a log link to a matching set of namespaces (e.g., `^openshift-`). The string is converted into a regular expression using the JavaScript RegExp constructor. If not specified, links will be displayed for all the namespaces.",
}

func (ConsoleExternalLogLinkSpec) SwaggerDoc() map[string]string {
	return map_ConsoleExternalLogLinkSpec
}

var map_ApplicationMenuSpec = map[string]string{
	"":         "ApplicationMenuSpec is the specification of the desired section and icon used for the link in the application menu.",
	"section":  "section is the section of the application menu in which the link should appear. This can be any text that will appear as a subheading in the application menu dropdown. A new section will be created if the text does not match text of an existing section.",
	"imageURL": "imageURL is the URL for the icon used in front of the link in the application menu. The URL must be an HTTPS URL or a Data URI. The image should be square and will be shown at 24x24 pixels.",
}

func (ApplicationMenuSpec) SwaggerDoc() map[string]string {
	return map_ApplicationMenuSpec
}

var map_ConsoleLink = map[string]string{
	"":         "ConsoleLink is an extension for customizing OpenShift web console links.\n\nCompatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsoleLink) SwaggerDoc() map[string]string {
	return map_ConsoleLink
}

var map_ConsoleLinkList = map[string]string{
	"":         "Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsoleLinkList) SwaggerDoc() map[string]string {
	return map_ConsoleLinkList
}

var map_ConsoleLinkSpec = map[string]string{
	"":                   "ConsoleLinkSpec is the desired console link configuration.",
	"location":           "location determines which location in the console the link will be appended to (ApplicationMenu, HelpMenu, UserMenu, NamespaceDashboard).",
	"applicationMenu":    "applicationMenu holds information about section and icon used for the link in the application menu, and it is applicable only when location is set to ApplicationMenu.",
	"namespaceDashboard": "namespaceDashboard holds information about namespaces in which the dashboard link should appear, and it is applicable only when location is set to NamespaceDashboard. If not specified, the link will appear in all namespaces.",
}

func (ConsoleLinkSpec) SwaggerDoc() map[string]string {
	return map_ConsoleLinkSpec
}

var map_NamespaceDashboardSpec = map[string]string{
	"":                  "NamespaceDashboardSpec is a specification of namespaces in which the dashboard link should appear. If both namespaces and namespaceSelector are specified, the link will appear in namespaces that match either",
	"namespaces":        "namespaces is an array of namespace names in which the dashboard link should appear.",
	"namespaceSelector": "namespaceSelector is used to select the Namespaces that should contain dashboard link by label. If the namespace labels match, dashboard link will be shown for the namespaces.",
}

func (NamespaceDashboardSpec) SwaggerDoc() map[string]string {
	return map_NamespaceDashboardSpec
}

var map_ConsoleNotification = map[string]string{
	"":         "ConsoleNotification is the extension for configuring openshift web console notifications.\n\nCompatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsoleNotification) SwaggerDoc() map[string]string {
	return map_ConsoleNotification
}

var map_ConsoleNotificationList = map[string]string{
	"":         "Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsoleNotificationList) SwaggerDoc() map[string]string {
	return map_ConsoleNotificationList
}

var map_ConsoleNotificationSpec = map[string]string{
	"":                "ConsoleNotificationSpec is the desired console notification configuration.",
	"text":            "text is the visible text of the notification.",
	"location":        "location is the location of the notification in the console. Valid values are: \"BannerTop\", \"BannerBottom\", \"BannerTopBottom\".",
	"link":            "link is an object that holds notification link details.",
	"color":           "color is the color of the text for the notification as CSS data type color.",
	"backgroundColor": "backgroundColor is the color of the background for the notification as CSS data type color.",
}

func (ConsoleNotificationSpec) SwaggerDoc() map[string]string {
	return map_ConsoleNotificationSpec
}

var map_ConsolePlugin = map[string]string{
	"":         "ConsolePlugin is an extension for customizing OpenShift web console by dynamically loading code from another service running on the cluster.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec contains the desired configuration for the console plugin.",
}

func (ConsolePlugin) SwaggerDoc() map[string]string {
	return map_ConsolePlugin
}

var map_ConsolePluginBackend = map[string]string{
	"":        "ConsolePluginBackend holds information about the endpoint which serves the console's plugin",
	"type":    "type is the backend type which servers the console's plugin. Currently only \"Service\" is supported.",
	"service": "service is a Kubernetes Service that exposes the plugin using a deployment with an HTTP server. The Service must use HTTPS and Service serving certificate. The console backend will proxy the plugins assets from the Service using the service CA bundle.",
}

func (ConsolePluginBackend) SwaggerDoc() map[string]string {
	return map_ConsolePluginBackend
}

var map_ConsolePluginCSP = map[string]string{
	"":          "ConsolePluginCSP holds configuration for a specific CSP directive",
	"directive": "directive specifies which Content-Security-Policy directive to configure. Available directive types are DefaultSrc, ScriptSrc, StyleSrc, ImgSrc, FontSrc, ObjectSrc and ConnectSrc. DefaultSrc directive serves as a fallback for the other CSP fetch directives. For more information about the DefaultSrc directive, see: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/default-src ScriptSrc directive specifies valid sources for JavaScript. For more information about the ScriptSrc directive, see: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/script-src StyleSrc directive specifies valid sources for stylesheets. For more information about the StyleSrc directive, see: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/style-src ImgSrc directive specifies a valid sources of images and favicons. For more information about the ImgSrc directive, see: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/img-src FontSrc directive specifies valid sources for fonts loaded using @font-face. For more information about the FontSrc directive, see: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/font-src ObjectSrc directive specifies valid sources for the <object> and <embed> elements. For more information about the ObjectSrc directive, see: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/object-src ConnectSrc directive restricts the URLs which can be loaded using script interfaces. For more information about the ConnectSrc directive, see: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/connect-src",
	"values":    "values defines an array of values to append to the console defaults for this directive. Each ConsolePlugin may define their own directives with their values. These will be set by the OpenShift web console's backend, as part of its Content-Security-Policy header. The array can contain at most 16 values. Each directive value must have a maximum length of 1024 characters and must not contain whitespace, commas (,), semicolons (;) or single quotes ('). The value '*' is not permitted. Each value in the array must be unique.",
}

func (ConsolePluginCSP) SwaggerDoc() map[string]string {
	return map_ConsolePluginCSP
}

var map_ConsolePluginI18n = map[string]string{
	"":         "ConsolePluginI18n holds information on localization resources that are served by the dynamic plugin.",
	"loadType": "loadType indicates how the plugin's localization resource should be loaded. Valid values are Preload, Lazy and the empty string. When set to Preload, all localization resources are fetched when the plugin is loaded. When set to Lazy, localization resources are lazily loaded as and when they are required by the console. When omitted or set to the empty string, the behaviour is equivalent to Lazy type.",
}

func (ConsolePluginI18n) SwaggerDoc() map[string]string {
	return map_ConsolePluginI18n
}

var map_ConsolePluginList = map[string]string{
	"":         "Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsolePluginList) SwaggerDoc() map[string]string {
	return map_ConsolePluginList
}

var map_ConsolePluginProxy = map[string]string{
	"":              "ConsolePluginProxy holds information on various service types to which console's backend will proxy the plugin's requests.",
	"endpoint":      "endpoint provides information about endpoint to which the request is proxied to.",
	"alias":         "alias is a proxy name that identifies the plugin's proxy. An alias name should be unique per plugin. The console backend exposes following proxy endpoint:\n\n/api/proxy/plugin/<plugin-name>/<proxy-alias>/<request-path>?<optional-query-parameters>\n\nRequest example path:\n\n/api/proxy/plugin/acm/search/pods?namespace=openshift-apiserver",
	"caCertificate": "caCertificate provides the cert authority certificate contents, in case the proxied Service is using custom service CA. By default, the service CA bundle provided by the service-ca operator is used. ",
	"authorization": "authorization provides information about authorization type, which the proxied request should contain",
}

func (ConsolePluginProxy) SwaggerDoc() map[string]string {
	return map_ConsolePluginProxy
}

var map_ConsolePluginProxyEndpoint = map[string]string{
	"":        "ConsolePluginProxyEndpoint holds information about the endpoint to which request will be proxied to.",
	"type":    "type is the type of the console plugin's proxy. Currently only \"Service\" is supported.",
	"service": "service is an in-cluster Service that the plugin will connect to. The Service must use HTTPS. The console backend exposes an endpoint in order to proxy communication between the plugin and the Service. Note: service field is required for now, since currently only \"Service\" type is supported.",
}

func (ConsolePluginProxyEndpoint) SwaggerDoc() map[string]string {
	return map_ConsolePluginProxyEndpoint
}

var map_ConsolePluginProxyServiceConfig = map[string]string{
	"":          "ProxyTypeServiceConfig holds information on Service to which console's backend will proxy the plugin's requests.",
	"name":      "name of Service that the plugin needs to connect to.",
	"namespace": "namespace of Service that the plugin needs to connect to",
	"port":      "port on which the Service that the plugin needs to connect to is listening on.",
}

func (ConsolePluginProxyServiceConfig) SwaggerDoc() map[string]string {
	return map_ConsolePluginProxyServiceConfig
}

var map_ConsolePluginService = map[string]string{
	"":          "ConsolePluginService holds information on Service that is serving console dynamic plugin assets.",
	"name":      "name of Service that is serving the plugin assets.",
	"namespace": "namespace of Service that is serving the plugin assets.",
	"port":      "port on which the Service that is serving the plugin is listening to.",
	"basePath":  "basePath is the path to the plugin's assets. The primary asset it the manifest file called `plugin-manifest.json`, which is a JSON document that contains metadata about the plugin and the extensions.",
}

func (ConsolePluginService) SwaggerDoc() map[string]string {
	return map_ConsolePluginService
}

var map_ConsolePluginSpec = map[string]string{
	"":                      "ConsolePluginSpec is the desired plugin configuration.",
	"displayName":           "displayName is the display name of the plugin. The dispalyName should be between 1 and 128 characters.",
	"backend":               "backend holds the configuration of backend which is serving console's plugin .",
	"proxy":                 "proxy is a list of proxies that describe various service type to which the plugin needs to connect to.",
	"i18n":                  "i18n is the configuration of plugin's localization resources.",
	"contentSecurityPolicy": "contentSecurityPolicy is a list of Content-Security-Policy (CSP) directives for the plugin. Each directive specifies a list of values, appropriate for the given directive type, for example a list of remote endpoints for fetch directives such as ScriptSrc. Console web application uses CSP to detect and mitigate certain types of attacks, such as cross-site scripting (XSS) and data injection attacks. Dynamic plugins should specify this field if need to load assets from outside the cluster or if violation reports are observed. Dynamic plugins should always prefer loading their assets from within the cluster, either by vendoring them, or fetching from a cluster service. CSP violation reports can be viewed in the browser's console logs during development and testing of the plugin in the OpenShift web console. Available directive types are DefaultSrc, ScriptSrc, StyleSrc, ImgSrc, FontSrc, ObjectSrc and ConnectSrc. Each of the available directives may be defined only once in the list. The value 'self' is automatically included in all fetch directives by the OpenShift web console's backend. For more information about the CSP directives, see: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy\n\nThe OpenShift web console server aggregates the CSP directives and values across its own default values and all enabled ConsolePlugin CRs, merging them into a single policy string that is sent to the browser via `Content-Security-Policy` HTTP response header.\n\nExample:\n  ConsolePlugin A directives:\n    script-src: https://script1.com/, https://script2.com/\n    font-src: https://font1.com/\n\n  ConsolePlugin B directives:\n    script-src: https://script2.com/, https://script3.com/\n    font-src: https://font2.com/\n    img-src: https://img1.com/\n\n  Unified set of CSP directives, passed to the OpenShift web console server:\n    script-src: https://script1.com/, https://script2.com/, https://script3.com/\n    font-src: https://font1.com/, https://font2.com/\n    img-src: https://img1.com/\n\n  OpenShift web console server CSP response header:\n    Content-Security-Policy: default-src 'self'; base-uri 'self'; script-src 'self' https://script1.com/ https://script2.com/ https://script3.com/; font-src 'self' https://font1.com/ https://font2.com/; img-src 'self' https://img1.com/; style-src 'self'; frame-src 'none'; object-src 'none'",
}

func (ConsolePluginSpec) SwaggerDoc() map[string]string {
	return map_ConsolePluginSpec
}

var map_ConsoleQuickStart = map[string]string{
	"":         "ConsoleQuickStart is an extension for guiding user through various workflows in the OpenShift web console.\n\nCompatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsoleQuickStart) SwaggerDoc() map[string]string {
	return map_ConsoleQuickStart
}

var map_ConsoleQuickStartList = map[string]string{
	"":         "Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsoleQuickStartList) SwaggerDoc() map[string]string {
	return map_ConsoleQuickStartList
}

var map_ConsoleQuickStartSpec = map[string]string{
	"":                      "ConsoleQuickStartSpec is the desired quick start configuration.",
	"displayName":           "displayName is the display name of the Quick Start.",
	"icon":                  "icon is a base64 encoded image that will be displayed beside the Quick Start display name. The icon should be an vector image for easy scaling. The size of the icon should be 40x40.",
	"tags":                  "tags is a list of strings that describe the Quick Start.",
	"durationMinutes":       "durationMinutes describes approximately how many minutes it will take to complete the Quick Start.",
	"description":           "description is the description of the Quick Start. (includes markdown)",
	"prerequisites":         "prerequisites contains all prerequisites that need to be met before taking a Quick Start. (includes markdown)",
	"introduction":          "introduction describes the purpose of the Quick Start. (includes markdown)",
	"tasks":                 "tasks is the list of steps the user has to perform to complete the Quick Start.",
	"conclusion":            "conclusion sums up the Quick Start and suggests the possible next steps. (includes markdown)",
	"nextQuickStart":        "nextQuickStart is a list of the following Quick Starts, suggested for the user to try.",
	"accessReviewResources": "accessReviewResources contains a list of resources that the user's access will be reviewed against in order for the user to complete the Quick Start. The Quick Start will be hidden if any of the access reviews fail.",
}

func (ConsoleQuickStartSpec) SwaggerDoc() map[string]string {
	return map_ConsoleQuickStartSpec
}

var map_ConsoleQuickStartTask = map[string]string{
	"":            "ConsoleQuickStartTask is a single step in a Quick Start.",
	"title":       "title describes the task and is displayed as a step heading.",
	"description": "description describes the steps needed to complete the task. (includes markdown)",
	"review":      "review contains instructions to validate the task is complete. The user will select 'Yes' or 'No'. using a radio button, which indicates whether the step was completed successfully.",
	"summary":     "summary contains information about the passed step.",
}

func (ConsoleQuickStartTask) SwaggerDoc() map[string]string {
	return map_ConsoleQuickStartTask
}

var map_ConsoleQuickStartTaskReview = map[string]string{
	"":               "ConsoleQuickStartTaskReview contains instructions that validate a task was completed successfully.",
	"instructions":   "instructions contains steps that user needs to take in order to validate his work after going through a task. (includes markdown)",
	"failedTaskHelp": "failedTaskHelp contains suggestions for a failed task review and is shown at the end of task. (includes markdown)",
}

func (ConsoleQuickStartTaskReview) SwaggerDoc() map[string]string {
	return map_ConsoleQuickStartTaskReview
}

var map_ConsoleQuickStartTaskSummary = map[string]string{
	"":        "ConsoleQuickStartTaskSummary contains information about a passed step.",
	"success": "success describes the succesfully passed task.",
	"failed":  "failed briefly describes the unsuccessfully passed task. (includes markdown)",
}

func (ConsoleQuickStartTaskSummary) SwaggerDoc() map[string]string {
	return map_ConsoleQuickStartTaskSummary
}

var map_ConsoleSample = map[string]string{
	"":         "ConsoleSample is an extension to customizing OpenShift web console by adding samples.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec contains configuration for a console sample.",
}

func (ConsoleSample) SwaggerDoc() map[string]string {
	return map_ConsoleSample
}

var map_ConsoleSampleContainerImportSource = map[string]string{
	"":        "ConsoleSampleContainerImportSource let the user import a container image.",
	"image":   "reference to a container image that provides a HTTP service. The service must be exposed on the default port (8080) unless otherwise configured with the port field.\n\nSupported formats:\n  - <repository-name>/<image-name>\n  - docker.io/<repository-name>/<image-name>\n  - quay.io/<repository-name>/<image-name>\n  - quay.io/<repository-name>/<image-name>@sha256:<image hash>\n  - quay.io/<repository-name>/<image-name>:<tag>",
	"service": "service contains configuration for the Service resource created for this sample.",
}

func (ConsoleSampleContainerImportSource) SwaggerDoc() map[string]string {
	return map_ConsoleSampleContainerImportSource
}

var map_ConsoleSampleContainerImportSourceService = map[string]string{
	"":           "ConsoleSampleContainerImportSourceService let the samples author define defaults for the Service created for this sample.",
	"targetPort": "targetPort is the port that the service listens on for HTTP requests. This port will be used for Service and Route created for this sample. Port must be in the range 1 to 65535. Default port is 8080.",
}

func (ConsoleSampleContainerImportSourceService) SwaggerDoc() map[string]string {
	return map_ConsoleSampleContainerImportSourceService
}

var map_ConsoleSampleGitImportSource = map[string]string{
	"":           "ConsoleSampleGitImportSource let the user import code from a public Git repository.",
	"repository": "repository contains the reference to the actual Git repository.",
	"service":    "service contains configuration for the Service resource created for this sample.",
}

func (ConsoleSampleGitImportSource) SwaggerDoc() map[string]string {
	return map_ConsoleSampleGitImportSource
}

var map_ConsoleSampleGitImportSourceRepository = map[string]string{
	"":           "ConsoleSampleGitImportSourceRepository let the user import code from a public git repository.",
	"url":        "url of the Git repository that contains a HTTP service. The HTTP service must be exposed on the default port (8080) unless otherwise configured with the port field.\n\nOnly public repositories on GitHub, GitLab and Bitbucket are currently supported:\n\n  - https://github.com/<org>/<repository>\n  - https://gitlab.com/<org>/<repository>\n  - https://bitbucket.org/<org>/<repository>\n\nThe url must have a maximum length of 256 characters.",
	"revision":   "revision is the git revision at which to clone the git repository Can be used to clone a specific branch, tag or commit SHA. Must be at most 256 characters in length. When omitted the repository's default branch is used.",
	"contextDir": "contextDir is used to specify a directory within the repository to build the component. Must start with `/` and have a maximum length of 256 characters. When omitted, the default value is to build from the root of the repository.",
}

func (ConsoleSampleGitImportSourceRepository) SwaggerDoc() map[string]string {
	return map_ConsoleSampleGitImportSourceRepository
}

var map_ConsoleSampleGitImportSourceService = map[string]string{
	"":           "ConsoleSampleGitImportSourceService let the samples author define defaults for the Service created for this sample.",
	"targetPort": "targetPort is the port that the service listens on for HTTP requests. This port will be used for Service created for this sample. Port must be in the range 1 to 65535. Default port is 8080.",
}

func (ConsoleSampleGitImportSourceService) SwaggerDoc() map[string]string {
	return map_ConsoleSampleGitImportSourceService
}

var map_ConsoleSampleList = map[string]string{
	"":         "Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsoleSampleList) SwaggerDoc() map[string]string {
	return map_ConsoleSampleList
}

var map_ConsoleSampleSource = map[string]string{
	"":                "ConsoleSampleSource is the actual sample definition and can hold different sample types. Unsupported sample types will be ignored in the web console.",
	"type":            "type of the sample, currently supported: \"GitImport\";\"ContainerImport\"",
	"gitImport":       "gitImport allows the user to import code from a git repository.",
	"containerImport": "containerImport allows the user import a container image.",
}

func (ConsoleSampleSource) SwaggerDoc() map[string]string {
	return map_ConsoleSampleSource
}

var map_ConsoleSampleSpec = map[string]string{
	"":            "ConsoleSampleSpec is the desired sample for the web console. Samples will appear with their title, descriptions and a badge in a samples catalog.",
	"title":       "title is the display name of the sample.\n\nIt is required and must be no more than 50 characters in length.",
	"abstract":    "abstract is a short introduction to the sample.\n\nIt is required and must be no more than 100 characters in length.\n\nThe abstract is shown on the sample card tile below the title and provider and is limited to three lines of content.",
	"description": "description is a long form explanation of the sample.\n\nIt is required and can have a maximum length of **4096** characters.\n\nIt is a README.md-like content for additional information, links, pre-conditions, and other instructions. It will be rendered as Markdown so that it can contain line breaks, links, and other simple formatting.",
	"icon":        "icon is an optional base64 encoded image and shown beside the sample title.\n\nThe format must follow the data: URL format and can have a maximum size of **10 KB**.\n\n  data:[<mediatype>][;base64],<base64 encoded image>\n\nFor example:\n\n  data:image;base64,             plus the base64 encoded image.\n\nVector images can also be used. SVG icons must start with:\n\n  data:image/svg+xml;base64,     plus the base64 encoded SVG image.\n\nAll sample catalog icons will be shown on a white background (also when the dark theme is used). The web console ensures that different aspect ratios work correctly. Currently, the surface of the icon is at most 40x100px.\n\nFor more information on the data URL format, please visit https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URLs.",
	"type":        "type is an optional label to group multiple samples.\n\nIt is optional and must be no more than 20 characters in length.\n\nRecommendation is a singular term like \"Builder Image\", \"Devfile\" or \"Serverless Function\".\n\nCurrently, the type is shown a badge on the sample card tile in the top right corner.",
	"provider":    "provider is an optional label to honor who provides the sample.\n\nIt is optional and must be no more than 50 characters in length.\n\nA provider can be a company like \"Red Hat\" or an organization like \"CNCF\" or \"Knative\".\n\nCurrently, the provider is only shown on the sample card tile below the title with the prefix \"Provided by \"",
	"tags":        "tags are optional string values that can be used to find samples in the samples catalog.\n\nExamples of common tags may be \"Java\", \"Quarkus\", etc.\n\nThey will be displayed on the samples details page.",
	"source":      "source defines where to deploy the sample service from. The sample may be sourced from an external git repository or container image.",
}

func (ConsoleSampleSpec) SwaggerDoc() map[string]string {
	return map_ConsoleSampleSpec
}

var map_ConsoleYAMLSample = map[string]string{
	"":         "ConsoleYAMLSample is an extension for customizing OpenShift web console YAML samples.\n\nCompatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsoleYAMLSample) SwaggerDoc() map[string]string {
	return map_ConsoleYAMLSample
}

var map_ConsoleYAMLSampleList = map[string]string{
	"":         "Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
	"metadata": "metadata is the standard list's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
}

func (ConsoleYAMLSampleList) SwaggerDoc() map[string]string {
	return map_ConsoleYAMLSampleList
}

var map_ConsoleYAMLSampleSpec = map[string]string{
	"":               "ConsoleYAMLSampleSpec is the desired YAML sample configuration. Samples will appear with their descriptions in a samples sidebar when creating a resources in the web console.",
	"targetResource": "targetResource contains apiVersion and kind of the resource YAML sample is representating.",
	"title":          "title of the YAML sample.",
	"description":    "description of the YAML sample.",
	"yaml":           "yaml is the YAML sample to display.",
	"snippet":        "snippet indicates that the YAML sample is not the full YAML resource definition, but a fragment that can be inserted into the existing YAML document at the user's cursor.",
}

func (ConsoleYAMLSampleSpec) SwaggerDoc() map[string]string {
	return map_ConsoleYAMLSampleSpec
}

// AUTO-GENERATED FUNCTIONS END HERE
