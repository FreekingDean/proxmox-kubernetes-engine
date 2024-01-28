SHELL=/bin/bash
BIN=bin
PROTODIR=protos
GENDIR=gen/go
GODIR=internal
CMDDIR=cmd
PROTOSOURCES=$(shell find $(PROTODIR) -name "*.proto")
GOSOURCES=$(shell find $(GODIR) -name "*.go") $(shell find $(CMDDIR) -name "*.go")

PROTOGO_GENERATED=$(patsubst $(PROTODIR)/%.proto,$(GENDIR)/%.pb.go, $(PROTOSOURCES))
PROTOGRPC_GENERATED=$(patsubst $(PROTODIR)/%.proto,$(GENDIR)/%_grpc.pb.go, $(PROTOSOURCES))
PROTOGW_GENERATED=$(patsubst $(PROTODIR)/%.proto,$(GENDIR)/%.pb.gw.go, $(PROTOSOURCES))
PROTOAIP_GENERATED=$(patsubst $(PROTODIR)/%.proto,$(GENDIR)/%_aip.go, $(PROTOSOURCES))
ALL_PROTO_GENERATED=$(PROTOGO_GENERATED) $(PROTOGRPC_GENERATED) $(PROTOGW_GENERATED) $(PROTOAIP_GENERATED)

PROTOGO_GENERATED_CLI=$(patsubst $(PROTODIR)/%.proto,cmd/pke/%.pb.go, $(PROTOSOURCES))
PROTOGOCLI_GENERATED_CLI=$(patsubst $(PROTODIR)/%.proto,cmd/pke/%_cli.pb.go, $(PROTOSOURCES))
ALL_PROTO_GENERATED_CLI=$(PROTOGO_GENERATED_CLI) $(PROTOGOCLI_GENERATED_CLI) cmd/pke/root.go

$(BIN)/proxmox-kubernetes-engine: $(ALL_PROTO_GENERATED) $(GOSOURCES) cmd/root.go go.sum
	go build -o bin/proxmox-kubernetes-engine

go.sum: go.mod $(GOSOURCES)
	go mod tidy

$(ALL_PROTO_GENERATED): $(PROTOSOURCES)
	rm -rf gen/go
	buf generate $(PROTODIR) --template $(PROTODIR)/buf.gen.yaml

$(ALL_PROTO_GENERATED_CLI): $(PROTOSOURCES)
	rm cmd/pke/root.go
	rm -rf cmd/pke/proxmox_kubernetes_engine/
	buf generate $(PROTODIR) --template $(PROTODIR)/buf.gen.cli.yaml

bin/goose: cmd/goose/main.go
	go build -o bin/goose ./cmd/goose

bin/pke: cmd/pke/main.go $(ALL_PROTO_GENERATED_CLI)
	go build -o bin/pke ./cmd/pke

.PHONY: test
test:
	@ go test ./...

.PHONY: clean
clean:
	@ echo "RUNNING CLEAN"
	rm -rf gen
	rm -rf bin

.PHONY: tools
tools: $(BIN)/goose

.PHONY: lint
lint:
	@ sh api-linter.sh

.PHONY: run.server
run.server: $(BIN)/proxmox-kubernetes-engine
	@ bash -c 'set -o allexport; source .env; set +o allexport; proxmox-kubernetes-engine server'

run.bgs: $(BIN)/proxmox-kubernetes-engine
	@ bash -c 'set -o allexport; source .env; set +o allexport; proxmox-kubernetes-engine bgs'

.PHONY: debug
debug:
	@ echo protodir: $(PROTODIR)
	@ echo
	@ echo protos: $(PROTOSOURCES)
	@ echo
	@ echo protogo: $(PROTOGO_GENERATED)
	@ echo
	@ echo protogrpc: $(PROTOGRPC_GENERATED)
	@ echo
	@ echo protogw: $(PROTOGW_GENERATED)
	@ echo
	@ echo protoaip: $(PROTOAIP_GENERATED)
	@ echo
	@ echo allgen: $(ALL_PROTO_GENERATED)
	@ echo
	@ echo protogocli: $(PROTOGO_GENERATED_CLI)
	@ echo
	@ echo protoaipcli: $(PROTOGOCLI_GENERATED_CLI)
	@ echo
	@ echo allgencli: $(ALL_PROTO_GENERATED_CLI)
