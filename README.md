# pg_lz
PostgreSQL Lazy App

This app intends to help DBAs with daily tasks like backups, minor and major version upgrades with minimum downtime among others.

## TODO

- [] criar scripts para:
    - [ ] criar cluster com arquivamento
 
## plano implementação

<!-- - [ ] Garantir WAL archive antes de iniciar o backup -->
- [x] Criar backup label no banco (pg_backup_start)
- [ ] Copiar os arquivos incrementalmente
	- [ ] Utilizar um rsync ou outra lib similar
	- [ ] Copiar blocos em paralelo 
	- [ ] Utilizar compressao pra enviar pela rede
- [ ] Finalizar o backup no PG com pg_backup_stop
- [ ] Fazer um segundo passe do rsync pra garantir que novos arquivos sejam copiados ou atualizados
- [ ] Criar um logical replication slot no primary a ser usado pela subscription
- [ ] Configurar a nova instancia como replica fisica do primary
- [ ] Subir a instancia pra forcar a replicacao e eliminar lag e avancar o LSN ate um ponto apos o LSN da subscription e entao promover
- [ ] Shutdown do banco de forma limpa
- [ ] Atualizar a versao do banco utilizando o pg_upgrade
- [ ] Criar uma subscription utilizando o slot criado anteriormente a partir da LSN utilizada para avancar a replica e promover
- [ ] Deixar os bancos ficarem em SYNC
- [ ] Desligar o antigo primary e configurar como read-only
- [ ] Criar um logical replication slot no novo primary para que o antigo primary fique em SYNC em caso de ter que voltar a versao
- [ ] Subir o antigo primary como read-only 
- [ ] Criar uma subscription no antigo primary