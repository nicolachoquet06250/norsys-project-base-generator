package configFiles

var ConfigFiles = func(configFiles tree) tree {
	new(treeGeneration).
		javaScript(&configFiles).
		react15(&configFiles).
		react16(&configFiles).
		react17(&configFiles).
		react18(&configFiles).
		vue2(&configFiles).
		vue3(&configFiles).
		angular(&configFiles).
		php(&configFiles)

	return configFiles
}(tree{})
