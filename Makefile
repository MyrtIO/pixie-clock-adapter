VENV_PATH = ./venv
VENV = . $(VENV_PATH)/bin/activate;

.PHONY: configure
configure:
	rm -rf "$(VENV_PATH)"
	python3.11 -m venv "$(VENV_PATH)"
	$(VENV) pip install -r requirements.txt