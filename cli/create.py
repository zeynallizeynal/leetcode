#!/usr/bin/env python3

import argparse
import os
import sys

def main():
    parser = argparse.ArgumentParser(description="Create a leetcode file")
    parser.add_argument("name", help="filename to create")
    parser.add_argument("-f", "--force", action="store_true", help="overwrite file if exists")
    parser.add_argument("-t", "--type", choices=["easy", "medium", "hard"], default="easy", help="difficulty type")
    parser.add_argument("-e", "--extension", choices=["py", "js"], default="py", help="select extension for the file. default is .py")
    
    args = parser.parse_args()
    
    filename: str = os.path.basename(args.name)
    filename = filename.lower().replace(" ", "_").replace("-", "_").replace(".", "") + f".{args.extension}"
    filepath = os.path.join(os.getcwd(), args.type, filename)
    
    if os.path.exists(filepath) and not args.force:
        print("file already exists. -f to overwrite", file=sys.stderr)
        sys.exit(1)
    
    with open(filepath, "w") as f:
        f.write("")
    
    print(f"created: {filepath}")

if __name__ == "__main__":
    main()
    