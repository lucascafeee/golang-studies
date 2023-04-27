# Buscador de IPs e Servidores

Este projeto é uma aplicação de linha de comando (CLI) desenvolvida em Go, que busca informações de IPs e servidores de nomes (nameservers) associados a um domínio específico.

## 🚀 Funcionalidades

- 🔎 Buscar endereços IP associados a um domínio
- 🔎 Buscar servidores de nomes (nameservers) associados a um domínio

## 🛠️ Requisitos

- Go 1.16 ou superior
## 📖 Uso

Para usar a aplicação, você pode executar os seguintes comandos:

- Buscar endereços IP associados a um domínio:

```bash
go run main.go ip --host seudominio.com // busca ips
go run main.go servers --host seudominio.com // busca servidores
