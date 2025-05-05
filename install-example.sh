THIS="$(readlink -f "$0")"
HERE="$(dirname "$THIS")"

cd "$HERE"

bash build.sh

## ----------------------------------------------------¬
# /!\ This section is not needed in usual scripts - this is here
#     because this present repo _is_ pathctl's own repo!

if [[ ! -f "$HOME/.PATH" ]]; then
    bin/pathctl help

    echo "== Install failed =="
    exit 1
fi
## ----------------------------------------------------/

## Actual usage example assuming pathctl is available - like in your own scripts
##

INSTALL_DIR="$(pathctl bin)"
cp bin/pathctl "$INSTALL_DIR/"
