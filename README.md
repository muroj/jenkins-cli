# TRON CI

A CLI for interacting with various systems maintained by the TRON team. 

# Build and run

```
go build -o bin/tronci -a ghestimator.go && ./ghestimator.bin --jenkinsUser jmuro --jenkinsAPIToken "$(cat ~/.creds/ghenkins-jmuro)" --jobURL job/watson-engagement-advisor/job/clu-algorithms-service/job/PR-1084 --instanaAPIKey "$(cat ~/.creds/instana-api-token.txt)"
```
