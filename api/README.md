# URL Shortener API Specification


- GET `/new_short_url?url={URL base64 encode} -> {“link”: “{Short link}”}`

- GET `/{short_link} -> переход по ссылке`

-  POST `/admin/set_ttl  body: {“ttl”: время в секундах}`

-  GET `/admin/get_all -> {“total”: количество всех хранимых сокращений, “result”: [ [short, url], [short, url], … ]}`
