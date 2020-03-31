<?php
// Source: http://github.com/bluerhinos/phpMQTT
require "phpMQTT.php";

$server = "localhost";            // change if necessary
$port = 1883;                     // change if necessary
$username = "username";           // set your username
$password = "password";           // set your password
$client_id = "phpMQTT-publisher"; // make sure this is unique for connecting to sever - you could use uniqid()

$date_format = "Y-m-d H:i:s";
$country = "Ukraine";

$fc = file_get_contents('https://corona.lmao.ninja/countries');

$date = new DateTime();
$mqtt = new Bluerhinos\phpMQTT($server, $port, $client_id);

$cv = json_decode($fc, true);

if ($mqtt->connect(true, NULL, $username, $password)) {
  foreach ($cv as $name => $value) {
    if ($value['country'] == $country) {
      $date->setTimestamp(substr($value['updated'], 0, 10));

      $mqtt->publish("/coronavirus/cases", $value['cases'], 0);
      $mqtt->publish("/coronavirus/todayCases", $value['todayCases'], 0);
      $mqtt->publish("/coronavirus/deaths", $value['deaths'], 0);
      $mqtt->publish("/coronavirus/todayDeaths", $value['todayDeaths'], 0);
      $mqtt->publish("/coronavirus/recovered", $value['recovered'], 0);
      $mqtt->publish("/coronavirus/updated", $date->format($date_format), 0);
    }
  }
  $mqtt->close();
}
