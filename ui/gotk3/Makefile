all: ui-builder

# generate file with UI builder XML stored in a variable
ui-builder:
	echo "package gotk3" > prompt_ui_builder.go
	echo "" >> prompt_ui_builder.go
	echo 'var promptUIXMLBuilder = `' >> prompt_ui_builder.go
	cat ui.xml >> prompt_ui_builder.go
	echo '`' >> prompt_ui_builder.go