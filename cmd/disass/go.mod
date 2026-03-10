module github.com/blacktop/arm64-cgo/cmd/disass

go 1.26

require (
	github.com/AlecAivazis/survey/v2 v2.3.7
	github.com/apex/log v1.9.0
	github.com/blacktop/arm64-cgo v1.0.70
	github.com/blacktop/go-macho v1.1.263
	github.com/spf13/cobra v1.10.2
	github.com/spf13/viper v1.21.0
)

replace github.com/blacktop/arm64-cgo => ../..

require (
	github.com/blacktop/go-dwarf v1.0.14 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sagikazarmark/locafero v0.11.0 // indirect
	github.com/sourcegraph/conc v0.3.1-0.20240121214520-5f936abd7ae8 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/term v0.27.0 // indirect
	golang.org/x/text v0.28.0 // indirect
)
