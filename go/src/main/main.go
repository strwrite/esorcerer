package main

import (
	"flag"
	"os"
	"os/signal"
	"plugin"
	"strings"
	//"bytes"
	"log"
	"path/filepath"
)

type esorcererVersion struct {
	Githash, Buildstamp, API string
}

var (
	Buildstamp = ""
	Githash = ""
	API = "v1"
	Version esorcererVersion
	plugin_loadings map[string]string = make(map[string]string)
	plugin_roots = flag.String("plugin_roots",
		"plugins/dynamic",
		"Names of root directories where to locate the plugins to load, " +
			"as a comma separated list. " +
			"Note that only default plugins located in 'plugins/dynamic' folder could access 'flags'")
)

func main() {
	Version = esorcererVersion{Githash, Buildstamp, API}
	log.Printf("Version of esorcerer is %+v", Version)

	loadPluginsFromDir("plugins/dynamic")
	flag.Parse()

	log.Printf("Load custom plugins from: %s \n", *plugin_roots)
	pluginRoots := strings.Split(*plugin_roots, ",")
	for _, pluginRoot := range pluginRoots {
		loadPluginsFromDir(pluginRoot)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	MainLoop:
	for {
		select {
		case <-signals:
			break MainLoop
		}
	}

}
func loadPluginsFromDir(dir string) {
	_, loaded := plugin_loadings[dir]
	if loaded {
		log.Printf("Location %s has been already processed\n", dir)
		return
	}
	plugin_loadings[dir] = ""
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Panicf("Failed to load plugins from %s: %v", dir, err)
		}
		if info.IsDir() {
			log.Printf("Found directory in %s: %+v \n", dir, info.Name())
		} else {
			log.Printf("Found file in %s: %+v \n", dir, info.Name())
			if strings.HasSuffix(info.Name(), ".so") {
				pluginName := strings.TrimSuffix(info.Name(), ".so")
				log.Printf("Found plugin: %+v \n", pluginName)
				the_plugin := loadPlugin(path, pluginName)
				tryEventLoop(the_plugin, path, pluginName)
			}
		}
		return nil
	})

	if err != nil {
		log.Printf("Error looking for plugins in the path %q: %v\n", dir, err)
	}

}

func loadPlugin(pluginLocation string, pluginName string) *plugin.Plugin {

	log.Printf("Loading plugin %s from %s ...", pluginName, pluginLocation)

	the_plugin, err := plugin.Open(pluginLocation)
	if err != nil {
		panic(err)
	}

	log.Printf("Lookup initialization for %s", pluginName)
	initFunc, err := the_plugin.Lookup("Plugin_init")
	if err != nil {
		log.Panicf("Failed to load plugin %s from %s : %v", pluginName, pluginLocation, err)
	}

	log.Printf("Run plugin %s initialization", pluginName)
	initFunc.(func())()

	version, err := the_plugin.Lookup("PluginVersion")
	if err != nil {
		log.Panicf("Failed to get version of %s from %s ...", pluginName, pluginLocation)
	}
	log.Printf("Version of %s is %+v", pluginName, version)

	return the_plugin

}
var conf = `
declare: something
`
func tryEventLoop(the_plugin *plugin.Plugin, pluginLocation string, pluginName string) {
	log.Printf("Look for event loop in the plugin %s from %s", pluginName, pluginLocation)
	spawnEventLoopFunc, err := the_plugin.Lookup("Spawn_event_loop")
	if err != nil {
		log.Printf("Not found Spawn_event_loop in the plugin %s from %s : %v", pluginName, pluginLocation, err)
	} else {
		log.Printf("Run Spawn_event_loop in the plugin %s from %s", pluginName, pluginLocation)
		go doTheThing(func() {
			spawnEventLoopFunc.(func([]byte))([]byte(conf))
		})
	}

}

func doTheThing(thing func()) {
	thing()
}

