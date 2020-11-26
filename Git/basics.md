# Git Basics
## Data Model
### File: 
```c
type blob = array<byte>
```

### Tree (snapshot): 
```c
type tree = map<string, tree | blob>
```

### Commit: 
```c
struct {
  parents: array<commit>,
  author: string,
  message: string,
  snapshot: tree
}
```
* Snapshot stores the changes to the folder/files
* Commit/tree are merely ptr/hash
* Commits are immutatable
* A merge commit has more than 1 parent

### Object:
Git basic storage unit
```c
type object = blob | tree | commit
type objects = map<string, object>
```
* string contains SHA-1 hash of the object

### Reference
```c
reference = map<string, string>
```
* First string -> is human-readable name. e.g. branch name
* Second string -> SHA-1 of commit
* References are human-readable labels given to commit for better identification
* Branches are in fact reference given to commit

## Commands
### Start of Repo
#### Init
Start report `git init`
Start bare repo (cannot edit/commit directly), for storage like github repo `git init --bare`

#### Clone
Clone repo `git clone <repo> <directory_path>`

### Analysis
#### Status
Show working tree status `git status`

#### Log
View commits in graphical `git log --all --graph --oneline`

#### cat-file
View git internal object (prettier) `git cat-file -p <object_id>`

#### Blame
View file authors by file `git blame <file>`

#### Diff
View diff in file between working directory and staged `git diff <file_path>`\
View diff in file between staged and latest-commit `git diff --cached <commit_id> <file_path>`\
View diff in file between working directory and head `git diff head`

### Branching and merging
#### Checkout
Reset file to last-committed state `git checkout <filename>`\
Move head to commit/ref `git checkout <commit/ref>`\
Checkout into new branch `git checkout -b <branch_name>`

#### Branch
List all branches `git branch -a`\
Create a branch `git branch <branch_name>`

#### Merge
Merge into current branch `git merge <commit_id/ref>`\
Abort conflicted merge `git merge -abort`

#### Rebase (TODO)

### Remote
#### Remote
View remote `git remote -v`\
Add remote `git remote add <name> <url`

#### Push
Update remote refs `git push <remote> <local_ref>/<remote_ref>` or `git push`, if up-stream is set

#### Fetch
Download objects and refs `git fetch`

#### Pull
Fetch and merge `git pull`

### Stash
