export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

os-archs=darwin:amd64 darwin:arm64 freebsd:386 freebsd:amd64 linux:386 linux:amd64 linux:arm linux:arm64 windows:386 windows:amd64 linux:mips64 linux:mips64le linux:mips:softfloat linux:mipsle:softfloat linux:riscv64

all: build

build: app

app:
	@$(foreach n, $(os-archs),\
		os=$(shell echo "$(n)" | cut -d : -f 1);\
		arch=$(shell echo "$(n)" | cut -d : -f 2);\
		gomips=$(shell echo "$(n)" | cut -d : -f 3);\
		name="spug-cli"\
		target_suffix=$${os}_$${arch};\
		echo "Build $${os}-$${arch}...";\
		env CGO_ENABLED=0 GOOS=$${os} GOARCH=$${arch} GOMIPS=$${gomips} go build -trimpath -ldflags "$(LDFLAGS)" -o ./release/$${name}_$${target_suffix} .;\
		echo "Build $${os}-$${arch} done";\
	)
	@mv ./release/spug-cli_windows_386 ./release/spug-cli_windows_386.exe
	@mv ./release/spug-cli_windows_amd64 ./release/spug-cli_windows_amd64.exe
