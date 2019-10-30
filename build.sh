#!bin/sh
# read version file
version=$(<config/version.txt)
echo "current version: "$version
IFS='.-' read -r -a array <<< "$version"
if [ ${#array[@]} -eq 4 ]
then
    majorver=${array[0]}
    minorver=${array[1]}
    buildver=${array[2]}
    stagever=${array[3]}
else
    echo "failed to use this version"
    exit -1
fi
buildver=$(echo $buildver | sed 's/^0*//')
buildver=$(printf "%03d" $((buildver+1)))
version="${majorver}.${minorver}.${buildver}-${stagever}"
package="main"
time=$(date -u +.%Y%m%d.%H%M%S)
flags="-X '$package.BuildTime=$time' -X '$package.BuildVersion=$version'"
# build
IFS=' '
read -ra ARGS <<< "$@"
for arg in "${ARGS[@]}"; do
    GOOS=${arg} GOARCH=amd64 go build -o out/historder_${arg} -ldflags="$flags" cmd/historder/*
done
echo "new version: "$version
echo $version> "config/version.txt"
