//
// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//    ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗
//    ████╗  ██║██╔═████╗██║   ██║██╔══██╗
//    ██╔██╗ ██║██║██╔██║██║   ██║███████║
//    ██║╚██╗██║████╔╝██║╚██╗ ██╔╝██╔══██║
//    ██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║
//    ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝

package yamyams

import (
	"github.com/kris-nova/logger"
	"os"
)

var registry = make(map[string]Deployable)

func Register(app Deployable) {

	// Validate the application
	if app == nil {
		logger.Critical("Unable to register NIL application.")
		os.Exit(1)
	}

	if app.About().Name == "" {
		logger.Critical("Empty name for application.")
		os.Exit(1)
	}

	if app.About().Command == "" {
		logger.Critical("Empty command line name for application %s.", app.About().Name)
		os.Exit(1)
	}

	if app.About().Version == "" {
		logger.Critical("Empty version for application %s.", app.About().Name)
		os.Exit(1)
	}

	if app.About().Description == "" {
		logger.Critical("Empty description for application %s.", app.About().Name)
		os.Exit(1)
	}

	registry[app.About().Command] = app
}

func Registry() map[string]Deployable {
	return registry
}

func Find(name string) Deployable {
	if app, ok := registry[name]; ok {
		return app
	}
	return nil
}
