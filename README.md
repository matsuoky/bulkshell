# bulkshell
Perform tasks in bulk!

## Usage
- go build
- edit: dir.txt
```
ex) three my repositories are targets
/Users/me/dev/tool/my_repo1
/Users/me/dev/tool/my_repo2
/Users/me/dev/tool/my_repo3
```

- edit: shell.txt
```
ex) two git grep command
git grep -nIi -5 -E "searchAmazon" -- ':!*test*' ':!vendor' ':!*/wp-admin/*' ':!*/wp-includes/*' ':!public' ':!*.sql' ':!*.bk' ':!*.bk.*' ':!*.BAK' ':!*.css' ':!*.scss' ':!*.svg' ':!*.log' | cut -c 1-300
git grep -nIi -5 -E "ItemLookup" -- ':!*test*' ':!vendor' ':!*/wp-admin/*' ':!*/wp-includes/*' ':!public' ':!*.sql' ':!*.bk' ':!*.bk.*' ':!*.BAK' ':!*.css' ':!*.scss' ':!*.svg' ':!*.log' | cut -c 1-300
```
- exec: ./bulkshell
#### Done!
