# TODO — автобаттлер

Легенда: ✅ сделано · ❌ не сделано · 🐛 есть, но с дефектом/недоделкой

Архитектура: `entity` — только типы · `catalog` — игровые таблицы · `methods` — операции · `events` — сценарии.
Зависимости: `catalog → entity`, `methods → entity, catalog`, `events → всё`.

Статус: игра проходится от выбора класса до «5 побед подряд». Все разделы закрыты; ниже — что осталось на ревью.

---

## 1. Создание персонажа

| ✓ | Файлы | Задача |
|---|-------|--------|
| ✅ | `entity/attribute.go` | Атрибуты как определённые типы + `Attributes` |
| ✅ | `entity/weapon.go`, `catalog/weapons.go` | `DamageType` + `Weapon`, таблица 6 видов оружия |
| ✅ | `entity/classes.go`, `catalog/heroes.go` | Статика класса (`ClassTemplate`) отделена от рандома; `RollAttributes` катает атрибуты 1–3 |
| ✅ | `catalog/heroes.go` | Бонусы классов по ТЗ, упорядочены по уровню (`[3]BonusName`) |
| ✅ | `entity/hero.go`, `methods/hero.go` | `Hero` c `ClassLevels`/`HP`/`MaxHP`; `NewHero` = пустой герой + `LevelUp` стартовым классом |

## 2. Бонусы

| ✓ | Файлы | Задача |
|---|-------|--------|
| ✅ | `entity/bonus.go` | `Bonus` (Title/Description/Income/Outcome/AttributeGain), `BonusName` как слаги |
| ✅ | `catalog/bonuses.go` | 9 бонусов + `damageChecker`; баланс в константах, описания через `Sprintf` |
| ✅ | `methods/hero.go` | `AttributeGain` вызывается в `LevelUp` |
| ✅ | `methods/hero.go` | Бонусы выдаются по одному за уровень (индекс = уровень класса) |

## 3. Модель прогрессии

| ✓ | Файлы | Задача |
|---|-------|--------|
| ✅ | `entity/hero.go` | Мультикласс: `ClassLevels map[HeroClass]uint8`, суммарный уровень = сумма |
| ✅ | `entity/classes.go`, `catalog/heroes.go` | Статика (`HpPerLevel`/`Weapon`/бонусы) отделена от рандомных атрибутов |
| ✅ | `methods/hero.go` | `LevelUp(hero, class)`: +уровень → бонус уровня → `AttributeGain` → `MaxHP += HpPerLevel + Stamina` → хил |
| ✅ | `methods/hero.go` | `TotalLevel`, `HeroTitle` («Воин 1 / Варвар 1») |

## 4. Бой

| ✓ | Файлы | Задача |
|---|-------|--------|
| ✅ | `entity/fight.go`, `entity/combatant.go` | `FightContext` + `Combatant`; `Hero`/`Enemy` приводятся конвертерами |
| ✅ | `entity/enemy.go`, `catalog/enemies.go` | Сущность `Enemy` + таблица 6 монстров |
| ✅ | `methods/fight.go` | Урон: `Weapon.Damage + Strength` → Outcome атакующего → Income защитника |
| ✅ | `methods/fight.go` | Типы урона: скелет ×2 от дробящего; слайм игнорит рубящее (сила и эффекты работают) |
| ✅ | `methods/fight.go` | Спецы через `Feature`: призрак/голем — бонусы, дракон — огонь +3 каждый 3-й ход |
| ✅ | `methods/fight.go` | `Hits` (шанс попадания), `HeroStrikesFirst` (порядок хода), `applyDamage` |
| ✅ | `events/fight.go` | Случайный монстр + пошаговый цикл до смерти одной из сторон |

## 5. После боя

| ✓ | Файлы | Задача |
|---|-------|--------|
| ✅ | `events/afterwards.go` | Победа → хил до максимума |
| ✅ | `events/afterwards.go` | Предложение апа любым классом (мультикласс), лимит суммарного уровня 3 |
| ✅ | `events/afterwards.go` | Предложение заменить оружие на дроп монстра |
| ✅ | `events/afterwards.go`, `main.go` | Поражение → новый персонаж; счётчик побед, 5 подряд → игра пройдена |

## 6. Точка входа

| ✓ | Файлы | Задача |
|---|-------|--------|
| ✅ | `main.go` | Игровой цикл: выбор класса → бой → исход → прокачка/новый герой → до 5 побед |

---

## На ревью (спорные места, решишь сам)

- **`events/beginning.go`** — не трогал (твой). Хардкод-меню `switch`, `os.RemoveAll` в ветке >5 невалидных вводов — оставил как есть.
- **`ImpulseToAction`** множит весь урон (`damage * 2`), а ТЗ говорит «двойной урон **оружием**». Через иммунитет слайма это заметно (порыв удваивает и прибавку силы). Оставил прежнюю семантику бонуса — если надо строго по ТЗ, надо умножать только оружейную часть.
- **Отображение** боя — простой текстовый лог в `events`. Формат на твой вкус.
