#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR

echo "--> Updating Github Gitignore Template"
rm -rf gitignore-master
if [ ! -e gitignore-master.zip ]; then
	curl -s -o gitignore-master.zip https://codeload.github.com/github/gitignore/zip/master
fi
unzip gitignore-master.zip 2>&1 1>/dev/null
echo "--> Done"
echo "--> Generating Go Codes & Packages"
rm -rf template/all template/github && mkdir -p template/all && mkdir -p template/github

PackageFile="template/all/template.go"
cat > $PackageFile << EOF
package template
import (
EOF
find gitignore-master -name '*.gitignore' | while read ignore; do
	DirName=$(dirname $ignore)
	FileName=$(basename $ignore)
	TemplateDirName=$(echo $DirName | sed "s/^gitignore-master/template\/github/")
	GoFileName="${FileName%.*}"
	GoFileName=$(echo $GoFileName | sed "s/+/p/g")
	GoFileName=$(echo $GoFileName | sed "s/-/_/g")
	KeyName=$(echo $DirName | sed "s/^gitignore-master/GitHub/")
	KeyName="$KeyName/$GoFileName"
	TemplateDirName="$TemplateDirName/$GoFileName"
	ImportPath="$TemplateDirName"
	mkdir -p "$TemplateDirName"
	echo "package $(basename $ImportPath)" > "$TemplateDirName/${GoFileName}.go"
	echo >> "$TemplateDirName/${GoFileName}.go"
	echo "import \"strings\"" >> "$TemplateDirName/${GoFileName}.go"
	echo "import \"github.com/imwithye/git_ignore/template\"" >> "$TemplateDirName/${GoFileName}.go"
	echo >> "$TemplateDirName/${GoFileName}.go"
	echo "func init () {" >> "$TemplateDirName/${GoFileName}.go"
	echo "	ignore := []string{" >> "$TemplateDirName/${GoFileName}.go"
	while IFS='' read -r line || [[ -n "$line" ]]; do
		line=$(echo $line | sed 's/\\/\\\\/g')
		line=$(echo $line | sed 's/\"/\\"/g')
		echo "		\"${line}\"," >> "$TemplateDirName/${GoFileName}.go"
	done < $ignore
	echo "	}" >> "$TemplateDirName/${GoFileName}.go"
	echo "	template.SetTemplate(\"$KeyName\", strings.Join(ignore, \"\\n\"))" >> "$TemplateDirName/${GoFileName}.go"
	echo "}" >> "$TemplateDirName/${GoFileName}.go"
	echo "	_ \"github.com/imwithye/git_ignore/$ImportPath\"" >> $PackageFile
done
echo ")" >> $PackageFile
echo "--> Done"

cd $PWD
