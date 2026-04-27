module github.com/Yandex-Practicum/tracker

go 1.24.1

require spentcalories v0.0.0-00010101000000-000000000000

require daysteps v0.0.0-00010101000000-000000000000 // indirect

replace spentcalories => ./internal/spentcalories

replace daysteps => ./internal/daysteps
