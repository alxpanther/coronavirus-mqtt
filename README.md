Скрипт берет данные с https://corona.lmao.ninja/countries в json формате, выбирает из них страну, которую задаем в начале скрипта и шлет по ней данные в MQTT.
Используется библиотека, которая есть в каталоге с Majordomo (mjdm.ru) - лежит по пути "корень MJD"/3rdparty/phpmqtt/phpMQTT.php
Если вдруг от туда файл phpMQTT.php пропадет, то вот исходник: http://github.com/bluerhinos/phpMQTT
За основу был взят файл publish.php, который лежит в этих исходниках.

Перед использованием подправить (при необходимости) переменные $server, $port, $username, $password, $client_id, $country.
У меня скрипт лежит в папке "корень MJD"/scripts и называется c-virus-mqtt.php. Отсюда и в директиве require (в начале скрипта) такой путь к phpMQTT.php библиотеке.

Сам скрипт поставить в крон (заходим через ssh на сервер majordomo. Далее crontab -e). Например так:

    */30 * * * * cd /var/www/md/scripts/; /usr/bin/php /var/www/md/scripts/c-virus-mqtt.php > /dev/null 2>&1

Запускать раз в 30 минут.
