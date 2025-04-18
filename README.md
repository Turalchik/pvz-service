# Тестовое задание для стажёра Backend-направления (весенняя волна 2025)

## Сервис для работы с ПВЗ

На ПВЗ несколько раз в день привозят новые товары, которые были заказаны через Авито. Прежде чем их отдавать заказчику, необходимо сначала проверить и внести информацию в базу. Из-за того, что ПВЗ много, а товаров ещё больше, нужно реализовать механизм, позволяющий в разрезе каждого ПВЗ увидеть, сколько раз в день к ним приезжали товары на приёмку и какие товары были получены.

## Описание задачи

Разработайте backend-сервис для сотрудников ПВЗ, который позволит вносить информацию по заказам в рамках приёмки товаров.

1. Авторизация пользователей:
   * Используя ручку /dummyLogin и передав в неё желаемый тип пользователя (client, moderator),
     сервис в ответе вернёт токен с соответствующим уровнем доступа — обычного пользователя или модератора.
     Этот токен нужно передавать во все endpoints, требующие авторизации.

2. Регистрация и авторизация пользователей по почте и паролю:
   * При регистрации используется endpoint /register. 
   В базе создаётся и сохраняется новый пользователь желаемого типа: обычный пользователь (client) или модератор (moderator).
   У созданного пользователя появляется токен endpoint /login. 
   При успешной авторизации по почте и паролю возвращается токен для пользователя с соответствующим ему уровнем доступа.

3. Заведение ПВЗ:
   * Только пользователь с ролью «модератор» может завести ПВЗ в системе.
   * В случае успешного запроса возвращается полная информация о созданном ПВЗ. Заведение ПВЗ возможно только в трёх городах: Москва, Санкт-Петербург и Казань. В других городах ПВЗ завести на         первых порах нельзя, в таком случае необходимо вернуть ошибку.
   * Результатом добавления ПВЗ должна стать новая запись в хранилище данных

4. Добавление информации о приёмке товаров:
   * Только авторизованный пользователь системы с ролью «сотрудник ПВЗ» может инициировать приём товара.
   * Результатом инициации приёма товаров должна стать новая запись в хранилище данных.
   * Если же предыдущая приёмка товара не была закрыта, то операция по созданию нового приёма товаров невозможна.

5. Добавление товаров в рамках одной приёмки:
   * Только авторизованный пользователь системы с ролью «сотрудник ПВЗ» может добавлять товары после его осмотра.
   * При этом товар должен привязываться к последнему незакрытому приёму товаров в рамках текущего ПВЗ.
   * Если же нет новой незакрытой приёмки товаров, то в таком случае должна возвращаться ошибка, и товар не должен добавляться в систему.
   * Если последняя приёмка товара все ещё не была закрыта, то результатом должна стать привязка товара к текущему ПВЗ и текущей приёмке с последующем добавлением данных в хранилище.

6. Удаление товаров в рамках не закрытой приёмки:
   * Только авторизованный пользователь системы с ролью «сотрудник ПВЗ» может удалять товары, которые были добавлены
   в рамках текущей приёмки на ПВЗ.
   * Удаление товара возможно только до закрытия приёмки, после этого уже невозможно изменить состав товаров, которые
   были приняты на ПВЗ.
   * Удаление товаров производится по принципу LIFO, т.е. возможно удалять товары только в том порядке, в котором
   они были добавлены в рамках текущей приёмки.

7. Закрытие приёмки:
   * Только авторизованный пользователь системы с ролью «сотрудник ПВЗ» может закрывать приём товаров.
   * В случае, если приёмка товаров уже была закрыта (или приёма товаров в данном ПВЗ ещё не было), 
   то следует вернуть ошибку.
   * Во всех остальных случаях необходимо обновить данные в хранилище и зафиксировать товары,
   которые были в рамках этой приёмки.

