# Octo-Manager
Is a simple tools that is designed to help automate tasks when managing servers yourself

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
 * 'local': Backs up the Server-Directory to the Local-Directory