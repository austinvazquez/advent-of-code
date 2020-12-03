PROJ  = advent_of_code
YEAR ?= 2020
DAY  ?= 0

.PHONY: day
day:
	@mkdir -p advent_of_code/$(YEAR)/$(DAY)
	@touch advent_of_code/$(YEAR)/$(DAY)/input.txt
	@cp advent_of_code/$(YEAR)/0/solution.py advent_of_code/$(YEAR)/$(DAY)/

.PHONY: format
format:
	@black advent_of_code/

.PHONY: lint
lint: format
	@pylint advent_of_code/

.PHONY: run
run:
	@rm -f sample.txt && ln -s advent_of_code/$(YEAR)/$(DAY)/sample.txt sample.txt
	@rm -f input.txt && ln -s advent_of_code/$(YEAR)/$(DAY)/input.txt input.txt
	@python advent_of_code/$(YEAR)/$(DAY)/solution.py

.PHONY: debug
	@rm -f sample.txt && ln -s advent_of_code/$(YEAR)/$(DAY)/sample.txt sample.txt
	@rm -f input.txt && ln -s advent_of_code/$(YEAR)/$(DAY)/input.txt input.txt
	@ipdb advent_of_code/$(YEAR)/$(DAY)/solution.py
