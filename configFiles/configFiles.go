package configFiles

var ConfigFiles = func(configFiles tree) tree {
	new(treeGeneration).
		javaScript(&configFiles).
		php(&configFiles).
		react15(&configFiles).
		react16(&configFiles).
		react17(&configFiles).
		react18(&configFiles).
		angular(&configFiles).
		vue2(&configFiles).
		vue3(&configFiles)
	return configFiles
}(tree{})
