LDFLAGS := "-s -w"
RUN := CGO_ENABLED=0 go run -ldflags $(LDFLAGS)

RUN_2023:= $(RUN) ./2023

.PHONY: 2023 2023_01 2023_02

2023: 2023_01 2023_02

2023_01:
	$(RUN_2023)/day01_trebuchet/day01_trebuchet.go
2023_02:
	$(RUN_2023)/day02_cube_conundrum/day02_cube_conundrum.go
2023_03:
	$(RUN_2023)/day03_gear_ratios/day03_gear_ratios.go