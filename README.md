# hugo-deploy
### Easily deploy your Hugo sites to Github Pages

-------------


# Install

Run: `go get github.com/danielbarbarito/hugo-deploy`

# Project Layout
Your source code for your Hugo project should all be pushed to the `source` branch. The "public" folder will be pushed to the `master` branch.
To achieve this, do the following:
- `cd` to the root of your project
- `echo public/ >> .gitignore`
- `git checkout -b source`
- `cd public` (if theres no public directory, run `hugo`)
- `git init`
- `git remote add origin <location of your hugo project>`

# Usage

Run `hugo-deploy` in the root of your Hugo project.
