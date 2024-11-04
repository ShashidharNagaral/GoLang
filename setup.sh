#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <folder_name> <module_name>"
    exit 1
fi

FOLDER_NAME=$1
MODULE_NAME=$2

echo $FOLDER_NAME
echo $MODULE_NAME

mkdir -p "$FOLDER_NAME"

cat <<EOF > "$FOLDER_NAME/main.go"
package main

import "fmt"

func main() {
    
}
EOF

cd "$FOLDER_NAME" || exit

go mod init "$MODULE_NAME"

echo "Go module '$MODULE_NAME' initialized in folder '$FOLDER_NAME'."