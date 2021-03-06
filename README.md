# REST API для тегов и заметок

Микросервис собран согласно DDD

---

## DDD структура

- **Interface**: Этот уровень отвечает за взаимодействие с пользователем, независимо от того, предоставляет ли программное обеспечение информацию или получает информацию от пользователя.
- **Application**: Это слой между интерфейсом и доменом, он может вызывать службы домена для обслуживания целей приложения.
- **Domain**: Ядро программного обеспечения, этот уровень содержит логику предметной области и бизнес-знания.
- **Infrastructure**: Опорный слой для остальных слоев. Этот уровень содержит вспомогательные библиотеки или внешние службы, такие как база данных или вспомогательная библиотека пользовательского интерфейса.

---

## Запуск проекта
```
make dev
```
