[configuration]

entry_symbol = "{{.ProjectName}}_extension_library_init"
compatibility_minimum = "{{.GodotVersion}}"
reloadable = true

[libraries]

macos.debug = "res://bin/{{.ProjectName}}.macos.template_debug.framework"
macos.release = "res://bin/{{.ProjectName}}.macos.template_release.framework"
ios.debug = "res://bin/{{.ProjectName}}.ios.template_debug.xcframework"
ios.release = "res://bin/{{.ProjectName}}.ios.template_release.xcframework"
windows.debug.x86_32 = "res://bin/{{.ProjectName}}.windows.template_debug.x86_32.dll"
windows.release.x86_32 = "res://bin/{{.ProjectName}}.windows.template_release.x86_32.dll"
windows.debug.x86_64 = "res://bin/{{.ProjectName}}.windows.template_debug.x86_64.dll"
windows.release.x86_64 = "res://bin/{{.ProjectName}}.windows.template_release.x86_64.dll"
linux.debug.x86_64 = "res://bin/{{.ProjectName}}.linux.template_debug.x86_64.so"
linux.release.x86_64 = "res://bin/{{.ProjectName}}.linux.template_release.x86_64.so"
linux.debug.arm64 = "res://bin/{{.ProjectName}}.linux.template_debug.arm64.so"
linux.release.arm64 = "res://bin/{{.ProjectName}}.linux.template_release.arm64.so"
linux.debug.rv64 = "res://bin/{{.ProjectName}}.linux.template_debug.rv64.so"
linux.release.rv64 = "res://bin/{{.ProjectName}}.linux.template_release.rv64.so"
android.debug.x86_64 = "res://bin/{{.ProjectName}}.android.template_debug.x86_64.so"
android.release.x86_64 = "res://bin/{{.ProjectName}}.android.template_release.x86_64.so"
android.debug.arm64 = "res://bin/{{.ProjectName}}.android.template_debug.arm64.so"
android.release.arm64 = "res://bin/{{.ProjectName}}.android.template_release.arm64.so"

[dependencies]
ios.debug = {
    "res://bin/libgodot-cpp.ios.template_debug.xcframework": ""
}
ios.release = {
    "res://bin/libgodot-cpp.ios.template_release.xcframework": ""
}

