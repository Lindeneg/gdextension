; Engine configuration file.
; It's best edited using the editor UI and not directly,
; since the parameters that go here are not all obvious.
;
; Format:
;   [section] ; section goes between []
;   param=value ; assign values to parameters

config_version=5

[application]

config/name="{{.ProjectName}}"
config/features=PackedStringArray("{{.MajorMinor}}", "Mobile")
config/icon="res://icon.svg"

[rendering]

renderer/rendering_method="mobile"
