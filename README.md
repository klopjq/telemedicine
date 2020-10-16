# telemedicine

```
docker build --no-cache --build-arg SSH_PRIVATE_KEY="$(cat ~/.ssh/id_rsa_github)" -t telemedecine:latest -f Dockerfile .
```