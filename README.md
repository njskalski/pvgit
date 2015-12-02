# pvgit

pvgit is a pipeviz message producer for working with git repositories. It has two basic operations:

1. `pvgit sync` will report all the refs (branches and tags) in a repository, as well as the current HEAD. If `-a` is passed, it will also send the repository’s entire commit history.
1. `pvgit instrument` is used to set up `post-commit` and/or `post-checkout` hooks on a git repository. These will automatically inform a pipeviz daemon whenever commits are made or branches are changed.

See the `help` output for each command for more details.

# Installation

This package uses [glide](https://github.com/Masterminds/glide) to manage dependencies. It also relies on [git2go](https://github.com/libgit2/git2go) for git interactions via libgit2, which means installation involves CGO. Installing libgit2 is an OS/distro-specific process, so see [git2go](https://github.com/libgit2/git2go#installing) docs on that.

Once libgit2 is findable by pkg-config and glide is on your $PATH, run `make && make install` to install pvgit.
