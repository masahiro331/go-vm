module github.com/masahiro331/go-vm

go 1.18

require (
	github.com/hashicorp/go-multierror v1.1.1
	github.com/masahiro331/go-disk v0.0.0-20220220182714-f5f832728b22
	github.com/masahiro331/go-vmdk-parser v0.0.0-20220810134757-d12ec8b771de
	github.com/masahiro331/go-xfs-filesystem v0.0.0-20220810140610-8d15d3d3a8d6
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

require (
	github.com/hashicorp/errwrap v1.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
)

replace github.com/masahiro331/go-vmdk-parser => ../go-vmdk-parser
