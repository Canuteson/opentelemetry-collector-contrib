// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testbed // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter/loggingexporter"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/exporter/otlphttpexporter"
	"go.opentelemetry.io/collector/extension/ballastextension"
	"go.opentelemetry.io/collector/extension/zpagesextension"
	"go.opentelemetry.io/collector/processor/batchprocessor"
	"go.opentelemetry.io/collector/processor/memorylimiterprocessor"
	"go.opentelemetry.io/collector/receiver/otlpreceiver"
	"go.uber.org/multierr"
)

// Components returns the set of components for tests
func Components() (
	component.Factories,
	error,
) {
	var errs error

	extensions, err := component.MakeExtensionFactoryMap(
		zpagesextension.NewFactory(),
		ballastextension.NewFactory(),
	)
	errs = multierr.Append(errs, err)

	receivers, err := component.MakeReceiverFactoryMap(
		otlpreceiver.NewFactory(),
	)
	errs = multierr.Append(errs, err)

	exporters, err := component.MakeExporterFactoryMap(
		loggingexporter.NewFactory(),
		otlpexporter.NewFactory(),
		otlphttpexporter.NewFactory(),
	)
	errs = multierr.Append(errs, err)

	processors, err := component.MakeProcessorFactoryMap(
		batchprocessor.NewFactory(),
		memorylimiterprocessor.NewFactory(),
	)
	errs = multierr.Append(errs, err)

	factories := component.Factories{
		Extensions: extensions,
		Receivers:  receivers,
		Processors: processors,
		Exporters:  exporters,
	}

	return factories, errs
}
