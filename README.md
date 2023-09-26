# url-shortener
Сделать web приложение URL-shortener (сервис сокращения размеров ссылки).

По URL http://127.0.0.1:8080/a/?url= должно осуществляться добавление ссылки, например: http://127.0.0.1:8080/a/?url=http%3A%2F%2Fgoogle.com%2F%3Fq%3Dgolang

Возвращается ответ с кодом 200 и 8-ми символьным кодом в теле ответа (/^[A-z0-9]{8}$/). При этом если мы добавляем одну и туже ссылку - код должен быть одним и тем же.

По URL http://127.0.0.1:8080/s/<код> должен осуществляться 302-й редирект на исходную ссылку, в нашем примере http://127.0.0.1:8080/s/pcur1sps перенаправит браузер на http://www.google.com/?q=golang.

Данные можно хранить в любом понравившемся key-value хранилище (persistent). Например http://ssdb.io/docs/go/index.html.

Постараться найти оптимальное по скорости решение. Подумать, что если нужно будет хранить 1M ссылок, 10M, 100M? Куда должно развиваться приложение? Наличие бенчмарка и тестов не обязательно, но крайне приветствуется.

Готовый к анализу код необходимо залить на свой аккаунт github или bitbucket.

**Сокращение ссылки**


<img width="650" alt="Screenshot 2023-09-26 at 2 09 07 PM" src="https://github.com/a-shdv/url-shortener/assets/54847558/80885245-9668-4241-b56a-87a528d5d4cd">


**Получение ссылки по коду**


<img width="650" alt="Screenshot 2023-09-26 at 2 11 00 PM" src="https://github.com/a-shdv/url-shortener/assets/54847558/bc443a8b-e560-43c1-9a2d-75ecb154ee17">


**Создание кастомной короткой ссылки**


<img width="650" alt="Screenshot 2023-09-26 at 2 21 48 PM" src="https://github.com/a-shdv/url-shortener/assets/54847558/2325cf11-bdc4-4f97-aff4-6daca617c897">


**Результат работы в Redis**


<img width="650" alt="Screenshot 2023-09-26 at 2 47 09 PM" src="https://github.com/a-shdv/url-shortener/assets/54847558/8d4482f1-2473-4cf7-9dc6-e23dafe74d43">


