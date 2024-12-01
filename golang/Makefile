LDFLAGS := "-s -w"
RUN := CGO_ENABLED=0 go run -ldflags $(LDFLAGS)

RUN_2023:= $(RUN) ./2023

.PHONY: 2023
2023: 2023_01 2023_02 2023_03 2023_04 2023_05 2023_06 2023_07 2023_09 2023_10 2023_11 2023_12 2023_13

.PHONY: 2023_01
2023_01:
	$(RUN_2023)/day01_trebuchet/day01_trebuchet.go

.PHONY: 2023_02
2023_02:
	$(RUN_2023)/day02_cube_conundrum/day02_cube_conundrum.go

.PHONY: 2023_03
2023_03:
	$(RUN_2023)/day03_gear_ratios/day03_gear_ratios.go

.PHONY: 2023_04
2023_04:
	$(RUN_2023)/day04_scratchcards/day04_scratchcards.go

.PHONY: 2023_05
2023_05:
	$(RUN_2023)/day05_if_you_give_a_seed_a_fertilizer/day05_if_you_give_a_seed_a_fertilizer.go

.PHONY: 2023_06
2023_06:
	$(RUN_2023)/day06_wait_for_it/day06_wait_for_it.go

.PHONY: 2023_07
2023_07:
	$(RUN_2023)/day07_camel_cards/day07_camel_cards.go

.PHONY: 2023_08
2023_08:
	$(RUN_2023)/day08_haunted_wasteland/day08_haunted_wasteland.go

.PHONY: 2023_09
2023_09:
	$(RUN_2023)/day09_mirage_maintenance/day09_mirage_maintenance.go

.PHONY: 2023_10
2023_10:
	$(RUN_2023)/day10_pipe_maze/day10_pipe_maze.go

.PHONY: 2023_11
2023_11:
	$(RUN_2023)/day11_cosmic_expansion/day11_cosmic_expansion.go

.PHONY: 2023_12
2023_12:
	$(RUN_2023)/day12_hot_springs/day12_hot_springs.

.PHONY: 2023_13
2023_13:
	$(RUN_2023)/day13_point_of_incidence/day13_point_of_incidence.go



RUN_2024:= $(RUN) ./2024

.PHONY: 2024
2024: 2024_01

.PHONY: 2024_01
2024_01:
	$(RUN_2024)/day01_historian_hysteria/day01_historian_hysteria.go
