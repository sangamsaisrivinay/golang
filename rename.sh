#!/bin/bash

# Root directory (can be passed as an argument)
ROOT_DIR="/Users/sangamsaisrivinayreddy/golang"

# Find all .md files except ones already named README.md
find "$ROOT_DIR" -type f -name "*.md" ! -iname "README.md" | while read -r file; do
    dir=$(dirname "$file")
    new_path="$dir/README.md"

    # If README.md already exists in this folder, skip to avoid overwrite
    if [ -e "$new_path" ]; then
        echo "Skipping '$file' (README.md already exists in '$dir')"
    else
        echo "Renaming '$file' to '$new_path'"
        mv "$file" "$new_path"
    fi
done
