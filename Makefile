export ROOT_MOD=github.com/beatpika/eshop

.PHONY: gen-getway-user
gen-getway-user:
	@cd app/api && cwgo server -I ../../idl --type HTTP --service api --module ${ROOT_MOD}/app/api --idl ../../idl/api/handler_user.proto

.PHONY: gen-user
gen-user: 
	@cd rpc_gen && cwgo client --type RPC --service user --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module ${ROOT_MOD}/app/user --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/user.proto