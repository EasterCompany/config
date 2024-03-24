# Go Module: sh

Go Module "config" allows the developer to maintain a standardised configuration per user for their application.

## Documentation

Below you will find documentation on the usage of the module.

### SetFilePath

Set the configuration file path.

**Arguments**:

  - `filePath` `string` absolute path to file.

**Returns**: `string` containing the new file path

```go
config.SetFilePath("~/.config/configuration.json")
```

### VerifyFileExists

Verifies if the configuration already exists and is accessible to the program.

**Arguments**: None

**Returns**: `boolean`

  - `true` if the configuration already exists & is readable.
  - `false` if the configuration does not exist or is not readable.

```go
configExists := config.Verify()
```

### GenerateFileTemplate

Generates a template configuration based on the current user & system accessing the program.

**Arguments**: None

**Returns**: String containing default configuration template.

```go
configTemplate := config.Generate()
```

### SaveFile

Saves the configuration to storage.

**Arguments**: `content` `string`

**Returns**: None

```go
config.Save()
```
