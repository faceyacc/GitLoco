
# GitLoco
GitLoco is Git implementation inspired by *Thibault Polge's Write yourself a Git!* GitLoco can initialize a repository, create commits and clone public repository from GitHub.



## Usage/Examples


Initalize a Git repository 
```bash
gitloco init
```
Store the data from `<file>` as a blob in `.git/objects` and print a 40-char SHA to stdout
```bash
gitloco hash-object --w=<file>
```
Print the raw contents of an object to stdout, uncompressed and header removed
```bash
gitloco cat-file <blob_sha>
```
Create a snapshot of your curreent git repository
```bash
gitloco write-tree
```
Inspect a tree object and list the contents of a tree object
```bash
gitloco ls-tree <tree_sha>
```
## Roadmap

- [ ] Publish and distribute with Homebrew.
- [x] Add `write-tree` command to create tree objects.
- [ ] Add  `commit-tree` command to allow users to create a commit using a tree_sha.
- [ ] Add `clone` command to allow users to clone public repos from Github.




## Support

For support, email me at justfacey@gmail.com.

