created - go mod init githubpath
created .gitignore
created - cmd > students-api > main.go 

Then added a local.yaml configuration file
<!-- Configuration management is very important :
-> Put all configuration variables in .env files
-> File based config is preferred in production becoz we can keep it in version control and we can easily change the config without changing the control, where its version can be tracked.
file based config is preferred -->

Created a storage folder -> storage.db (added this storage folder to .gitignore becoz we don't want out database to go to Repository)