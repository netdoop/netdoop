set -e

export BUILD=$(git rev-list HEAD -n 1 | cut -c 1-8 | tr "a-z" "A-Z")
echo "version:$1 build:$BUILD"
sed -i "s/\(\s*\)Version\(\s*\)=\s*\".*\"/\1Version\2= \"$1\"/g" server/global/global.go
sed -i "s/\(\s*\)Build\(\s*\)=\s*\".*\"/\1Build\2= \"$BUILD\"/g" server/global/global.go


