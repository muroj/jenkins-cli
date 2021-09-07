# ghestimator
Estimate resource usage of a Ghenkins job

# Build and run

```
go build -o ghestimator.bin -a ghestimator.go && ./ghestimator.bin --jenkinsUser jmuro --jenkinsAPIToken "$(cat ~/.creds/ghenkins-jmuro)" --jobURL job/watson-engagement-advisor/job/clu-algorithms-service/job/PR-1084 --instanaAPIKey "$(cat ~/.creds/instana-api-token.txt)"
```