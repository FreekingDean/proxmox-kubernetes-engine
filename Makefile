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

$(BIN)/proxmox-kubernetes-engine: $(ALL_PROTO_GENERATED) $(GOSOURCES) cmd/root.go
	@ go build -o bin/proxmox-kubernetes-engine

$(ALL_PROTO_GENERATED): $(PROTOSOURCES)
	@ rm -rf gen/go
	@ buf generate $(PROTODIR) --template $(PROTODIR)/buf.gen.yaml

bin/goose: cmd/goose/main.go
	@ go build -o bin/goose ./cmd/goose

.PHONY: test
test:
	@ go test ./...

.PHONY: clean
clean:
	@ echo "RUNNING CLEAN"
	@ rm -rf gen
	@ rm -rf bin

.PHONY: debug
tools: $(BIN)/goose

.PHONY: run
run: $(BIN)/proxmox-kubernetes-engine
	@ bash -c 'set -o allexport; source .env; set +o allexport; proxmox-kubernetes-engine'

.PHONY: debug
debug:
	@ echo protodir: $(PROTODIR)
	@ echo protos: $(PROTOSOURCES)
	@ echo protogo: $(PROTOGO_GENERATED)
	@ echo protogrpc: $(PROTOGRPC_GENERATED)
	@ echo protogw: $(PROTOGW_GENERATED)
	@ echo protoaip: $(PROTOAIP_GENERATED)
	@ echo allgen: $(ALL_PROTO_GENERATED)
