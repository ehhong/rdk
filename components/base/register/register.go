// Package register registers all relevant bases
package register

import (
	// register bases.
	_ "go.viam.com/rdk/components/base/agilex"
	_ "go.viam.com/rdk/components/base/boat"
	_ "go.viam.com/rdk/components/base/fake"
	_ "go.viam.com/rdk/components/base/wheeled"
)