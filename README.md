# Notification

# Config

- `DB_DSN`

  ```
  root:notification123456@tcp(127.0.0.1:3306)/notification?parseTime=true&loc=Local&charset=utf8mb4,utf8
  ```

- `SECURITY_JWT`

  Generate with `genpass --jwt` or `openssl rand --hex 32`

- `SECURITY_TELEGRAM_BOT_TOKEN`

  e.g.

  ```
  8800000001:AAXXXXXXXXXXXX-XXXXXXXXX-XXXXXXXXXX
  ```

- `SECURITY_TELEGRAM_CHAT_ID`

  e.g.

  ```
  500000002
  ```

- `PROXY_SOCKS5`

  e.g. `127.0.0.1:1080`

- `PROXY_SOCKS5_USERNAME`

- `PROXY_SOCKS5_PASSWORD`

