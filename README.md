# Desafio Kubernetes + HPA

## Repositório do container 'go_hpa' no [DockerHub](https://hub.docker.com/r/pereiraze/go_hpa).

## O teste do Horizontal AutoScaler pode ser realizado conforme abaixo:
```
*kubectl run -it loader --image=busybox /bin/sh*

### Dentro do executar o comando:

*while true; do wget -q -O- http://go-hpa; done;*
```