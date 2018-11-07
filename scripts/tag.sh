# Usage
#   bash scripts/tag.sh 0.3.2

if [ $# -gt 1 ]; then
  echo "too many arguments" > /dev/stderr
  echo 'Usage tag.sh $TAG' > /dev/stderr
  exit 1
fi

if [ $# -lt 1 ]; then
  echo "TAG argument is required" > /dev/stderr
  echo 'Usage tag.sh $TAG' > /dev/stderr
  exit 1
fi

echo "cd `dirname $0`/.." || exit 1
cd `dirname $0`/.. || exit 1

TAG=$1 || exit 1
echo "TAG: $TAG" || exit 1
echo "create internal/domain/version.go" || exit 1
cat << EOS > internal/domain/version.go
package domain

// Version is the git-rm-branch's version.
const Version = "$TAG"
EOS

git add internal/domain/version.go || exit 1
git commit -m "build: update version to $TAG" || exit 1
echo "git tag $TAG" || exit 1
git tag $TAG || exit 1

# update CHANGELOG
git push --tags || exit 1
git fetch --all --prune || exit 1
npm run changelog || exit 1
git add CHANGELOG.md || exit 1
git commit -m "docs: update CHANGELOG" || exit 1
git push origin master || exit 1
