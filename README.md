# GoCleaner 🧹

**GoCleaner** é uma ferramenta CLI escrita em Go que automatiza a limpeza de arquivos e diretórios com mais de N dias, com suporte a:

- Remoção segura e paralela com **Goroutines**
- Configuração via arquivo **YAML**
- Envio de **relatório por e-mail** via Microsoft 365 (SMTP)
- Log de auditoria em arquivo `log.txt`
- Build para **Windows, Linux e macOS**

## ⚙️ Exemplo do `config.yaml`

```yaml
directory: "E:\\Transfer"
days_threshold: 30

schedule:
  enabled: true
  cron: "0 0 * * *"
  
smtp:
  host: "smtp.office365.com"
  port: 587
  username: "e-mail"
  password: "senha"
  to: "destinatario"