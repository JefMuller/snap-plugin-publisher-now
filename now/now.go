/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package file

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"bytes"
	"encoding/json"
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

// OIMetric London metric structure.
type OIMetric struct {
	Metric    string                 `json:"metric_type"`
	Resource  string                 `json:"resource"`
	Node      string                 `json:"node"`
	Value     interface{}            `json:"value"`
	Timestamp int64                  `json:"timestamp"`
	CiMapping map[string]interface{} `json:"ci2metric_id"`
	Source    string                 `json:"source"`
}

/*
	GetConfigPolicy() returns the configPolicy for your plugin.
*/
func (f OIMetric) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	return *policy, nil
}

// Publish test publish function
func (f OIMetric) Publish(mts []plugin.Metric, cfg plugin.Config) error {
	
	if len(mts) > 0 {
		log.Println("Source")
		f.Source = "snap"

		log.Println("Determining hostname/nodename")
		hostname, err = os.Hostname()
		if err != nil {
			hostname = "localhost"
		}
		f.Node = hostname
	}

	// Iterate over the supplied metrics
	for _, m := range mts {

		log.Println("Setting Timestamp")
		f.Timestamp = m.Timestamp

		/*
		// Metric contains all info related to a Snap Metric
		type Metric struct {
			Namespace   Namespace
			Version     int64
			Config      Config
			Data        interface{}
			Tags        map[string]string
			Timestamp   time.Time
			Unit        string
			Description string
			//Unexported but passed through for legacy reasons
			lastAdvertisedTime time.Time
		}
		*/

		// Do some type conversion and send the data
		switch v := m.Data.(type) {
		case uint:
			s.sendIntValue(int64(v))
		case uint32:
			s.sendIntValue(int64(v))
		case uint64:
			s.sendIntValue(int64(v))
		case int:
			s.sendIntValue(int64(v))
		case int32:
			s.sendIntValue(int64(v))
		case int64:
			s.sendIntValue(int64(v))
		case float32:
			s.sendFloatValue(float64(v))
		case float64:
			s.sendFloatValue(float64(v))
		default:
			log.Printf("Ignoring %T: %v\n", v, v)
			log.Printf("Contact the plugin author if you think this is an error")
		}
	}
	
	/*
	file, err := cfg.GetString("file")
	if err != nil {
		return err
	}
	if val, err := cfg.GetBool("return_error"); err == nil && val {
		return errors.New("Houston we have a problem")
	}
	fileHandle, _ := os.Create(file)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	for _, m := range mts {
		fmt.Fprintf(writer, "%s|%v|%d|%s|%s|%s|%v|%v\n",
			m.Namespace.Strings(),
			m.Data,
			m.Version,
			m.Unit,
			m.Description,
			m.Timestamp,
			m.Tags,
			m.Config,
		)
	}
	writer.Flush()
	*/
	return nil
}