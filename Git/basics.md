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
### 1.  Start of Repo
#### Init
Start report `git init`
Start bare repo (cannot edit/commit directly), for storage like github repo `git init --bare`

#### Clone
Clone repo `git clone <repo> <directory_path>`

### 2.  Analysis
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

### 3.  Branching and merging
#### Checkout
Discard unstaged changes `git checkout <filename>`\
Move head to commit/ref `git checkout <commit/ref>`\
Checkout into new branch `git checkout -b <branch_name>`

#### Branch
List all branches `git branch -a`\
Create a branch `git branch <branch_name>`

#### Merge
Merge into current branch `git merge <commit_id/ref>`\
Abort conflicted merge `git merge -abort`

#### Rebase
(TODO)

### 4.  Remote
#### Remote
View remote `git remote -v`\
Add remote `git remote add <name> <url>`

#### Push
Update remote refs `git push <remote> <local_ref>/<remote_ref>` or `git push`, if up-stream is set

#### Fetch
Download objects and refs `git fetch`

#### Pull
Fetch and merge `git pull`

### 5.  Stash

### 6.  Undo / Remove
#### Revert
Revert to parent of commit `git revert <commit>`,underlyingly it is merging the parent of the commit to the head

#### Reset
Reset current head to a specific state/commit. (Alters commit history)

Unstage all staged files (mixed, back to head as its state) `git reset`\
Unstage a file (mixed, back to head as its state) `git reset <file_path>`

Discard all uncommitted changes `git reset --hard`, (see checkout to remove per file)

Uncommit all previous commits (that has commit_id has parent) into staging `git reset --soft <commit>`\
Uncommit and unstage all previous commits `git reset [--mixed] <commit>`\
Remove all previous commits `git reset --hard <commit>`
* If commit has already been pushed, you cannot reset and remove the record

#### Restore (Experimental)
(TODO)

# References:
https://missing.csail.mit.edu/2020/version-control/
