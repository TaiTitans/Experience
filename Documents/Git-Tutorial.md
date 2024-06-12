
---

### Importing a new project

```
cd project

git init
```

Next, tell Git to take a snapshot of the contents of all files under the current.

`git add <option>`

For option: `.` all file or `name folder`

This snapshot is now stored in a temporary staging area. Next, you can permanently store the contents of the index in repository with :
`git commit -m <message>`


### Making changes
Modify some files, then add their updated contents to the index:
`git add <option>`

Seeing about commited :

`git diff --cached`


You can also get a brief summary of the situation with `git status`


### Viewing project history

```
git log
```


### Managing branches


Create new branch:
```
git checkout -b <name new branch>

```
Show list branch:
```
git branch
```

Switching branch:

```
git switch <name branch>
or
git checkout <name branch> 
```

 To merge the changes made:
```
git merge <branch A> <branch B>
```

### Git command

Clone new project:

```
cd <folder>
git clone <url>
```

Pull data from original project:

```
git pull origin <branch name>
```

Delete file from reponsitory:

```
git rm <name file>
```

**Reset HEAD:**

If dont want commit with file and remove file out staging area:

```
git reset HEAD <file name>
```

---

## More example:

### Add remote url

```bash
git remote add origin <YOUR URI (HTTPS/SSH)>
```

### Pull project with user name and password in command-line

```bash
url Ex: https://Abc.git
    git pull https://username:password@Abc.git
```

### amend commit
 this command let you change the message of previous commit or push the latest change into previous commit.
```bash
    git commit --amend -m "message"
```


### Delete/Remove all of file is not added with git

```bash
git clean -f -d
```

## Stash

### Save stash to local

```bash
    git stash save (save task ,that is not commited yet)
    git stash push -m "message" (save task with the message)
```

### Show stash at local

```bash
    git stash list (show list stash)
```

### Apply stash from local with index

```bash
    git stash apply stash@{0} (apply stash at position is 0)
```


## Git Setting

### Remove your user settings:

```bash
    git config --global --unset user.name
    git config --global --unset user.email
    git config --global --unset credential.helper
```


### Or all your global settings:

```bash
 git config --global --unset-all
```


## Generateig SSH

Using ssh to push/clone code without email/username and password

### Windows Command Line:

```bash
    ssh-keygen -t rsa
```

### macOS:

```bash
    pbcopy < ~/.ssh/id_rsa.pub
```

### GNU/Linux :

```bash
    ssh-keygen -t rsa -b 4096 -C "YOUR_EMAIL_HERE"
```


### To set repository-specific username/email configuration:

```bash
    git config user.name "USERNAME"
    git config user.email "YOUR_EMAIL_HERE@gmail.com"
```

### To fix corrupted interactive rebase

```bash
    git rebase --quit
```

### Move a commit to another branch in git with Cherry pick

```bash
git cherry-pick <commit-code>
```



### Pushing code on two different account from the same system
Step 1: Now generate an SSH key for second accounts

```bash
ssh-keygen -t rsa -C "Github-email-address"
```
Caution: 
It will ask you file to save SSH key, **Don't forget to change**
```
{Home Directory}/.ssh/id_rsa -> {Home Directory}/.ssh/id_rsa_example (Any file name)
```

Step 2: Attach the SSH key at user setting with your second account.
Step 3: Setup GitHub Host

We need to configure when we want to push/clone code on our first account and when we want to push on our second account. Let's create a config file and edit it.

```bash
touch ~/.ssh/config
```
Open it and edit file with your specific file like this :

```
#COMPANY account
Host github.com/taiphan-COMPANY
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_rsa
#personal account
Host github.com-taititans
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_rsa_meonix
```
Step 4: Push/Clone code from each one of two account

Go to the working directory of your project that you want to push on second GitHub account and clone it with the second host name:
Must Use:
```bash
git clone git@github.com-taititans:taititans/Experiences.git
```
Instead of:
```bash
git clone git@github.com:taititans/Experiences.git
```

### Reapplying .gitignore to an existing Git repository

```bash
git rm -r --cached .
```