8. Получение данных:
   * Только авторизованный пользователь системы с ролью «сотрудник ПВЗ» или «модератор» может получать эти данные.
   * Необходимо получить список ПВЗ и всю информацию по ним при помощи пагинации.
   * При этом добавить фильтр по дате приёмки товаров, т.е. выводить только те ПВЗ и всю информацию по ним,  которые в указанный диапазон времени 
     проводили приёмы товаров.

## Общие вводные

У сущности «Пункт приёма заказов (ПВЗ)» есть:
* Уникальный идентификатор
* Дата регистрации в системе
* Город

У сущности «Приёмка товара» есть:
* Уникальный идентификатор
* Дата и время проведения приёмки
* ПВЗ, в котором была осуществлена приёмка
* Товары, которые были приняты в рамках данной приёмки
* Статус (in_progress, close)

У сущности «Товар» есть:
* Уникальный идентификатор
* Дата и время приёма товара (дата и время, когда товар был добавлен в систему в рамках приёмки товаров)
* Тип (мы работаем с тремя типами товаров: электроника, одежда, обувь)

## Условия

1. Используйте этот [API](swagger.yaml).
2. Реализуйте все требования, которые прописаны в условиях задания.
3. Сервер должен быть запущен на порту 8080.
4. Реализация пользовательской авторизаций не является обязательным условием. В таком случае токен авторизации можно получить из метода /dummyLogin.

   В параметрах запроса можно выбрать роль пользователя: модератор или обычный пользователь.
   В зависимости от роли будет сгенерирован токен с определённым уровнем доступа.
6. Нефункциональные требования:
   * RPS — 1000 
   * SLI времени ответа — 100 мс 
   * SLI успешности ответа — 99.99%
7. Код обязательно должен быть покрыт unit-тестами. Тестовое покрытие не менее 75%.
8. Должен быть разработан один интеграционный тест, который:
   * Первым делом создает новый ПВЗ
   * Добавляет новую приёмку заказов
   * Добавляет 50 товаров в рамках текущей приёмки заказов
   * Закрывает приёмку заказов

## Дополнительные задания

Не являются обязательными, но дадут вам преимущество перед другими кандидатами.

1. Реализовать пользовательскую авторизацию по методам /register и /login 
   (при этом метод /dummyLogin все равно должен быть реализован)
2. Реализовать gRPC-метод, который просто вернёт все добавленные в систему ПВЗ. Для него не
   требуется проверка авторизации и валидация ролей пользователей. Сервер для gRPC должен быть запущен на порту 3000.
   Обратите внимание, что в файле [pvz.proto](pvz.proto) необходимо прописать go_package под вашу структуру проекта
3. Добавить в проект prometheus и собирать следующие метрики:
   * Технические: 
     * Количество запросов
     * Время ответа
   * Бизнесовые:
     * Количество созданных ПВЗ
     * Количество созданных приёмок заказов
     * Количество добавленных товаров
   Сервер для prometheus должен быть поднят на порту 9000 и отдавать данные по ручке /metrics.
4. Настроить логирование в проекте
5. Настроить кодогенерацию DTO endpoint'ов по openapi схеме

## Требования по стеку

* **Язык сервиса:** предпочтительным является Go, но также допустимы следующие языки: PHP, Java, Python, C#.
* **База данных:** предпочтительно PostgreSQL, но можно выбрать другую удобную вам. Нельзя использовать ORM для взаимодействия с базой 
* Допустимо использовать **билдеры для запросов**, например, такой: https://github.com/Masterminds/squirrel
* Для деплоя зависимостей и самого сервиса нужно использовать Docker или Docker & DockerCompose

## Дополнения к решению
* Если у вас возникнут вопросы, ответов на которые нет в условиях, то принимать решения можете самостоятельно 
* В таком случае приложите к проекту Readme-файл со списком вопросов и объяснениями своих решений

## Оформление и отправка решения

Создайте публичный git-репозиторий на любом хосте (GitHub, GitLab и другие), содержащий в master/main ветке:

1. Код сервиса
2. Docker или Docker & DockerCompose или описанную в Readme.md инструкцию по запуску
3. Описанные в Readme.md вопросы или проблемы, с которыми вы столкнулись, и описание своих решений
