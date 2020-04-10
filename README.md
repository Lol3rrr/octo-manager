# Octo-Manager
Is a simple tools that is designed to help automate tasks when managing servers yourself

## Download
To download the latest Release you can just use the link `https://github.com/Lol3rrr/octo-manager/releases/latest/download/{distro}`.
Simply replace the `{distro}` with the one you need.
Possible Distros:
 * 'linux_x64': 64bit Linux

## Setup

There are 2 types of configs needed for this to work
 * `config.json` contains all the Server-Configuration, so address, username etc.
 * `job.json` contains all the Job-Configuration, so the Name and all the different Stages for the Job

### config.json
 * `address`: The Address of the Server
 * `port`: The Port used for SSH
 * `username`: The Username that should execute all commands
 * `sshkeyfile`: The File-Path to the fitting SSH-Key

### job.json
 * `name`: The Name of this Job
 * `stages`: An array of Stages

#### Stage
 * `name`: The Name for this Stage
 * `description`: A brief description of this job
 * `category`: The Category of this Stage (docker, backups, etc.)
 * `action`: The Specific Action that should be executed
 * `variables`: Any needed hard-coded variables for the Action

## Modules

### Docker ('docker')
Actions:
 * 'pull': Pulls the newest Image
 * 'compose up': Simply starts up Docker-Compose

### Backup ('backup')
Actions:
 * 'save-local': Backs up the Server-Directory to the Local-Directory
 * 'restore-local': Restores the Server-Directory from the Local-Directory
 * 'delete-local': Deletes all Backup-Directorys in the Local-Directory older than X hours
 * 'save-googleDrive': Backs up the Server-Directory to a new Directory in Google Drive
 * 'restore-googleDrive': Restores the Server-Directory from the Directory in Google Drive
 * 'delete-local': Deletes all Backup-Directorys older than X hours

## Command Line Flags
 * 'config': Used to specify another file for the config-file
 * 'job': Used to specify another file for the job-file
 * 'auth': Used if you want to auth to use different modules, like googleDrive