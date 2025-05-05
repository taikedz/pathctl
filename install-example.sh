# /!\ This section is not needed in usual scripts - this is here
#     because this present repo _is_ pathctl's own repo!
#     In fact, this is the kind of boilerplate we want to avoid,
#     like guessing where the user wants stuf installed... !!
if [[ ! -f "$HOME/.PATH" ]]; then
    BINDIR="$HOME/.local/bin" # making a bold assumption!
    mkdir -p "$BINDIR"
    echo "%bin=$BINDIR" >> "$HOME/.PATH"
    echo "$BINDIR" >> "$HOME/.PATH"
fi

bash build.sh
export PATH="$PATH:$PWD/bin"

## ----------------------------------------------------
## Actual usage example assuming pathctl is available - like in your own scripts
##

INSTALL_DIR="$(pathctl bin)"
cp bin/pathctl "$INSTALL_DIR/"

