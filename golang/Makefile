LDFLAGS := "-s -w"
RUN := CGO_ENABLED=0 go run -ldflags $(LDFLAGS)

RUN_2023:= $(RUN) ./2023

.PHONY: 2023 2023_01 2023_02 2023_03 2023_04 2023_05 2023_06

2023: 2023_01 2023_02 2023_03 2023_04 2023_05 2023_06

2023_01:
	$(RUN_2023)/day01_trebuchet/day01_trebuchet.go
2023_02:
	$(RUN_2023)/day02_cube_conundrum/day02_cube_conundrum.go
2023_03:
	$(RUN_2023)/day03_gear_ratios/day03_gear_ratios.go
2023_04:
	$(RUN_2023)/day04_scratchcards/day04_scratchcards.go
2023_05:
	$(RUN_2023)/day05_if_you_give_a_seed_a_fertilizer/day05_if_you_give_a_seed_a_fertilizer.go
2023_06:
	$(RUN_2023)/day_6_wait_for_it/day_6_wait_for_it.go