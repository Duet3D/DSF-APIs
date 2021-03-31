package commands

// InstallPlugin is used to install or upgrade a plugin
type InstallPlugin struct {
	BaseCommand
	// Absolute file path to the plugin ZIP bundle
	PluginFile string
}

// NewInstallPlugin creates a new InstallPlugin instance for the given path
func NewInstallPlugin(pluginFile string) *InstallPlugin {
	return &InstallPlugin{
		BaseCommand: *NewBaseCommand("InstallPlugin"),
		PluginFile:  pluginFile,
	}
}

// PluginControl is used to start/stop/uninstall plugins
type PluginControl struct {
	BaseCommand
	// Plugin is the identifier of the plugin
	Plugin string
}

// NewStartPlugin creates a new start command for the given plugin
func NewStartPlugin(plugin string) *PluginControl {
	return &PluginControl{
		BaseCommand: *NewBaseCommand("StartPlugin"),
		Plugin:      plugin,
	}
}

// NewStopPlugin creates a new stop command for the given plugin
func NewStopPlugin(plugin string) *PluginControl {
	return &PluginControl{
		BaseCommand: *NewBaseCommand("StopPlugin"),
		Plugin:      plugin,
	}
}

// NewUninstallPlugin creates a new uninstall command for the given plugin
func NewUninstallPlugin(plugin string) *PluginControl {
	return &PluginControl{
		BaseCommand: *NewBaseCommand("UninstallPlugin"),
		Plugin:      plugin,
	}
}

// SetPluginData updates custom plugin data in the object model
// May be used to update only the own plugin data unless the plugin has the
// SbcPermissions.ManagePlugins permission.
// Note that the corresponding key must already exist in the plugin data!
type SetPluginData struct {
	BaseCommand
	// Plugin is the identifier of the plugin (optional)
	Plugin string
	// Key to set
	Key string
	// Value custom value to set
	Value string
}

// New SetPluginData creates a new command to set plugin data
func NewSetPluginData(plugin, key, value string) *SetPluginData {
	return &SetPluginData{
		BaseCommand: *NewBaseCommand("SetPluginData"),
		Plugin:      plugin,
		Key:         key,
		Value:       value,
	}
}

var startPlugins = NewBaseCommand("StartPlugins")

// NewStartPlugins starts all previously started plugins again
func NewStartPlugins() *BaseCommand {
	return startPlugins
}

var stopPlugins = NewBaseCommand("StopPlugins")

// NewStopPlugins returns a command to stop all plugins and save which plugins were running.
// This command is intended for shutdown or update requests
func NewStopPlugins() *BaseCommand {
	return stopPlugins
}
