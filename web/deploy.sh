set -e

yarn add yarn build

yarn build

cd dist

git init

git config user.name ${GITHUB_USER}

git add -A
git commit -m 'deploy'
git push -f https://${GITHUB_USER}:${GITHUB_TOKEN}@github.com/ahsayde/yapl-playground master:gh-pages
