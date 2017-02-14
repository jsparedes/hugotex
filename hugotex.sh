#!/bin/bash
rm -rf public/
hugo
go run multi.go
go run rmblankspace.go
cd public/
find . -type f -name "*.tex" -exec xelatex {} \;
find . -type f -name "*.aux" -exec bibtex {} \;
find . -type f -name "*.tex" -exec xelatex {} \;
