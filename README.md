[![Build Status](https://travis-ci.com/Spendesk/pm2-exporter.svg?token=sPBEAwpUkF9bZVsWNMyJ&branch=master)](https://travis-ci.com/Spendesk/pm2-exporter)

# pm2-exporter
Exporte CPU Usage/Memory Usage and Restart Time PM2 process

## Options
| Name | Env | Parameter | Default | Description | 
|--|--|--|--|--|
| Pm2 Path | PM2_PATH | pm2_path, pp | pm2 | Path to PM2 command if not present in $PATH |
| Refresh | REFRESH | refresh, r | 30 | PM2 status refresh interval |
| Port | PORT | port, p | 10100 | Exporter listening port |
