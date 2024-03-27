package configobserver

import (
	"fmt"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"

	"github.com/openshift/insights-operator/pkg/config"
)

// Config defines the configuration loaded from cluster secret
type Config struct {
	config.Controller
}

// MinDuration defines the minimal report interval
const MinDuration = 10 * time.Second

// LoadConfigFromSecret loads the controller config with given secret data
func LoadConfigFromSecret(secret *v1.Secret) (config.Controller, error) {
	var cfg Config
	var err error

	cfg.loadEndpoint(secret.Data)
	cfg.loadConditionalGathererEndpoint(secret.Data)
	cfg.loadHTTP(secret.Data)
	cfg.loadReport(secret.Data)
	cfg.loadOCM(secret.Data)
	cfg.loadProcessingStatusEndpoint(secret.Data)
	cfg.loadReportEndpointTechPreview(secret.Data)

	if intervalString, ok := secret.Data["interval"]; ok {
		var duration time.Duration
		duration, err = time.ParseDuration(string(intervalString))
		if err == nil && duration < MinDuration {
			err = fmt.Errorf("too short")
		}
		if err == nil {
			cfg.Interval = duration
		} else {
			err = fmt.Errorf(
				"insights secret interval must be a duration (1h, 10m) greater than or equal to ten seconds: %v",
				err)
			cfg.Report = false
		}
	}

	return cfg.Controller, err
}

func (c *Config) loadEndpoint(data map[string][]byte) {
	if endpoint, ok := data["endpoint"]; ok {
		c.Endpoint = string(endpoint)
	}
}

func (c *Config) loadConditionalGathererEndpoint(data map[string][]byte) {
	if endpoint, ok := data["conditionalGathererEndpoint"]; ok {
		c.ConditionalGathererEndpoint = string(endpoint)
	}
}

func (c *Config) loadHTTP(data map[string][]byte) {
	if httpProxy, ok := data["httpProxy"]; ok {
		c.HTTPConfig.HTTPProxy = string(httpProxy)
	}
	if httpsProxy, ok := data["httpsProxy"]; ok {
		c.HTTPConfig.HTTPSProxy = string(httpsProxy)
	}
	if noProxy, ok := data["noProxy"]; ok {
		c.HTTPConfig.NoProxy = string(noProxy)
	}
}

func (c *Config) loadReport(data map[string][]byte) {
	c.loadEnableGlobalObfuscation(data)
	c.loadReportEndpoint(data)
	c.loadReportPullingDelay(data)
	c.loadReportPullingTimeout(data)
	c.loadReportMinRetryTime(data)
	c.loadReportSettings(data)
}

func (c *Config) loadOCM(data map[string][]byte) {
	if scaEndpoint, ok := data["scaEndpoint"]; ok {
		c.OCMConfig.SCAEndpoint = string(scaEndpoint)
	}
	if scaInterval, ok := data["scaInterval"]; ok {
		if newInterval, err := time.ParseDuration(string(scaInterval)); err == nil {
			c.OCMConfig.SCAInterval = newInterval
		} else {
			klog.Warningf(
				"secret contains an invalid value (%s) for scaInterval. Using previous value",
				scaInterval,
			)
		}
	}
	if scaDisabled, ok := data["scaPullDisabled"]; ok {
		c.OCMConfig.SCADisabled = strings.EqualFold(string(scaDisabled), "true")
	}

	if clusterTransferEndpoint, ok := data["clusterTransferEndpoint"]; ok {
		c.OCMConfig.ClusterTransferEndpoint = string(clusterTransferEndpoint)
	}
	if clusterTransferInterval, ok := data["clusterTransferInterval"]; ok {
		if newInterval, err := time.ParseDuration(string(clusterTransferInterval)); err == nil {
			c.OCMConfig.ClusterTransferInterval = newInterval
		} else {
			klog.Warningf(
				"secret contains an invalid value (%s) for clusterTransferInterval. Using previous value",
				clusterTransferInterval,
			)
		}
	}
}

func (c *Config) loadProcessingStatusEndpoint(data map[string][]byte) {
	if endpoint, ok := data["processingStatusEndpoint"]; ok {
		c.ProcessingStatusEndpoint = string(endpoint)
	}
}

func (c *Config) loadReportEndpointTechPreview(data map[string][]byte) {
	if endpoint, ok := data["reportEndpointTechPreview"]; ok {
		c.ReportEndpointTechPreview = string(endpoint)
	}
}

func (c *Config) loadEnableGlobalObfuscation(data map[string][]byte) {
	if enableGlobalObfuscation, ok := data["enableGlobalObfuscation"]; ok {
		c.EnableGlobalObfuscation = strings.EqualFold(string(enableGlobalObfuscation), "true")
	}
}

func (c *Config) loadReportEndpoint(data map[string][]byte) {
	if reportEndpoint, ok := data["reportEndpoint"]; ok {
		c.ReportEndpoint = string(reportEndpoint)
	}
}

func (c *Config) loadReportPullingDelay(data map[string][]byte) {
	c.loadReportDurationSetting(data, "reportPullingDelay", &c.ReportPullingDelay)
}

func (c *Config) loadReportPullingTimeout(data map[string][]byte) {
	c.loadReportDurationSetting(data, "reportPullingTimeout", &c.ReportPullingTimeout)
}

func (c *Config) loadReportMinRetryTime(data map[string][]byte) {
	c.loadReportDurationSetting(data, "reportMinRetryTime", &c.ReportMinRetryTime)
}

func (c *Config) loadReportDurationSetting(data map[string][]byte, key string, target *time.Duration) {
	if value, ok := data[key]; ok {
		if v, err := time.ParseDuration(string(value)); err == nil {
			*target = v
		} else {
			klog.Warningf("%s secret contains an invalid value (%s). Using previous value", key, value)
		}
	} else {
		if key == "reportPullingDelay" {
			*target = time.Duration(-1)
		}
	}
}

func (c *Config) loadReportSettings(data map[string][]byte) {
	c.Report = len(c.Endpoint) > 0
	if disableInsightsAlerts, ok := data["disableInsightsAlerts"]; ok {
		c.DisableInsightsAlerts = strings.EqualFold(string(disableInsightsAlerts), "true")
	}
}
