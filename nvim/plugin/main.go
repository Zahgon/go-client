// Package plugin is a Nvim remote plugin host.
package plugin

// Main implements the main function for a Nvim remote plugin.
//
// Plugin applications call the Main function to run the plugin. The Main
// function creates a Nvim client, calls the supplied function to register
// handlers with the plugin and then runs the server loop to handle requests
// from Nvim.
//
// Applications should use the default logger in the standard log package to
// write to Nvim's log.
//
// Run the plugin application with the command line option --manifest=hostName
// to print the plugin manifest to stdout. Add the manifest manually to a
// Vimscript file. The :UpdateRemotePlugins command is not supported at this
// time.
//
// If the --manifest=host command line flag is specified, then Main prints the
// plugin manifest to stdout insead of running the application as a plugin.
// If the --location=vimfile command line flag is specified, then plugin
// manifest will be automatically written to .vim file.
func Main(registerHandlers func(p *Plugin) error) { _ = "STUB: not implemented"; return }

func overwriteManifest(path, host string, manifest []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func replaceManifest(host string, input, manifest []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// No need for trailing \n if in middle of file.
