// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package services

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/config"

	"github.com/percona/pmm-managed/utils/logger"
)

// Prometheus service is responsible for interactions with Prometheus.
// It assumes the following:
//   * Prometheus API is accessible;
//   * Prometheus configuration and rule files are accessible;
//   * promtool is available.
type Prometheus struct {
	ConfigPath     string
	URL            *url.URL
	PromtoolPath   string
	AlertRulesPath string
	lock           sync.RWMutex
}

// loadConfig returns current Prometheus configuration.
func (p *Prometheus) loadConfig() (*config.Config, error) {
	c, err := config.LoadFile(p.ConfigPath)
	if err != nil {
		return nil, errors.Wrap(err, "can't load Prometheus configuration file")
	}

	if len(c.RuleFiles) == 0 {
		return nil, errors.New("no RuleFiles patterns")
	}
	// TODO check p.AlertRulesPath is in c.RuleFiles patterns

	return c, nil
}

// reload causes Prometheus to reload configuration, including alert rules files.
func (p *Prometheus) reload() error {
	u := *p.URL
	u.Path = filepath.Join(u.Path, "-", "reload")
	resp, err := http.Post(u.String(), "", nil)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}
	return errors.Errorf("%d: %s", resp.StatusCode, b)
}

// loadAlertRules returns all Prometheus alert rules.
func (p *Prometheus) loadAlertRules(ctx context.Context) ([]AlertRule, error) {
	files, err := filepath.Glob(filepath.Join(p.AlertRulesPath, "*"))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	names := make(map[string]struct{})
	rules := make([]AlertRule, 0, len(files))
	for _, f := range files {
		// extract rule name and disabled status from filename
		var disabled bool
		base := filepath.Base(f)
		ext := filepath.Ext(base)
		name := strings.TrimSuffix(base, ext)
		if ext == ".disabled" {
			disabled = true
			ext = filepath.Ext(name)
			name = strings.TrimSuffix(name, ext)
		}
		if ext != ".rule" {
			logger.Get(ctx).Warnf("unexpected file %q, skipped", f)
			continue
		}
		if _, ok := names[name]; ok {
			return nil, errors.Errorf("duplicate alert rule name %q", name)
		}
		names[name] = struct{}{}

		// load file and make rule
		b, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		rule := AlertRule{
			Name:     name,
			FilePath: f,
			Text:     string(b),
			Disabled: disabled,
		}
		rules = append(rules, rule)
	}
	return rules, nil
}

// Check returns error if configuration is not right or Prometheus is not available.
func (p *Prometheus) Check(ctx context.Context) error {
	l := logger.Get(ctx)

	if _, err := p.loadConfig(); err != nil {
		return err
	}

	if p.URL == nil {
		return errors.New("URL is not set")
	}
	u := *p.URL
	u.Path = filepath.Join(u.Path, "version")
	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	l.Infof("Prometheus: %s", b)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.Errorf("expected 200, got %d", resp.StatusCode)
	}

	b, err = exec.Command(p.PromtoolPath, "version").CombinedOutput()
	if err != nil {
		return errors.Wrap(err, string(b))
	}
	l.Infof("%s", b)

	return nil
}
