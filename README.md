# JWTCLI

small binary for generating JWT tokens in cli.

```bash
curl -o jwtcli https://github.com/goblain/jwtcli/releases/download/0.0.1/jwtcli-linux-amd64
chmod +x jwtcli
./jwtcli create --iss issuer_uri --sub johnny@bravo.toons --aud aud1,aud2 --exp 300s 
```