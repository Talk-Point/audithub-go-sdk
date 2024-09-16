# AuditHub Go SDK

The audit package provides a simple way to log audit entries in a structured format. It supports creating entries with various fields such as service name, event type, user information, and metadata. The package is designed to be flexible and extensible, making it suitable for various auditing needs.

## Features

- Create structured audit logs with a consistent format.
- Chainable methods for easy and intuitive API usage.
- Supports adding GIDs, labels, user information, and metadata.
- Environment awareness with dynamic environment variable reading.

## Installation

To install the package, run:

```bash
go get github.com/Talk-Point/audithub-go-sdk
```

## Usage

Here is an example of how to use the audit package:

```go
package main

import (
	"github.com/Talk-Point/audithub-go-sdk/pkg/v1/audit"
)

func main() {
	audit.AuditLog("de.talk-point.sagehub", "create").
		AddGids([]string{"a", "b"}).
		AddLabels([]string{"label-1", "label-2"}).
		ByUser("test").
		AddMetadata("key1", "value1").
		AddMetadata("key2", "value2").
		Log()
}
```

__Example Output:__  

When running the above code, you will see an output similar to:

```
::{"env":"test","timestamp":1631518212,"service":"de.talk-point.sagehub","event":"create","gids":["a","b"],"labels":["label-1","label-2"],"by_user":"test","metadata":{"key1":"value1","key2":"value2"}}::
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

Make sure to replace github.com/Talk-Point/audithub-go-sdk with the actual path to your repository, and add any additional details or sections that are relevant to your package.