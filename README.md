Скрипт берет данные с https://corona.lmao.ninja/countries в json формате, выбирает из них страну, которую задаем в начале скрипта и шлет по ней данные в MQTT.

Используется библиотека, которая есть в каталоге с Majordomo (mjdm.ru) - лежит по пути "корень MJD"/3rdparty/phpmqtt/phpMQTT.php
Если вдруг от туда файл phpMQTT.php пропадет, то вот исходник: http://github.com/bluerhinos/phpMQTT
За основу был взят файл publish.php, который лежит в этих исходниках.

Перед использованием подправить (при необходимости) переменные $server, $port, $username, $password, $client_id, $country.
У меня скрипт лежит в папке "корень MJD"/scripts и называется c-virus-mqtt.php. Отсюда и в директиве require (в начале скрипта) такой путь к phpMQTT.php библиотеке.

Сам скрипт поставить в крон (заходим через ssh на сервер majordomo. Далее crontab -e). Например так:

    */30 * * * * cd /var/www/md/scripts/; /usr/bin/php /var/www/md/scripts/c-virus-mqtt.php > /dev/null 2>&1

Запускать раз в 30 минут.

---

Переписал полностью на Go.

Запускать так:

    c-virus.exe -user <имя пользователя на mqtt брокере> -password <пароль для этого пользователя> -broker tcp://10.1.3.35:1883

Где: `tcp://10.1.3.35:1883` - адрес где у меня стоит mqtt брокер.

`-user, -password` можно не указывать если на брокере не настроена аутентификации.

Есть справка. Вызывать `c-virus-mqtt.exe -h`

    Usage of c-virus-mqtt:
        -broker strins
            The broker URI. ex: tcp://localhost:1883 (default "tcp://localhost:1883")
        -country string
            For witch country (optional) (default "Ukraine")
        -id string
            The ClientID (optional) (default "CV-Stats")
        -password string
            The password (optional)
        -timezone string
            Timezone for updated date (optional) (default "Europe/Kiev")
        -topic string
            Topics start at (optional) (default "/coronavirus")
        -user string
            The User (optional)

_Да-да-да, статистику можно получать для разных стран и указывать разные timezone для получаемого поля updated._
