/**
* Copyright 2018 Comcast Cable Communications Management, LLC
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package influxdb

import (
	"encoding/json"
	"time"

	"github.com/Comcast/trickster/internal/timeseries"
	"github.com/influxdata/influxdb/models"
)

// SeriesEnvelope ...
type SeriesEnvelope struct {
	Results      []Result            `json:"results"`
	Err          string              `json:"error,omitempty"`
	StepDuration time.Duration       `json:"step,omitempty"`
	ExtentList   []timeseries.Extent `json:"extents,omitempty"`
}

// Result ...
type Result struct {
	StatementID int          `json:"statement_id"`
	Series      []models.Row `json:"series,omitempty"`
	Err         string       `json:"error,omitempty"`
}

// MarshalTimeseries ...
func (c Client) MarshalTimeseries(ts timeseries.Timeseries) ([]byte, error) {
	// Marshal the Envelope back to a json object for Cache Storage
	return json.Marshal(ts)
}

// UnmarshalTimeseries ...
func (c Client) UnmarshalTimeseries(data []byte) (timeseries.Timeseries, error) {
	se := &SeriesEnvelope{}
	err := json.Unmarshal(data, se)
	return se, err
}