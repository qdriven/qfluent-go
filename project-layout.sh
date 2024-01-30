mkdir -p cmd api docs
# An internal directory showcases to anyone looking at your service that this code is private.
# External services should not import any packages within this directory.
mkdir -p pkg/database pkg/app pkg/base pkg/config #external
mkdir -p internal # internal
mkdir -p config
mkdir -p vendor
#touch Dockerfile
#touch Makefile_tpl
touch main.go