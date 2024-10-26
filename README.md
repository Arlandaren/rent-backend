# Apartment Rental Service

Этот проект представляет собой бэкенд-решение для сервиса аренды квартир посуточно, реализованное с использованием современных технологий и подходов к архитектуре.

## Стек технологий

- **Go**: Язык программирования, используемый для разработки высокопроизводительных и масштабируемых приложений.
- **gRPC**: Протокол удаленного вызова процедур, обеспечивающий высокую производительность и низкую задержку для коммуникации между микросервисами.
- **Kafka**: Платформа потоковой обработки данных, используемая для реализации событийно-ориентированной архитектуры.
- **PostgreSQL**: Реляционная база данных, обеспечивающая надежное и эффективное хранение данных.
- **Docker**: Платформа контейнеризации, позволяющая упрощать развертывание и масштабирование приложений.

## Архитектурные паттерны

- **Event-Driven Architecture**: Событийно-ориентированная архитектура, обеспечивающая асинхронное взаимодействие между компонентами системы.
- **CQRS**: Паттерн разделения команд и запросов для оптимизации чтения и записи данных.
- **Client-Transport-Service-Repository Representation**: Четкое разделение ответственности между слоями приложения для упрощения модификации и тестирования.

## Возможности

- **Регистрация и аутентификация пользователей**: Безопасное управление пользователями с использованием современных методов аутентификации.
- **Управление объявлениями**: Создание, обновление и удаление объявлений о сдаче квартир в аренду.
- **Поиск и фильтрация**: Возможность поиска квартир по различным критериям, таким как цена, местоположение и доступные удобства.
- **Уведомления**: Реализация уведомлений о изменениях состояния объектов через Kafka.

## Запуск проекта

1. **Клонируйте репозиторий:**

    ```bash
    git clone https://github.com/your-username/apartment-rental-service.git
    cd apartment-rental-service
    ```

2. **Соберите и запустите контейнеры Docker:**

    ```bash
    docker-compose up --build
    ```

3. **Конфигурация:** Убедитесь, что все переменные окружения и конфигурационные файлы настроены должным образом для вашего окружения.

## Разработка

- **Тестирование:** Используйте встроенные тесты для проверки функциональности перед развертыванием.
- **Документация API:** gRPC интерфейсы детально описаны в соответствующих протобуф-файлах.

## Вклад

Принимаются pull requests и обсуждения. Пожалуйста, открывайте issues, если у вас есть предложения или вы обнаружили баги.

## Лицензия

Этот проект лицензируется под лицензией MIT. Подробнее см. файл LICENSE.

