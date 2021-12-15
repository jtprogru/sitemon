# sitemon

Мониторинг сайтов с уведомлением

Основано на статье с Habr:
- [Мониторинг демон на Asyncio + Dependency Injector — руководство по применению dependency injection](https://habr.com/ru/post/514384/)

Каналы уведомлений:
- Telegram
- Email
- etc

Пример конфигурационного файла `config.yml`:
```yaml
log:
  level: "DEBUG"
  format: "[%(asctime)s] [%(levelname)s] [%(name)s]: %(message)s"

sentry:
  dsn: "https://1231231231231231231238cc0375b556@o412493.ingest.sentry.io/5383803"

telegram:
  token: "123456789:qwertyuiopasdfghjkzxcvbnm"
  chat: "-12123123123"

monitors:
  jtprog:
    method: "GET"
    url: "https://jtprog.ru"
    timeout: 15
    check_every: 60

  httpbin:
    method: "GET"
    url: "https://httpbin.org/get"
    timeout: 5
    check_every: 90
```
## Задачи
Вводные: В качестве вводных данных имеется исключительно вышеупомянутый конфиг.

Необходимо реализовать:
- автоматическое перечитываение конфига (по таймеру, по триггеру ФС, по команде боту в телеграм, any way);
- автоматический парсинг объектов мониторинга – секция `monitors` – объекты должны именоваться так же как в конфиге, иметь те же поля и настройки;
- etc
