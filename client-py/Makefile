PYTHON = venv/bin/python
PIP = venv/bin/pip
PACKAGE_NAME = godis
VENV = venv
DIST_DIR = dist

.PHONY: venv
venv:
	$(PYTHON) -m venv $(VENV)

.PHONY: install
install: venv
	$(PIP) install -r requirements.txt

.PHONY: test
test:
	$(PYTHON) -m pytest tests/

.PHONY: build
build:
	rm -rf $(DIST_DIR)/
	$(PYTHON) -m build

.PHONY: upload
upload: build
	$(PIP) install twine
	$(VENV)/bin/twine upload $(DIST_DIR)/*

.PHONY: clean
clean:
	rm -rf $(DIST_DIR) build/ *.egg-info __pycache__ .pytest_cache

.PHONY: test-install
test-install:
	$(PIP) install $(DIST_DIR)/*.whl
