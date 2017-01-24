#!/bin/bash
rm -rf public/
hugo
mv public/index.html public/index.tex
mv public/es/index.html public/index_es.tex
mv public/pt/index.html public/index_pt.tex
go run rmblankspace.go
cd public/
# xelatex index.tex
# xelatex index_es.tex
# xelatex index_pt.tex
find . -type f -name "*.tex" -exec xelatex {} \;
