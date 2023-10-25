package storage

import (
	"fmt"
)

var ErrNoRows = fmt.Errorf("db: no rows found")
var ErrNoUpdateData = fmt.Errorf("db: nothing to update")
