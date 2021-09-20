# TRON CI

A CLI for interacting with various systems maintained by the TRON team. 

## Build 

```
go install
```

## Run

```
tronci jenkins --url my-jenkins.host.com --user jmuro --api-token "$(cat ~/.creds/ghenkins-jmuro)" get build "job/watson-engagement-advisor/job/clu-algorithms-service/job/master" 
```