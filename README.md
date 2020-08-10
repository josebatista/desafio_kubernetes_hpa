# Desafio Kubernetes + HPA

O desafio consiste em criar uma aplicação em GoLang, para que essa possa "sobrecarregar" o processamento e fazer com o que seja realizado o auto escalonamento horizontal dos pods de acordo com o que foi configurado. E retornar como estava inicialmente quando o "stress" acabar.

## Repositório do container 'go-hpa' no [DockerHub](https://hub.docker.com/r/pereiraze/go-hpa).

## Publicando deployment

```
kubectl apply -f deployment.yaml
```

## Publicando service

```
kubectl apply -f services.yaml
```

## Publicando hpa

```
kubectl apply -f hpa.yaml
```

## O teste do HorizontalPodAutoScaler pode ser realizado conforme abaixo:

```
kubectl run -it loader --image=busybox /bin/sh
```

### Dentro do executar o comando:

```
while true; do wget -q -O - http://go-hpa.default.svc.cluster.local; done;
```
