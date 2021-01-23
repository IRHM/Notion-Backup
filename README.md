# Notion-Backup
Backup all of your notion notes, and optionally upload them to a git repo.

## Usage

1. Create a folder where you want to export your notion notes to and optionally make it a git repo
2. Run Notion-Backup executable, if you don't want to backup to a git repo append `nogit` to the end of the command (eg. `./Notion-Backup nogit`)
3. (Happens only first time) Enter your API Token (look below on how to get it)
4. (Happens only first time) Enter the path to the folder you created in step 1
5. Wait for Notion-Backup to download your notes
6. If asked for authentication:
  - Windows
    On windows a prompt will open allowing you to login to Git through the web browser or by using a personal access token.
  - Linux
    On Linux you will be asked for a username and password in the command line. If you have 2FA enabled, use a personal access token
    as your password with the `repo` privilege set.

  Optionally you can save your login credentials for Git so you don't have to authenticate everytime: 
  ```
  # More info: https://stackoverflow.com/a/35942890/11025032
  git config --global credential.helper store
  git pull
  ```
7. Done.. providing you have Git installed and it's in your PATH variable (if using git), Notion-Backup will do everything else for you

## Getting your notion API Key

1. Open `Inspect Element` (Ctrl+Shift+I) and go to the `Network tab`
2. Refresh the page (whilst logged in) so network requests show up
3. Click on the `first request` (most of the other requests will also work) and in the new menu click on `Cookies`
4. In `Request Cookies` find `token_v2` and copy its value (eg. `token_v2:	"myTokenHere"` in this case, your token is: `myTokenHere`)
5. If you don't see `token_v2`, click on a different request and then repeat step 4
6. Done.. copy your API Token into Notion-Backup
