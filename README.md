# Flygreeting

An example app which provides a selection of country, language and greeting lookups as endpoints. Designed to be integrated into other examples.


## https://flygreeting.fly.dev/v1/countries/ 

returns a JSON list of country codes the application knows

```
❯ curl flygreeting.fly.dev/v1/countries/
{"countries":["AD","AE","AF","AG","AI","AL","AM","AO","AQ","AR","AS","AT","AU","AW","AX","AZ","BA","BB","BD","BE","BF","BG","BH","BI","BJ","BL","BM","BN","BO","BQ","BR","BS","BT","BV","BW","BY","BZ","CA","CC","CD","CF","CG","CH","CI","CK","CL","CM","CN","CO","CR","CU","CV","CW","CX","CY","CZ","DE","DJ","DK","DM","DO","DZ","EC","EE","EG","EH","ER","ES","ET","FI","FJ","FK","FM","FO","FR","GA","GB","GD","GE","GF","GG","GH","GI","GL","GM","GN","GP","GQ","GR","GS","GT","GU","GW","GY","HK","HM","HN","HR","HT","HU","ID","IE","IL","IM","IN","IO","IQ","IR","IS","IT","JE","JM","JO","JP","KE","KG","KH","KI","KM","KN","KP","KR","KW","KY","KZ","LA","LB","LC","LI","LK","LR","LS","LT","LU","LV","LY","MA","MC","MD","ME","MF","MG","MH","MK","ML","MM","MN","MO","MP","MQ","MR","MS","MT","MU","MV","MW","MX","MY","MZ","NA","NC","NE","NF","NG","NI","NL","NO","NP","NR","NU","NZ","OM","PA","PE","PF","PG","PH","PK","PL","PM","PN","PR","PS","PT","PW","PY","QA","RE","RO","RS","RU","RW","SA","SB","SC","SD","SE","SG","SH","SI","SJ","SK","SL","SM","SN","SO","SR","SS","ST","SV","SX","SY","SZ","TC","TD","TF","TG","TH","TJ","TK","TL","TM","TN","TO","TR","TT","TV","TW","TZ","UA","UG","UM","US","UY","UZ","VA","VC","VE","VG","VI","VN","VU","WF","WS","YE","YT","ZA","ZM","ZW"]}
```

## https://flygreeting.fly.dev/v1/languages/ 

returns a JSON list of languages which the application may know

```
curl flygreeting.fly.dev/v1/languages/
{"languages":["aa","af","am","ar","ast","ay","az","be","bg","bi","bn","bs","ca","ch","crs","cs","cy","da","de","dv","dz","el","en","es","et","eu","fa","fi","fo","fr","ga","gd","gl","gn","he","hi","ho","hr","ht","hu","hy","id","is","it","ja","ka","kaa","kk","kl","km","ko","ku","kw","ky","lb","lo","lt","lv","mfe","mg","mh","mi","mk","mn","ms","mt","my","na","nb","nd","ne","niu","nl","nn","no","ny","om","pap","pau","pih","pl","ps","pt","qu","rar","rm","ro","ru","rw","se","sf","sg","si","sk","sl","sm","sn","so","sov","sq","sr","sr-latn","srp","ss","st","sv","sw","ta","tet","tg","th","ti","tk","tkl","tl","tn","tox","tpi","tr","uk","ur","uz","vi","xh","zgh","zh-hans","zh-hant","zu"]}
```

## https://flygreeting.fly.dev/v1/language/:countrycode 

returns the language (or languages) spoke in the country

```
❯ curl flygreeting.fly.dev/v1/language/ch
{"country":"ch","languages":["de","fr","it","rm"]}
```

## https://flygreeting.fly.dev/v1/country/:countrycode 

returns the country details of the country

```
❯ curl flygreeting.fly.dev/v1/country/ch
{"country":{"CountryCode":"CH","CountrynameEnglish":"Switzerland","CountrynameLocal":"Schweiz,Suisse,Svizzera,Svizra","Languages":"de,fr,it,rm"}}
```

## https://flygreeting.fly.dev/v1/greeting/:languagecode 

returns an informal/formal greeting in that language

```
❯ curl flygreeting.fly.dev/v1/greeting/fr
{"language":"fr","formal":"Bonjour","informal":"Salut"}
```

Note that currently the greetings only cover a small number of languages. We welcome pull requests which will increase the coverage or make any of the data in the JSON files more accurate.


