# Example for how to add services records to consul
```
#retrun "0" means ok
curl -H 'Content-Type: application/json' -X PUT -d '{
  "ID": "redis1",
  "Name": "redis",
  "Tags": [
    "redis1"
  ],
  "Address": "127.0.0.1",
  "Port": 8000,
  "check": {
    "id": "api",
    "name": "HTTP API on port 4567",
    "http": "http://127.0.0.1:4567/health",
    "interval": "10s",
    "timeout": "1s"
  }
}'  localhost:8500/v1/agent/service/register
```
